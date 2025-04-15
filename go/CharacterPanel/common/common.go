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
	Proficient               = "Proficient"               // 精通
	DefenseBreak             = "DefenseBreak"             // 破防
	PenetrationValue         = "PenetrationValue"         // 穿透值
)

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
	Different    = "Different"    // 异放
)

// 异常伤害倍率固定
const (
	Fire        = "Fire"        // 火
	Electricity = "Electricity" // 电
	Physical    = "Physical"    // 物理
	Ice         = "Ice"         // 冰
	Ether       = "Ether"       // 以太
)

var TimeTotal float64 = 10 // 固定异常时间

var AbnormalMagnification = map[string]float64{
	Fire:        50,
	Electricity: 125,
	Physical:    713,
	Ice:         500,
	Ether:       62.5,
}

var DisorderMagnification = map[string]float64{
	Fire:        50,
	Electricity: 125,
	Physical:    7.5,
	Ice:         7.5,
	Ether:       62.5,
}

var DifferentMagnification = map[string]float64{
	Fire:        8,
	Electricity: 3.2,
	Physical:    0.75,
	Ice:         1.08,
	Ether:       6.15,
}

func FireArea(totalTime, usedTime, rate float64) float64 {
	// 将 rate 转换为小数，如 50 -> 0.5
	r := rate / 100
	remaining := totalTime - usedTime

	// 计算 (totalTime - usedTime) / 0.5
	value := remaining / r

	// 按 2 的倍数向下取整，等效于 Excel 的 FLOOR(value, 2)
	floored := 2 * math.Floor(value/2)

	// 返回结果：4.5 + (取整后的值 * 0.5)
	return 4.5 + floored*r
}

func PhysicalArea(totalTime, usedTime, rate float64) float64 {
	significance := 2.0
	rate = rate / 100
	// 计算剩余时间
	remaining := totalTime - usedTime
	// 将 remaining 向下取整到最接近的 2 的倍数
	floored := math.Floor(remaining/significance) * significance
	// 最终结果：4.5 + (floored * multiplier)
	return 4.5 + floored*rate
}

func EtherArea(totalTime, usedTime, rate float64) float64 {
	rate = rate / 100
	// 第一步：计算 (totalTime - usedTime) / step
	ratio := (totalTime - usedTime) / 0.5
	// 第二步：向下取整到最接近的 2 的倍数
	floored := math.Floor(ratio/2) * 2
	// 第三步：乘以倍率再加上 4.5
	return 4.5 + floored*rate
}

func IceArea(totalTime, usedTime, rate float64) float64 {
	rate = rate / 100
	significance := 2.0
	// 计算剩余时间
	remaining := totalTime - usedTime
	// 将 remaining 向下取整到最接近的 2 的倍数
	floored := math.Floor(remaining/significance) * significance
	// 最终结果：4.5 + (floored * multiplier)
	return 4.5 + floored*rate
}

func ElectricityArea(totalTime, usedTime, rate float64) float64 {
	rate = rate / 100
	significance := 2.0
	remaining := totalTime - usedTime
	// 向下取整到最近的 2 的倍数
	floored := math.Floor(remaining/significance) * significance
	return 4.5 + floored*rate
}
