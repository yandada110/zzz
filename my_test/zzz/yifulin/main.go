package main

import (
	"fmt"
	"math"
)

func main() {
	// 初始化角色数据和主词条Map
	initialization := NewInitialization()
	mainArticleDamageMap := make(map[string][]*Result)
	MainArticleMap := map[string]int{
		AttackPercentage: 0,
		Critical:         0,
		ExplosiveInjury:  0,
		IncreasedDamage:  0,
	}
	// 通过主词条数量，循环递增词条，记录面板和伤害量
	for key, _ := range MainArticleMap {
		// 初始面板
		initialization.CharacterPanel("", 0)
		OldDamage := initialization.CalculatingTotalDamage()
		MainArticleMap[key] = 1
		// 循环递增词条
		for MainArticleMap[key] < initialization.Condition.MainArticle {
			// 更新面板
			initialization.CharacterPanel(key, MainArticleMap[key])
			// 计算伤害
			NewDamage := initialization.CalculatingTotalDamage()
			// 深拷贝当前面板和输出，避免指针污染
			currentPanelCopy := *initialization.CurrentPanel
			outputCopy := *initialization.Output
			// 记录结果
			mainArticleDamageMap[key] = append(mainArticleDamageMap[key], &Result{
				CurrentPanel:         &currentPanelCopy,
				Output:               &outputCopy,
				Damage:               NewDamage,
				PercentageDifference: DecimalToPercentage(NewDamage, OldDamage),
			})
			// 打印日志
			fmt.Println(fmt.Sprintf("输出总伤害: %.6f万", NewDamage/10000))
			fmt.Println(fmt.Sprintf("伤害差距: %.2f%%", DecimalToPercentage(NewDamage, OldDamage)))
			// 更新旧伤害值
			OldDamage = NewDamage
			// 递增词条数量
			MainArticleMap[key]++
		}
	}
	initialization.GetMaxFloat(mainArticleDamageMap)
}

