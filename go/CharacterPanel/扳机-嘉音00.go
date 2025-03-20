package main

import "zzz/CharacterPanel/common"

func 安比01扳机00嘉音00() *Initializations {
	init := &Initializations{
		Name:       "安比01扳机00嘉音00",
		NumberFour: common.Critical,
	}
	// 初始化 Initializations（可包含多套计算方式，这里举例两个）
	init.Initializations = []*Initialization{
		安比01扳机00嘉音00站场(), // 模型A
		安比01扳机00嘉音00失衡(), // 模型B
	}
	return init
}
func 安比01扳机01嘉音00() *Initializations {
	init := &Initializations{
		Name:       "安比01扳机01嘉音00",
		NumberFour: common.Critical,
	}
	// 初始化 Initializations（可包含多套计算方式，这里举例两个）
	init.Initializations = []*Initialization{
		安比01扳机01嘉音00站场(), // 模型A
		安比01扳机01嘉音00失衡(), // 模型B
	}
	return init
}
func 安比21扳机00嘉音00() *Initializations {

	init := &Initializations{
		Name:       "安比21扳机00嘉音00",
		NumberFour: common.Critical,
	}
	// 初始化 Initializations（可包含多套计算方式，这里举例两个）
	init.Initializations = []*Initialization{
		安比21扳机00嘉音00站场(), // 模型A
		安比21扳机00嘉音00失衡(), // 模型B
	}
	return init
}

func 安比21扳机01嘉音00() *Initializations {
	init := &Initializations{
		Name:       "安比21扳机01嘉音00",
		NumberFour: common.Critical,
	}
	// 初始化 Initializations（可包含多套计算方式，这里举例两个）
	init.Initializations = []*Initialization{
		安比21扳机01嘉音00站场(), // 模型A
		安比21扳机01嘉音00失衡(), // 模型B
	}
	return init
}

func MagnificationBase1() []*Magnification {
	return []*Magnification{
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

}

func MagnificationBase2() []*Magnification {
	return []*Magnification{
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
}

func (i *Initializations) InitializationBase() {
	i.Basic = &Basic{
		BasicAttack:              713 + 929, // 基础攻击力（角色+专武）
		BasicCritical:            19.4 + 24, // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     50 + 48,   // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     25,        // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: 0,         // 基础减抗（角色+武器+驱动盘）
	}
	i.Gain = &Gain{
		AttackValue:              316,               // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
		AttackValue2:             1200,              // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
		AttackPowerPercentage:    30,                // 攻击力百分比加成(主词条，副词条，2件套,6号位)
		AttackInternalPercentage: 12 + 12,           // 局内攻击力百分比(武器，驱动盘绿字攻击力)
		Critical:                 12 + 10,           // 增加暴击（角色+武器+2件套+4号位）
		ExplosiveInjury:          60 + 25,           // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          25 + 20 + 20 + 24, // 增伤（队友百分比）
		ReductionResistance:      0,                 // 减抗（百分比）
		Vulnerable:               35,                // 易伤（百分比）
		SpecialDamage:            0,                 // 特殊增伤（百分比）
	}
	i.Defense = &Defense{
		Penetration:      0, // 穿透率（百分比）
		DefenseBreak:     0, // 破防百分比（百分比）
		PenetrationValue: 0, // 穿透值（固定值）
	}
}

func 安比01扳机00嘉音00站场() *Initialization {
	init := &Initialization{
		Magnifications: MagnificationBase1(),
		Name:           "安比01扳机00嘉音00站场",

		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}

	return init
}

func 安比01扳机00嘉音00失衡() *Initialization {
	// 此模型可与模型A类似，仅作少量数值调整
	init := &Initialization{
		Name:         "安比01扳机00嘉音00失衡",
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}

	return init
}

/*
*
0+1的扳机，有破防25%
*/
func 安比01扳机01嘉音00站场() *Initialization {
	init := &Initialization{
		Name:         "安比01扳机01嘉音00站场",
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}
	return init
}

func 安比01扳机01嘉音00失衡() *Initialization {
	init := &Initialization{
		Name:         "安比01扳机01嘉音00失衡",
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}
	return init
}

/**
2+1的安比，有12%暴击率
*/

func 安比21扳机00嘉音00站场() *Initialization {
	init := &Initialization{
		Name:         "安比21扳机00嘉音00站场",
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}
	return init
}

func 安比21扳机00嘉音00失衡() *Initialization {
	init := &Initialization{
		Name:         "安比21扳机00嘉音00失衡",
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}
	return init
}

func 安比21扳机01嘉音00站场() *Initialization {
	init := &Initialization{
		Name:         "安比21扳机01嘉音00站场",
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}

	return init
}

func 安比21扳机01嘉音00失衡() *Initialization {
	init := &Initialization{
		Name:         "安比21扳机01嘉音00失衡",
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}
	return init
}
