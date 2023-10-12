package goset

import "fmt"

func TestifySetImplement() {
	set := NewSet[int64]()
	set.Add(1)
	set.Add(2)
	set.Add(4)
	set.Add(2)

	fmt.Println(set.Values())
}
