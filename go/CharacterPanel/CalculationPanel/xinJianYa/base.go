package main

import (
	"math"
	"zzz/CharacterPanel/common"
)

// CalculateExternalPanel 根据当前模型和词条分配计算局外面板
func (i *Initializations) CalculateExternalPanel(distribution map[string]int) *ExternalPanel {
	attack := i.Basic.BasicAttack*(1+(i.Gain.AttackPowerPercentage+float64(distribution[common.AttackPowerPercentage])*3)/100) + i.Gain.AttackValue
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
		ReductionResistance: i.Basic.BasicReductionResistance + i.Gain.ReductionResistance + initialization.Gain.ReductionResistance,
		Vulnerable:          i.Basic.BasicVulnerable + i.Gain.Vulnerable + initialization.Gain.Vulnerable,
		SpecialDamage:       i.Basic.BasicSpecialDamage + i.Gain.SpecialDamage + initialization.Gain.SpecialDamage,
	}
	//i.HandleBasicAttack(initialization, common.AttackPowerPercentage, distribution[common.AttackPowerPercentage], 0)
	i.HandleBasicCritical(initialization, distribution[common.Critical])
	i.HandleBasicExplosiveInjury(initialization, distribution[common.ExplosiveInjury])
	i.HandleBasicIncreasedDamage(initialization, distribution[common.IncreasedDamage])
	i.HandlePenetrateDamage(initialization, distribution[common.Penetrate])
	i.HandleBasicProficient(initialization, distribution[common.Proficient])
}

// ===== 以下各函数处理词条加成 =====

func (i *Initializations) HandleBasicAttack(initialization *Initialization, key string, count int, attackInternalPercentage float64) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage + initialization.Gain.AttackPowerPercentage
	if key == common.AttackPowerPercentage {
		attackPowerPercentage += 3 * float64(count)
	}
	initialization.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100)+i.Gain.AttackValue+float64(i.Condition.AttackValueMin*19))*(1+(i.Gain.AttackInternalPercentage+attackInternalPercentage)/100) + i.Gain.AttackValue2
}

// HandleBasicCritical 根据暴击词条更新暴击率，并计算转换为爆伤的词条数
func (i *Initializations) HandleBasicCritical(initialization *Initialization, count int) {
	critical := i.Gain.Critical + initialization.Gain.Critical
	critical += 2.4 * float64(count)
	initialization.CurrentPanel.Critical = i.Basic.BasicCritical + critical
}

// HandleBasicExplosiveInjury 根据爆伤词条更新爆伤
func (i *Initializations) HandleBasicExplosiveInjury(initialization *Initialization, count int) {
	explosiveInjury := i.Gain.ExplosiveInjury + initialization.Gain.ExplosiveInjury
	explosiveInjury += 4.8 * float64(count)
	initialization.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + explosiveInjury
}

// HandleBasicIncreasedDamage 根据增伤词条更新增伤
func (i *Initializations) HandleBasicIncreasedDamage(initialization *Initialization, count int) {
	increasedDamage := i.Gain.IncreasedDamage + initialization.Gain.IncreasedDamage
	if count == 3 {
		increasedDamage += 10
	}
	if count == 10 {
		increasedDamage += 30
	}
	if count == 13 {
		increasedDamage += 40
	}
	initialization.CurrentPanel.IncreasedDamage = i.Basic.BasicIncreasedDamage + increasedDamage
}

// HandlePenetrateDamage 根据穿透词条更新穿透率
func (i *Initializations) HandlePenetrateDamage(initialization *Initialization, count int) {
	increasedDamage := i.Defense.Penetration
	if count == 3 {
		increasedDamage += 8
	}
	if count == 10 {
		increasedDamage += 24
	}
	if count == 13 {
		increasedDamage += 32
	}
	initialization.CurrentPanel.Penetration = i.Basic.Penetration + increasedDamage
	initialization.CurrentPanel.DefenseBreak = i.Defense.DefenseBreak
	initialization.CurrentPanel.PenetrationValue = i.Defense.PenetrationValue
}

// HandleBasicProficient 处理递增精通
func (i *Initializations) HandleBasicProficient(initialization *Initialization, count int) {
	critical := i.Gain.Proficient + initialization.Gain.Proficient
	critical += 9 * float64(count)
	initialization.CurrentPanel.Proficient = i.Basic.BasicProficient + critical
}

// ===== 以下各函数计算各分区加成 =====

func (i *Initializations) CalculatingTotalDamage(initialization *Initialization, distribution map[string]int) float64 {
	totalDamage := 0.0
	for _, mag := range initialization.Magnifications {
		i.InitializationArea(initialization, mag, distribution)
		switch mag.DamageType {
		case common.Disorder:
			totalDamage += i.DisorderDamage(initialization)
		case common.Abnormal:
			totalDamage += i.AbnormalDamage(initialization)
		default:
			totalDamage += i.DirectInjuryDamage(initialization)
		}
	}
	return totalDamage
}

