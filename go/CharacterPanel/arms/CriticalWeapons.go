package arms

import "zzz/CharacterPanel/common"

/*
*
暴击武器系统
*/
func XinXianEeXiang(type1, type2 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 713
	m.MainArticle = 24
	m.Type = common.Critical
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    50,                         // 被动效果增益值
			Type:     common.ExplosiveInjury,     // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	if type2 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    25,                         // 被动效果增益值
			Type:     common.ReductionResistance, // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	return m
}
