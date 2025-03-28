package main

import (
	"fmt"
	"math"
)

// ------------------------ 常量定义 ------------------------
const (
	AttackPercentage = "AttackPercentage"
	Critical         = "Critical"
	ExplosiveInjury  = "ExplosiveInjury"
	IncreasedDamage  = "IncreasedDamage"
	Penetrate        = "Penetrate"
	Proficient       = "BasicProficient" // 精通

	// 每个队伍可分配的词条数（示例值）
	GlobalMainArticle = 45
	// 暴击上限（示例值）异常角色，暴击率基本不给分配
	GlobalCritical         = 95
	DamageTypeDirectInjury = 1             // 直伤
	DamageTypeAbnormal     = 2             // 异常
	Physics                = "Physics"     // 物理
	fire                   = "fire"        // 火
	electricity            = "electricity" // 电
	ice                    = "ice"         // 冰
	ether                  = "ether"       // 以太
)

var AbnormalMagnificationValue = map[string]float64{
	"Physics":     713,
	"fire":        50,
	"electricity": 125,
	"ice":         500,
	"ether":       62.5,
}

// allowedGroupB 定义允许的增伤、穿透原始分配组合
var allowedGroupB = [][2]int{
	{0, 13},
	{3, 10},
	{10, 3},
	{13, 0},
	{0, 0},
	{3, 0},
	{0, 3},
}

// ------------------------ 主函数 ------------------------
func main() {
	// 初始化各套队伍（示例，具体初始化函数需自行实现）
	initializations := []*Initialization{
		柳柏妮思00露西65站场(),
		柳柏妮思00露西65失衡(),
	}

	// 针对每套队伍进行计算
	for idx, initialization := range initializations {
		fmt.Printf("====== 队伍组合 %d: %s ======\n", idx+1, initialization.Name)
		_, bestDistribution, _, _, bestCritConverted, _ := initialization.FindOptimalDistribution()

		// 输出整体最佳方案
		fmt.Println("【整体最佳方案】")
		fmt.Println("最佳词条分配方案:")
		fmt.Printf("  攻击力词条: %d, 暴击词条: %d, 爆伤词条: %d, 增伤词条: %d, 穿透词条: %d, 暴击转换爆伤词条: %d\n",
			bestDistribution[AttackPercentage],
			bestDistribution[Critical],
			bestDistribution[ExplosiveInjury],
			bestDistribution[IncreasedDamage],
			bestDistribution[Penetrate],
			bestCritConverted,
		)
		//fmt.Printf("最高总伤害: %.6f\n", bestDamage)
		//fmt.Println("最佳局内面板（综合最优方案）:")
		//fmt.Printf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%\n",
		//	bestPanel.Attack,
		//	bestPanel.Critical,
		//	bestPanel.ExplosiveInjury,
		//	bestPanel.IncreasedDamage,
		//	bestPanel.Penetration,
		//)
		//fmt.Println("==================================================")
		// 后续各模型、技能伤害等输出保持不变…
		//fmt.Printf("最高总伤害: %.6f\n", bestDamage)
		//fmt.Println("最佳局内面板（综合最优方案）:")
		//fmt.Printf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%\n",
		//	bestPanel.Attack,
		//	bestPanel.Critical,
		//	bestPanel.ExplosiveInjury,
		//	bestPanel.IncreasedDamage,
		//	bestPanel.Penetration,
		//)
		fmt.Println("--------------------------------------------------")
		// 输出各计算模型（模型可能代表不同计算方法）的局内、局外面板及【单模型】的技能伤害明细
		for _, model := range initialization.CalculationModels {
			// 更新当前模型的局内面板
			model.CharacterPanelWithDistribution(bestDistribution)
			internalPanel := model.CurrentPanel
			externalPanel := model.CalculateExternalPanel(bestDistribution, model.CritConverted)
			fmt.Println("模型: " + model.Name)
			fmt.Println("局内面板:")
			fmt.Printf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%\n",
				internalPanel.Attack,
				internalPanel.Critical,
				internalPanel.ExplosiveInjury,
				internalPanel.IncreasedDamage,
				internalPanel.Penetration,
			)
			fmt.Println("局外面板:")
			fmt.Printf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%\n",
				externalPanel.Attack,
				externalPanel.Critical,
				externalPanel.ExplosiveInjury,
			)
			// ---------------- 新增部分 ----------------
			// 根据当前模型的最终参数，计算并输出各技能的伤害明细
			fmt.Println("各技能最终伤害:")
			totalModelSkillDamage := 0.0
			// 注意：由于每次调用 InitializationArea 会更新 Output，因此建议对每个技能分别调用
			for _, mag := range model.Magnifications {
				model.InitializationArea(mag)
				skillDmg := model.Output.BasicDamageArea *
					model.Output.IncreasedDamageArea *
					model.Output.ExplosiveInjuryArea *
					model.Output.DefenseArea *
					model.Output.ReductionResistanceArea *
					model.Output.VulnerableArea *
					model.Output.SpecialDamageArea *
					(1 + mag.SpecialDamage/100)
				fmt.Printf("  技能 [%s] 伤害: %.6f\n", mag.Name, skillDmg)
				totalModelSkillDamage += skillDmg
			}
			fmt.Printf("  技能总伤害: %.6f\n", totalModelSkillDamage)
			fmt.Println("--------------------------------------------------")
		}
	}
}

