package main

import (
	"fmt"
)

// 定义常量词条类型
const (
	AttackPercentage = "AttackPercentage"
	Critical         = "Critical"
	ExplosiveInjury  = "ExplosiveInjury"
	IncreasedDamage  = "IncreasedDamage"
	Penetrate        = "Penetrate"
)

func main() {
	// 初始化角色数据
	initializations := []*Initialization{
		NewInitialization(), NewInitialization1(),
	}
	for _, initialization := range initializations {
		//// 初始化角色数据
		//initialization := NewInitialization()

		// 穷举所有词条分配方案，得到最佳方案
		bestDamage, bestDistribution, bestPanel, _, bestCritConverted := initialization.FindOptimalDistribution()

		// 输出最佳分配方案和伤害
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
		fmt.Println("最佳局内面板（来自综合最优方案）:")
		fmt.Println(fmt.Sprintf(" 攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%",
			bestPanel.Attack,
			bestPanel.Critical,
			bestPanel.ExplosiveInjury,
			bestPanel.IncreasedDamage,
			bestPanel.Penetration,
		))
		fmt.Println("--------------------------------------------------")
		// 遍历 CalculationModels，输出每个模型对应的局内和局外面板信息
		for _, model := range initialization.CalculationModels {
			// 使用最佳分配方案构建面板
			model.CharacterPanelWithDistribution(bestDistribution)
			internalPanel := model.CurrentPanel
			externalPanel := model.CalculateExternalPanel(bestDistribution, model.CritConverted)
			fmt.Println("模型: " + model.Name)
			fmt.Println("局内面板:")
			fmt.Println(fmt.Sprintf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%",
				internalPanel.Attack,
				internalPanel.Critical,
				internalPanel.ExplosiveInjury,
				internalPanel.IncreasedDamage,
				internalPanel.Penetration,
			))
			fmt.Println("局外面板:")
			fmt.Println(fmt.Sprintf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%",
				externalPanel.Attack,
				externalPanel.Critical,
				externalPanel.ExplosiveInjury,
			))
			fmt.Println("--------------------------------------------------")
		}

	}
}

// CalculateExternalPanel 根据当前模型和词条分配计算局外面板
// 局外面板公式：
//
//	攻击力 = Basic.BasicAttack * (1 + (Gain.AttackPowerPercentage + 攻击力词条数*3)/100) + Gain.AttackValue
//	暴击率 = Basic.BasicCritical + 暴击词条数*2.4
//	爆伤   = Basic.BasicExplosiveInjury + (爆伤词条数 + 暴击转换爆伤词条数量)*4.8
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

// FindOptimalDistribution 穷举所有词条分配方案，返回最佳方案对应的伤害、分配方案、局内面板及暴击转换数
func (i *Initialization) FindOptimalDistribution() (bestDamage float64, bestDistribution map[string]int, bestPanel *CurrentPanel, bestOutput *Output, bestCritConverted int) {
	totalTokens := i.MainArticle
	bestDamage = -1.0
	bestDistribution = make(map[string]int)
	bestCritConverted = 0

	// 穷举所有组合：a+b+c+d+e = totalTokens（5个词条类型）
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
					var damage float64
					var lastSim *Initialization
					// 对所有模型分别计算伤害总和
					for _, model := range i.CalculationModels {
						sim := model.Clone()
						sim.ResetCondition()
						sim.CharacterPanelWithDistribution(distribution)
						damage += sim.CalculatingTotalDamage()
						lastSim = sim
					}
					if damage > bestDamage {
						bestDamage = damage
						bestDistribution = make(map[string]int)
						for k, v := range distribution {
							bestDistribution[k] = v
						}
						bestPanel = lastSim.ClonePanel()
						bestOutput = lastSim.CloneOutput()
						bestCritConverted = lastSim.CritConverted
					}
				}
			}
		}
	}
	return bestDamage, bestDistribution, bestPanel, bestOutput, bestCritConverted
}

