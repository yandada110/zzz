package arms

import "zzz/CharacterPanel/common"

/*
*
暴击武器系统
*/
func XinXianEeXiang(type1 bool) *MainArticle {
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
	//if type2 {
	//	m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
	//		Value:    25,                         // 被动效果增益值
	//		Type:     common.ReductionResistance, // 被动效果增益值
	//		GainForm: common.GainFormInsideFixed, // 被动效果增益值
	//	})
	//}
	return m
}

/*
*
强音
*/
func QiangYinReWang(type1, type2 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 594
	m.MainArticle = 20
	m.Type = common.Critical
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    9.6,                             // 被动效果增益值
			Type:     common.AttackInternalPercentage, // 被动效果增益值
			GainForm: common.GainFormInsideFixed,      // 被动效果增益值
		})
	}
	if type2 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    9.6,                             // 被动效果增益值
			Type:     common.AttackInternalPercentage, // 被动效果增益值
			GainForm: common.GainFormInsideFixed,      // 被动效果增益值
		})
	}
	return m
}

/*
*
千面日陨
*/
func QianMianRiYun(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 713
	m.MainArticle = 24
	m.Type = common.Critical
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    20,                              // 被动效果增益值
			Type:     common.AttackInternalPercentage, // 被动效果增益值
			GainForm: common.GainFormInsideTheBureau,  // 被动效果增益值
		})
	}
	//if type2 {
	//	m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
	//		Value:    25,                         // 被动效果增益值
	//		Type:     common.ReductionResistance, // 被动效果增益值
	//		GainForm: common.GainFormInsideFixed, // 被动效果增益值
	//	})
	//}
	return m
}

/*
*
加农转子
*/
func JiaNongZhuanZi(type1 bool, number float64) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 594
	m.MainArticle = 20
	m.Type = common.Critical
	value := (number - 1) * 1.1
	if number == 5 {
		value = 4.5
	}
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    7.5 + value,                     // 被动效果增益值
			Type:     common.AttackInternalPercentage, // 被动效果增益值
			GainForm: common.GainFormInsideFixed,      // 被动效果增益值
		})
	}
	return m
}
