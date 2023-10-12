package strategy

import (
	"encoding/json"
	"fmt"
)

func CallAllTestifyFunc() {
	// TestifyCreateTicketWhenAddMP4Rule(shengchao, zhangyusf)
	// TestifyCreateTicketWhenAddMp3Rule(zhangyusf, shengchao)
	// TestifyCreateTicketWhenAddPreloadRule(shengchao, shengchao)
	// TestifyCreateTicketWhenEditMKVRule(shengchao, zhangyusf)
	// TestifyAddMKVRuleStrategy(shengchao)
	// TestifyCreateTicket4TiktokDataMiddleWareV3(zhangyusf, shengchao)
	// TestifyCreateTicket4TiktokSeekingOptimization(zhangyusf, zhangyusf)
	TestifyCreateTicket4DecodingDoblySettings(shengchao, shengchao)
	// TestifyCreateTicket4TiktokfeedFullCacheSettings(zhangyusf, zhangyusf)
}

// 测试新增工单
func TestifyCreateTicketWhenAddMP4Rule(creator, reviewer string) {
	ruleGroupJSON, err := ruleGroups2JSON(hardwareDecodeMp4Ability())
	if err != nil {
		panic(err)
	}

	fmt.Sprintln(string(ruleGroupJSON))

	desc := "to test create new rule group"
	req := &CreateTicketRequest{
		TenantID:    3,
		ModuleName:  "vod",
		Reviewer:    reviewer,
		Description: &desc,
		AfterRules:  ruleGroupJSON,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	url := "https://vip-boe-i18n.byted.org/ker/api/ticket/create_ticket"
	resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	fmt.Println(string(resp), err)
}

func TestifyCreateTicketWhenAddMp3Rule(creator, reviewer string) {
	ruleGroupJSON, err := ruleGroups2JSON(hardwareDecodeMp3Ability())
	if err != nil {
		panic(err)
	}

	fmt.Sprintln(ruleGroupJSON)

	desc := "to test create new rule group"
	req := &CreateTicketRequest{
		TenantID:    3,
		ModuleName:  "vod",
		Reviewer:    reviewer,
		Description: &desc,
		AfterRules:  ruleGroupJSON,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	url := "https://vip-boe-i18n.byted.org/ker/api/ticket/create_ticket"
	resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	fmt.Println(string(resp), err)
}

func TestifyCreateTicketWhenAddPreloadRule(creator, reviewer string) {
	ruleGroupJSON, err := ruleGroups2JSON(mainFeedPreloadSettings())
	if err != nil {
		panic(err)
	}

	fmt.Println(ruleGroupJSON)

	desc := "new rule about main feed preload"
	req := &CreateTicketRequest{
		TenantID:    3,
		ModuleName:  "vod",
		Reviewer:    reviewer,
		Description: &desc,
		AfterRules:  ruleGroupJSON,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	url := "https://vip-boe-i18n.byted.org/ker/api/ticket/create_ticket"
	resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	fmt.Println(string(resp), err)
}

func TestifyCreateTicketWhenEditMKVRule(creator, reviewer string) {
	Before, err := ruleGroups2JSON(hardwareDecodeMKVAbility(1, 1, 1, 1))
	if err != nil {
		panic(err)
	}

	After, err := ruleGroups2JSON(hardwareDecodeMKVAbilityV1(1, 1, 1, 1))
	if err != nil {
		panic(err)
	}

	fmt.Println(Before)
	fmt.Println(After)

	description := "ticket with after rules and before rules"
	req := &CreateTicketRequest{
		TenantID:    3,
		ModuleName:  "vod",
		Reviewer:    reviewer,
		Description: &description,
		BeforeRules: Before,
		AfterRules:  After,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(reqJSON)

	url := "https://vip-boe-i18n.byted.org/ker/api/ticket/create_ticket"
	resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	fmt.Println(string(resp), err)
}

// 测试添加vod策略
func TestifyAddMKVRuleStrategy(creator string) {
	req := &AddVodStrategyRequest{
		Rule: hardwareDecodeMKVAbility(0, 0, 0, 0)[0],
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reqJSON))

	// url := "https://vip-boe-i18n.byted.org/ker/api/vod/add_strategy"
	// resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	// fmt.Println(string(resp), err)
}

func TestifyCreateTicket4TiktokDataMiddleWareV3(creator, reviewer string) {
	ruleGroupJSON, err := ruleGroups2JSON(dataMiddleWareV3Settings())
	if err != nil {
		panic(err)
	}

	fmt.Println(ruleGroupJSON)

	desc := "tiktok data middleware v3"
	req := &CreateTicketRequest{
		TenantID:    3,
		ModuleName:  "vod",
		Reviewer:    reviewer,
		Description: &desc,
		AfterRules:  ruleGroupJSON,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(reqJSON)

	url := "https://vip-boe-i18n.byted.org/ker/api/ticket/create_ticket"
	resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	fmt.Println(string(resp), err)
}

func TestifyCreateTicket4TiktokSeekingOptimization(creator, reviewer string) {
	ruleGroupJSON, err := ruleGroups2JSON(tiktokSeekOptSettings())
	if err != nil {
		panic(err)
	}

	fmt.Println(ruleGroupJSON)

	desc := "tiktok seeking optimization"
	req := &CreateTicketRequest{
		TenantID:    3,
		ModuleName:  "vod",
		Reviewer:    reviewer,
		Description: &desc,
		AfterRules:  ruleGroupJSON,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reqJSON))

	url := "https://vip-boe-i18n.byted.org/ker/api/ticket/create_ticket"
	resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	fmt.Println(string(resp), err)
}

func TestifyCreateTicket4TiktokfeedFullCacheSettings(creator, reviewer string) {
	ruleGroupJSON, err := ruleGroups2JSON(feedFullCacheV2Settings())
	if err != nil {
		panic(err)
	}

	fmt.Println(ruleGroupJSON)

	desc := "tiktok feed cache optimization"
	req := &CreateTicketRequest{
		TenantID:    3,
		ModuleName:  "vod",
		Reviewer:    reviewer,
		Description: &desc,
		AfterRules:  ruleGroupJSON,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reqJSON))

	url := "https://vip-boe-i18n.byted.org/ker/api/ticket/create_ticket"
	resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	fmt.Println(string(resp), err)
}

func TestifyCreateTicket4DecodingDoblySettings(creator, reviewer string) {
	ruleGroupJSON, err := ruleGroups2JSON(hardwareDecodeDoblyAbility())
	if err != nil {
		panic(err)
	}

	fmt.Println(ruleGroupJSON)

	desc := "tiktok dobly decoding ablity"
	req := &CreateTicketRequest{
		TenantID:    3,
		ModuleName:  "vod",
		Reviewer:    reviewer,
		Description: &desc,
		AfterRules:  ruleGroupJSON,
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reqJSON))

	url := "https://vip-boe-i18n.byted.org/ker/api/ticket/create_ticket"
	resp, err := doHttpPost(url, name2jwt[creator], "boe_clouddev5663", reqJSON)
	fmt.Println(string(resp), err)
}
