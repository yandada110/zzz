package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

// ------------------------ 常量定义 ------------------------
const (
	// 每个队伍可分配的词条数（示例值）
	GlobalMainArticle = 58
	// 词条类型数量
	GlobalMainArticleTypeCount = 6
)

func MagnificationBase1() []*Magnification {
	return []*Magnification{
		{
			MagnificationValue: 160.5,
			TriggerTimes:       6,
			Name:               "群居浮游",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 440,
			TriggerTimes:       12,
			Name:               "落雨生花",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 1185.5,
			TriggerTimes:       4,
			Name:               "强化特殊技",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 747.9,
			TriggerTimes:       4,
			Name:               "格挡反击",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 3374.2 + 1317.8,
			TriggerTimes:       1,
			Name:               "终结技+连携技",
			DamageType:         common.DirectInjury,
		},
		{
			MagnificationValue: 55,
			TriggerTimes:       48 / 0.55,
			Name:               "薇薇安的语言",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		// 这里需要处理一下，如果是简，那么这个就跟强击收益有关
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       12,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             720067,
		},
		{
			MagnificationValue: common.DisorderMagnification[common.Ether],
			TriggerTimes:       2,
			Name:               "以太结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Ether,
			TimeConsumption:    2,
			IncreasedDamage:    12 + 25,
		},
	}
}

func MagnificationBase双生() []*Magnification {
	return []*Magnification{
		{
			MagnificationValue: 160.5,
			TriggerTimes:       6,
			Name:               "群居浮游",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 440,
			TriggerTimes:       12,
			Name:               "落雨生花",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 1185.5,
			TriggerTimes:       4,
			Name:               "强化特殊技",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 747.9,
			TriggerTimes:       4,
			Name:               "格挡反击",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 3374.2 + 1317.8,
			TriggerTimes:       1,
			Name:               "终结技+连携技",
			DamageType:         common.DirectInjury,
		},
		{
			MagnificationValue: 55,
			TriggerTimes:       48 / 0.55,
			Name:               "薇薇安的语言",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		// 这里需要处理一下，如果是简，那么这个就跟强击收益有关
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       1,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             720067,
			Proficient:         48,
		},
		// 这里需要处理一下，如果是简，那么这个就跟强击收益有关
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       1,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             650000,
			Proficient:         48 * 2,
		},
		{
			MagnificationValue: common.DisorderMagnification[common.Ether],
			TriggerTimes:       1,
			Name:               "以太结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Ether,
			TimeConsumption:    2,
			IncreasedDamage:    12 + 25,
			Proficient:         48 * 2,
		},
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       1,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             720067,
			Proficient:         48 * 3,
		},
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       3,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             720067,
			Proficient:         48 * 4,
		},
		{
			MagnificationValue: common.DisorderMagnification[common.Ether],
			TriggerTimes:       1,
			Name:               "以太结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Ether,
			TimeConsumption:    2,
			IncreasedDamage:    12 + 25,
			Proficient:         48 * 4,
		},
	}
}

func MagnificationBase双生1命() []*Magnification {
	return []*Magnification{
		{
			MagnificationValue: 160.5,
			TriggerTimes:       6,
			Name:               "群居浮游",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 440,
			TriggerTimes:       12,
			Name:               "落雨生花",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 1185.5,
			TriggerTimes:       4,
			Name:               "强化特殊技",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 747.9,
			TriggerTimes:       4,
			Name:               "格挡反击",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 3374.2 + 1317.8,
			TriggerTimes:       1,
			Name:               "终结技+连携技",
			DamageType:         common.DirectInjury,
		},
		{
			MagnificationValue: 55,
			TriggerTimes:       48 / 0.55,
			Name:               "薇薇安的语言",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		// 这里需要处理一下，如果是简，那么这个就跟强击收益有关
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       1,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             779150,
			Proficient:         48,
		},
		// 这里需要处理一下，如果是简，那么这个就跟强击收益有关
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       1,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             779150,
			Proficient:         48 * 2,
		},
		{
			MagnificationValue: common.DisorderMagnification[common.Ether],
			TriggerTimes:       1,
			Name:               "以太结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Ether,
			TimeConsumption:    2,
			IncreasedDamage:    12 + 25 + 16,
			Proficient:         48 * 2,
		},
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       1,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             779150,
			Proficient:         48 * 3,
		},
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       3,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             779150,
			Proficient:         48 * 4,
		},
		{
			MagnificationValue: common.DisorderMagnification[common.Ether],
			TriggerTimes:       1,
			Name:               "以太结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Ether,
			TimeConsumption:    2,
			IncreasedDamage:    12 + 25 + 16,
			Proficient:         48 * 4,
		},
	}
}