// CharacterPanelWithDistribution 根据词条分配更新局内面板
func (i *Initialization) CharacterPanelWithDistribution(distribution map[string]int) *Initialization {
	i.CurrentPanel = &CurrentPanel{
		ReductionResistance: i.Basic.BasicReductionResistance + i.Gain.ReductionResistance,
		Vulnerable:          i.Basic.BasicVulnerable + i.Gain.Vulnerable,
		SpecialDamage:       i.Basic.BasicSpecialDamage + i.Gain.SpecialDamage,
	}
	i.HandleBasicAttack(AttackPercentage, distribution[AttackPercentage])
	i.HandleBasicCritical(Critical, distribution[Critical])
	i.HandleBasicExplosiveInjury(ExplosiveInjury, distribution[ExplosiveInjury])
	i.HandleBasicIncreasedDamage(IncreasedDamage, distribution[IncreasedDamage])
	i.HandlePenetrateDamage(Penetrate, distribution[Penetrate])
	return i
}

// ResetCondition 重置暴击转换计数
func (i *Initialization) ResetCondition() {
	i.CritConverted = 0
}

// Clone 克隆 Initialization 对象（共享 CalculationModels）
func (i *Initialization) Clone() *Initialization {
	return &Initialization{
		MainArticle:    i.MainArticle,
		Magnifications: i.Magnifications,
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
			Critical: i.Condition.Critical,
		},
		Output:        &Output{},
		CurrentPanel:  &CurrentPanel{},
		CritConverted: i.CritConverted,
		Name:          i.Name,
	}
}

// ClonePanel 克隆当前局内面板
func (i *Initialization) ClonePanel() *CurrentPanel {
	cp := *i.CurrentPanel
	return &cp
}

// CloneOutput 克隆当前输出
func (i *Initialization) CloneOutput() *Output {
	op := *i.Output
	return &op
}

// 以下各函数处理词条加成效果

func (i *Initialization) HandleBasicAttack(key string, count int) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage
	if key == AttackPercentage {
		attackPowerPercentage += 3 * float64(count)
	}
	i.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100) + i.Gain.AttackValue + i.Gain.AttackValue2) * (1 + i.Gain.AttackInternalPercentage/100)
}

func (i *Initialization) HandleBasicCritical(key string, count int) {
	if key == Critical {
		baseCrit := i.Basic.BasicCritical + i.Gain.Critical
		maxCritTokens := int((i.Condition.Critical - baseCrit) / 2.4)
		if maxCritTokens < 0 {
			maxCritTokens = 0
		}
		if count <= maxCritTokens {
			i.CurrentPanel.Critical = baseCrit + 2.4*float64(count)
			i.CritConverted = 0
		} else {
			i.CurrentPanel.Critical = i.Condition.Critical
			overflowTokens := count - maxCritTokens
			i.CritConverted = overflowTokens
		}
	} else {
		i.CurrentPanel.Critical = i.Basic.BasicCritical + i.Gain.Critical
	}
}

func (i *Initialization) HandleBasicExplosiveInjury(key string, count int) {
	explosiveInjury := i.Gain.ExplosiveInjury
	if key == ExplosiveInjury {
		explosiveInjury += 4.8 * float64(count)
	}
	convertedBonus := 4.8 * float64(i.CritConverted)
	i.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + explosiveInjury + convertedBonus
}

func (i *Initialization) HandleBasicIncreasedDamage(key string, count int) {
	if key == IncreasedDamage {
		var effectiveTokens int
		if count < 3 {
			effectiveTokens = 0
		} else if count < 10 {
			effectiveTokens = 3
		} else if count == 10 {
			effectiveTokens = 10
		} else if count < 13 {
			effectiveTokens = 10
		} else {
			effectiveTokens = 13
		}
		i.CurrentPanel.IncreasedDamage = i.Basic.BasicIncreasedDamage + (i.Gain.IncreasedDamage + 3*float64(effectiveTokens))
	}
}

