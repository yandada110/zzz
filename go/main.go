package main

import (
	"fmt"
	"math"
)

func main() {
	// 初始化角色数据
	initialization := NewInitialization()
	// 通过遍历所有可能分配方案寻找最佳词条分配
	bestDamage, bestDistribution, bestPanel, bestOutput := initialization.FindOptimalDistribution()

	// 输出最佳方案及对应面板和伤害信息
	fmt.Println("最佳词条分配方案:")
	fmt.Println(fmt.Sprintf(" 攻击力: %d, 暴击: %d, 爆伤: %d, 增伤: %d, 穿透: %d",
		bestDistribution[AttackPercentage],
		bestDistribution[Critical],
		bestDistribution[ExplosiveInjury],
		bestDistribution[IncreasedDamage],
		bestDistribution[Penetrate],
	))
	fmt.Println(fmt.Sprintf("最高总伤害: %.6f", bestDamage))
	fmt.Println("最佳面板:")
	fmt.Println(fmt.Sprintf(" 攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%",
		bestPanel.Attack,
		bestPanel.Critical,
		bestPanel.ExplosiveInjury,
		bestPanel.IncreasedDamage,
		bestPanel.Penetration,
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

// FindOptimalDistribution 遍历所有可能的词条分配，找出最佳方案
func (i *Initialization) FindOptimalDistribution() (bestDamage float64, bestDistribution map[string]int, bestPanel *CurrentPanel, bestOutput *Output) {
	totalTokens := i.Condition.MainArticle
	bestDamage = -1.0
	bestDistribution = make(map[string]int)
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
					// 克隆一个新的初始化（保证各项状态从初始值开始）
					sim := i.Clone()
					sim.ResetCondition() // 重置暴击溢出等计数
					// 根据当前分配方案构建面板
					sim.CharacterPanelWithDistribution(distribution)
					damage := sim.CalculatingTotalDamage()
					if damage > bestDamage {
						bestDamage = damage
						bestDistribution = distribution
						bestPanel = sim.ClonePanel()
						bestOutput = sim.CloneOutput()
					}
				}
			}
		}
	}
	return bestDamage, bestDistribution, bestPanel, bestOutput
}

// CharacterPanelWithDistribution 根据完整的词条分配更新面板
func (i *Initialization) CharacterPanelWithDistribution(distribution map[string]int) *Initialization {
	i.CurrentPanel = &CurrentPanel{
		ReductionResistance: i.Basic.BasicReductionResistance + i.Gain.ReductionResistance,
		Vulnerable:          i.Basic.BasicVulnerable + i.Gain.Vulnerable,
		SpecialDamage:       i.Basic.BasicSpecialDamage + i.Gain.SpecialDamage,
	}
	// 调用各属性处理函数（顺序可根据属性依赖关系调整）
	i.HandleBasicAttack(AttackPercentage, distribution[AttackPercentage])
	i.HandleBasicExplosiveInjury(ExplosiveInjury, distribution[ExplosiveInjury])
	i.HandleBasicCritical(Critical, distribution[Critical])
	i.HandleBasicIncreasedDamage(IncreasedDamage, distribution[IncreasedDamage])
	i.HandlePenetrateDamage(Penetrate, distribution[Penetrate])
	return i
}

// ResetCondition 重置 Condition 中的计数状态
func (i *Initialization) ResetCondition() {
	i.Condition.CriticalStatus = 0
	i.Condition.CriticalCount = 0
	i.CritConverted = 0
}

// Clone 克隆 Initialization 对象（部分字段采用深拷贝，假设 Magnifications 为常量）
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
		CritConverted: i.CritConverted,
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

// 以下函数保持原有处理逻辑，仅做少量注释调整

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
			// 无溢出
			i.CritConverted = 0
		} else {
			// 超出部分转换为额外爆伤
			i.CurrentPanel.Critical = i.Condition.Critical
			overflowTokens := count - maxCritTokens
			i.CritConverted = overflowTokens
			// 基础爆伤（不含溢出转换）由 HandleBasicExplosiveInjury 处理，此处只累加转换部分的额外爆伤
			baseExplosive := i.Basic.BasicExplosiveInjury + i.Gain.ExplosiveInjury
			i.CurrentPanel.ExplosiveInjury = baseExplosive + 4.8*float64(overflowTokens)
			fmt.Println(fmt.Sprintf("暴击溢出处理: 基础暴击=%.2f%%, 阈值=%.2f%%, 可用词条数=%d, 溢出词条数=%d, 最终暴击率=%.2f%%, 爆伤增加后=%.2f%%",
				baseCrit, i.Condition.Critical, maxCritTokens, overflowTokens, i.CurrentPanel.Critical, i.CurrentPanel.ExplosiveInjury))
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
	i.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + explosiveInjury
}

func (i *Initialization) HandleBasicIncreasedDamage(key string, count int) {
	increasedDamage := i.Gain.IncreasedDamage
	if key == IncreasedDamage {
		fmt.Println(fmt.Sprintf("词条使用数量：增伤：%d", count))
		increasedDamage += 3 * float64(count)
	}
	i.CurrentPanel.IncreasedDamage = i.Basic.BasicIncreasedDamage + increasedDamage
}

func (i *Initialization) HandlePenetrateDamage(key string, count int) {
	increasedDamage := i.Defense.Penetration
	if key == Penetrate {
		fmt.Println(fmt.Sprintf("词条使用数量：穿透：%d", count))
		increasedDamage += 2.4 * float64(count)
		if increasedDamage >= 100 {
			increasedDamage = 100
		}
	}
	i.CurrentPanel.Penetration = increasedDamage
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
	MainArticle    int
	Critical       float64
	CriticalStatus int
	CriticalCount  int
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

// NewInitialization2 构造函数，根据实际需求初始化各字段
func NewInitialization2() *Initialization {
	return &Initialization{
		Magnifications: []*Magnification{
			{
				MagnificationValue:  100,
				TriggerTimes:        1,
				Name:                "技能1",
				IncreasedDamage:     0,
				ReductionResistance: 0,
				DefenseBreak:        0,
				Penetration:         0,
				SpecialDamage:       0,
				ExplosiveInjury:     0,
			},
			// 可添加更多技能配置
		},
		Basic: &Basic{
			BasicAttack:              1000,
			BasicCritical:            50,
			BasicExplosiveInjury:     100,
			BasicIncreasedDamage:     0,
			BasicReductionResistance: 0,
			BasicVulnerable:          0,
			BasicSpecialDamage:       0,
			Penetration:              0,
		},
		Gain: &Gain{
			AttackValue:              0,
			AttackValue2:             0,
			AttackPowerPercentage:    0,
			AttackInternalPercentage: 0,
			Critical:                 0,
			ExplosiveInjury:          0,
			IncreasedDamage:          0,
			ReductionResistance:      0,
			Vulnerable:               0,
			SpecialDamage:            0,
		},
		Defense: &Defense{
			Penetration:      0,
			DefenseBreak:     0,
			PenetrationValue: 0,
		},
		Condition: &Condition{
			MainArticle: 6,
			Critical:    100,
		},
		Output:       &Output{},
		CurrentPanel: &CurrentPanel{},
	}
}