func MagnificationBase1薇薇安11() []*Magnification {
	return []*Magnification{
		{
			MagnificationValue: 160.5,
			TriggerTimes:       6,
			Name:               "群居浮游",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 440,
			TriggerTimes:       12,
			Name:               "落雨生花",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 1185.5,
			TriggerTimes:       4,
			Name:               "强化特殊技",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 747.9,
			TriggerTimes:       4,
			Name:               "格挡反击",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 3374.2 + 1317.8,
			TriggerTimes:       1,
			Name:               "终结技+连携技",
			DamageType:         common.DirectInjury,
		},
		{
			MagnificationValue: 55,
			TriggerTimes:       48 / 0.55,
			Name:               "薇薇安的语言",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		// 这里需要处理一下，如果是简，那么这个就跟强击收益有关
		{
			MagnificationValue: common.DifferentMagnification[common.Physical],
			TriggerTimes:       12,
			Name:               "异放",
			DamageType:         common.Different,
			DisorderType:       common.Physical,
			Damage:             720067,
		},
		{
			MagnificationValue: common.DisorderMagnification[common.Ether],
			TriggerTimes:       2,
			Name:               "以太结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Ether,
			TimeConsumption:    2,
		},
	}
}

func MagnificationBase1薇薇安21() []*Magnification {
	return []*Magnification{
		{
			MagnificationValue: 160.5,
			TriggerTimes:       6,
			Name:               "群居浮游",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 440,
			TriggerTimes:       12,
			Name:               "落雨生花",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 1185.5,
			TriggerTimes:       4,
			Name:               "强化特殊技",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 747.9,
			TriggerTimes:       4,
			Name:               "格挡反击",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		{
			MagnificationValue: 3374.2 + 1317.8,
			TriggerTimes:       1,
			Name:               "终结技+连携技",
			DamageType:         common.DirectInjury,
		},
		{
			MagnificationValue: 55,
			TriggerTimes:       48 / 0.55,
			Name:               "薇薇安的语言",
			DamageType:         common.DirectInjury,
			IncreasedDamage:    12 + 25,
		},
		// 这里需要处理一下，如果是简，那么这个就跟强击收益有关
		{
			MagnificationValue:  common.DifferentMagnification[common.Physical] * 1.4,
			TriggerTimes:        12,
			Name:                "异放",
			DamageType:          common.Different,
			DisorderType:        common.Physical,
			Damage:              720067,
			ReductionResistance: 20,
		},
		{
			MagnificationValue: common.DisorderMagnification[common.Ether],
			TriggerTimes:       2,
			Name:               "以太结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Ether,
			TimeConsumption:    2,
			IncreasedDamage:    12 + 25,
		},
	}
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
		i.Defense.Penetration += buffCharacter.Penetration
		i.Defense.DefenseBreak += buffCharacter.DefenseBreak
	}
}

// 0命
func (i *Initializations) InitializationBase0命(role *Role.BaseRole, article *arms.MainArticle) {
	i.Basic = &Basic{
		BasicAttack:              role.AttackValue + article.BaseAttackValue, // 基础攻击力（角色+专武）
		BasicCritical:            role.Critical,                              // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     role.ExplosiveInjury,                       // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     role.IncreasedDamage,                       // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: role.ReductionResistance,                   // 基础减抗（角色+武器+驱动盘）
		BasicProficient:          role.Proficient,                            // 精通
	}
	if i.NumberFour == common.Critical {
		i.Basic.BasicCritical += 24
	}
	if i.NumberFour == common.ExplosiveInjury {
		i.Basic.BasicExplosiveInjury += 48
	}
	if i.NumberFour == common.Proficient {
		i.Basic.BasicProficient += 90
	}

	if article.Type == common.Critical {
		i.Basic.BasicCritical += article.MainArticle
	}
	if article.Type == common.ExplosiveInjury {
		i.Basic.BasicExplosiveInjury += article.MainArticle
	}
	i.Gain = &Gain{
		AttackValue:              316, // 攻击力值增加(固定2号位数值)
		AttackValue2:             0,   // 攻击力值增加(局内加的固定攻击力)
		AttackPowerPercentage:    0,   // 局外攻击力百分比(6号位+武器主词条+5号位+4号位+副词条)
		AttackInternalPercentage: 0,   // 局内攻击力百分比(武器，4件套)
		Critical:                 0,   // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          0,   // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          0,   // 增伤（队友百分比）
		ReductionResistance:      0,   // 减抗（百分比）
		Vulnerable:               0,   // 易伤（百分比）
		SpecialDamage:            0,   // 特殊增伤（百分比）
	}
	if i.NumberSix == common.AttackPowerPercentage {
		i.Gain.AttackPowerPercentage = 30
	}
	if article.Type == common.AttackPowerPercentage {
		i.Gain.AttackPowerPercentage += article.MainArticle
	}
	for _, OtherBenefit := range article.OtherBenefits {
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
	i.Defense = &Defense{
		Penetration:      role.Penetration,  // 穿透率（百分比）
		DefenseBreak:     role.DefenseBreak, // 破防百分比（百分比）
		PenetrationValue: 0,                 // 穿透值（固定值）
	}
}
