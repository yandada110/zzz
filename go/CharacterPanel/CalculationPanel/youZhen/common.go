package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

// ------------------------ 常量定义 ------------------------
const (
	// 每个队伍可分配的词条数（示例值）
	GlobalMainArticle = 55
	// 词条类型数量
	GlobalMainArticleTypeCount = 5
)

var AttackPercentageEntriesLimit = map[string]int{
	common.Critical:        15,
	common.ExplosiveInjury: 15,
}

var ExplosiveInjuryEntriesLimit = map[string]int{
	common.Critical:        30,
	common.ExplosiveInjury: 25,
}

func MagnificationBase如影专武6命() []*Magnification {
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue:       325.1,
			TriggerTimes:             1,
			AttackInternalPercentage: 0,
			ExplosiveInjury:          0,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 15 + 40,
		},
		&Magnification{
			MagnificationValue:       333.8,
			TriggerTimes:             1,
			AttackInternalPercentage: 4,
			ExplosiveInjury:          12,
			Critical:                 4,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 15 + 40,
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 2,
			AttackInternalPercentage: 4 * 2,
			Critical:                 4 * 2,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 15 + 40,
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             4,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 15 + 40,
		},
		&Magnification{
			MagnificationValue:       899.2,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			Name:                     "强化特殊技",
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9*5,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 15 + 40,
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9*5,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 15 + 40,
		},
		&Magnification{
			MagnificationValue:       3908.6,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			Name:                     "终结技",
		},
		&Magnification{
			MagnificationValue:       1500,
			TriggerTimes:             2,
			AttackInternalPercentage: 4 * 3,
			Name:                     "6命伤害",
		},
	}
}
func MagnificationBase如影专武() []*Magnification {
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue:       325.1,
			TriggerTimes:             1,
			AttackInternalPercentage: 0,
			ExplosiveInjury:          0,
			Name:                     "飞弦斩",
			IncreasedDamage:          15 + 40,
		},
		&Magnification{
			MagnificationValue:       333.8,
			TriggerTimes:             1,
			AttackInternalPercentage: 4,
			ExplosiveInjury:          12,
			Critical:                 4,
			Name:                     "飞弦斩",
			IncreasedDamage:          15 + 40,
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 2,
			AttackInternalPercentage: 4 * 2,
			Critical:                 4 * 2,
			Name:                     "飞弦斩",
			IncreasedDamage:          15 + 40,
		},
		&Magnification{
			MagnificationValue:       899.2,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 2,
			Critical:                 4 * 2,
			Name:                     "强化特殊技",
		},
		&Magnification{
			MagnificationValue:       325.1,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          15 + 40,
		},
		&Magnification{
			MagnificationValue:       333.8,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          15 + 40,
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          15 + 40,
		},
		&Magnification{
			MagnificationValue:       899.2,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			Name:                     "强化特殊技",
		},
		&Magnification{
			MagnificationValue:       3908.6,
			TriggerTimes:             1,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			Name:                     "终结技",
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9,
			TriggerTimes:             2,
			AttackInternalPercentage: 4 * 3,
			Critical:                 4 * 3,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          15 + 40,
		},
	}
}
func MagnificationBase如雷专武6命() []*Magnification {
	var aa float64 = 0
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    0,
			Name:               "飞弦斩",
			IncreasedDamage:    50 + 40,
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12,
			Name:               "飞弦斩",
			IncreasedDamage:    50 + 40,
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 2,
			Name:               "飞弦斩",
			IncreasedDamage:    50 + 40,
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       4,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
			IncreasedDamage:    50 + 40,
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9*5,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: aa,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 40,
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9*5,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: aa,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 40,
		},
		&Magnification{
			MagnificationValue:       3908.6,
			TriggerTimes:             1,
			AttackInternalPercentage: aa,
			Name:                     "终结技",
		},
		&Magnification{
			MagnificationValue:       1500,
			AttackInternalPercentage: aa,
			TriggerTimes:             2,
			Name:                     "6命伤害",
		},
	}
}
func MagnificationBase如雷专武() []*Magnification {
	var aa float64 = 0
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    0,
			Name:               "飞弦斩",
			IncreasedDamage:    40,
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12,
			Name:               "飞弦斩",
			IncreasedDamage:    40,
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 2,
			Name:               "飞弦斩",
			IncreasedDamage:    40,
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
			IncreasedDamage:    40,
		},
		&Magnification{
			MagnificationValue:       325.1,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       333.8,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       899.2,
			TriggerTimes:             1,
			AttackInternalPercentage: aa,
			Name:                     "强化特殊技",
			IncreasedDamage:          40,
		},
		&Magnification{
			MagnificationValue:       3908.6,
			TriggerTimes:             1,
			AttackInternalPercentage: aa,
			Name:                     "终结技",
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9,
			TriggerTimes:             2,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: aa,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
		},
	}
}
func MagnificationBase朋克专武6命() []*Magnification {
	var aa float64 = 0
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    0,
			Name:               "飞弦斩",
			IncreasedDamage:    50 + 40,
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12,
			Name:               "飞弦斩",
			IncreasedDamage:    50 + 40,
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 2,
			Name:               "飞弦斩",
			IncreasedDamage:    50 + 40,
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       4,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
			IncreasedDamage:    50 + 40,
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9*5,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: aa,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 40,
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9*5,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: aa,
			Name:                     "飞弦斩",
			IncreasedDamage:          50 + 40,
		},
		&Magnification{
			MagnificationValue:       3908.6,
			TriggerTimes:             1,
			AttackInternalPercentage: aa,
			Name:                     "终结技",
		},
		&Magnification{
			MagnificationValue:       1500,
			AttackInternalPercentage: aa,
			TriggerTimes:             2,
			Name:                     "6命伤害",
		},
	}
}
func MagnificationBase朋克专武() []*Magnification {
	var aa float64 = 0
	return []*Magnification{
		&Magnification{
			MagnificationValue:       1035.7,
			TriggerTimes:             1,
			Name:                     "连携技",
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       325.1,
			TriggerTimes:             1,
			ExplosiveInjury:          0,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       333.8,
			TriggerTimes:             1,
			ExplosiveInjury:          12,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 2,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       899.2,
			TriggerTimes:             1,
			Name:                     "强化特殊技",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       3908.6,
			TriggerTimes:             1,
			AttackInternalPercentage: aa,
			Name:                     "终结技",
		},
		&Magnification{
			MagnificationValue:       325.1,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       333.8,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			Name:                     "飞弦斩",
			IncreasedDamage:          40,
			AttackInternalPercentage: aa,
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
			IncreasedDamage:    40,
		},
		&Magnification{
			MagnificationValue: 325.1 + 333.8 + 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
			IncreasedDamage:    40,
		},
	}
}
func MagnificationBase硫磺石() []*Magnification {
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue:       325.1,
			TriggerTimes:             1,
			AttackInternalPercentage: 0,
			ExplosiveInjury:          0,
			Name:                     "飞弦斩",
		},
		&Magnification{
			MagnificationValue:       333.8,
			TriggerTimes:             1,
			AttackInternalPercentage: 3.5,
			ExplosiveInjury:          12,
			Name:                     "飞弦斩",
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 2,
			AttackInternalPercentage: 3.5 * 2,
			Name:                     "飞弦斩",
		},
		&Magnification{
			MagnificationValue:       899.2,
			TriggerTimes:             1,
			AttackInternalPercentage: 3.5 * 3,
			Name:                     "强化特殊技",
		},
		&Magnification{
			MagnificationValue:       325.1,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: 3.5 * 4,
			Name:                     "飞弦斩",
		},
		&Magnification{
			MagnificationValue:       333.8,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: 3.5 * 5,
			Name:                     "飞弦斩",
		},
		&Magnification{
			MagnificationValue:       379.9,
			TriggerTimes:             1,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: 3.5 * 6,
			Name:                     "飞弦斩",
		},
		&Magnification{
			MagnificationValue:       899.2,
			TriggerTimes:             1,
			AttackInternalPercentage: 3.5 * 7,
			Name:                     "强化特殊技",
		},
		&Magnification{
			MagnificationValue:       3908.6,
			TriggerTimes:             1,
			AttackInternalPercentage: 3.5 * 8,
			Name:                     "终结技",
		},
		&Magnification{
			MagnificationValue:       325.1 + 333.8 + 379.9,
			TriggerTimes:             2,
			ExplosiveInjury:          12 * 3,
			AttackInternalPercentage: 3.5 * 8,
			Name:                     "飞弦斩",
		},
	}
}

