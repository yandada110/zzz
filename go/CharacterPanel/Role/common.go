package Role

type BaseRole struct {
	AttackValue         float64 // 基础攻击力
	Critical            float64 // 基础暴击
	ExplosiveInjury     float64 // 基础爆伤
	Proficient          float64 // 基础精通
	ReductionResistance float64 // 基础减抗
	IncreasedDamage     float64 // 基础全局增伤
	BasicVulnerable     float64 // 基础易伤
	BasicSpecialDamage  float64 // 基础特殊伤害乘区
	Penetration         float64 // 基础穿透率
	DefenseBreak        float64 // 基础破防
}

type BuffCharacter struct {
	AttackValue              float64 // 提供攻击力
	AttackInternalPercentage float64 // 局内攻击力
	Critical                 float64 // 提供暴击
	ExplosiveInjury          float64 // 提供爆伤
	Proficient               float64 // 提供精通
	ReductionResistance      float64 // 提供减抗
	IncreasedDamage          float64 // 提供全局增伤
	Vulnerable               float64 // 提供易伤
	SpecialDamage            float64 // 提供特殊伤害乘区
	Penetration              float64 // 提供穿透率
	DefenseBreak             float64 // 提供破防
}
