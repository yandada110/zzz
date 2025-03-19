package main

import (
	"fmt"
	"math"
)

// ------------------------ 常量定义 ------------------------

// 属性名称常量
const (
	AttackPercentage = "AttackPercentage" // 攻击力词条
	Critical         = "Critical"         // 暴击词条
	ExplosiveInjury  = "ExplosiveInjury"  // 爆伤词条
	IncreasedDamage  = "IncreasedDamage"  // 增伤词条
	Penetrate        = "Penetrate"        // 穿透词条
)

// 全局参数
const (
	GlobalMainArticle = 55 // 队伍可分配的总词条数
	GlobalCritical    = 97 // 暴击率上限（百分比值）
)

// allowedGroupB 定义允许的 (增伤, 穿透) 原始分配组合
var allowedGroupB = [][2]int{
	{0, 13},
	{3, 10},
	{10, 3},
	{13, 0},
	{0, 0},
	{3, 0},
	{0, 3},
}

// ------------------------ 辅助函数 ------------------------

// effectiveGroupB 根据原始分配计算有效词条值：
// 小于 3 -> 0；3<=count<10 -> 3；10<=count<13 -> 10；count>=13 -> 13
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

// reallocateABC 将候选分配中攻击、暴击、爆伤（a、b、c）超出上限的部分
// 重新分配给这些属性中尚未达到上限的项，保持 a+b+c 不变。
// 参数：distribution 包含各属性分配；attackLimit、critLimit、explLimit 分别为 a、b、c 的允许最大值。
// 返回调整后的分配及 ok 标志（false 表示无法完成重分配）。
func reallocateABC(distribution map[string]int, attackLimit, critLimit, explLimit int) (map[string]int, bool) {
	a := distribution[AttackPercentage]
	b := distribution[Critical]
	c := distribution[ExplosiveInjury]

	surplus := 0
	if a > attackLimit {
		surplus += a - attackLimit
		a = attackLimit
	}
	if b > critLimit {
		surplus += b - critLimit
		b = critLimit
	}
	if c > explLimit {
		surplus += c - explLimit
		c = explLimit
	}

	// 将剩余词条尝试分配到未达上限的属性中
	for surplus > 0 {
		allocated := false
		if a < attackLimit {
			a++
			surplus--
			allocated = true
			if surplus == 0 {
				break
			}
		}
		if b < critLimit {
			b++
			surplus--
			allocated = true
			if surplus == 0 {
				break
			}
		}
		if c < explLimit {
			c++
			surplus--
			allocated = true
			if surplus == 0 {
				break
			}
		}
		if !allocated {
			// 若三个属性均已到上限但仍有剩余，则无法重分配
			return nil, false
		}
	}
	distribution[AttackPercentage] = a
	distribution[Critical] = b
	distribution[ExplosiveInjury] = c
	return distribution, true
}

// ------------------------ 数据结构定义 ------------------------

// Initialization 保存角色（队伍）配置、计算模型及词条分配相关参数。
type Initialization struct {
	// CalculationModels：用于计算伤害的各个模型（可采用不同计算公式）
	CalculationModels []*Initialization
	// MainArticle：可分配的总词条数
	MainArticle int
	// 新增参数：控制攻击、暴击、爆伤三个属性的基础词条上限
	AttackCount          int // 攻击力词条基础上限
	CriticalCount        int // 暴击词条基础上限
	ExplosiveInjuryCount int // 爆伤词条基础上限

	Name string // 队伍或模型名称

	// 以下为角色基础属性及增益数据
	Basic        *Basic
	Gain         *Gain
	Defense      *Defense
	Condition    *Condition
	CurrentPanel *CurrentPanel
	Output       *Output

	// 用于记录暴击转换为爆伤的词条数量（超出 GlobalCritical 限制部分）
	CritConverted int

	// Magnifications：各技能倍率数据列表
	Magnifications []*Magnification
}

// ExternalPanel 表示局外面板数据（经过词条分配后的外部效果）
type ExternalPanel struct {
	Attack          float64
	Critical        float64
	ExplosiveInjury float64
}

