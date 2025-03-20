package main

import (
	"fmt"
	"zzz/CharacterPanel/common"
)

// ------------------------ 常量定义 ------------------------
const (

	// 每个队伍可分配的词条数（示例值）
	GlobalMainArticle = 55
	// 词条类型数量
	GlobalMainArticleTypeCount = 5
)

var AttackPercentageEntriesLimit = map[string]int{
	common.Critical:        16,
	common.ExplosiveInjury: 16,
}

var ExplosiveInjuryEntriesLimit = map[string]int{
	common.Critical:        24,
	common.ExplosiveInjury: 20,
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
	initializations := []*Initializations{
		安比01扳机00嘉音00(),
		安比01扳机01嘉音00(),
	}
	// 针对每套队伍进行计算
	for idx, initialization := range initializations {
		fmt.Printf("====== 队伍组合 %d: %s ======\n", idx+1, initialization.Name)
		_, bestDistribution, bestPanel, _, bestCritConverted, _ := initialization.FindOptimalDistribution()

		// 输出整体最佳方案
		fmt.Println("【整体最佳方案】")
		fmt.Println("最佳词条分配方案:")
		fmt.Printf("  攻击力词条: %d, 暴击词条: %d, 爆伤词条: %d, 增伤词条: %d, 穿透词条: %d, 暴击转换爆伤词条: %d\n",
			bestDistribution[common.AttackPercentage],
			bestDistribution[common.Critical],
			bestDistribution[common.ExplosiveInjury],
			bestDistribution[common.IncreasedDamage],
			bestDistribution[common.Penetrate],
			bestCritConverted,
		)

		// 输出最佳局内面板（综合最优方案）
		fmt.Println("最佳局内面板（综合最优方案）:")
		fmt.Printf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%\n",
			bestPanel.Attack,
			bestPanel.Critical,
			bestPanel.ExplosiveInjury,
			bestPanel.IncreasedDamage,
			bestPanel.Penetration,
		)

		// 为了计算局外面板，这里选取队伍中第一个模型（也可按实际需求处理）
		externalPanel := initialization.CalculateExternalPanel(bestDistribution)
		fmt.Println("最佳局外面板:")
		fmt.Printf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%\n",
			externalPanel.Attack,
			externalPanel.Critical,
			externalPanel.ExplosiveInjury,
		)
		fmt.Println("--------------------------------------------------")

		// 输出各模型（不同计算方法）的局内、局外面板及【单模型】的技能伤害明细
		for _, model := range initialization.Initializations { // 注意这里使用的是 Initialization 集合
			internalPanel := model.CurrentPanel
			externalPanel := initialization.CalculateExternalPanel(bestDistribution)
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
			// 对每个技能分别调用 InitializationArea 更新输出数据后计算伤害
			for _, mag := range model.Magnifications {
				initialization.InitializationArea(model, mag)
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

func copyMap(m map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range m {
		res[k] = v
	}
	return res
}

// FindOptimalDistribution 核心分配逻辑
func (i *Initializations) FindOptimalDistribution() (bestDamage float64, bestDistribution map[string]int, bestPanel *CurrentPanel, bestOutput *Output, bestCritConverted int, skillDamages map[string]float64) {
	distributions := generateDistributions(GlobalMainArticle, GlobalMainArticleTypeCount)
	bestDamage = -1.0
	bestDistribution = make(map[string]int)
	bestCritConverted = 0
	var bestSim []*Initialization
	// 初始化基础数值
	i.InitializationBase()
	//初始化属性词条的上限
	i.initializationCount()
	// 遍历所有分配方案
	for _, dist := range distributions {
		distribution := map[string]int{
			common.AttackPercentage: dist[0],
			common.Critical:         dist[1],
			common.ExplosiveInjury:  dist[2],
			common.IncreasedDamage:  dist[3],
			common.Penetrate:        dist[4],
		}
		var damage = 0.0
		var lastSim []*Initialization
		// 根据本分配方案，各模型计算伤害，并在计算前做必要的条件校验
		for _, initialization := range i.Initializations {
			if !i.checkCondition(distribution) {
				continue
			}
			i.CharacterPanelWithDistribution(initialization, distribution)
			// 模拟：如果暴击超过阈值（例如 GlobalCritical），则设为 100
			if initialization.CurrentPanel.Critical > 100 {
				initialization.CurrentPanel.Critical = 100
			}
			// 模拟：如果穿透大于 100，则设为 100
			if initialization.CurrentPanel.Penetration > 100 {
				initialization.CurrentPanel.Penetration = 100
			}
			damage += i.CalculatingTotalDamage(initialization)
			lastSim = i.Initializations
		}
		if lastSim == nil {
			continue
		}
		if damage > bestDamage {
			bestDamage = damage
			bestDistribution = copyMap(distribution)
			bestSim = lastSim
		}
	}
	if bestSim == nil {
		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")
		fmt.Println("出现错误，并没有获得最佳的面板")
		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")
		return bestDamage, bestDistribution, bestPanel, bestOutput, bestCritConverted, skillDamages
	}
	// 使用最佳模拟器计算各技能最终伤害及获取局内面板、输出数据
	// 这里选取 bestSim[0] 作为代表（也可以根据需求合并多个模型的结果）
	bestPanel = bestSim[0].CurrentPanel
	bestOutput = bestSim[0].Output

	skillDamages = make(map[string]float64)
	// 对每个模拟器中的每个技能倍率进行计算
	for _, sim := range bestSim {
		for _, mag := range sim.Magnifications {
			// 每次计算前调用 InitializationArea 更新输出数据
			i.InitializationArea(sim, mag)
			skillDamage := sim.Output.BasicDamageArea *
				sim.Output.IncreasedDamageArea *
				sim.Output.ExplosiveInjuryArea *
				sim.Output.DefenseArea *
				sim.Output.ReductionResistanceArea *
				sim.Output.VulnerableArea *
				sim.Output.SpecialDamageArea *
				(1 + mag.SpecialDamage/100)
			// 累加相同技能的伤害
			skillDamages[mag.Name] += skillDamage
		}
	}
	return bestDamage, bestDistribution, bestPanel, bestOutput, bestCritConverted, skillDamages
}

func (i *Initializations) initializationCount() {
	// 暴击词条上限
	critical := i.Basic.BasicCritical + i.Gain.Critical
	remaining := 100 - critical
	// 计算最大可用词条数，向下取整
	i.CriticalCount = int(remaining / 2.4)

	// 计算最终暴击率
	finalCritical := critical + float64(i.CriticalCount)*2.4

	// 如果最终暴击率超过 100，则减少 1 个词条
	if finalCritical > 100 {
		i.CriticalCount--
	}
	// 爆伤词条上限
	i.ExplosiveInjuryCount = ExplosiveInjuryEntriesLimit[i.NumberFour]
	i.AttackCount = AttackPercentageEntriesLimit[i.NumberFour]
}

func (i *Initializations) checkCondition(slots map[string]int) bool {
	pairStatus := false
	// 例如：校验（穿透+增伤）词条数量是否满足要求
	for _, pair := range allowedGroupB {
		if slots[common.IncreasedDamage] == pair[0] && slots[common.Penetrate] == pair[1] {
			pairStatus = true
			break
		}
	}
	if !pairStatus {
		return pairStatus
	}
	fiveStatus := false
	// 校验穿透固定之后，其他属性条件是否满足
	// 处理阈值问题，每个磁盘单个属性，最多5个有效词条，其中其中3号位可以+3词条，5号位+10词条，需要校验穿透和增伤词条数量，才可以确定阈值的上限是否有提升，判断穿透率的数值，提升不同属性的阈值上限
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 0 {
		// 三个词条总数必须大于13
		if slots[common.Critical]+slots[common.ExplosiveInjury]+slots[common.AttackPercentage] >= 13 {
			if slots[common.Critical] >= 3 || slots[common.ExplosiveInjury] >= 3 {
				// 2件套选择暴击或者爆伤的情况,5号位必须是攻击力
				if slots[common.AttackPercentage] >= 10 {
					if (i.AttackCount+10 >= slots[common.AttackPercentage]) && (i.CriticalCount+i.ExplosiveInjuryCount+3 >= slots[common.Critical]+slots[common.ExplosiveInjury]) {
						fiveStatus = true
					}
				}
			} else {
				// 2件套，5号位都是攻击力
				if slots[common.AttackPercentage] >= 13 {
					if i.AttackCount+13 >= slots[common.AttackPercentage] {
						fiveStatus = true
					}
				}
			}
		}
	}
	// 增伤+穿透 =3，说明2件套无法分配，攻击力必须>=10
	if (slots[common.IncreasedDamage]+slots[common.Penetrate] == 3) && slots[common.AttackPercentage] >= 10 {
		if i.AttackCount+10 >= slots[common.AttackPercentage] {
			fiveStatus = true
		}
	}
	// 增伤+穿透 =10，说明5号位无法分配，2件套可以是攻击，暴击，爆伤任意一个
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 10 {
		if slots[common.Critical] >= 3 || slots[common.ExplosiveInjury] >= 3 || slots[common.AttackPercentage] >= 3 {
			if i.AttackCount+i.CriticalCount+i.ExplosiveInjuryCount+3 >= slots[common.AttackPercentage]+slots[common.Critical]+slots[common.ExplosiveInjury] {
				fiveStatus = true
			}
		}
	}
	// 增伤+穿透 =13，说明承包了2件套和5号位选择，其他词条随意
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 13 {
		if i.AttackCount >= slots[common.AttackPercentage] && i.CriticalCount >= slots[common.Critical] && i.ExplosiveInjuryCount >= slots[common.ExplosiveInjury] {
			fiveStatus = true
		}
	}
	return fiveStatus
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
func (i *Initializations) CalculateExternalPanel(distribution map[string]int) *ExternalPanel {
	attack := i.Basic.BasicAttack*(1+(i.Gain.AttackPowerPercentage+float64(distribution[common.AttackPercentage])*3)/100) + i.Gain.AttackValue
	critical := i.Basic.BasicCritical + float64(distribution[common.Critical])*2.4
	explosiveInjury := i.Basic.BasicExplosiveInjury + (float64(distribution[common.ExplosiveInjury]))*4.8
	return &ExternalPanel{
		Attack:          attack,
		Critical:        critical,
		ExplosiveInjury: explosiveInjury,
	}
}

// CharacterPanelWithDistribution 根据词条分配更新局内面板
func (i *Initializations) CharacterPanelWithDistribution(initialization *Initialization, distribution map[string]int) {
	initialization.CurrentPanel = &CurrentPanel{
		ReductionResistance: i.Basic.BasicReductionResistance + i.Gain.ReductionResistance,
		Vulnerable:          i.Basic.BasicVulnerable + i.Gain.Vulnerable,
		SpecialDamage:       i.Basic.BasicSpecialDamage + i.Gain.SpecialDamage,
	}
	i.HandleBasicAttack(initialization, common.AttackPercentage, distribution[common.AttackPercentage])
	i.HandleBasicCritical(initialization, common.Critical, distribution[common.Critical])
	i.HandleBasicExplosiveInjury(initialization, common.ExplosiveInjury, distribution[common.ExplosiveInjury])
	i.HandleBasicIncreasedDamage(initialization, common.IncreasedDamage, distribution[common.IncreasedDamage])
	i.HandlePenetrateDamage(initialization, common.Penetrate, distribution[common.Penetrate])
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
func (i *Initializations) HandleBasicAttack(initialization *Initialization, key string, count int) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage
	if key == common.AttackPercentage {
		attackPowerPercentage += 3 * float64(count)
	}
	initialization.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100) + i.Gain.AttackValue + i.Gain.AttackValue2) * (1 + i.Gain.AttackInternalPercentage/100)
}

// HandleBasicCritical 根据暴击词条更新暴击率，并计算转换为爆伤的词条数
func (i *Initializations) HandleBasicCritical(initialization *Initialization, key string, count int) {
	critical := i.Gain.Critical
	if key == common.Critical {
		critical += 2.4 * float64(count)
	}
	initialization.CurrentPanel.Critical = i.Basic.BasicCritical + critical
}

// HandleBasicExplosiveInjury 根据爆伤词条更新爆伤
func (i *Initializations) HandleBasicExplosiveInjury(initialization *Initialization, key string, count int) {
	explosiveInjury := i.Gain.ExplosiveInjury
	if key == common.ExplosiveInjury {
		explosiveInjury += 4.8 * float64(count)
	}
	initialization.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + explosiveInjury
}

// HandleBasicIncreasedDamage 根据增伤词条更新增伤
func (i *Initializations) HandleBasicIncreasedDamage(initialization *Initialization, key string, count int) {
	increasedDamage := i.Gain.ExplosiveInjury
	if key == common.IncreasedDamage {
		increasedDamage += 3 * float64(count)
	}
	initialization.CurrentPanel.IncreasedDamage = i.Basic.BasicIncreasedDamage + increasedDamage
}

// HandlePenetrateDamage 根据穿透词条更新穿透率
func (i *Initializations) HandlePenetrateDamage(initialization *Initialization, key string, count int) {
	increasedDamage := i.Gain.Penetration
	if key == common.IncreasedDamage {
		increasedDamage += 2.4 * float64(count)
	}
	initialization.CurrentPanel.Penetration = i.Basic.Penetration + increasedDamage
}

// ===== 以下各函数计算各分区加成 =====

func (i *Initializations) CalculatingTotalDamage(initialization *Initialization) float64 {
	totalDamage := 0.0
	for _, mag := range initialization.Magnifications {
		i.InitializationArea(initialization, mag)
		damage := initialization.Output.BasicDamageArea *
			initialization.Output.IncreasedDamageArea *
			initialization.Output.ExplosiveInjuryArea *
			initialization.Output.DefenseArea *
			initialization.Output.ReductionResistanceArea *
			initialization.Output.VulnerableArea *
			initialization.Output.SpecialDamageArea *
			(1 + mag.SpecialDamage/100)
		totalDamage += damage
	}
	return totalDamage
}

func (i *Initializations) InitializationArea(initialization *Initialization, magnification *Magnification) {
	initialization.BasicDamageArea(magnification)
	initialization.IncreasedDamageArea(magnification)
	initialization.ExplosiveInjuryArea(magnification)
	i.DefenseArea(initialization, magnification)
	initialization.ReductionResistanceArea(magnification)
	initialization.VulnerableArea()
	initialization.SpecialDamageArea()
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

func (i *Initializations) DefenseArea(initialization *Initialization, magnification *Magnification) {
	characterBase, TotalDefense := 793.783, 873.1613
	penetration := (initialization.CurrentPanel.Penetration + magnification.Penetration) / 100
	defenseBreak := (i.Defense.DefenseBreak + magnification.DefenseBreak) / 100
	initialization.Output.DefenseArea = characterBase / (TotalDefense*(1-penetration)*(1-defenseBreak) - i.Defense.PenetrationValue + characterBase)
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

// DeepCopy 深拷贝方法（结构体方法）
func (i *Initializations) DeepCopy(initialization *Initialization) *Initialization {
	// 手动拷贝嵌套结构体，防止修改原始对象影响拷贝对象
	copyData := Initialization{
		Name:           initialization.Name,
		Magnifications: initialization.Magnifications,
		CurrentPanel:   initialization.CurrentPanel,
		Output:         initialization.Output,
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