func MagnificationBase星辉引擎() []*Magnification {
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    0,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 2,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue: 3908.6,
			TriggerTimes:       1,
			Name:               "终结技",
		},
		&Magnification{
			MagnificationValue: 325.1 + 333.8 + 379.9,
			TriggerTimes:       2,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
	}
}

func MagnificationBase强音() []*Magnification {
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    0,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 2,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue: 3908.6,
			TriggerTimes:       1,
			Name:               "终结技",
		},
		&Magnification{
			MagnificationValue: 325.1 + 333.8 + 379.9,
			TriggerTimes:       2,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
	}
}

func MagnificationBase加农转子() []*Magnification {
	return []*Magnification{
		&Magnification{
			MagnificationValue: 1035.7,
			TriggerTimes:       1,
			Name:               "连携技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    0,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 2,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue: 325.1,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 333.8,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 379.9,
			TriggerTimes:       1,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
		},
		&Magnification{
			MagnificationValue: 899.2,
			TriggerTimes:       1,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue: 3908.6,
			TriggerTimes:       1,
			Name:               "终结技",
		},
		&Magnification{
			MagnificationValue: 325.1 + 333.8 + 379.9,
			TriggerTimes:       2,
			ExplosiveInjury:    12 * 3,
			Name:               "飞弦斩",
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

func (i *Initializations) InitializationBase1(role *Role.BaseRole, article *arms.MainArticle) {
	i.Basic = &Basic{
		BasicAttack:              role.AttackValue + article.BaseAttackValue, // 基础攻击力（角色+专武）
		BasicCritical:            role.Critical,                              // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     role.ExplosiveInjury,                       // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     role.IncreasedDamage,                       // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: role.ReductionResistance,                   // 基础减抗（角色+武器+驱动盘）
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
		AttackValue:              316,     // 攻击力值增加(固定2号位数值)
		AttackValue2:             0,       // 攻击力值增加(局内加的固定攻击力)
		AttackPowerPercentage:    30 + 10, // 局外攻击力百分比(6号位+武器主词条+5号位+4号位+副词条)
		AttackInternalPercentage: 25,      // 局内攻击力百分比(武器，4件套)
		Critical:                 25 + 10, // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          0,       // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          40,      // 增伤（队友百分比）
		ReductionResistance:      15,      // 减抗（百分比）
		Vulnerable:               0,       // 易伤（百分比）
		SpecialDamage:            0,       // 特殊增伤（百分比）
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

// 6命
func (i *Initializations) InitializationBase2(role *Role.BaseRole, article *arms.MainArticle) {
	i.Basic = &Basic{
		BasicAttack:              role.AttackValue + article.BaseAttackValue, // 基础攻击力（角色+专武）
		BasicCritical:            role.Critical,                              // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     role.ExplosiveInjury,                       // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     role.IncreasedDamage,                       // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: role.ReductionResistance,                   // 基础减抗（角色+武器+驱动盘）
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
		AttackValue:              316,     // 攻击力值增加(固定2号位数值)
		AttackValue2:             0,       // 攻击力值增加(局内加的固定攻击力)
		AttackPowerPercentage:    30,      // 局外攻击力百分比(6号位+武器主词条+5号位+4号位+副词条)
		AttackInternalPercentage: 0,       // 局内攻击力百分比(武器，4件套)
		Critical:                 25 + 10, // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          0,       // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          40,      // 增伤（队友百分比）
		ReductionResistance:      15,      // 减抗（百分比）
		Vulnerable:               0,       // 易伤（百分比）
		SpecialDamage:            0,       // 特殊增伤（百分比）
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

// 6命
func (i *Initializations) InitializationBase3(role *Role.BaseRole, article *arms.MainArticle) {
	i.Basic = &Basic{
		BasicAttack:              role.AttackValue + article.BaseAttackValue, // 基础攻击力（角色+专武）
		BasicCritical:            role.Critical,                              // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     role.ExplosiveInjury,                       // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     role.IncreasedDamage,                       // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: role.ReductionResistance,                   // 基础减抗（角色+武器+驱动盘）
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
		AttackValue:              316,     // 攻击力值增加(固定2号位数值)
		AttackValue2:             0,       // 攻击力值增加(局内加的固定攻击力)
		AttackPowerPercentage:    30,      // 局外攻击力百分比(6号位+武器主词条+5号位+4号位+副词条)
		AttackInternalPercentage: 28,      // 局内攻击力百分比(武器，4件套)
		Critical:                 25 + 10, // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          0,       // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          40 + 10, // 增伤（队友百分比）
		ReductionResistance:      15,      // 减抗（百分比）
		Vulnerable:               0,       // 易伤（百分比）
		SpecialDamage:            0,       // 特殊增伤（百分比）
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

// 0命
func (i *Initializations) InitializationBase朋克(role *Role.BaseRole, article *arms.MainArticle) {
	i.Basic = &Basic{
		BasicAttack:              role.AttackValue + article.BaseAttackValue, // 基础攻击力（角色+专武）
		BasicCritical:            role.Critical,                              // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     role.ExplosiveInjury,                       // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     role.IncreasedDamage,                       // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: role.ReductionResistance,                   // 基础减抗（角色+武器+驱动盘）
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
		AttackValue:              316,     // 攻击力值增加(固定2号位数值)
		AttackValue2:             0,       // 攻击力值增加(局内加的固定攻击力)
		AttackPowerPercentage:    30 + 10, // 局外攻击力百分比(6号位+武器主词条+5号位+4号位+副词条)
		AttackInternalPercentage: 25,      // 局内攻击力百分比(武器，4件套)
		Critical:                 25,      // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          0,       // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          40,      // 增伤（队友百分比）
		ReductionResistance:      0,       // 减抗（百分比）
		Vulnerable:               0,       // 易伤（百分比）
		SpecialDamage:            0,       // 特殊增伤（百分比）
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

func (i *Initializations) InitializationBase如影(role *Role.BaseRole, article *arms.MainArticle) {
	i.Basic = &Basic{
		BasicAttack:              role.AttackValue + article.BaseAttackValue, // 基础攻击力（角色+专武）
		BasicCritical:            role.Critical,                              // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     role.ExplosiveInjury,                       // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     role.IncreasedDamage,                       // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: role.ReductionResistance,                   // 基础减抗（角色+武器+驱动盘）
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
		AttackPowerPercentage:    30,  // 局外攻击力百分比(6号位+武器主词条+5号位+4号位+副词条)
		AttackInternalPercentage: 0,   // 局内攻击力百分比(武器，4件套)
		Critical:                 25,  // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          0,   // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          40,  // 增伤（队友百分比）
		ReductionResistance:      0,   // 减抗（百分比）
		Vulnerable:               0,   // 易伤（百分比）
		SpecialDamage:            0,   // 特殊增伤（百分比）
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

func (i *Initializations) InitializationBase雷暴(role *Role.BaseRole, article *arms.MainArticle) {
	i.Basic = &Basic{
		BasicAttack:              role.AttackValue + article.BaseAttackValue, // 基础攻击力（角色+专武）
		BasicCritical:            role.Critical,                              // 基础暴击（角色+武器+2件套+4号位）
		BasicExplosiveInjury:     role.ExplosiveInjury,                       // 基础爆伤（角色+武器+2件套+4号位）
		BasicIncreasedDamage:     role.IncreasedDamage,                       // 基础增伤（角色+武器+驱动盘）
		BasicReductionResistance: role.ReductionResistance,                   // 基础减抗（角色+武器+驱动盘）
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
		AttackValue:              316,     // 攻击力值增加(固定2号位数值)
		AttackValue2:             0,       // 攻击力值增加(局内加的固定攻击力)
		AttackPowerPercentage:    30,      // 局外攻击力百分比(6号位+武器主词条+5号位+4号位+副词条)
		AttackInternalPercentage: 28,      // 局内攻击力百分比(武器，4件套)
		Critical:                 25,      // 增加暴击（角色+武器+4件套）
		ExplosiveInjury:          0,       // 增加爆伤（角色+武器+2件套+4号位）
		IncreasedDamage:          40 + 10, // 增伤（队友百分比）
		ReductionResistance:      0,       // 减抗（百分比）
		Vulnerable:               0,       // 易伤（百分比）
		SpecialDamage:            0,       // 特殊增伤（百分比）
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
