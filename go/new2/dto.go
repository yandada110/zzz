package main

// NewInitialization 构造一个初始化对象，并设置相关参数
func NewInitialization() *Initialization {
	// 在最外层设置 MainArticle（有效词条数）为 45
	init := &Initialization{
		MainArticle: 45,
		Condition: &Condition{
			Critical: 95,
		},
	}
	// 将计算模型放入集合中（这里添加了两个模型）
	init.CalculationModels = append(init.CalculationModels, 安比01扳机00嘉音00站场())
	init.CalculationModels = append(init.CalculationModels, 安比01扳机00嘉音00失衡())
	return init
}

// 安比01扳机00嘉音00站场 模型
func 安比01扳机00嘉音00站场() *Initialization {
	magnifications := []*Magnification{
		{
			MagnificationValue: 65.7 + 127.9 + 160.4 + 80.6 + 309.6*5 + 297,
			TriggerTimes:       11,
			Name:               "普攻",
		},
		{
			MagnificationValue: 334.4,
			TriggerTimes:       6 * 3,
			Name:               "白雷",
			IncreasedDamage:    25 + 15,
			ExplosiveInjury:    30,
		},
		{
			MagnificationValue: 376.2,
			TriggerTimes:       6,
			Name:               "落雷",
			IncreasedDamage:    25 + 15,
			ExplosiveInjury:    30,
		},
		{
			MagnificationValue: 87,
			TriggerTimes:       6 * 3,
			Name:               "白雷直伤",
		},
		{
			MagnificationValue: 760.2 + 120.1,
			TriggerTimes:       2,
			Name:               "强化特殊技+特殊技",
		},
	}
	return &Initialization{
		MainArticle:    45,
		Magnifications: magnifications,
		Basic: &Basic{
			BasicAttack:              713 + 929,
			BasicCritical:            19.4 + 24,
			BasicExplosiveInjury:     50 + 48,
			BasicIncreasedDamage:     25,
			BasicReductionResistance: 0,
		},
		Gain: &Gain{
			AttackValue:              316,
			AttackValue2:             1200,
			AttackPowerPercentage:    30,
			AttackInternalPercentage: 12 + 12,
			Critical:                 12 + 10,
			ExplosiveInjury:          60 + 25,
			IncreasedDamage:          25 + 20 + 20 + 24,
			ReductionResistance:      0,
			Vulnerable:               35,
			SpecialDamage:            0,
		},
		Defense: &Defense{
			Penetration:      0,
			DefenseBreak:     0,
			PenetrationValue: 0,
		},
		Condition: &Condition{
			Critical: 95,
		},
		Output:        &Output{},
		CritConverted: 0,
	}
}

// 安比01扳机00嘉音00失衡 模型
func 安比01扳机00嘉音00失衡() *Initialization {
	magnifications := []*Magnification{
		{
			MagnificationValue: 334.4,
			TriggerTimes:       4 * 3,
			Name:               "白雷",
			IncreasedDamage:    25 + 15,
			ExplosiveInjury:    30,
		},
		{
			MagnificationValue: 376.2,
			TriggerTimes:       4,
			Name:               "落雷",
			IncreasedDamage:    25 + 15,
			ExplosiveInjury:    30,
		},
		{
			MagnificationValue: 87,
			TriggerTimes:       4 * 3,
			Name:               "白雷直伤",
		},
		{
			MagnificationValue: 760.2 + 120.1,
			TriggerTimes:       2,
			Name:               "强化特殊技+特殊技",
		},
		{
			MagnificationValue: 1128.1 + 3470.7,
			TriggerTimes:       1,
			Name:               "终结技+连携技",
		},
	}
	return &Initialization{
		MainArticle:    45,
		Magnifications: magnifications,
		Basic: &Basic{
			BasicAttack:              713 + 929,
			BasicCritical:            19.4 + 24,
			BasicExplosiveInjury:     50 + 48,
			BasicIncreasedDamage:     25,
			BasicReductionResistance: 0,
		},
		Gain: &Gain{
			AttackValue:              316,
			AttackValue2:             1200,
			AttackPowerPercentage:    30,
			AttackInternalPercentage: 12 + 12,
			Critical:                 12 + 10,
			ExplosiveInjury:          60 + 25,
			IncreasedDamage:          25 + 20 + 20 + 24,
			ReductionResistance:      0,
			Vulnerable:               35 + 25,
			SpecialDamage:            0,
		},
		Defense: &Defense{
			Penetration:      0,
			DefenseBreak:     0,
			PenetrationValue: 0,
		},
		Condition: &Condition{
			Critical: 95,
		},
		Output:        &Output{},
		CritConverted: 0,
	}
}