// GetMaxFloat 函数用于找出四个浮点数中的最大值
func (i *Initialization) GetMaxFloat(mainArticleDamageMap map[string][]*Result) {
	entriesMap := make(map[string]int)
	var criticalStatus = false
	for a := 0; a < i.Condition.MainArticle; a++ {
		optimal := Critical
		optimalKey := Critical + "--"
		// 判断，暴击率和攻击力，哪个更高
		if mainArticleDamageMap[optimal][entriesMap[optimalKey]].PercentageDifference < mainArticleDamageMap[AttackPercentage][entriesMap[AttackPercentage]].PercentageDifference {
			optimal, optimalKey = AttackPercentage, AttackPercentage
		}
		// 判断，上一个最大值和爆伤比较
		if mainArticleDamageMap[optimal][entriesMap[optimalKey]].PercentageDifference < mainArticleDamageMap[ExplosiveInjury][entriesMap[ExplosiveInjury]].PercentageDifference {
			optimal, optimalKey = ExplosiveInjury, ExplosiveInjury
		}
		// 判断，上一个最大值和增伤比较
		if mainArticleDamageMap[optimal][entriesMap[optimalKey]].PercentageDifference < mainArticleDamageMap[IncreasedDamage][entriesMap[IncreasedDamage]].PercentageDifference {
			optimal, optimalKey = IncreasedDamage, IncreasedDamage
		}
		var explosiveInjuryStatus = false
		// 如果选择的是暴击，检查是否超过阈值
		if optimal == Critical {
			entriesMap[optimalKey]++
			if criticalStatus {
				explosiveInjuryStatus = true
				optimal, optimalKey = ExplosiveInjury, ExplosiveInjury
			} else {
				if i.Condition.CriticalCount == entriesMap[optimal] {
					explosiveInjuryStatus = true
					criticalStatus = true
					optimal, optimalKey = ExplosiveInjury, ExplosiveInjury
				}
			}
		}
		// 记录本次递增词条
		entriesMap[optimal]++
		if explosiveInjuryStatus {
			fmt.Println(fmt.Sprintf("\n词条分配：攻击力：%d个，暴击：%d个，爆伤：%d个，增伤：%d个 差距：%f\n 总伤：%f\n",
				entriesMap[AttackPercentage],
				entriesMap[Critical],
				entriesMap[ExplosiveInjury],
				entriesMap[IncreasedDamage],
				mainArticleDamageMap[Critical][entriesMap[Critical+"--"]].PercentageDifference,
				mainArticleDamageMap[Critical][entriesMap[Critical+"--"]].Damage,
			))
		} else {
			fmt.Println(fmt.Sprintf("词条分配：攻击力：%d个，暴击：%d个，爆伤：%d个，增伤：%d个 差距：%f 总伤：%f",
				entriesMap[AttackPercentage],
				entriesMap[Critical],
				entriesMap[ExplosiveInjury],
				entriesMap[IncreasedDamage],
				mainArticleDamageMap[optimal][entriesMap[optimal]].PercentageDifference,
				mainArticleDamageMap[optimal][entriesMap[optimal]].Damage,
			))
		}
	}
	i.CurrentPanel = &CurrentPanel{
		ReductionResistance: i.Basic.BasicReductionResistance + i.Gain.ReductionResistance,
		Vulnerable:          i.Basic.BasicVulnerable + i.Gain.Vulnerable,
		SpecialDamage:       i.Basic.BasicSpecialDamage + i.Gain.SpecialDamage,
	}
	i.HandleBasicAttack("AttackPercentage", entriesMap[AttackPercentage])
	i.HandleBasicCritical("Critical", entriesMap[Critical])
	i.HandleBasicExplosiveInjury("ExplosiveInjury", entriesMap[ExplosiveInjury])
	i.HandleBasicIncreasedDamage("IncreasedDamage", entriesMap[IncreasedDamage])

	for _, mag := range i.Magnifications {
		i.IncreasedDamageArea(mag)
		i.ReductionResistanceArea(mag)
	}
	i.VulnerableArea()
	i.SpecialDamageArea()
	// 打印最终词条分配结果
	fmt.Println("最终面板：")
	fmt.Println(fmt.Sprintf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%",
		i.CurrentPanel.Attack,
		i.CurrentPanel.Critical,
		i.CurrentPanel.ExplosiveInjury,
		(i.Output.IncreasedDamageArea-1)*100, // 转换为百分比
	))
	fmt.Println(fmt.Sprintf("  抗性区: %.2f%%, 易伤区: %.2f%%, 特殊乘区: %.2f%%",
		(i.Output.ReductionResistanceArea-1)*100,
		(i.Output.VulnerableArea-1)*100,
		(i.Output.SpecialDamageArea-1)*100,
	))
}

// CharacterPanel 计算角色面板
func (i *Initialization) CharacterPanel(key string, count int) *Initialization {
	i.CurrentPanel = &CurrentPanel{
		ReductionResistance: i.Basic.BasicReductionResistance + i.Gain.ReductionResistance,
		Vulnerable:          i.Basic.BasicVulnerable + i.Gain.Vulnerable,
		SpecialDamage:       i.Basic.BasicSpecialDamage + i.Gain.SpecialDamage,
	}
	i.HandleBasicAttack(key, count)
	i.HandleBasicExplosiveInjury(key, count)
	i.HandleBasicCritical(key, count)
	i.HandleBasicIncreasedDamage(key, count)
	return i
}

// HandleBasicAttack 处理攻击力递增
func (i *Initialization) HandleBasicAttack(key string, count int) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage
	if key == AttackPercentage {
		fmt.Println(fmt.Sprintf("词条使用数量：攻击力：%d", count))
		attackPowerPercentage += 3 * float64(count)
	}
	i.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100) + i.Gain.AttackValue) * (1 + i.Gain.AttackInternalPercentage/100)
}