// Basic 表示角色的基础数值
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

// CurrentPanel 表示角色经过词条分配后更新的局内面板数据
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

// Magnification 表示技能倍率数据
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

// Gain 表示角色的各项增益数值
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

// Defense 表示防御相关数值
type Defense struct {
	Penetration      float64
	DefenseBreak     float64
	PenetrationValue float64
}

// Condition 表示角色条件（如暴击率上限等）
type Condition struct {
	Critical float64
}

// Output 表示计算后输出的各区加成数据
type Output struct {
	BasicDamageArea         float64
	IncreasedDamageArea     float64
	ExplosiveInjuryArea     float64
	DefenseArea             float64
	ReductionResistanceArea float64
	VulnerableArea          float64
	SpecialDamageArea       float64
}

// ------------------------ 方法实现 ------------------------

// Clone 返回 Initialization 的深拷贝（CalculationModels 等共享）
func (i *Initialization) Clone() *Initialization {
	if i == nil {
		return nil
	}
	return &Initialization{
		MainArticle:          i.MainArticle,
		CalculationModels:    i.CalculationModels,
		AttackCount:          i.AttackCount,
		CriticalCount:        i.CriticalCount,
		ExplosiveInjuryCount: i.ExplosiveInjuryCount,
		Name:                 i.Name,
		Basic:                i.Basic,
		Gain:                 i.Gain,
		Defense:              i.Defense,
		Condition:            i.Condition,
		CurrentPanel:         i.CurrentPanel,
		Output:               i.Output,
		CritConverted:        i.CritConverted,
		Magnifications:       i.Magnifications,
	}
}

// ClonePanel 克隆当前局内面板数据
func (i *Initialization) ClonePanel() *CurrentPanel {
	cp := *i.CurrentPanel
	return &cp
}

// CloneOutput 克隆输出数据
func (i *Initialization) CloneOutput() *Output {
	op := *i.Output
	return &op
}

// ResetCondition 重置动态状态（例如暴击转换计数）
func (i *Initialization) ResetCondition() {
	i.CritConverted = 0
}

// CharacterPanelWithDistribution 根据词条分配方案更新局内面板数据。
// 示例计算：
//   - 攻击力 = 基础攻击 + (AttackPercentage 词条数×3)
//   - 暴击率 = 基础暴击 + (Critical 词条数×2.4)
//   - 爆伤 = 基础爆伤 + (ExplosiveInjury 词条数×4.8)
//
// 注意：实际项目中应调用更多 HandleXXX 方法更新各区数值。
func (i *Initialization) CharacterPanelWithDistribution(distribution map[string]int) *Initialization {
	i.CurrentPanel = &CurrentPanel{}
	i.CurrentPanel.Attack = i.Basic.BasicAttack + float64(distribution[AttackPercentage])*3
	i.CurrentPanel.Critical = i.Basic.BasicCritical + float64(distribution[Critical])*2.4
	i.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + float64(distribution[ExplosiveInjury])*4.8
	// 此处未处理增伤与穿透，实际项目中应调用对应 Handle 函数
	return i
}

// CalculatingTotalDamage 根据当前局内面板和技能倍率计算总伤害
func (i *Initialization) CalculatingTotalDamage() float64 {
	totalDamage := 0.0
	for _, mag := range i.Magnifications {
		// 示例计算：伤害 = 攻击力 × (倍率/100) × 触发次数
		totalDamage += i.CurrentPanel.Attack * mag.MagnificationValue / 100 * mag.TriggerTimes
	}
	return totalDamage
}

// CalculateExternalPanel 根据词条分配计算局外面板数据
func (i *Initialization) CalculateExternalPanel(distribution map[string]int, critConverted int) *ExternalPanel {
	attack := i.Basic.BasicAttack + float64(distribution[AttackPercentage])*3
	critical := i.Basic.BasicCritical + float64(distribution[Critical])*2.4
	explosiveInjury := i.Basic.BasicExplosiveInjury + float64(distribution[ExplosiveInjury])*4.8
	return &ExternalPanel{Attack: attack, Critical: critical, ExplosiveInjury: explosiveInjury}
}