// ------------------------ 辅助函数 ------------------------

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func copyMap(m map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range m {
		res[k] = v
	}
	return res
}

// effectiveGroupB 根据原始分配计算有效值：
// count < 3  -> 0, 3<=count<10 -> 3, 10<=count<13 -> 10, count>=13 -> 13
func effectiveGroupB(count int) int {
	if count < 3 {
		return 0
	} else if count < 10 {
		return 3
	} else if count < 13 {
		return 10
	} else {
		return 13
	}
}

// ------------------------ 修改后的 FindOptimalDistribution ------------------------
// 枚举所有 5 项分配方案（总和 = GlobalMainArticle），
// 仅接受满足以下条件的方案：
// 1. (IncreasedDamage, Penetrate) 必须属于 allowedGroupB 列表
// 2. AttackPercentage + effectiveGroupB(IncreasedDamage) + effectiveGroupB(Penetrate) >= 13
func (i *Initialization) FindOptimalDistribution() (bestDamage float64, bestDistribution map[string]int, bestPanel *CurrentPanel, bestOutput *Output, bestCritConverted int, skillDamages map[string]float64) {
	distributions := generateDistributions(i.MainArticle, 5)
	bestDamage = -1.0
	bestDistribution = make(map[string]int)
	bestCritConverted = 0
	var bestSim *Initialization

	for _, dist := range distributions {
		distribution := map[string]int{
			AttackPercentage: dist[0],
			Critical:         dist[1],
			ExplosiveInjury:  dist[2],
			IncreasedDamage:  dist[3],
			Penetrate:        dist[4],
			Proficient:       dist[5],
		}
		// 检查 (IncreasedDamage, Penetrate) 是否在允许范围内
		allowed := false
		for _, pair := range allowedGroupB {
			if distribution[IncreasedDamage] == pair[0] && distribution[Penetrate] == pair[1] {
				allowed = true
				break
			}
		}
		if !allowed {
			continue
		}
		// 计算攻击力词条和增伤、穿透的有效值之和
		totalEffective := distribution[AttackPercentage] + effectiveGroupB(distribution[IncreasedDamage]) + effectiveGroupB(distribution[Penetrate])
		if totalEffective < 13 {
			// 不满足至少13词条的要求，跳过
			continue
		}

		var damage float64 = 0.0
		var lastSim *Initialization
		for _, model := range i.CalculationModels {
			sim := model.Clone()
			sim.ResetCondition()
			sim.CharacterPanelWithDistribution(distribution)
			damage += sim.CalculatingTotalDamage()
			lastSim = sim
		}
		if damage > bestDamage {
			bestDamage = damage
			bestDistribution = copyMap(distribution)
			bestPanel = lastSim.ClonePanel()
			bestOutput = lastSim.CloneOutput()
			bestCritConverted = lastSim.CritConverted
			bestSim = lastSim
		}
	}

	// 在最佳模拟器下，计算各技能伤害（整体最佳方案）
	bestSimClone := bestSim.Clone()
	bestSimClone.CharacterPanelWithDistribution(bestDistribution)
	skillDamages = make(map[string]float64)
	for _, mag := range bestSimClone.Magnifications {
		bestSimClone.InitializationArea(mag)
		skillDamage := bestSimClone.Output.BasicDamageArea *
			bestSimClone.Output.IncreasedDamageArea *
			bestSimClone.Output.ExplosiveInjuryArea *
			bestSimClone.Output.DefenseArea *
			bestSimClone.Output.ReductionResistanceArea *
			bestSimClone.Output.VulnerableArea *
			bestSimClone.Output.SpecialDamageArea *
			(1 + mag.SpecialDamage/100)
		skillDamages[mag.Name] = skillDamage
	}
	return bestDamage, bestDistribution, bestPanel, bestOutput, bestCritConverted, skillDamages
}

