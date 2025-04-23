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

/*
*
飞鸟星梦
*/
func FeiNiaoXingMeng(type1 bool, number float64) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 713
	m.MainArticle = 90
	m.Type = common.Proficient
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    20 * number,                // 被动效果增益值
			Type:     common.Proficient,          // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	return m
}

/*
*
双生泣星
*/
func ShuangShengQiXing(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 594
	m.MainArticle = 25
	m.Type = common.AttackPowerPercentage
	return m
}

/*
*
灼心摇壶
*/
func ZhuoXinYaoHu(type1 bool, type2 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 713
	m.MainArticle = 30
	m.Type = common.AttackPowerPercentage
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    35,                         // 被动效果增益值
			Type:     common.IncreasedDamage,     // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	if type2 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    50,                         // 被动效果增益值
			Type:     common.Proficient,          // 被动效果增益值
			GainForm: common.GainFormInsideFixed, // 被动效果增益值
		})
	}
	return m
}
