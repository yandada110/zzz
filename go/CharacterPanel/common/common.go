package common

const (
	Critical                 = "Critical"              // 暴击
	ExplosiveInjury          = "ExplosiveInjury"       // 爆伤
	AttackPowerPercentage    = "AttackPowerPercentage" // 局外百分比攻击力
	AttackInternalPercentage = "AttackPowerPercentage" // 局内百分比攻击力
	ReductionResistance      = "ReductionResistance"   // 减抗
	IncreasedDamage          = "IncreasedDamage"       // 增伤
	Penetrate                = "Penetrate"             // 穿透
	AttackValue              = "AttackValue"           //攻击力
	Proficient               = "Proficient"            //精通
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
