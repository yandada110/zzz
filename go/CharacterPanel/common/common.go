package common

import "math"

const (
	Critical                 = "Critical"                 // 暴击
	ExplosiveInjury          = "ExplosiveInjury"          // 爆伤
	AttackPowerPercentage    = "AttackPowerPercentage"    // 局外百分比攻击力
	AttackInternalPercentage = "AttackInternalPercentage" // 局内百分比攻击力
	ReductionResistance      = "ReductionResistance"      // 减抗
	IncreasedDamage          = "IncreasedDamage"          // 增伤
	Penetrate                = "Penetrate"                // 穿透
	AttackValue              = "AttackValue"              // 攻击力值
	Proficient               = "BasicProficient"          // 精通
	DefenseBreak             = "DefenseBreak"             // 破防
	PenetrationValue         = "PenetrationValue"         // 穿透值
)

var AttackPercentageEntriesLimit = map[string]int{
	Critical:        15,
	ExplosiveInjury: 15,
	Proficient:      15,
}

var ExplosiveInjuryEntriesLimit = map[string]int{
	Critical:        30,
	ExplosiveInjury: 25,
	Proficient:      30,
}

var ProficientEntriesLimit = map[string]int{
	Critical:        30,
	ExplosiveInjury: 30,
	Proficient:      25,
}

var AttackValueEntriesLimit = map[string]int{
	Critical:        30,
	ExplosiveInjury: 30,
	Proficient:      30,
}

var PenetrationValueEntriesLimit = map[string]int{
	Critical:        30,
	ExplosiveInjury: 30,
	Proficient:      30,
}

const (
	GainFormInsideTheBureau = 1 // 局内增益
	GainFormInsideFixed     = 2 // 固定值增益
)

// AllowedGroupB 定义允许的增伤、穿透原始分配组合
var AllowedGroupB = [][2]int{
	{0, 13},
	{3, 10},
	{10, 3},
	{13, 0},
	{0, 0},
	{3, 0},
	{0, 3},
}

const (
	DirectInjury = "DirectInjury" // 直伤
	Abnormal     = "Abnormal"     // 异常
	Disorder     = "Disorder"     // 紊乱
)

// 异常伤害倍率固定
const (
	Fire                = "Fire"        // 火
	Electricity         = "Electricity" // 电
	Physical            = "Physical"    // 物理
	Ice                 = "Ice"         // 冰
	Ether               = "Ether"       // 以太
	TimeTotal   float64 = 10            // 固定异常时间
)

var AbnormalMagnification = map[string]float64{
	Fire:        50,
	Electricity: 125,
	Physical:    713,
	Ice:         500,
	Ether:       62.5,
}

var DisorderMagnification = map[string]float64{
	Fire:        4.5+FLOOR((10-C21)/0.5,2)*0.5),
	Electricity: 125,
	Physical:    713,
	Ice:         500,
	Ether:       62.5,
}

func Floor(number,timeTotal, significance float64) float64 {
	return math.Floor(number/significance) * significance
}