func (i *Initialization) HandlePenetrateDamage(key string, count int) {
	if key == Penetrate {
		var effectiveTokens int
		if count < 3 {
			effectiveTokens = 0
		} else if count < 10 {
			effectiveTokens = 3
		} else if count == 10 {
			effectiveTokens = 10
		} else if count < 13 {
			effectiveTokens = 10
		} else {
			effectiveTokens = 13
		}
		penetrationValue := i.Defense.Penetration + 2.4*float64(effectiveTokens)
		if penetrationValue >= 100 {
			penetrationValue = 100
		}
		i.CurrentPanel.Penetration = penetrationValue
	}
}

// CalculatingTotalDamage 遍历所有倍率区，计算总伤害
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
	}
	return totalDamage
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
	// 存放多套模型
	Magnifications    []*Magnification  // 倍率列表
	CalculationModels []*Initialization // 计算模型集合（每套面板）
	MainArticle       int               // 有效主词条总数
	CritConverted     int               // 记录暴击转换为爆伤的词条数量
	Name              string            // 模型名称

	Basic        *Basic
	Gain         *Gain
	Defense      *Defense
	Condition    *Condition
	CurrentPanel *CurrentPanel
	Output       *Output
}

type ExternalPanel struct {
	Attack          float64
	Critical        float64
	ExplosiveInjury float64
}

type Basic struct {
	BasicAttack              float64 // 基础攻击力
	BasicCritical            float64 // 基础暴击
	BasicExplosiveInjury     float64 // 基础爆伤
	BasicIncreasedDamage     float64 // 基础增伤
	BasicReductionResistance float64 // 基础减抗
	BasicVulnerable          float64 // 基础易伤（百分比）
	BasicSpecialDamage       float64 // 基础特殊增伤（百分比）
	Penetration              float64 // 穿透率（百分比）
}

type CurrentPanel struct {
	Attack              float64 // 攻击力
	Critical            float64 // 暴击率（百分比）
	ExplosiveInjury     float64 // 爆伤（百分比）
	IncreasedDamage     float64 // 增伤（百分比）
	ReductionResistance float64 // 减抗（百分比）
	Vulnerable          float64 // 易伤（百分比）
	SpecialDamage       float64 // 特殊增伤（百分比）
	Penetration         float64 // 穿透率（百分比）
}

type Magnification struct {
	MagnificationValue  float64 // 倍率值（百分比）
	TriggerTimes        float64 // 触发次数
	Name                string  // 伤害名称
	IncreasedDamage     float64 // 指定增伤（百分比）
	ReductionResistance float64 // 指定减抗（百分比）
	DefenseBreak        float64 // 指定破防（百分比）
	Penetration         float64 // 指定穿透（百分比）
	SpecialDamage       float64 // 指定特殊增益（百分比）
	ExplosiveInjury     float64 // 局内爆伤计算（百分比）
}

type Gain struct {
	AttackValue              float64 // 攻击力增加值
	AttackValue2             float64 // 攻击力增加值2
	AttackPowerPercentage    float64 // 攻击力百分比加成
	AttackInternalPercentage float64 // 局内攻击力加成
	Critical                 float64 // 增加暴击
	ExplosiveInjury          float64 // 增加爆伤
	IncreasedDamage          float64 // 增伤（百分比）
	ReductionResistance      float64 // 减抗（百分比）
	Vulnerable               float64 // 易伤（百分比）
	SpecialDamage            float64 // 特殊增伤（百分比）
}

type Defense struct {
	Penetration      float64 // 穿透率（百分比）
	DefenseBreak     float64 // 破防（百分比）
	PenetrationValue float64 // 穿透固定值
}

// Condition 只保留真正使用的字段
type Condition struct {
	Critical float64 // 暴击率上限
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

type Result struct {
	CurrentPanel         *CurrentPanel
	Output               *Output
	Damage               float64
	PercentageDifference float64
}
