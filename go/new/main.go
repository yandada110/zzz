package main

import (
	"fmt"
	"math"
)

func main() {
	// 初始化角色数据
	initialization := NewInitialization()
	// 通过遍历所有可能分配方案寻找最佳词条分配
	bestDamage, bestDistribution, bestPanel, bestOutput, bestCritConverted := initialization.FindOptimalDistribution()

	// 计算局外面板
	externalPanel := initialization.CalculateExternalPanel(bestDistribution, bestCritConverted)

	// 输出最佳方案及对应面板和伤害信息
	fmt.Println("最佳词条分配方案:")
	fmt.Println(fmt.Sprintf(" 攻击力: %d, 暴击: %d, 爆伤: %d, 增伤: %d, 穿透: %d, 暴击转换爆伤: %d",
		bestDistribution[AttackPercentage],
		bestDistribution[Critical],
		bestDistribution[ExplosiveInjury],
		bestDistribution[IncreasedDamage],
		bestDistribution[Penetrate],
		bestCritConverted,
	))
	fmt.Println(fmt.Sprintf("最高总伤害: %.6f", bestDamage))
	fmt.Println("局内面板:")
	fmt.Println(fmt.Sprintf(" 攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%",
		bestPanel.Attack,
		bestPanel.Critical,
		bestPanel.ExplosiveInjury,
		bestPanel.IncreasedDamage,
		bestPanel.Penetration,
	))
	fmt.Println("局外面板:")
	fmt.Println(fmt.Sprintf(" 攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%",
		externalPanel.Attack,
		externalPanel.Critical,
		externalPanel.ExplosiveInjury,
	))
	fmt.Println("增益输出详情:")
	fmt.Println(fmt.Sprintf(" 基础伤害区: %.2f, 增伤区: %.2f, 爆伤区: %.2f, 防御区: %.2f, 减抗区: %.2f, 易伤区: %.2f, 特殊乘区: %.2f",
		bestOutput.BasicDamageArea,
		bestOutput.IncreasedDamageArea,
		bestOutput.ExplosiveInjuryArea,
		bestOutput.DefenseArea,
		bestOutput.ReductionResistanceArea,
		bestOutput.VulnerableArea,
		bestOutput.SpecialDamageArea,
	))
}

// CalculateExternalPanel 根据最佳词条分配方案计算局外面板
// 局外面板公式：
//
//	攻击力 = Basic.BasicAttack * (1 + (Gain.AttackPowerPercentage + 攻击力词条数*3) / 100)
//	暴击率 = Basic.BasicCritical + 暴击词条数*2.4
//	爆伤   = Basic.BasicExplosiveInjury + (爆伤词条数 + 暴击转换爆伤词条数量) * 2.4
func (i *Initialization) CalculateExternalPanel(distribution map[string]int, critConverted int) *ExternalPanel {
	attack := i.Basic.BasicAttack*(1+(i.Gain.AttackPowerPercentage+float64(distribution[AttackPercentage])*3)/100) + i.Gain.AttackValue
	critical := i.Basic.BasicCritical + float64(distribution[Critical])*2.4
	explosiveInjury := i.Basic.BasicExplosiveInjury + (float64(distribution[ExplosiveInjury])+float64(critConverted))*4.8
	return &ExternalPanel{
		Attack:          attack,
		Critical:        critical,
		ExplosiveInjury: explosiveInjury,
	}
}

// ExternalPanel 局外面板数据结构
type ExternalPanel struct {
	Attack          float64
	Critical        float64
	ExplosiveInjury float64
}

// --------------------- 以下代码与之前保持一致 ---------------------

