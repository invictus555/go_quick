// Copyright 2021 ByteDance Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gopool

import (
	"context"
	"sync"
	"sync/atomic"
)

type Pool interface {
	Go(f func())                                          // Go executes f.
	Name() string                                         // Name returns the corresponding pool name.
	SetCap(cap int32)                                     // SetCap sets the goroutine capacity of the pool.
	WorkerCount() int32                                   // WorkerCount returns the number of running workers
	CtxGo(ctx context.Context, f func())                  // CtxGo executes f and accepts the context.
	SetPanicHandler(f func(context.Context, interface{})) // SetPanicHandler sets the panic handler.
}

var taskPool sync.Pool // 使用后的task必须回到对象池

func init() {
	taskPool.New = newTask // 设置New函数
}

// 任务信息
type task struct {
	ctx  context.Context
	f    func()
	next *task
}

func (t *task) zero() {
	t.ctx = nil
	t.f = nil
	t.next = nil
}

func (t *task) Recycle() {
	t.zero()
	taskPool.Put(t)
}

func newTask() interface{} {
	return &task{}
}

type pool struct {
	name         string                             // The name of the pool
	cap          int32                              // capacity of the pool
	config       *Config                            // Configuration information
	taskHead     *task                              // task list head
	taskTail     *task                              // task list tail
	taskLock     sync.Mutex                         // mutex lock for sync
	taskCount    int32                              // record total count of task
	workerCount  int32                              // Record the number of running workers
	panicHandler func(context.Context, interface{}) // This method will be called when the worker panic
}

// NewPool creates a new pool with the given name, cap and config.
func NewPool(name string, cap int32, config *Config) Pool {
	p := &pool{
		name:   name,
		cap:    cap,
		config: config,
	}
	return p
}

func (p *pool) Name() string {
	return p.name
}

func (p *pool) SetCap(cap int32) {
	atomic.StoreInt32(&p.cap, cap)
}

func (p *pool) Go(f func()) {
	p.CtxGo(context.Background(), f)
}

func (p *pool) CtxGo(ctx context.Context, f func()) {
	t := taskPool.Get().(*task) // 从池中闲置的task信息
	t.ctx = ctx
	t.f = f

	p.taskLock.Lock()
	if p.taskHead == nil { // 任务入队列
		p.taskHead = t
		p.taskTail = t
	} else {
		p.taskTail.next = t
		p.taskTail = t
	}
	p.taskLock.Unlock()

	atomic.AddInt32(&p.taskCount, 1) // 任务数目+1
	// The following two conditions are met:
	// 1. the number of tasks is greater than the threshold.
	// 2. The current number of workers is less than the upper limit p.cap.
	// or there are currently no workers.
	if (atomic.LoadInt32(&p.taskCount) >= p.config.ScaleThreshold && p.WorkerCount() < atomic.LoadInt32(&p.cap)) || p.WorkerCount() == 0 {
		p.incWorkerCount()
		w := workerPool.Get().(*worker) // 获取一个工作对象
		w.pool = p
		w.run() // 启动一个协程
	}
}

// SetPanicHandler the func here will be called after the panic has been recovered.
func (p *pool) SetPanicHandler(f func(context.Context, interface{})) {
	p.panicHandler = f
}

func (p *pool) WorkerCount() int32 {
	return atomic.LoadInt32(&p.workerCount)
}

func (p *pool) incWorkerCount() {
	atomic.AddInt32(&p.workerCount, 1)
}

func (p *pool) decWorkerCount() {
	atomic.AddInt32(&p.workerCount, -1)
}
