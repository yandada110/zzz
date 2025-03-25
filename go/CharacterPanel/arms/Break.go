package arms

import "zzz/CharacterPanel/common"

/*
*
锁魂影眸
*/
func SuoHunYingMou(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 713
	m.MainArticle = 24
	m.Type = common.Critical
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    25,                         // 被动效果增益值
			Type:     common.DefenseBreak,        // 被动效果增益值
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
拘缚者
*/
func JuFuZhe(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 684
	m.MainArticle = 0
	m.Type = common.ExplosiveInjury
	//if type1 {
	//	m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
	//		Value:    25,                         // 被动效果增益值
	//		Type:     common.DefenseBreak,        // 被动效果增益值
	//		GainForm: common.GainFormInsideFixed, // 被动效果增益值
	//	})
	//}
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
德玛拉电池
*/
func DeMaLaDianChi(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 624
	m.MainArticle = 0
	m.Type = common.ExplosiveInjury
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    24,                         // 被动效果增益值
			Type:     common.ReductionResistance, // 被动效果增益值
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
裁纸刀
*/
func CaiZhiDao(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 624
	m.MainArticle = 0
	m.Type = common.ExplosiveInjury
	//if type1 {
	//	m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
	//		Value:    24,                         // 被动效果增益值
	//		Type:     common.ReductionResistance, // 被动效果增益值
	//		GainForm: common.GainFormInsideFixed, // 被动效果增益值
	//	})
	//}
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
焰心桂冠
*/
func YanXinGuiGuan(type1 bool) *MainArticle {
	m := &MainArticle{}
	m.BaseAttackValue = 713
	m.MainArticle = 0
	m.Type = common.ExplosiveInjury
	if type1 {
		m.OtherBenefits = append(m.OtherBenefits, &OtherBenefits{
			Value:    30,                         // 被动效果增益值
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