// FindOptimalDistribution 遍历所有可能的词条分配，找出最佳方案，新增返回 bestCritConverted
func (i *Initialization) FindOptimalDistribution() (bestDamage float64, bestDistribution map[string]int, bestPanel *CurrentPanel, bestOutput *Output, bestCritConverted int) {
	totalTokens := i.Condition.MainArticle
	bestDamage = -1.0
	bestDistribution = make(map[string]int)
	bestCritConverted = 0
	// 穷举所有组合：a+b+c+d+e = totalTokens（5 个词条类型）
	for a := 0; a <= totalTokens; a++ {
		for b := 0; a+b <= totalTokens; b++ {
			for c := 0; a+b+c <= totalTokens; c++ {
				for d := 0; a+b+c+d <= totalTokens; d++ {
					e := totalTokens - (a + b + c + d)
					distribution := map[string]int{
						AttackPercentage: a,
						Critical:         b,
						ExplosiveInjury:  c,
						IncreasedDamage:  d,
						Penetrate:        e,
					}
					// 克隆一个新的初始化，保证各项状态从初始值开始
					sim := i.Clone()
					sim.ResetCondition() // 重置暴击溢出、流转等计数
					// 根据当前分配方案构建面板
					sim.CharacterPanelWithDistribution(distribution)
					damage := sim.CalculatingTotalDamage()
					if damage > bestDamage {
						bestDamage = damage
						bestDistribution = distribution
						bestPanel = sim.ClonePanel()
						bestOutput = sim.CloneOutput()
						bestCritConverted = sim.CritConverted
					}
				}
			}
		}
	}
	return bestDamage, bestDistribution, bestPanel, bestOutput, bestCritConverted
}

// CharacterPanelWithDistribution 根据完整的词条分配更新面板
func (i *Initialization) CharacterPanelWithDistribution(distribution map[string]int) *Initialization {
	i.CurrentPanel = &CurrentPanel{
		ReductionResistance: i.Basic.BasicReductionResistance + i.Gain.ReductionResistance,
		Vulnerable:          i.Basic.BasicVulnerable + i.Gain.Vulnerable,
		SpecialDamage:       i.Basic.BasicSpecialDamage + i.Gain.SpecialDamage,
	}
	i.HandleBasicAttack(AttackPercentage, distribution[AttackPercentage])
	// 先计算暴击（包括转换溢出）
	i.HandleBasicCritical(Critical, distribution[Critical])
	// 再计算爆伤，利用暴击转换结果
	i.HandleBasicExplosiveInjury(ExplosiveInjury, distribution[ExplosiveInjury])
	i.HandleBasicIncreasedDamage(IncreasedDamage, distribution[IncreasedDamage])
	i.HandlePenetrateDamage(Penetrate, distribution[Penetrate])
	return i
}

// ResetCondition 重置 Condition 中的计数状态，并初始化流转字段
func (i *Initialization) ResetCondition() {
	i.Condition.CriticalStatus = 0
	i.Condition.CriticalCount = 0
	i.CritConverted = 0
	i.ExtraPenetrateTokens = 0
	i.ExtraIncreasedDamageTokens = 0
}

// Clone 克隆 Initialization 对象（部分字段采用深拷贝，假定 Magnifications 为常量）
func (i *Initialization) Clone() *Initialization {
	return &Initialization{
		Magnifications: i.Magnifications, // 假定不变
		Basic: &Basic{
			BasicAttack:              i.Basic.BasicAttack,
			BasicCritical:            i.Basic.BasicCritical,
			BasicExplosiveInjury:     i.Basic.BasicExplosiveInjury,
			BasicIncreasedDamage:     i.Basic.BasicIncreasedDamage,
			BasicReductionResistance: i.Basic.BasicReductionResistance,
			BasicVulnerable:          i.Basic.BasicVulnerable,
			BasicSpecialDamage:       i.Basic.BasicSpecialDamage,
			Penetration:              i.Basic.Penetration,
		},
		Gain: &Gain{
			AttackValue:              i.Gain.AttackValue,
			AttackValue2:             i.Gain.AttackValue2,
			AttackPowerPercentage:    i.Gain.AttackPowerPercentage,
			AttackInternalPercentage: i.Gain.AttackInternalPercentage,
			Critical:                 i.Gain.Critical,
			ExplosiveInjury:          i.Gain.ExplosiveInjury,
			IncreasedDamage:          i.Gain.IncreasedDamage,
			ReductionResistance:      i.Gain.ReductionResistance,
			Vulnerable:               i.Gain.Vulnerable,
			SpecialDamage:            i.Gain.SpecialDamage,
		},
		Defense: &Defense{
			Penetration:      i.Defense.Penetration,
			DefenseBreak:     i.Defense.DefenseBreak,
			PenetrationValue: i.Defense.PenetrationValue,
		},
		Condition: &Condition{
			MainArticle:    i.Condition.MainArticle,
			Critical:       i.Condition.Critical,
			CriticalStatus: 0,
			CriticalCount:  0,
		},
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
		// 新增字段
		CritConverted:              i.CritConverted,
		ExtraPenetrateTokens:       i.ExtraPenetrateTokens,
		ExtraIncreasedDamageTokens: i.ExtraIncreasedDamageTokens,
	}
}