// HandleBasicCritical 处理暴击率递增
func (i *Initialization) HandleBasicCritical(key string, count int) {
	if key == Critical {
		fmt.Println(fmt.Sprintf("词条使用数量：暴击率：%d", count))
		// 计算当前暴击率
		critical := i.Basic.BasicCritical + i.Gain.Critical + 2.4*float64(count)

		// 如果暴击率超出阈值
		if critical > i.Condition.Critical {
			i.Condition.CriticalStatus++
			i.CurrentPanel.Critical = i.Basic.BasicCritical + i.Gain.Critical + 2.4*float64(count-i.Condition.CriticalStatus)
			i.CurrentPanel.ExplosiveInjury = i.CurrentPanel.ExplosiveInjury + float64(i.Condition.CriticalStatus)*4.8

			// 打印调试信息
			fmt.Println(fmt.Sprintf("暴击溢出处理: 原始暴击率=%.2f%%, 阈值=%.2f%%, 转换词条数=%d, 最终暴击率=%.2f%%, 爆伤增加=%.2f%%",
				critical, i.Condition.Critical, i.Condition.CriticalStatus, i.CurrentPanel.Critical, i.CurrentPanel.ExplosiveInjury))
		} else {
			i.Condition.CriticalCount = count
			// 如果未超出阈值，直接更新暴击率
			i.CurrentPanel.Critical = critical
		}
	} else {
		// 如果不是暴击词条，保持基础暴击率
		i.CurrentPanel.Critical = i.Basic.BasicCritical + i.Gain.Critical
	}
}

// HandleBasicExplosiveInjury 处理爆伤递增
func (i *Initialization) HandleBasicExplosiveInjury(key string, count int) {
	explosiveInjury := i.Gain.ExplosiveInjury
	if key == ExplosiveInjury {
		fmt.Println(fmt.Sprintf("词条使用数量：爆伤：%d", count))
		// 一个词条加4.8%
		explosiveInjury += 4.8 * float64(count)
	}
	i.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + explosiveInjury
}

// HandleBasicIncreasedDamage 处理增伤递增
func (i *Initialization) HandleBasicIncreasedDamage(key string, count int) {
	increasedDamage := i.Gain.IncreasedDamage
	if key == IncreasedDamage {
		fmt.Println(fmt.Sprintf("词条使用数量：增伤：%d", count))
		// 一个词条加4.8%
		increasedDamage += 3 * float64(count)
	}
	i.CurrentPanel.IncreasedDamage = i.Basic.BasicIncreasedDamage + increasedDamage
}

