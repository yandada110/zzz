package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

// ------------------------ 常量定义 ------------------------
const (
	// 每个队伍可分配的词条数（示例值）
	GlobalMainArticle = 50
	// 词条类型数量
	GlobalMainArticleTypeCount = 6
	// 暴击值上限
	CriticalCount = 100
)

func MagnificationBase1() []*Magnification {
	return []*Magnification{
		&Magnification{
			MagnificationValue: 72.4 + 125 + 167.1 + 327.3 + 197.8 + 582.8,
			TriggerTimes:       6,
			Name:               "普攻",
			DamageType:         common.DirectInjury,
		},
		&Magnification{
			MagnificationValue: 602.2 + 323,
			TriggerTimes:       3,
			Name:               "萨霍夫跳",
			DamageType:         common.DirectInjury,
		},
		&Magnification{
			MagnificationValue: 683.3,
			TriggerTimes:       5,
			Name:               "闪避反击",
			DamageType:         common.DirectInjury,
		},
		&Magnification{
			MagnificationValue: 1150,
			TriggerTimes:       2,
			Name:               "强化特殊技",
			DamageType:         common.DirectInjury,
		},
		&Magnification{
			MagnificationValue: common.AbnormalMagnification[common.Physical],
			TriggerTimes:       3.5,
			Name:               "强击",
			DamageType:         common.Abnormal,
			DisorderType:       common.Physical,
			SpecialDamage:      50,
		},
		&Magnification{
			MagnificationValue: common.DisorderMagnification[common.Physical],
			TriggerTimes:       1,
			Name:               "畏缩结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Physical,
			TimeConsumption:    0,
		},
	}
}

func MagnificationBase2() []*Magnification {
	return []*Magnification{
		&Magnification{
			MagnificationValue: 602.2 + 323,
			TriggerTimes:       3,
			Name:               "萨霍夫跳",
			DamageType:         common.DirectInjury,
		},
		&Magnification{
			MagnificationValue: 1150,
			TriggerTimes:       1,
			Name:               "强化特殊技",
			DamageType:         common.DirectInjury,
		},
		&Magnification{
			MagnificationValue: 1266.2 + 1266.2 + 2941.3,
			TriggerTimes:       1,
			Name:               "双连携+终结技",
			DamageType:         common.DirectInjury,
		},
		&Magnification{
			MagnificationValue: common.AbnormalMagnification[common.Physical],
			TriggerTimes:       2.5,
			Name:               "强击",
			DamageType:         common.Abnormal,
			SpecialDamage:      50,
		},
	}
}

func (i *Initializations) InitializationBase1(role *Role.BaseRole, article *arms.MainArticle) {
	i.Basic = &Basic{
		BasicAttack:              role.AttackValue + article.BaseAttackValue, // 基础攻击力（角色+专武）
		BasicCritical:            role.Critical,                              // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     role.ExplosiveInjury,                       // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     role.IncreasedDamage,                       // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: role.ReductionResistance,                   // 基础减抗（角色+武器+驱动盘）
		BasicProficient:          role.Proficient,                            // 基础精通（角色）
	}
	i.Gain = &Gain{
		AttackValue:              316, // 攻击力值增加(固定2号位数值)
		AttackValue2:             600, // 攻击力值增加(局内加的固定攻击力)
		AttackPowerPercentage:    0,   // 局外攻击力百分比(6号位+武器主词条+5号位+4号位+副词条)
		AttackInternalPercentage: 0,   // 局内攻击力百分比(武器，4件套)
		Critical:                 0,   // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          0,   // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          0,   // 增伤（角色自身）
		ReductionResistance:      0,   // 减抗（百分比）
		Vulnerable:               0,   // 易伤（百分比）
		SpecialDamage:            0,   // 特殊增伤（百分比）
	}
	i.Defense = &Defense{
		Penetration:      role.Penetration,  // 穿透率（百分比）
		DefenseBreak:     role.DefenseBreak, // 破防百分比（百分比）
		PenetrationValue: 0,                 // 穿透值（固定值）
	}
	i.HandleNumberFour()
	i.HandleArticleType(article)
}

func (i *Initializations) InitializationRole(buffCharacters []*Role.BuffCharacter) {
	for _, buffCharacter := range buffCharacters {
		i.Gain.AttackValue2 += buffCharacter.AttackValue
		i.Gain.AttackInternalPercentage += buffCharacter.AttackInternalPercentage
		i.Gain.Critical += buffCharacter.Critical
		i.Gain.ExplosiveInjury += buffCharacter.ExplosiveInjury
		i.Gain.IncreasedDamage += buffCharacter.IncreasedDamage
		i.Gain.ReductionResistance += buffCharacter.ReductionResistance
		i.Gain.Vulnerable += buffCharacter.Vulnerable
		i.Gain.SpecialDamage += buffCharacter.SpecialDamage
		i.Gain.Proficient += buffCharacter.Proficient
		i.Defense.Penetration += buffCharacter.Penetration
		i.Defense.DefenseBreak += buffCharacter.DefenseBreak
	}
}

func (i *Initializations) HandleNumberFour() {
	if i.NumberFour == common.Critical {
		i.Basic.BasicCritical += 24
	}
	if i.NumberFour == common.ExplosiveInjury {
		i.Basic.BasicExplosiveInjury += 48
	}
	if i.NumberFour == common.Proficient {
		i.Basic.BasicProficient += 92
	}
}

func (i *Initializations) HandleArticleType(article *arms.MainArticle) {
	if article.Type == common.Critical {
		i.Basic.BasicCritical += article.MainArticle
	}
	if article.Type == common.ExplosiveInjury {
		i.Basic.BasicExplosiveInjury += article.MainArticle
	}
	if article.Type == common.Proficient {
		i.Basic.BasicProficient += article.MainArticle
	}
	if article.Type == common.AttackPowerPercentage {
		i.Gain.AttackPowerPercentage += article.MainArticle
	}
	i.HandleOtherBenefit(article.OtherBenefits)
}

func (i *Initializations) HandleOtherBenefit(OtherBenefits []*arms.OtherBenefits) {
	for _, OtherBenefit := range OtherBenefits {
		if OtherBenefit.Type == common.Critical {
			i.Gain.Critical += OtherBenefit.Value
		}
		if OtherBenefit.Type == common.ExplosiveInjury {
			i.Gain.ExplosiveInjury += OtherBenefit.Value
		}
		if OtherBenefit.Type == common.IncreasedDamage {
			i.Gain.IncreasedDamage += OtherBenefit.Value
		}
		if OtherBenefit.Type == common.AttackPowerPercentage {
			i.Gain.AttackPowerPercentage += OtherBenefit.Value
		}
		if OtherBenefit.Type == common.AttackInternalPercentage {
			i.Gain.AttackInternalPercentage += OtherBenefit.Value
		}
		if OtherBenefit.Type == common.ReductionResistance {
			i.Gain.ReductionResistance += OtherBenefit.Value
		}
	}
}
