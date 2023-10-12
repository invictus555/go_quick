package strategy

// 构建解码MKV的规则例子
func hardwareDecodeMKVAbility(ruleGroupID, CountryGroupID, PlatformGroupID, RuleID int64) []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          ruleGroupID,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "hardware_decode_mkv_ability",
			Description: "this is a test rule",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          CountryGroupID,
					CountryList: []string{"US", "CN"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       PlatformGroupID,
							Platform: "android",
							Rules: []RuleInfo{
								{
									ID:        RuleID,
									Status:    true,
									Name:      "mkv video",
									Priority:  1,
									Condition: "platform=='android'",
									DecisionInfoList: []*DecisionInfo{
										{
											// 省略了Name
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "hardware",
													Value:    "'mediacodec'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "video",
													Value:    "'mkv'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "max_size",
													Value:    "409600000",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// 构建解码MKV的规则例子
func hardwareDecodeMKVAbilityV1(ruleGroupID, CountryGroupID, PlatformGroupID, RuleID int64) []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          ruleGroupID,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "hardware_decode_mkv_ability",
			Description: "this is a test rule",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          CountryGroupID,
					CountryList: []string{"US", "CN", "GE", "ID"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       PlatformGroupID,
							Platform: "android",
							Rules: []RuleInfo{
								{
									ID:        RuleID,
									Status:    true,
									Name:      "mkv video",
									Priority:  1,
									Condition: "platform=='android'",
									DecisionInfoList: []*DecisionInfo{
										{
											// 省略了Name
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "hardware",
													Value:    "'mediacodec'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "video",
													Value:    "'mkv'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "subtitle",
													Value:    "'srt'",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// 构建解码Mp4的规则例子
func hardwareDecodeMp4Ability() []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          0,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "hardware_decode_mp4_ability",
			Description: "this is a test rule",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          0,
					CountryList: []string{"US", "CN", "ID"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       0,
							Platform: "IOS",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "mkv video",
									Priority:  1,
									Condition: "platform=='IOS'",
									DecisionInfoList: []*DecisionInfo{
										{
											// 省略了Name
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "hardware",
													Value:    "'mediacodec'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "video",
													Value:    "'mp4'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "encoder",
													Value:    "'h264'",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// 构建解码Mp3的规则例子
func hardwareDecodeMp3Ability() []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          0,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "hardware_decode_mp3_ability",
			Description: "this is a test rule",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          0,
					CountryList: []string{"US", "CN", "ID"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       0,
							Platform: "IOS",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "mp3 audio",
									Priority:  1,
									Condition: "platform=='IOS'",
									DecisionInfoList: []*DecisionInfo{
										{
											// 省略了Name
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "hardware",
													Value:    "'mediacodec'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "audio",
													Value:    "'mp3'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "encoder",
													Value:    "'mp3-lame'",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// 构建解码Mp4的规则例子
func mainFeedPreloadSettings() []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          0,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "main_feed_preload_settings",
			Description: "this is a test rule",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          0,
					CountryList: []string{"US", "CN"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       0,
							Platform: "IOS",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "preload",
									Priority:  1,
									Condition: "platform=='IOS'",
									DecisionInfoList: []*DecisionInfo{
										{
											// 省略了Name
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "buffersize",
													Value:    "1024",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "enable_preload",
													Value:    "'mp4'",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataMiddleWareV3Settings() []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          0,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "tiktok_data_middleware_v3",
			Description: "tiktok playback data log v3",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          0,
					CountryList: []string{"US", "CN"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       0,
							Platform: "IOS",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "preload",
									Priority:  1,
									Condition: "appid==1180",
									DecisionInfoList: []*DecisionInfo{
										{
											// 省略了Name
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "enable_v3",
													Value:    "true",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "delay_load",
													Value:    "'true'",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func tiktokSeekOptSettings() []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          0,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "tiktok_seek_optimization",
			Description: "tiktok seek optimization",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          0,
					CountryList: []string{"US"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       0,
							Platform: "IOS",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "ios_seek",
									Priority:  1,
									Condition: "appid==1180",
									DecisionInfoList: []*DecisionInfo{
										{
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "enable_seek",
													Value:    "true",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "threshold",
													Value:    "400",
												},
											},
										},
									},
								},
							},
						},
						{
							ID:       0,
							Platform: "Android",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "android_seek",
									Priority:  1,
									Condition: "appid==1180",
									DecisionInfoList: []*DecisionInfo{
										{
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "enable_seek",
													Value:    "true",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "threshold",
													Value:    "300",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func feedFullCacheV2Settings() []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          0,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "tiktok_feed_full_cache",
			Description: "full cache strategy of tiktok main feed",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          0,
					CountryList: []string{"US"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       0,
							Platform: "IOS",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "ios_strategy",
									Priority:  1,
									Condition: "appid==1180",
									DecisionInfoList: []*DecisionInfo{
										{
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "enable_cache",
													Value:    "true",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "cache_size",
													Value:    "20480000",
												},
											},
										},
									},
								},
							},
						},
						{
							ID:       0,
							Platform: "Android",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "anroid_strategy",
									Priority:  1,
									Condition: "appid==1180",
									DecisionInfoList: []*DecisionInfo{
										{
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "enable_cache",
													Value:    "true",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "cache_size",
													Value:    "40960000",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func hardwareDecodeDoblyAbility() []*RuleGroupInfo {
	return []*RuleGroupInfo{
		{
			ID:          0,
			TenantID:    3,
			RuleStageID: &RuleStageID,
			GroupName:   "hardware_decode_dobly_ability",
			Description: "test rule group for dobly",
			CountryGroup: []CountryGroupInfo{
				{
					ID:          0,
					CountryList: []string{"US"},
					PlatformGroup: []PlatformGroupInfo{
						{
							ID:       0,
							Platform: "IOS",
							Rules: []RuleInfo{
								{
									ID:        0,
									Status:    true,
									Name:      "dobly decoding",
									Priority:  1,
									Condition: "platform=='IOS'",
									DecisionInfoList: []*DecisionInfo{
										{
											Decisions: []*Decision{
												{
													ID:       0,
													Operator: "SET",
													Key:      "hardware",
													Value:    "'mediacodec'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "video",
													Value:    "'mp4'",
												},
												{
													ID:       0,
													Operator: "SET",
													Key:      "decoder",
													Value:    "'dobly'",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
