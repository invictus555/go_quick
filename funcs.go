package main

import (
	"encoding/json"
	"fmt"
	"time"

	"quick_practices/service"
	"quick_practices/utils"
)

const (
	resourceKey = "1180_vod_5"
)

func Permission() {
	service.CheckPermission("shengchao", utils.NameSpace, resourceKey, utils.Region, "Read")
}

func Permissions() {
	var req []*service.PermissionRequest
	req = append(req, &service.PermissionRequest{
		User:       "shenhchao",
		Attributes: []string{"Read", "Read_Write"},
	})
	service.CheckBatchPermissions(utils.NameSpace, resourceKey, utils.Region, req)
}

func AddAdmins() {
	reqs := []*service.AuthorizationRequest{
		{
			User:       "wangdarui",
			Attributes: []string{"Read", "Read_write"},
			Duration:   time.Hour * 24 * 180,
		},
		{
			User:       "gaohonglei",
			Attributes: []string{"Read", "Read_write"},
			Duration:   time.Hour * 24 * 180,
		},
	}

	service.AddAdmins(utils.NameSpace, utils.Region, resourceKey, reqs)
}

func RegisterResource() {
	region := "cn"
	ns := "kani_3541"
	if region == "boe" {
		ns = "kani_34710"
	}
	resourceKeyCn := "1180_vod_7"
	service.RegisterResource(resourceKeyCn, ns, region, []string{"shengchao"}, []string{})
}

func Serialization() {
	reqs := []*service.AuthorizationRequest{
		{
			User:       "wangdarui",
			Attributes: []string{"Read", "Read_write"},
			Duration:   time.Hour * 24 * 180,
		},
		{
			User:       "gaohonglei",
			Attributes: []string{"Read", "Read_write"},
			Duration:   time.Hour * 24 * 180,
		},
	}

	body, err := json.Marshal(reqs)
	if err != nil {
		fmt.Sprintln(nil)
		return
	}

	fmt.Println(string(body))

	var sers []service.AuthorizationRequest
	err = json.Unmarshal(body, &sers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range sers {
		fmt.Println(v)
	}
}