// CalculatingTotalDamage 计算总伤害
func (i *Initialization) CalculatingTotalDamage() float64 {
	// 初始化总伤害
	var totalDamage float64

	// 遍历所有伤害倍率
	for _, mag := range i.Magnifications {
		// 初始化乘区数值
		i.InitializationArea(mag)

		// 计算当前伤害
		damage := i.Output.BasicDamageArea *
			i.Output.IncreasedDamageArea *
			i.Output.ExplosiveInjuryArea *
			i.Output.DefenseArea *
			i.Output.ReductionResistanceArea *
			i.Output.VulnerableArea *
			i.Output.SpecialDamageArea *
			(1 + mag.SpecialDamage/100)

		// 异常角色，计算精通伤害

		// 累加总伤害
		totalDamage += damage

		// 打印当前伤害详情
		fmt.Println(fmt.Sprintf("[伤害详情] 技能: %s", mag.Name))
		fmt.Println(fmt.Sprintf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%",
			i.CurrentPanel.Attack,
			i.CurrentPanel.Critical,
			i.CurrentPanel.ExplosiveInjury,
			(i.Output.IncreasedDamageArea-1)*100, // 转换为百分比
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
	// 返回总伤害
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
	// 初始化乘区数值
	i.BasicDamageArea(magnification)
	i.IncreasedDamageArea(magnification)
	i.ExplosiveInjuryArea()
	i.DefenseArea(magnification)
	i.ReductionResistanceArea(magnification)
	i.VulnerableArea()
	i.SpecialDamageArea()
}

// BasicDamageArea 基础伤害区=基础攻击力*伤害倍率*次数
func (i *Initialization) BasicDamageArea(magnification *Magnification) {
	i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
}

// IncreasedDamageArea 增伤区
func (i *Initialization) IncreasedDamageArea(magnification *Magnification) {
	i.Output.IncreasedDamageArea = 1 + (magnification.IncreasedDamage+i.CurrentPanel.IncreasedDamage)/100
}

// ExplosiveInjuryArea 爆伤区
func (i *Initialization) ExplosiveInjuryArea() {
	i.Output.ExplosiveInjuryArea = 1 + (i.CurrentPanel.Critical*i.CurrentPanel.ExplosiveInjury)/10000
}

// DefenseArea 计算防御区
func (i *Initialization) DefenseArea(magnification *Magnification) {
	//防御区 = 角色基数/(总防御力*(1-穿透率)*(1-破防百分比)-穿透值+角色基数)
	//角色基数：794,总防御力：873
	characterBase, TotalDefense := 793.783, 873.1613
	penetration := (i.Defense.Penetration - magnification.Penetration) / 100
	defenseBreak := (i.Defense.DefenseBreak - magnification.DefenseBreak) / 100
	i.Output.DefenseArea = characterBase / (TotalDefense*(1-penetration)*(1-defenseBreak) - i.Defense.PenetrationValue + characterBase)
}

// ReductionResistanceArea 减抗区
func (i *Initialization) ReductionResistanceArea(magnification *Magnification) {
	i.Output.ReductionResistanceArea = 1 + (magnification.ReductionResistance+i.CurrentPanel.ReductionResistance)/100
}

// VulnerableArea 易伤区
func (i *Initialization) VulnerableArea() {
	i.Output.VulnerableArea = 1 + (i.CurrentPanel.Vulnerable)/100
}

// SpecialDamageArea 特殊乘区
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
}

// Basic 角色基础区
type Basic struct {
	BasicAttack              float64 // 基础攻击力（角色+专武）
	BasicCritical            float64 // 基础暴击（角色+武器+2件套+4号位）
	BasicExplosiveInjury     float64 // 基础爆伤（角色+武器+2件套+4号位）
	BasicIncreasedDamage     float64 // 基础增伤（角色+武器+驱动盘）
	BasicReductionResistance float64 // 基础减抗（角色+武器+驱动盘）
	BasicVulnerable          float64 // 基础易伤（百分比）
	BasicSpecialDamage       float64 // 基础特殊增伤（百分比）
}

// CurrentPanel 角色当前面板
type CurrentPanel struct {
	Attack              float64 // 攻击力（角色+专武）
	Critical            float64 // 暴击（角色+武器+2件套+4号位）
	ExplosiveInjury     float64 // 爆伤（角色+武器+2件套+4号位）
	IncreasedDamage     float64 // 增伤（角色+武器+驱动盘）
	ReductionResistance float64 // 减抗（角色+武器+驱动盘）
	Vulnerable          float64 // 易伤（百分比）
	SpecialDamage       float64 // 特殊增伤（百分比）
}

// Magnification 伤害倍率区
type Magnification struct {
	MagnificationValue  float64 // 倍率值 百分比
	TriggerTimes        float64 // 触发次数
	Name                string  // 伤害名称
	IncreasedDamage     float64 // 指定增伤，使用这个增伤+基础增伤  百分比
	ReductionResistance float64 // 指定减抗，使用这个减抗+基础减抗  百分比
	DefenseBreak        float64 // 指定破防，防御乘区，需要通过这个计算  百分比
	Penetration         float64 // 指定穿透，防御乘区，需要通过这个计算  百分比
	SpecialDamage       float64 // 指定独立的增益倍率区 百分比
}

// Gain 增益区
type Gain struct {
	AttackValue              float64 // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
	AttackPowerPercentage    float64 // 攻击力百分比加成(主词条，副词条，2件套)
	AttackInternalPercentage float64 // 局内攻击力(武器，驱动盘绿字攻击力)
	Critical                 float64 // 增加暴击（角色+武器+2件套+4号位）
	ExplosiveInjury          float64 // 增加爆伤（角色+武器+2件套+4号位）
	IncreasedDamage          float64 // 增伤（百分比）
	ReductionResistance      float64 // 减抗（百分比）
	Vulnerable               float64 // 易伤（百分比）
	SpecialDamage            float64 // 特殊增伤（百分比）
}

// Defense 防御区
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
	AttackPercentage = "AttackPercentage" // 百分比攻击力
	Critical         = "Critical"         // 暴击
	ExplosiveInjury  = "ExplosiveInjury"  // 爆伤
	IncreasedDamage  = "IncreasedDamage"  // 增伤
)

