package main

import (
	"quick_practices/gorm"
	"quick_practices/service"
	"quick_practices/utils"
	"time"
)

const (
	resourceKey = "1180_vod_5"
)

func main() {
	gorm.TestifyHasManyUpdate()
}

/*-----------------------------------------------------------------------------------------*/
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
		&service.AuthorizationRequest{
			User:       "wangdarui",
			Attributes: []string{"Read", "Read_write"},
			Duration:   time.Hour * 24 * 180,
		},
		&service.AuthorizationRequest{
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
