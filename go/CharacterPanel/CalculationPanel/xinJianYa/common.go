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

func MagnificationBase月城柳() []*Magnification {
	return []*Magnification{
		{
			MagnificationValue: 126.6 + 193.3 + 258.8,
			TriggerTimes:       3,
			Critical:           12,
			Name:               "普攻",
			IncreasedDamage:    40,
		},
		{
			MagnificationValue: 788.3,
			TriggerTimes:       1,
			Critical:           12,
			Name:               "强化特殊技",
			IncreasedDamage:    40,
		},
		{
			MagnificationValue: 4282.8,
			TriggerTimes:       2,
			Critical:           12,
			Name:               "拔刀",
			IncreasedDamage:    40 + 60,
		},
		&Magnification{
			MagnificationValue: common.AbnormalMagnification[common.Ice],
			TriggerTimes:       1,
			Name:               "碎冰",
			DamageType:         common.Abnormal,
			DisorderType:       common.Ice,
		},
		&Magnification{
			MagnificationValue: common.DisorderMagnification[common.Ice],
			TriggerTimes:       1,
			Name:               "冰结算",
			DamageType:         common.Disorder,
			DisorderType:       common.Ice,
			TimeConsumption:    -5,
		},
		{
			MagnificationValue: 4776.1,
			TriggerTimes:       1,
			Critical:           12,
			Name:               "终结技",
			IncreasedDamage:    40 + 60 + 30,
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
		Critical:                 12,  // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          46,  // 增加爆伤（角色+武器+2件套+4号位）
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