type Result struct {
	CurrentPanel         *CurrentPanel // 当前面板
	Output               *Output       // 当前增益
	Damage               float64       // 当前总伤害
	PercentageDifference float64       // 与上一次差距
}

func NewInitialization() *Initialization {
	magnifications := []*Magnification{
		&Magnification{
			MagnificationValue: 453 + 490.5,
			TriggerTimes:       2,
			Name:               "普攻",
		},
		&Magnification{
			MagnificationValue:  1658.7,
			TriggerTimes:        4.5,
			Name:                "连携技",
			IncreasedDamage:     30,
			ReductionResistance: 25,
			SpecialDamage:       25,
		},
		&Magnification{
			MagnificationValue: 1082.4 + 120.1,
			TriggerTimes:       2,
			Name:               "强化特殊技",
		},
		&Magnification{
			MagnificationValue:  3977.3,
			TriggerTimes:        1,
			Name:                "终结技",
			IncreasedDamage:     30,
			ReductionResistance: 25,
			SpecialDamage:       25,
		},
	}
	return &Initialization{
		Magnifications: magnifications, // 伤害倍率
		Basic: &Basic{
			BasicAttack:              929 + 713, // 基础攻击力（角色+专武）
			BasicCritical:            19.4 + 24, // 基础暴击（角色+武器+2件套+4号位）
			BasicExplosiveInjury:     50 + 48,   // 基础爆伤（角色+武器+2件套+4号位）
			BasicIncreasedDamage:     15,        // 基础增伤（角色+武器+驱动盘）
			BasicReductionResistance: 0,         // 基础减抗（角色+武器+驱动盘）
		},
		Gain: &Gain{
			AttackValue:              316 + 1200, // 攻击力值增加(2号位，副词条，嘉音，露西，凯撒增益)
			AttackPowerPercentage:    30,         // 攻击力百分比加成(主词条，副词条，2件套)
			AttackInternalPercentage: 37,         // 局内攻击力百分比(武器，驱动盘绿字攻击力)
			Critical:                 25,         // 增加暴击（角色+武器+2件套+4号位）
			ExplosiveInjury:          50 + 25,    // 增加爆伤（角色+武器+2件套+4号位）
			IncreasedDamage:          24 + 20,    // 增伤（队友百分比）
			ReductionResistance:      0,          // 减抗（百分比）
			Vulnerable:               25,         // 易伤（百分比）
			SpecialDamage:            0,          // 特殊增伤（百分比）
		},
		Defense: &Defense{
			Penetration:      0, // 穿透率（百分比）
			DefenseBreak:     0, // 破防百分比（百分比）
			PenetrationValue: 0, // 穿透值（固定值）
		},
		Condition: &Condition{
			MainArticle: 44, // 有效词条
			Critical:    95, // 最高暴击率
		},
		Output: &Output{},
	}
}