func (i *Initializations) DirectInjuryDamage(initialization *Initialization) float64 {
	return initialization.Output.BasicDamageArea *
		initialization.Output.IncreasedDamageArea *
		initialization.Output.ExplosiveInjuryArea *
		initialization.Output.DefenseArea *
		initialization.Output.ReductionResistanceArea *
		initialization.Output.VulnerableArea *
		initialization.Output.SpecialDamageArea
}

func (i *Initializations) AbnormalDamage(initialization *Initialization) float64 {
	return initialization.Output.BasicDamageArea *
		initialization.Output.IncreasedDamageArea *
		initialization.Output.DefenseArea *
		initialization.Output.ReductionResistanceArea *
		initialization.Output.VulnerableArea *
		initialization.Output.SpecialDamageArea *
		initialization.Output.GradeArea *
		initialization.Output.ProficientArea
}

func (i *Initializations) DisorderDamage(initialization *Initialization) float64 {
	return initialization.Output.BasicDamageArea *
		initialization.Output.IncreasedDamageArea *
		initialization.Output.DefenseArea *
		initialization.Output.ReductionResistanceArea *
		initialization.Output.VulnerableArea *
		initialization.Output.SpecialDamageArea *
		initialization.Output.GradeArea *
		initialization.Output.ProficientArea
}

func (i *Initializations) InitializationArea(initialization *Initialization, magnification *Magnification, distribution map[string]int) {
	initialization.BasicDamageArea(i, magnification, distribution)
	initialization.IncreasedDamageArea(magnification)
	initialization.ExplosiveInjuryArea(magnification)
	i.DefenseArea(initialization, magnification)
	initialization.ReductionResistanceArea(magnification)
	initialization.VulnerableArea()
	initialization.SpecialDamageArea()
	initialization.ProficientArea(initialization)
	initialization.GradeArea(initialization)
}

func (i *Initialization) BasicDamageArea(initializations *Initializations, magnification *Magnification, distribution map[string]int) {
	i.HandleDamageArea(initializations, magnification, distribution)
}

func (i *Initialization) HandleDamageArea(initializations *Initializations, magnification *Magnification, distribution map[string]int) {
	initializations.HandleBasicAttack(i, common.AttackPowerPercentage, distribution[common.AttackPowerPercentage], magnification.AttackInternalPercentage)

	switch magnification.DamageType {
	case common.Disorder:
		i.HandleDisorderArea(magnification)
	case common.Abnormal:
		i.handleAbnormalArea(magnification)
	default:
		i.handleDamageArea(magnification)
	}
}

func (i *Initialization) handleDamageArea(magnification *Magnification) {
	i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
}

func (i *Initialization) handleAbnormalArea(magnification *Magnification) {
	// 强击-碎冰只有次数，异常，感电，计算剩余秒的伤害量
	var number float64 = 1
	// 火和以太，0.5秒一次伤害
	if magnification.DisorderType == common.Ether || magnification.DisorderType == common.Fire {
		number = (common.TimeTotal - magnification.TimeConsumption) * 2
	}
	// 火和以太，0.5秒一次伤害
	if magnification.DisorderType == common.Electricity {
		number = common.TimeTotal - magnification.TimeConsumption
	}
	i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes * number
}

func (i *Initialization) HandleDisorderArea(magnification *Magnification) (basicDamageArea float64) {
	switch magnification.DisorderType {
	case common.Fire:
		basicDamageArea = common.FireArea(common.TimeTotal, magnification.TimeConsumption, magnification.MagnificationValue)
	case common.Electricity:
		basicDamageArea = common.ElectricityArea(common.TimeTotal, magnification.TimeConsumption, magnification.MagnificationValue)
	case common.Physical:
		basicDamageArea = common.PhysicalArea(common.TimeTotal, magnification.TimeConsumption, magnification.MagnificationValue)
	case common.Ice:
		basicDamageArea = common.IceArea(common.TimeTotal, magnification.TimeConsumption, magnification.MagnificationValue)
	case common.Ether:
		basicDamageArea = common.EtherArea(common.TimeTotal, magnification.TimeConsumption, magnification.MagnificationValue)
	}
	i.Output.BasicDamageArea = i.CurrentPanel.Attack * basicDamageArea
	return basicDamageArea
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
	defenseBreak := (initialization.CurrentPanel.DefenseBreak + magnification.DefenseBreak) / 100
	initialization.Output.DefenseArea = characterBase / (TotalDefense*(1-penetration)*(1-defenseBreak) - initialization.CurrentPanel.PenetrationValue + characterBase)
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

func (i *Initialization) ProficientArea(initialization *Initialization) {
	initialization.Output.ProficientArea = i.CurrentPanel.Proficient / 100
}

func (i *Initialization) GradeArea(initialization *Initialization) {
	x := 1 + 1.0/59*(60-1)
	factor := math.Pow(10, float64(4))
	result := math.Trunc(x*factor) / factor
	initialization.Output.GradeArea = result
}
