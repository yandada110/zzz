package arms

import "zzz/CharacterPanel/common"

/*
*
硫磺石
*/
func LiuHuangShi(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 684
	m.MainArticle = 30
	m.Type = common.AttackPowerPercentage
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    25,                              // 被动效果增益值
			Type:     common.AttackInternalPercentage, // 被动效果增益值
			GainForm: common.GainFormInsideFixed,      // 被动效果增益值
		})
	}
	return m
}