// ------------------------ generateDistributions ------------------------
// generateDistributions 递归生成将 total 个词条分配到 slots 个属性上的所有方案（和等于 total）
func generateDistributions(total, slots int) [][]int {
	var results [][]int
	var helper func(remaining, slots int, current []int)
	helper = func(remaining, slots int, current []int) {
		if slots == 1 {
			newDist := append([]int{}, current...)
			newDist = append(newDist, remaining)
			results = append(results, newDist)
			return
		}
		for i := 0; i <= remaining; i++ {
			newCurrent := append([]int{}, current...)
			newCurrent = append(newCurrent, i)
			helper(remaining-i, slots-1, newCurrent)
		}
	}
	helper(total, slots, []int{})
	return results
}

// ------------------------ 以下各函数保持不变 ------------------------

// CalculateExternalPanel 根据当前模型和词条分配计算局外面板
// 公式：
//
//	攻击力 = BasicAttack * (1 + (AttackPowerPercentage + 攻击力词条数*3)/100) + AttackValue
//	暴击率 = BasicCritical + 暴击词条数*2.4
//	爆伤   = BasicExplosiveInjury + (爆伤词条数 + 暴击转换爆伤词条数量)*4.8
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
	i.HandleBasicProficient(Proficient, distribution[Proficient])
	return i
}

// ResetCondition 重置暴击转换计数
func (i *Initialization) ResetCondition() {
	i.CritConverted = 0
}

// Clone 克隆 Initialization 对象（CalculationModels 共享）
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
		Output:            &Output{},
		CurrentPanel:      &CurrentPanel{},
		CritConverted:     i.CritConverted,
		Name:              i.Name,
		CalculationModels: i.CalculationModels,
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

// ===== 以下各函数处理词条加成 =====

// HandleBasicAttack 根据攻击力词条增加攻击力
func (i *Initialization) HandleBasicAttack(key string, count int) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage
	if key == AttackPercentage {
		attackPowerPercentage += 3 * float64(count)
	}
	i.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100) + i.Gain.AttackValue + i.Gain.AttackValue2) * (1 + i.Gain.AttackInternalPercentage/100)
}

// HandleBasicProficient 根据精通词条增加精通
func (i *Initialization) HandleBasicProficient(key string, count int) {
	attackPowerPercentage := i.Gain.Proficient
	if key == Proficient {
		attackPowerPercentage += 9 * float64(count)
	}
	i.CurrentPanel.Proficient = i.Basic.Proficient + attackPowerPercentage
}

// HandleBasicCritical 根据暴击词条更新暴击率，并计算转换为爆伤的词条数
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

// HandleBasicExplosiveInjury 根据爆伤词条更新爆伤
func (i *Initialization) HandleBasicExplosiveInjury(key string, count int) {
	explosiveInjury := i.Gain.ExplosiveInjury
	if key == ExplosiveInjury {
		explosiveInjury += 4.8 * float64(count)
	}
	convertedBonus := 4.8 * float64(i.CritConverted)
	i.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + explosiveInjury + convertedBonus
}

