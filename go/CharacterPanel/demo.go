package main

import (
	"fmt"
)

// ------------------------ 常量定义 ------------------------
const (
	AttackPercentage = "AttackPercentage"
	Critical         = "Critical"
	ExplosiveInjury  = "ExplosiveInjury"
	IncreasedDamage  = "IncreasedDamage"
	Penetrate        = "Penetrate"

	// 每个队伍可分配的词条数（示例值）
	GlobalMainArticle = 45
	// 暴击上限（示例值）
	GlobalCritical = 95
)

// allowedGroupB 定义允许的增伤、穿透原始分配组合（保持不变）
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
		安比21扳机21嘉音61(),
		安比21扳机21凯撒21(),
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
		fmt.Println("--------------------------------------------------")
		// 输出各计算模型（模型可能代表不同计算方法）的局内、局外面板及技能伤害明细……
		for _, model := range initialization.CalculationModels {
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
			// 输出各技能伤害（略）
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

// -----------------------------------------------------------------------------
// 新增部分：根据“磁盘”规则验证候选全局词条分配是否可行
// -----------------------------------------------------------------------------
// 假设游戏中共有 6 个磁盘（副词条槽），每个磁盘最多 5 个有效词条
// 每个磁盘允许的副词条类型取决于其主词条（例如：1,3,4,5,6 号盘通常允许【攻击、暴击、穿透】；
// 但特殊磁盘（如2号、4号、5号、6号）会根据主词条变化而改变可选范围）
//
// 此处提供一个简化示例：
// 对于“暴击-爆伤-穿透”情形（例如计算结果要求暴击=14, 爆伤=19, 穿透=13）：
//   - 假设磁盘 1、2、3、5、6 均可提供最多（3 暴击 + 2 爆伤）
//   - 磁盘 4因主词条选择（例如爆伤）则只允许最多 3 暴击（不提供爆伤）
//
// 对于“增伤-攻击”情形（例如计算结果要求增伤=13, 攻击力=13）：
//   - 根据规则，不同盘允许的分布不同，此处示例给出预设上限
//
// 实际实现时，可根据详细规则枚举每个盘所有可能分配，再判断是否存在满足候选全局分配的方案。
// 下面示例仅采用各盘“预设最大贡献值”判断是否可能达到候选全局需求。
//
// scenario 参数：
//
//	"crit_explosive" —— 针对暴击、爆伤、穿透组合的情况
//	"attack_increased" —— 针对增伤、攻击力组合的情况
//
// -----------------------------------------------------------------------------
func ValidateSubstatDistributionWithScenario(distribution map[string]int, scenario string) bool {
	if scenario == "crit_explosive" {
		// 以下示例数值仅为演示，实际数值需根据游戏规则调整
		maxCrit := 0
		maxExplosive := 0
		// 磁盘1：允许暴击和爆伤，预设分配：最多暴击3，爆伤2
		maxCrit += 3
		maxExplosive += 2
		// 磁盘2：允许暴击和爆伤，预设分配同上
		maxCrit += 3
		maxExplosive += 2
		// 磁盘3：允许暴击和爆伤
		maxCrit += 3
		maxExplosive += 2
		// 磁盘4：特殊盘，根据主词条判断
		// 假设本盘主词条选择为【爆伤】，则仅能在副词条中补足暴击（上限3），爆伤不计
		maxCrit += 3
		// 磁盘5：假设主词条为穿透（或增伤），允许分配暴击和爆伤（预设：3和2）
		maxCrit += 3
		maxExplosive += 2
		// 磁盘6：固定主词条为攻击，允许暴击和爆伤（预设：3和2）
		maxCrit += 3
		maxExplosive += 2

		// 获取候选全局分配中要求的副词条总量（注意：这里要求的值可能为经过某种换算的“有效词条”数）
		requiredCrit := distribution[Critical]
		requiredExplosive := distribution[ExplosiveInjury]

		// 若候选要求超过各盘预设上限，则判定该方案不可行
		if requiredCrit > maxCrit || requiredExplosive > maxExplosive {
			return false
		}
		// 若需要更细致的枚举分配，则在此处进行二重循环穷举各盘分配组合，判断是否存在满足要求的方案
		// 为简化示例，这里直接返回 true
		return true
	} else if scenario == "attack_increased" {
		// 针对增伤=13, 攻击力=13的情况，预设各盘上限（示例数据）
		maxAttack := 0
		maxCrit := 0
		// 假设磁盘1、3、5允许分配攻击和暴击，预设分配：攻击2，暴击3
		maxAttack += 2
		maxCrit += 3
		maxAttack += 2
		maxCrit += 3
		maxAttack += 2
		maxCrit += 3
		// 磁盘2：只允许暴击（预设上限3）
		maxCrit += 3
		// 磁盘4：特殊盘，根据主词条，假设主词条为攻击，则允许分配：攻击0，暴击3
		maxAttack += 0
		maxCrit += 3
		// 磁盘6：固定主词条为增伤，允许分配：攻击2，暴击3
		maxAttack += 2
		maxCrit += 3

		requiredAttack := distribution[AttackPercentage]
		requiredCrit := distribution[Critical]

		if requiredAttack > maxAttack || requiredCrit > maxCrit {
			return false
		}
		return true
	}
	// 默认返回 true
	return true
}

// -----------------------------------------------------------------------------
// 修改后的 FindOptimalDistribution：
// 在枚举全局 5 项分配方案后，新增调用验证函数，只有当候选方案满足“磁盘分配规则”时才进入后续计算。
// -----------------------------------------------------------------------------
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
		// 检查 攻击力词条 + 有效增伤、穿透换算后的有效值 >= 13
		totalEffective := distribution[AttackPercentage] + effectiveGroupB(distribution[IncreasedDamage]) + effectiveGroupB(distribution[Penetrate])
		if totalEffective < 13 {
			continue
		}

		// 新增验证：根据新的磁盘分配规则判断该候选全局分配是否可行
		// 若候选方案同时满足“暴击-爆伤-穿透”与“增伤-攻击”两种情形，则可分别验证
		// 这里以一个简单判定：如果增伤为 13 且攻击力为 13，则采用 "attack_increased" 情形，否则使用 "crit_explosive"
		scenario := "crit_explosive"
		if distribution[IncreasedDamage] == 13 && distribution[AttackPercentage] == 13 {
			scenario = "attack_increased"
		}
		if !ValidateSubstatDistributionWithScenario(distribution, scenario) {
			continue
		}

		// 模型计算总伤害（保持原有逻辑）
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
}

type Defense struct {
	Penetration      float64
	DefenseBreak     float64
	PenetrationValue float64
}

type Condition struct {
	Critical float64
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
