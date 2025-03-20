package arms

import "zzz/CharacterPanel/common"

/*
*
牺牲纯洁
*/
func XiShengChunJie(type1, type2, type3 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 713
	m.MainArticle = 48
	m.Type = common.ExplosiveInjury
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    50,                         // 被动效果增益值
			Type:     common.ExplosiveInjury,     // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	if type2 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    10,                         // 被动效果增益值
			Type:     common.ExplosiveInjury,     // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	if type3 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    20,                         // 被动效果增益值
			Type:     common.IncreasedDamage,     // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	return m
}