// HandleBasicIncreasedDamage 根据增伤词条更新增伤
// 有效值只允许取 0, 3, 10, 或 13
func (i *Initialization) HandleBasicIncreasedDamage(key string, count int) {
	if key == IncreasedDamage {
		var effectiveTokens int
		if count >= 13 {
			effectiveTokens = 13
		} else if count >= 10 {
			effectiveTokens = 10
		} else if count >= 3 {
			effectiveTokens = 3
		} else {
			effectiveTokens = 0
		}
		i.CurrentPanel.IncreasedDamage = i.Basic.BasicIncreasedDamage + (i.Gain.IncreasedDamage + 3*float64(effectiveTokens))
	}
}

// HandlePenetrateDamage 根据穿透词条更新穿透率
// 有效值只允许取 0, 3, 10, 或 13
func (i *Initialization) HandlePenetrateDamage(key string, count int) {
	if key == Penetrate {
		var effectiveTokens int
		if count >= 13 {
			effectiveTokens = 13
		} else if count >= 10 {
			effectiveTokens = 10
		} else if count >= 3 {
			effectiveTokens = 3
		} else {
			effectiveTokens = 0
		}
		penetrationValue := i.Defense.Penetration + 2.4*float64(effectiveTokens)
		if penetrationValue >= 100 {
			penetrationValue = 100
		}
		i.CurrentPanel.Penetration = penetrationValue
	}
}

// ===== 以下各函数计算各分区加成 =====

func (i *Initialization) CalculatingTotalDamage() float64 {
	totalDamage := 0.0
	for _, mag := range i.Magnifications {
		i.InitializationArea(mag)
		if mag.DamageType == DamageTypeDirectInjury {
			// 直伤计算
			damage := i.Output.BasicDamageArea *
				i.Output.IncreasedDamageArea *
				i.Output.ExplosiveInjuryArea *
				i.Output.DefenseArea *
				i.Output.ReductionResistanceArea *
				i.Output.VulnerableArea *
				i.Output.SpecialDamageArea *
				(1 + mag.SpecialDamage/100)
			totalDamage += damage
		} else {
			// 异常伤害计算
			damage := i.Output.BasicAbnormalArea *
				i.Output.IncreasedDamageArea *
				i.Output.GradeArea *
				i.Output.DefenseArea *
				i.Output.ReductionResistanceArea *
				i.Output.VulnerableArea *
				i.Output.SpecialDamageArea *
				(1 + mag.SpecialDamage/100)
			totalDamage += damage
		}
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
	i.ProficientArea()
	i.GradeArea()
}

func (i *Initialization) BasicDamageArea(magnification *Magnification) {
	i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
}

func (i *Initialization) BasicAbnormalArea(magnification *Magnification) {
	i.Output.BasicAbnormalArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
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

func (i *Initialization) ProficientArea() {
	i.Output.ProficientArea = i.CurrentPanel.SpecialDamage / 100
}

func (i *Initialization) GradeArea() {
	power := math.Pow(10, float64(4))
	i.Output.GradeArea = math.Trunc(60*power) / power
}

// ------------------------ 数据结构定义 ------------------------

type Initialization struct {
	Magnifications    []*Magnification  // 各技能倍率列表
	CalculationModels []*Initialization // 计算模型集合（可以包含不同模型）
	MainArticle       int               // 有效词条总数
	CritConverted     int               // 记录暴击转换为爆伤的词条数量
	Name              string            // 队伍或模型名称

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
	Proficient               float64 // 精通
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
	Proficient          float64 // 精通
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
	DamageType          int32   // 伤害类型，直伤还是异常伤害
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
	Proficient               float64 // 精通
}

type Defense struct {
	Penetration      float64 // 穿透率（百分比）
	DefenseBreak     float64 // 破防（百分比）
	PenetrationValue float64 // 穿透固定值
}

type Condition struct {
	Critical float64
}

type Output struct {
	BasicDamageArea         float64 // 直伤倍率区
	IncreasedDamageArea     float64 // 增伤区
	ExplosiveInjuryArea     float64 // 双爆区
	DefenseArea             float64 // 防御区
	ReductionResistanceArea float64 // 减抗区
	VulnerableArea          float64 // 易伤区
	SpecialDamageArea       float64 // 特殊乘区
	ProficientArea          float64 // 精通区
	GradeArea               float64 // 等级区
	BasicAbnormalArea       float64 // 异常倍率区
}
