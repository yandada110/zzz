package main

// DeepCopy 深拷贝方法（结构体方法）
func (i *Initialization) DeepCopy() *Initialization {
	// 手动拷贝嵌套结构体，防止修改原始对象影响拷贝对象
	copyData := Initialization{
		Magnifications: i.Magnifications,
		CurrentPanel:   i.CurrentPanel,
		Output:         i.Output,
		Gain:           i.Gain,
		Name:           i.Name,
	}
	return &copyData
}

// DeepCopy 深拷贝方法（结构体方法）
func (i *Initializations) DeepCopyData(list []*Initialization) *Initializations {
	// 手动拷贝嵌套结构体，防止修改原始对象影响拷贝对象
	copyData := Initializations{
		Name:                 i.Name,
		NumberFour:           i.NumberFour,
		AttackCount:          i.AttackCount,
		CriticalCount:        i.CriticalCount,
		ExplosiveInjuryCount: i.ExplosiveInjuryCount,
		Basic:                i.Basic,
		Gain:                 i.Gain,
		Defense:              i.Defense,
		Initializations:      list,
	}
	return &copyData
}

// ------------------------ 数据结构定义 ------------------------
type Initializations struct {
	Name                 string            // 队伍名称
	NumberFour           string            // 4号位固定属性 暴击或者爆伤
	AttackCount          int               // 攻击力词条基础上限
	CriticalCount        int               // 暴击词条基础上限
	ExplosiveInjuryCount int               // 爆伤词条基础上限
	Basic                *Basic            // 角色基础面板，不变
	Gain                 *Gain             // 队友增益，不变
	Defense              *Defense          // 破防收益，不变
	Initializations      []*Initialization // 计算不同模型集合（可以包含不同模型）
}

type Initialization struct {
	Magnifications []*Magnification // 各技能倍率列表
	CurrentPanel   *CurrentPanel
	Output         *Output
	Gain           *Gain  // 不同模型下存在改动：比如失衡状态，加攻击力，加易伤倍率
	Name           string // 模型名称
}

type ExternalPanel struct {
	Attack          float64
	Critical        float64
	ExplosiveInjury float64
}

type Basic struct {
	BasicAttack              float64
	BasicCritical            float64
	BasicExplosiveInjury     float64
	BasicIncreasedDamage     float64
	BasicReductionResistance float64
	BasicVulnerable          float64
	BasicSpecialDamage       float64
	Penetration              float64
}

type CurrentPanel struct {
	Attack              float64
	Critical            float64
	ExplosiveInjury     float64
	IncreasedDamage     float64
	ReductionResistance float64
	Vulnerable          float64
	SpecialDamage       float64
	Penetration         float64
}

type Magnification struct {
	MagnificationValue  float64
	TriggerTimes        float64
	Name                string
	IncreasedDamage     float64
	ReductionResistance float64
	DefenseBreak        float64
	Penetration         float64
	SpecialDamage       float64
	ExplosiveInjury     float64
}

// 队友提供的局内增益效果
type Gain struct {
	AttackValue              float64
	AttackValue2             float64
	AttackPowerPercentage    float64
	AttackInternalPercentage float64
	Critical                 float64
	ExplosiveInjury          float64
	IncreasedDamage          float64
	ReductionResistance      float64
	Vulnerable               float64
	SpecialDamage            float64
	Penetration              float64
}

// 失衡状态下，额外提供的增益效果
type ImbalanceGain struct {
	AttackValue              float64
	AttackValue2             float64
	AttackPowerPercentage    float64
	AttackInternalPercentage float64
	Critical                 float64
	ExplosiveInjury          float64
	IncreasedDamage          float64
	ReductionResistance      float64
	Vulnerable               float64
	SpecialDamage            float64
}

type Defense struct {
	Penetration      float64
	DefenseBreak     float64
	PenetrationValue float64
}

type Output struct {
	BasicDamageArea         float64
	IncreasedDamageArea     float64
	ExplosiveInjuryArea     float64
	DefenseArea             float64
	ReductionResistanceArea float64
	VulnerableArea          float64
	SpecialDamageArea       float64
}