// ClonePanel 克隆当前面板
func (i *Initialization) ClonePanel() *CurrentPanel {
	cp := *i.CurrentPanel
	return &cp
}

// CloneOutput 克隆当前输出
func (i *Initialization) CloneOutput() *Output {
	op := *i.Output
	return &op
}

// 以下函数保持原有处理逻辑，仅作少量注释调整

func (i *Initialization) HandleBasicAttack(key string, count int) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage
	if key == AttackPercentage {
		fmt.Println(fmt.Sprintf("词条使用数量：攻击力：%d", count))
		attackPowerPercentage += 3 * float64(count)
	}
	i.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100) + i.Gain.AttackValue + i.Gain.AttackValue2) * (1 + i.Gain.AttackInternalPercentage/100)
}

func (i *Initialization) HandleBasicCritical(key string, count int) {
	if key == Critical {
		fmt.Println(fmt.Sprintf("词条使用数量：暴击率：%d", count))
		baseCrit := i.Basic.BasicCritical + i.Gain.Critical
		// 每个词条增加 2.4% 暴击率
		maxCritTokens := int((i.Condition.Critical - baseCrit) / 2.4)
		if maxCritTokens < 0 {
			maxCritTokens = 0
		}
		if count <= maxCritTokens {
			// 未超出阈值
			i.CurrentPanel.Critical = baseCrit + 2.4*float64(count)
			i.CritConverted = 0
		} else {
			// 超出部分转换为额外爆伤
			i.CurrentPanel.Critical = i.Condition.Critical
			overflowTokens := count - maxCritTokens
			i.CritConverted = overflowTokens
			fmt.Println(fmt.Sprintf("暴击溢出处理: 基础暴击=%.2f%%, 阈值=%.2f%%, 可用词条数=%d, 溢出词条数=%d, 最终暴击率=%.2f%%",
				baseCrit, i.Condition.Critical, maxCritTokens, overflowTokens, i.CurrentPanel.Critical))
		}
	} else {
		// 非暴击词条，保持基础暴击率
		i.CurrentPanel.Critical = i.Basic.BasicCritical + i.Gain.Critical
	}
}

func (i *Initialization) HandleBasicExplosiveInjury(key string, count int) {
	explosiveInjury := i.Gain.ExplosiveInjury
	if key == ExplosiveInjury {
		fmt.Println(fmt.Sprintf("词条使用数量：爆伤：%d", count))
		explosiveInjury += 4.8 * float64(count)
	}
	// 计算因暴击溢出转换得到的额外爆伤加成
	convertedBonus := 4.8 * float64(i.CritConverted)
	i.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + explosiveInjury + convertedBonus
}

func (i *Initialization) HandleBasicIncreasedDamage(key string, count int) {
	if key == IncreasedDamage {
		fmt.Println(fmt.Sprintf("词条使用数量：增伤：%d", count))
		var effectiveTokens, extraTokens int
		if count < 3 {
			effectiveTokens = 0
			extraTokens = count
		} else if count >= 3 && count < 10 {
			effectiveTokens = 3
			extraTokens = count - 3
		} else if count == 10 {
			effectiveTokens = 10
			extraTokens = 0
		} else if count > 10 && count < 13 {
			effectiveTokens = 10
			extraTokens = count - 10
		} else if count == 13 {
			effectiveTokens = 13
			extraTokens = 0
		} else { // count > 13
			effectiveTokens = 13
			extraTokens = count - 13
		}
		i.ExtraIncreasedDamageTokens = extraTokens
		// 每个有效增伤词条提供 3% 的增伤加成
		increasedDamageValue := i.Gain.IncreasedDamage + 3*float64(effectiveTokens)
		i.CurrentPanel.IncreasedDamage = i.Basic.BasicIncreasedDamage + increasedDamageValue
		fmt.Println("流转增伤词条数:", extraTokens)
	}
}

