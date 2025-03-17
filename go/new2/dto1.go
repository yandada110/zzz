package main

// NewInitialization1 构造队伍组合二（例如：安比01扳机00嘉音00失衡）
func 扳机妮可() *Initialization {
	init := &Initialization{
		MainArticle: 45,
		Name:        "计算扳机妮可",
		Condition: &Condition{
			Critical: 95,
		},
	}
	init.CalculationModels = []*Initialization{
		teamModelC(), // 模型C
		teamModelD(), // 模型D
	}
	return init
}

func teamModelC() *Initialization {
	mags := []*Magnification{
		&Magnification{
			MagnificationValue: 65.7 + 127.9 + 160.4 + 80.6 + 309.6*5 + 297,
			TriggerTimes:       11,
			Name:               "普攻",
		},
		&Magnification{
			MagnificationValue: 334.4,
			TriggerTimes:       6 * 3,
			Name:               "白雷",
			IncreasedDamage:    25 + 15,
			ExplosiveInjury:    30,
		},
		&Magnification{
			MagnificationValue: 376.2,
			TriggerTimes:       6,
			Name:               "落雷",
			IncreasedDamage:    25 + 15,
			ExplosiveInjury:    30,
		},
		&Magnification{
			MagnificationValue: 87,
			TriggerTimes:       6 * 3,
			Name:               "白雷直伤",
		},
		&Magnification{
			MagnificationValue: 760.2 + 120.1,
			TriggerTimes:       2,
			Name:               "强化特殊技+特殊技",
		},
	}
	return &Initialization{
		MainArticle:    45,
		Magnifications: mags,
		Name:           "安比01扳机00妮可65站场",
		Basic: &Basic{
			BasicAttack:              713 + 929, // 基础攻击力（角色+专武）
			BasicCritical:            19.4 + 24, // 基础暴击（角色+武器+2件套+4号位）
			BasicExplosiveInjury:     50 + 48,   // 基础爆伤（角色+武器+2件套+4号位）
			BasicIncreasedDamage:     25,        // 基础增伤（角色+武器+驱动盘）
			BasicReductionResistance: 0,         // 基础减抗（角色+武器+驱动盘）
		},
		Gain: &Gain{
			AttackValue:              316,          // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
			AttackValue2:             0,            // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
			AttackPowerPercentage:    30,           // 攻击力百分比加成(主词条，副词条，2件套,6号位)
			AttackInternalPercentage: 12,           // 局内攻击力百分比(武器，驱动盘绿字攻击力)
			Critical:                 12 + 10,      // 增加暴击（角色+武器+2件套+4号位）
			ExplosiveInjury:          60,           // 增加爆伤（角色+武器+2件套+4号位）
			IncreasedDamage:          25 + 20 + 20, // 增伤（队友百分比）
			ReductionResistance:      0,            // 减抗（百分比）
			Vulnerable:               35,           // 易伤（百分比）
			SpecialDamage:            0,            // 特殊增伤（百分比）
		},
		Defense: &Defense{
			Penetration:      0,  // 穿透率（百分比）
			DefenseBreak:     40, // 破防百分比（百分比）
			PenetrationValue: 0,  // 穿透值（固定值）
		},
		Condition: &Condition{
			Critical: 95, // 最高暴击率
		},
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}
}

func teamModelD() *Initialization {
	// 此模型可与模型A类似，仅作少量数值调整
	mags := []*Magnification{
		&Magnification{
			MagnificationValue: 334.4,
			TriggerTimes:       4 * 3,
			Name:               "白雷",
			IncreasedDamage:    25 + 15,
			ExplosiveInjury:    30,
		},
		&Magnification{
			MagnificationValue: 376.2,
			TriggerTimes:       4,
			Name:               "落雷",
			IncreasedDamage:    25 + 15,
			ExplosiveInjury:    30,
		},
		&Magnification{
			MagnificationValue: 87,
			TriggerTimes:       4 * 3,
			Name:               "白雷直伤",
		},
		&Magnification{
			MagnificationValue: 760.2 + 120.1,
			TriggerTimes:       2,
			Name:               "强化特殊技+特殊技",
		},
		&Magnification{
			MagnificationValue: 1128.1 + 3470.7,
			TriggerTimes:       1,
			Name:               "终结技+连携技",
		},
	}
	return &Initialization{
		MainArticle:    45,
		Magnifications: mags,
		Name:           "安比01扳机00妮可65失衡",
		Basic: &Basic{
			BasicAttack:              713 + 929, // 基础攻击力（角色+专武）
			BasicCritical:            19.4 + 24, // 基础暴击（角色+武器+2件套+4号位）
			BasicExplosiveInjury:     50 + 48,   // 基础爆伤（角色+武器+2件套+4号位）
			BasicIncreasedDamage:     25,        // 基础增伤（角色+武器+驱动盘）
			BasicReductionResistance: 0,         // 基础减抗（角色+武器+驱动盘）
		},
		Gain: &Gain{
			AttackValue:              316,          // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
			AttackValue2:             0,            // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
			AttackPowerPercentage:    30,           // 攻击力百分比加成(主词条，副词条，2件套,6号位)
			AttackInternalPercentage: 12,           // 局内攻击力百分比(武器，驱动盘绿字攻击力)
			Critical:                 12 + 10,      // 增加暴击（角色+武器+2件套+4号位）
			ExplosiveInjury:          60,           // 增加爆伤（角色+武器+2件套+4号位）
			IncreasedDamage:          25 + 20 + 20, // 增伤（队友百分比）
			ReductionResistance:      0,            // 减抗（百分比）
			Vulnerable:               35 + 25,      // 易伤（百分比）
			SpecialDamage:            0,            // 特殊增伤（百分比）
		},
		Defense: &Defense{
			Penetration:      0,  // 穿透率（百分比）
			DefenseBreak:     40, // 破防百分比（百分比）
			PenetrationValue: 0,  // 穿透值（固定值）
		},
		Condition: &Condition{
			Critical: 95, // 最高暴击率
		},
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}
}