// InitializationArea 更新 Output 各区数值（示例实现）
func (i *Initialization) InitializationArea(magnification *Magnification) {
	// 示例：仅更新基本伤害区域
	i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
	// 实际中应计算增伤、爆伤、穿透等各区加成
}

// ------------------------ FindOptimalDistribution 方法 ------------------------
// FindOptimalDistribution 穷举所有分配方案（a+b+c+d+e = MainArticle，其中
//
//	a: AttackPercentage, b: Critical, c: ExplosiveInjury, d: IncreasedDamage, e: Penetrate），
//
// 计算各方案的总伤害，并返回最佳方案。
// 新增逻辑说明：
// 1. (IncreasedDamage, Penetrate) 的组合必须在 allowedGroupB 列表中。
// 2. 根据穿透和增伤词条的有效值（effectiveGroupB）总和，给攻击、暴击、爆伤三个属性增加奖金：
//   - sumEffB == 0：攻击上限 = AttackCount+13, 暴击上限 = CriticalCount+3, 爆伤上限 = ExplosiveInjuryCount+3
//   - sumEffB == 3：攻击上限 = AttackCount+10（暴击、爆伤无奖金）
//   - sumEffB == 10：攻击上限 = AttackCount+3, 暴击上限 = CriticalCount+3, 爆伤上限 = ExplosiveInjuryCount+3
//   - sumEffB == 13：各上限均为基础值（AttackCount, CriticalCount, ExplosiveInjuryCount）
//     3. 若 a、b、c 超出上限，则调用 reallocateABC 进行重分配（其中暴击分配额外不超过 GlobalCritical 允许的上限）。
//     这里先计算 maxCritTokens = floor((GlobalCritical - 基础暴击)/2.4)，
//     并令 effectiveCriticalLimit = min(根据 bonus 得到的值, maxCritTokens)。
//     4. 返回最佳方案及其对应的总伤害、局内面板、输出数据以及暴击转换计数。
func (i *Initialization) FindOptimalDistribution() (bestDamage float64, bestDistribution map[string]int, bestPanel *CurrentPanel, bestOutput *Output, bestCritConverted int) {
	totalTokens := i.MainArticle
	bestDamage = -1.0
	bestDistribution = make(map[string]int)
	bestCritConverted = 0

	// 穷举 a+b+c+d+e = totalTokens
	for a := 0; a <= totalTokens; a++ {
		for b := 0; a+b <= totalTokens; b++ {
			for c := 0; a+b+c <= totalTokens; c++ {
				for d := 0; a+b+c+d <= totalTokens; d++ {
					e := totalTokens - (a + b + c + d)
					// 构造候选方案
					candidate := map[string]int{
						AttackPercentage: a,
						Critical:         b,
						ExplosiveInjury:  c,
						IncreasedDamage:  d,
						Penetrate:        e,
					}
					// 检查 (IncreasedDamage, Penetrate) 是否在允许列表内
					allowedPair := false
					for _, pair := range allowedGroupB {
						if candidate[IncreasedDamage] == pair[0] && candidate[Penetrate] == pair[1] {
							allowedPair = true
							break
						}
					}
					if !allowedPair {
						continue
					}

					// 计算穿透和增伤的有效词条数
					effD := effectiveGroupB(candidate[IncreasedDamage])
					effE := effectiveGroupB(candidate[Penetrate])
					sumEffB := effD + effE

					// 根据 sumEffB 确定攻击、暴击、爆伤的有效上限（基础上限加奖金）
					var effectiveAttackLimit, effectiveCriticalLimit, effectiveExplosiveLimit int
					switch sumEffB {
					case 0:
						effectiveAttackLimit = i.AttackCount + 13
						effectiveCriticalLimit = i.CriticalCount + 3
						effectiveExplosiveLimit = i.ExplosiveInjuryCount + 3
					case 3:
						effectiveAttackLimit = i.AttackCount + 10
						effectiveCriticalLimit = i.CriticalCount
						effectiveExplosiveLimit = i.ExplosiveInjuryCount
					case 10:
						effectiveAttackLimit = i.AttackCount + 3
						effectiveCriticalLimit = i.CriticalCount + 3
						effectiveExplosiveLimit = i.ExplosiveInjuryCount + 3
					case 13:
						effectiveAttackLimit = i.AttackCount
						effectiveCriticalLimit = i.CriticalCount
						effectiveExplosiveLimit = i.ExplosiveInjuryCount
					default:
						effectiveAttackLimit = i.AttackCount
						effectiveCriticalLimit = i.CriticalCount
						effectiveExplosiveLimit = i.ExplosiveInjuryCount
					}
					// 限制暴击分配不能超过 GlobalCritical 限制计算的最大 tokens
					maxCritTokens := int(math.Floor((GlobalCritical - i.Basic.BasicCritical) / 2.4))
					if effectiveCriticalLimit > maxCritTokens {
						effectiveCriticalLimit = maxCritTokens
					}

					// 对 a、b、c 超出上限的部分进行重分配
					adjusted, ok := reallocateABC(candidate, effectiveAttackLimit, effectiveCriticalLimit, effectiveExplosiveLimit)
					if !ok {
						continue
					}
					candidate[AttackPercentage] = adjusted[AttackPercentage]
					candidate[Critical] = adjusted[Critical]
					candidate[ExplosiveInjury] = adjusted[ExplosiveInjury]

					// 遍历每个计算模型，累计总伤害
					var damage float64 = 0.0
					var simLast *Initialization
					for _, model := range i.CalculationModels {
						if model == nil {
							continue
						}
						sim := model.Clone()
						sim.ResetCondition()
						sim.CharacterPanelWithDistribution(candidate)
						damage += sim.CalculatingTotalDamage()
						simLast = sim
					}
					if damage < 0 {
						continue
					}
					// 记录最佳方案
					if damage > bestDamage {
						bestDamage = damage
						bestDistribution = make(map[string]int)
						for k, v := range candidate {
							bestDistribution[k] = v
						}
						bestPanel = simLast.ClonePanel()
						bestOutput = simLast.CloneOutput()
						bestCritConverted = simLast.CritConverted
					}
				}
			}
		}
	}
	return bestDamage, bestDistribution, bestPanel, bestOutput, bestCritConverted
}