func (i *Initialization) HandlePenetrateDamage(key string, count int) {
	if key == Penetrate {
		fmt.Println(fmt.Sprintf("词条使用数量：穿透：%d", count))
		var effectiveTokens, extraTokens int
		if count < 3 {
			effectiveTokens = 0
			extraTokens = count
		} else if count >= 3 && count < 10 {
			effectiveTokens = 3
			extraTokens = count - 3
		} else if count == 10 {
			effectiveTokens = 10
			extraTokens = 0
		} else if count > 10 && count < 13 {
			effectiveTokens = 10
			extraTokens = count - 10
		} else if count == 13 {
			effectiveTokens = 13
			extraTokens = 0
		} else { // count > 13
			effectiveTokens = 13
			extraTokens = count - 13
		}
		i.ExtraPenetrateTokens = extraTokens
		// 每个有效穿透词条增加 2.4%
		penetrationValue := i.Defense.Penetration + 2.4*float64(effectiveTokens)
		if penetrationValue >= 100 {
			penetrationValue = 100
		}
		i.CurrentPanel.Penetration = penetrationValue
		fmt.Println("流转穿透词条数:", extraTokens)
	}
}

func (i *Initialization) CalculatingTotalDamage() float64 {
	var totalDamage float64
	for _, mag := range i.Magnifications {
		i.InitializationArea(mag)
		damage := i.Output.BasicDamageArea *
			i.Output.IncreasedDamageArea *
			i.Output.ExplosiveInjuryArea *
			i.Output.DefenseArea *
			i.Output.ReductionResistanceArea *
			i.Output.VulnerableArea *
			i.Output.SpecialDamageArea *
			(1 + mag.SpecialDamage/100)
		totalDamage += damage
		fmt.Println(fmt.Sprintf("[伤害详情] 技能: %s", mag.Name))
		fmt.Println(fmt.Sprintf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%",
			i.CurrentPanel.Attack,
			i.CurrentPanel.Critical,
			i.CurrentPanel.ExplosiveInjury,
			(i.Output.IncreasedDamageArea-1)*100,
		))
		fmt.Println(fmt.Sprintf("  抗性区: %.2f%%, 易伤区: %.2f%%, 特殊乘区: %.2f%%",
			(i.Output.ReductionResistanceArea-1)*100,
			(i.Output.VulnerableArea-1)*100,
			(i.Output.SpecialDamageArea-1)*100,
		))
		fmt.Println(fmt.Sprintf("  当前伤害: %.2f万, 累计总伤害: %.2f万",
			damage/10000,
			totalDamage/10000,
		))
		fmt.Println("----------------------------------------")
	}
	return totalDamage
}

func DecimalToPercentage(newNumber, oldNumber float64) float64 {
	if oldNumber == 0 {
		return 0 // 避免除零错误
	}
	decimalPart := (newNumber - oldNumber) / oldNumber
	percentage := decimalPart * 100
	multiplier := math.Pow(10, 3)
	return math.Floor(percentage*multiplier+0.5) / multiplier
}

func (i *Initialization) InitializationArea(magnification *Magnification) {
	i.BasicDamageArea(magnification)
	i.IncreasedDamageArea(magnification)
	i.ExplosiveInjuryArea(magnification)
	i.DefenseArea(magnification)
	i.ReductionResistanceArea(magnification)
	i.VulnerableArea()
	i.SpecialDamageArea()
}

func (i *Initialization) BasicDamageArea(magnification *Magnification) {
	i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
}

func (i *Initialization) IncreasedDamageArea(magnification *Magnification) {
	i.Output.IncreasedDamageArea = 1 + (magnification.IncreasedDamage+i.CurrentPanel.IncreasedDamage)/100
}

func (i *Initialization) ExplosiveInjuryArea(magnification *Magnification) {
	i.Output.ExplosiveInjuryArea = 1 + (i.CurrentPanel.Critical*(i.CurrentPanel.ExplosiveInjury)*(1+magnification.ExplosiveInjury/100))/10000
}

func (i *Initialization) DefenseArea(magnification *Magnification) {
	characterBase, TotalDefense := 793.783, 873.1613
	penetration := (i.CurrentPanel.Penetration + magnification.Penetration) / 100
	defenseBreak := (i.Defense.DefenseBreak + magnification.DefenseBreak) / 100
	i.Output.DefenseArea = characterBase / (TotalDefense*(1-penetration)*(1-defenseBreak) - i.Defense.PenetrationValue + characterBase)
}

func (i *Initialization) ReductionResistanceArea(magnification *Magnification) {
	i.Output.ReductionResistanceArea = 1 + (magnification.ReductionResistance+i.CurrentPanel.ReductionResistance)/100
}

