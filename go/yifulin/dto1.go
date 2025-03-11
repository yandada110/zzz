package main

func NewInitialization2() *Initialization {
	magnifications := []*Magnification{
		&Magnification{
			MagnificationValue: 473.3 + 477.8,
			TriggerTimes:       12,
			Name:               "普攻",
		},
		&Magnification{
			MagnificationValue: 81.1 + 61.1,
			TriggerTimes:       4,
			Name:               "特殊技-束缚式",
		},
		&Magnification{
			MagnificationValue: 453,
			TriggerTimes:       9,
			Name:               "缴勒式1",
		},
		&Magnification{
			MagnificationValue: 490.5,
			TriggerTimes:       4,
			Name:               "缴勒式2",
		},
		&Magnification{
			MagnificationValue: 1082.4 + 120.1,
			TriggerTimes:       2,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue:  1658.7,
			TriggerTimes:        4,
			Name:                "连携技",
			IncreasedDamage:     30,
			ReductionResistance: 25,
			SpecialDamage:       25,
		},
		&Magnification{
			MagnificationValue:  3977.3,
			TriggerTimes:        1,
			Name:                "终结技",
			IncreasedDamage:     30,
			ReductionResistance: 25,
			SpecialDamage:       25,
		},
	}
	return &Initialization{
		Magnifications: magnifications, // 伤害倍率
		Basic: &Basic{
			BasicAttack:              3035, // 基础攻击力（角色+专武）
			BasicCritical:            100,  // 基础暴击（角色+武器+2件套+4号位）
			BasicExplosiveInjury:     160,  // 基础爆伤（角色+武器+2件套+4号位）
			BasicIncreasedDamage:     0,    // 基础增伤（角色+武器+驱动盘）
			BasicReductionResistance: 0,    // 基础减抗（角色+武器+驱动盘）
		},
		Gain: &Gain{
			AttackValue:              1200,         // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
			AttackPowerPercentage:    0,            // 攻击力百分比加成(主词条，副词条，2件套)
			AttackInternalPercentage: 12 + 18,      // 局内攻击力百分比(武器，驱动盘绿字攻击力)
			Critical:                 0,            // 增加暴击（角色+武器+2件套+4号位）
			ExplosiveInjury:          50 + 25,      // 增加爆伤（角色+武器+2件套+4号位）
			IncreasedDamage:          24 + 20 + 40, // 增伤（队友百分比）
			ReductionResistance:      0,            // 减抗（百分比）
			Vulnerable:               0,            // 易伤（百分比）
			SpecialDamage:            0,            // 特殊增伤（百分比）
		},
		Defense: &Defense{
			Penetration:      0,  // 穿透率（百分比）
			DefenseBreak:     12, // 破防百分比（百分比）
			PenetrationValue: 36, // 穿透值（固定值）
		},
		Condition: &Condition{
			MainArticle: 0,   // 有效词条
			Critical:    100, // 最高暴击率
		},
		Output: &Output{},
	}
}

func NewInitialization3() *Initialization {
	magnifications := []*Magnification{
		//&Magnification{
		//	MagnificationValue: 473.3 + 477.8,
		//	TriggerTimes:       12,
		//	Name:               "普攻",
		//},
		//&Magnification{
		//	MagnificationValue: 81.1 + 61.1,
		//	TriggerTimes:       4,
		//	Name:               "特殊技-束缚式",
		//},
		&Magnification{
			MagnificationValue: 453,
			TriggerTimes:       2,
			Name:               "缴勒式1",
		},
		&Magnification{
			MagnificationValue: 490.5,
			TriggerTimes:       1,
			Name:               "缴勒式2",
		},
		&Magnification{
			MagnificationValue: 1082.4 + 120.1,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue:  1658.7,
			TriggerTimes:        5,
			Name:                "连携技",
			IncreasedDamage:     30,
			ReductionResistance: 25,
			SpecialDamage:       25,
		},
		&Magnification{
			MagnificationValue:  3977.3,
			TriggerTimes:        1,
			Name:                "终结技",
			IncreasedDamage:     30,
			ReductionResistance: 25,
			SpecialDamage:       25,
		},
	}
	return &Initialization{
		Magnifications: magnifications, // 伤害倍率
		Basic: &Basic{
			BasicAttack:              3035, // 基础攻击力（角色+专武）
			BasicCritical:            100,  // 基础暴击（角色+武器+2件套+4号位）
			BasicExplosiveInjury:     160,  // 基础爆伤（角色+武器+2件套+4号位）
			BasicIncreasedDamage:     0,    // 基础增伤（角色+武器+驱动盘）
			BasicReductionResistance: 0,    // 基础减抗（角色+武器+驱动盘）
		},
		Gain: &Gain{
			AttackValue:              1200,         // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
			AttackPowerPercentage:    0,            // 攻击力百分比加成(主词条，副词条，2件套)
			AttackInternalPercentage: 12 + 18,      // 局内攻击力百分比(武器，驱动盘绿字攻击力)
			Critical:                 0,            // 增加暴击（角色+武器+2件套+4号位）
			ExplosiveInjury:          50 + 25,      // 增加爆伤（角色+武器+2件套+4号位）
			IncreasedDamage:          24 + 20 + 40, // 增伤（队友百分比）
			ReductionResistance:      0,            // 减抗（百分比）
			Vulnerable:               25,           // 易伤（百分比）
			SpecialDamage:            0,            // 特殊增伤（百分比）
		},
		Defense: &Defense{
			Penetration:      0,  // 穿透率（百分比）
			DefenseBreak:     12, // 破防百分比（百分比）
			PenetrationValue: 36, // 穿透值（固定值）
		},
		Condition: &Condition{
			MainArticle: 0,   // 有效词条
			Critical:    100, // 最高暴击率
		},
		Output: &Output{},
	}
}
