package arms

import "zzz/CharacterPanel/common"

/*
*
淬锋钳刺
*/
func CuiFengQianCi(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 713
	m.MainArticle = 90
	m.Type = common.Proficient
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    36,                         // 被动效果增益值
			Type:     common.IncreasedDamage,     // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	return m
}