func (i *Initialization) VulnerableArea() {
	i.Output.VulnerableArea = 1 + (i.CurrentPanel.Vulnerable)/100
}

func (i *Initialization) SpecialDamageArea() {
	i.Output.SpecialDamageArea = 1 + (i.CurrentPanel.SpecialDamage)/100
}

// --------------------- 数据结构定义 ---------------------

type Initialization struct {
	Magnifications []*Magnification
	CurrentPanel   *CurrentPanel
	Basic          *Basic
	Gain           *Gain
	Defense        *Defense
	Condition      *Condition
	Output         *Output
	// 用于记录暴击转换为爆伤的词条数量
	CritConverted int
	// 记录穿透和增伤流转的词条数量（可在整体优化中重新分配）
	ExtraPenetrateTokens       int
	ExtraIncreasedDamageTokens int
}

type Basic struct {
	BasicAttack              float64 // 基础攻击力（角色+专武）
	BasicCritical            float64 // 基础暴击（角色+武器+2件套+4号位）
	BasicExplosiveInjury     float64 // 基础爆伤（角色+武器+2件套+4号位）
	BasicIncreasedDamage     float64 // 基础增伤（角色+武器+驱动盘）
	BasicReductionResistance float64 // 基础减抗（角色+武器+驱动盘）
	BasicVulnerable          float64 // 基础易伤（百分比）
	BasicSpecialDamage       float64 // 基础特殊增伤（百分比）
	Penetration              float64 // 穿透率（百分比）
}

type CurrentPanel struct {
	Attack              float64 // 攻击力（角色+专武）
	Critical            float64 // 暴击（角色+武器+2件套+4号位）
	ExplosiveInjury     float64 // 爆伤（角色+武器+2件套+4号位）
	IncreasedDamage     float64 // 增伤（角色+武器+驱动盘）
	ReductionResistance float64 // 减抗（角色+武器+驱动盘）
	Vulnerable          float64 // 易伤（百分比）
	SpecialDamage       float64 // 特殊增伤（百分比）
	Penetration         float64 // 穿透率（百分比）
}

type Magnification struct {
	MagnificationValue  float64 // 倍率值 百分比
	TriggerTimes        float64 // 触发次数
	Name                string  // 伤害名称
	IncreasedDamage     float64 // 指定增伤，使用这个增伤+基础增伤  百分比
	ReductionResistance float64 // 指定减抗，使用这个减抗+基础减抗  百分比
	DefenseBreak        float64 // 指定破防，防御乘区，需要通过这个计算  百分比
	Penetration         float64 // 指定穿透，防御乘区，需要通过这个计算  百分比
	SpecialDamage       float64 // 指定独立的增益倍率区 百分比
	ExplosiveInjury     float64 // 局内爆伤计算
}

type Gain struct {
	AttackValue              float64 // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
	AttackValue2             float64 // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
	AttackPowerPercentage    float64 // 攻击力百分比加成(主词条，副词条，2件套)
	AttackInternalPercentage float64 // 局内攻击力(武器，驱动盘绿字攻击力)
	Critical                 float64 // 增加暴击（角色+武器+2件套+4号位）
	ExplosiveInjury          float64 // 增加爆伤（角色+武器+2件套+4号位）
	IncreasedDamage          float64 // 增伤（百分比）
	ReductionResistance      float64 // 减抗（百分比）
	Vulnerable               float64 // 易伤（百分比）
	SpecialDamage            float64 // 特殊增伤（百分比）
}

type Defense struct {
	Penetration      float64 // 穿透率（百分比）
	DefenseBreak     float64 // 破防百分比（百分比）
	PenetrationValue float64 // 穿透值（固定值）
}

type Condition struct {
	MainArticle    int     // 有效主词条总数
	Critical       float64 // 暴击率最大值,超出，其他给爆伤
	CriticalStatus int     // 已经超出阈值，直接递增+1个词条即可
	CriticalCount  int     // 记录第几个词条溢出暴击率
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

const (
	AttackPercentage = "AttackPercentage"
	Critical         = "Critical"
	ExplosiveInjury  = "ExplosiveInjury"
	IncreasedDamage  = "IncreasedDamage"
	Penetrate        = "Penetrate"
)

type Result struct {
	CurrentPanel         *CurrentPanel
	Output               *Output
	Damage               float64
	PercentageDifference float64
}
