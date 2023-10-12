package gorm

type DBVodRuleGroup struct {
}

func (d *DBVodRuleGroup) TableName() string {
	return "vod_t_rule_group"
}
