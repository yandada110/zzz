package main

// DeepCopy 对 Initialization 进行深拷贝
func (i *Initialization) DeepCopy() *Initialization {
	// 拷贝 Magnifications 切片
	var copyMagnifications []*Magnification
	if i.Magnifications != nil {
		copyMagnifications = make([]*Magnification, len(i.Magnifications))
		for idx, m := range i.Magnifications {
			copyMagnifications[idx] = m.DeepCopy()
		}
	}
	return &Initialization{
		Magnifications: copyMagnifications,
		CurrentPanel:   i.CurrentPanel.DeepCopy(),
		Output:         i.Output.DeepCopy(),
		Gain:           i.Gain.DeepCopy(),
		Name:           i.Name,
	}
}

// DeepCopyData 对 Initializations 进行深拷贝（包括其内部 Initialization 列表）
func (i *Initializations) DeepCopyData(list []*Initialization) *Initializations {
	var copyList []*Initialization
	if list != nil {
		copyList = make([]*Initialization, len(list))
		for idx, init := range list {
			copyList[idx] = init.DeepCopy()
		}
	}
	return &Initializations{
		Name:            i.Name,
		NumberFour:      i.NumberFour,
		Basic:           i.Basic.DeepCopy(),
		Gain:            i.Gain.DeepCopy(),
		Defense:         i.Defense.DeepCopy(),
		Condition:       i.Condition.DeepCopy(),
		Initializations: copyList,
	}
}

// 以下是各结构体的深拷贝辅助方法

func (b *Basic) DeepCopy() *Basic {
	if b == nil {
		return nil
	}
	copyB := *b
	return &copyB
}

func (c *CurrentPanel) DeepCopy() *CurrentPanel {
	if c == nil {
		return nil
	}
	copyC := *c
	return &copyC
}

func (m *Magnification) DeepCopy() *Magnification {
	if m == nil {
		return nil
	}
	copyM := *m
	return &copyM
}

func (g *Gain) DeepCopy() *Gain {
	if g == nil {
		return nil
	}
	copyG := *g
	return &copyG
}

func (d *Defense) DeepCopy() *Defense {
	if d == nil {
		return nil
	}
	copyD := *d
	return &copyD
}

func (c *Condition) DeepCopy() *Condition {
	if c == nil {
		return nil
	}
	copyD := *c
	return &copyD
}

func (o *Output) DeepCopy() *Output {
	if o == nil {
		return nil
	}
	copyO := *o
	return &copyO
}

type Initializations struct {
	Name            string            // 队伍名称
	NumberFour      string            // 4号位固定属性 暴击或者爆伤
	NumberSix       string            // 6号位固定属性攻击，或者其他
	NumberFive      string            // 5号位固定属性攻击，或者其他
	Basic           *Basic            // 角色基础面板，不变
	Gain            *Gain             // 队友增益，不变
	Defense         *Defense          // 破防收益，不变
	Initializations []*Initialization // 计算不同模型集合（可以包含不同模型）
	Condition       *Condition        // 词条限制条件
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
	BasicAttack              float64 // 攻击
	BasicCritical            float64 // 暴击
	BasicExplosiveInjury     float64 // 爆伤
	BasicIncreasedDamage     float64 // 增伤
	BasicReductionResistance float64 // 减抗
	BasicVulnerable          float64 // 易伤
	BasicSpecialDamage       float64 // 特殊乘区增伤
	Penetration              float64 // 穿透
	BasicProficient          float64 // 精通
	AddTime                  float64 // 角色自身延长异常时间
	BasicAttackValue         float64 // 攻击力值
	PenetrationValue         float64 // 穿透值
}

type CurrentPanel struct {
	Attack              float64 // 攻击
	Critical            float64 // 暴击
	ExplosiveInjury     float64 // 爆伤
	IncreasedDamage     float64 // 增伤
	ReductionResistance float64 // 减抗
	Vulnerable          float64 // 易伤
	SpecialDamage       float64 // 特殊乘区增伤
	Penetration         float64 // 穿透
	DefenseBreak        float64 // 破防
	PenetrationValue    float64 // 穿透值
	Proficient          float64 // 精通
}

type Magnification struct {
	MagnificationValue       float64 // 基础倍率 如技能倍率，异常倍率，紊乱倍率
	TriggerTimes             float64 // 计算次数
	Name                     string  // 技能名称
	IncreasedDamage          float64 // 增伤
	ReductionResistance      float64 // 减抗
	DefenseBreak             float64 // 破防
	Penetration              float64 // 穿透
	SpecialDamage            float64 // 特殊增伤
	ExplosiveInjury          float64 // 爆伤
	AttackInternalPercentage float64 // 局内攻击力
	Critical                 float64 // 暴击率
	DamageType               string  // 伤害类型，默认直伤类型
	DisorderType             string  // 紊乱伤害类型
	TimeConsumption          float64 // 异常消耗时间量
	Damage                   float64 // 针对异常伤害-计算异放伤害
	Proficient               float64 // 精通
}

// 队友提供的局内增益效果
type Gain struct {
	AttackValue              float64 // 攻击-局外攻击力2号位，副词条
	AttackValue2             float64 // 攻击-局内攻击力值-嘉音，凯撒
	AttackPowerPercentage    float64 // 攻击百分比-局外-5号位，音擎主词条-2件套，副词条
	AttackInternalPercentage float64 // 攻击百分比-局内-驱动盘四件套-阿炮
	Critical                 float64 // 暴击 包括角色局内暴击率-队友暴击率
	ExplosiveInjury          float64 // 爆伤 队友爆伤，角色局内爆伤
	IncreasedDamage          float64 // 增伤
	ReductionResistance      float64 // 减抗
	Vulnerable               float64 // 易伤
	SpecialDamage            float64 // 特殊城区
	Penetration              float64 // 穿透
	DefenseBreak             float64 // 破防
	PenetrationValue         float64 // 穿透值
	Proficient               float64 // 精通
	AddTime                  float64 // 异常时间延长
}

type Defense struct {
	Penetration      float64
	DefenseBreak     float64
	PenetrationValue float64
}

type Output struct {
	BasicDamageArea         float64 // 基础区
	IncreasedDamageArea     float64 // 增伤区
	ExplosiveInjuryArea     float64 // 爆伤区
	DefenseArea             float64 // 防御区
	ReductionResistanceArea float64 // 减抗区
	VulnerableArea          float64 // 易伤区
	SpecialDamageArea       float64 // 特殊区
	GradeArea               float64 // 等级区
	ProficientArea          float64 // 精通区
	DifferentArea           float64 // 异放倍率
}