// ------------------------ main 函数 ------------------------

func main() {
	// 初始化多个队伍配置
	teams := []*Initialization{
		安比21扳机21嘉音61(),
		安比21扳机21凯撒21(),
	}

	// 针对每个队伍进行计算
	for idx, initialization := range teams {
		// 如果 CalculationModels 为空，则使用自身作为计算模型
		if len(initialization.CalculationModels) == 0 {
			initialization.CalculationModels = []*Initialization{initialization}
		}
		_, bestDist, _, _, bestCritConv := initialization.FindOptimalDistribution()

		// 输出整体方案
		fmt.Printf("====== 队伍组合 %d: %s ======\n", idx+1, initialization.Name)
		fmt.Println("【整体最佳方案】")
		fmt.Println("最佳词条分配方案:")
		fmt.Printf("  攻击力词条: %d, 暴击词条: %d, 爆伤词条: %d, 增伤词条: %d, 穿透词条: %d, 暴击转换爆伤词条: %d\n",
			bestDist[AttackPercentage],
			bestDist[Critical],
			bestDist[ExplosiveInjury],
			bestDist[IncreasedDamage],
			bestDist[Penetrate],
			bestCritConv,
		)
		fmt.Println("--------------------------------------------------")
		// 输出各计算模型的局内、局外面板及【单模型】技能伤害明细
		for _, model := range initialization.CalculationModels {
			// 更新当前模型的局内面板
			model.CharacterPanelWithDistribution(bestDist)
			internalPanel := model.CurrentPanel
			externalPanel := model.CalculateExternalPanel(bestDist, model.CritConverted)
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
			// 输出各技能最终伤害明细
			fmt.Println("各技能最终伤害:")
			totalModelSkillDamage := 0.0
			// 注意：每次调用 InitializationArea 会更新 Output，因此建议对每个技能分别调用
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
