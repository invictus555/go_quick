package utils

import "errors"

var (
	DBCreateAppInfoFailed    = errors.New("create app info failed")
	DBCreateStrategyFailed   = errors.New("create strategy info failed")
	DBCreateTenantInfoFailed = errors.New("create tenant info failed")
)

var (
	KaniInputParametersError       = errors.New("input parameters error")
	KaniAuthorizationFailed        = errors.New("kani authorization failed")
	KaniResourceRegistrationFailed = errors.New("kani resource registration failed")
)
