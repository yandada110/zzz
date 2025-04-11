package main

import (
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
	i.HandleBasicCritical(initialization, common.Critical, distribution[common.Critical])
	i.HandleBasicExplosiveInjury(initialization, common.ExplosiveInjury, distribution[common.ExplosiveInjury])
	i.HandleBasicIncreasedDamage(initialization, common.IncreasedDamage, distribution[common.IncreasedDamage])
	i.HandlePenetrateDamage(initialization, distribution[common.Penetrate])
}

// ===== 以下各函数处理词条加成 =====

func (i *Initializations) HandleBasicAttack(initialization *Initialization, key string, count int, attackInternalPercentage float64) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage + initialization.Gain.AttackPowerPercentage
	if key == common.AttackPowerPercentage {
		attackPowerPercentage += 3 * float64(count)
	}
	initialization.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100)+i.Gain.AttackValue)*(1+(i.Gain.AttackInternalPercentage+attackInternalPercentage)/100) + i.Gain.AttackValue2
}

// HandleBasicCritical 根据暴击词条更新暴击率，并计算转换为爆伤的词条数
func (i *Initializations) HandleBasicCritical(initialization *Initialization, key string, count int) {
	critical := i.Gain.Critical + initialization.Gain.Critical
	if key == common.Critical {
		critical += 2.4 * float64(count)
	}
	initialization.CurrentPanel.Critical = i.Basic.BasicCritical + critical
}

// HandleBasicExplosiveInjury 根据爆伤词条更新爆伤
func (i *Initializations) HandleBasicExplosiveInjury(initialization *Initialization, key string, count int) {
	explosiveInjury := i.Gain.ExplosiveInjury + initialization.Gain.ExplosiveInjury
	if key == common.ExplosiveInjury {
		explosiveInjury += 4.8 * float64(count)
	}
	initialization.CurrentPanel.ExplosiveInjury = i.Basic.BasicExplosiveInjury + explosiveInjury
}

// HandleBasicIncreasedDamage 根据增伤词条更新增伤
func (i *Initializations) HandleBasicIncreasedDamage(initialization *Initialization, key string, count int) {
	increasedDamage := i.Gain.IncreasedDamage + initialization.Gain.IncreasedDamage
	if key == common.IncreasedDamage {
		if count == 3 {
			increasedDamage += 10
		}
		if count == 10 {
			increasedDamage += 30
		}
		if count == 13 {
			increasedDamage += 40
		}
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

// ===== 以下各函数计算各分区加成 =====

func (i *Initializations) CalculatingTotalDamage(initialization *Initialization, distribution map[string]int) float64 {
	totalDamage := 0.0
	for _, mag := range initialization.Magnifications {
		i.InitializationArea(initialization, mag, distribution)
		damage := initialization.Output.BasicDamageArea *
			initialization.Output.IncreasedDamageArea *
			initialization.Output.ExplosiveInjuryArea *
			initialization.Output.DefenseArea *
			initialization.Output.ReductionResistanceArea *
			initialization.Output.VulnerableArea *
			initialization.Output.SpecialDamageArea
		totalDamage += damage
	}
	return totalDamage
}

func (i *Initializations) InitializationArea(initialization *Initialization, magnification *Magnification, distribution map[string]int) {
	initialization.BasicDamageArea(i, magnification, distribution)
	initialization.IncreasedDamageArea(magnification)
	initialization.ExplosiveInjuryArea(magnification)
	i.DefenseArea(initialization, magnification)
	initialization.ReductionResistanceArea(magnification)
	initialization.VulnerableArea()
	initialization.SpecialDamageArea()
}

func (i *Initialization) BasicDamageArea(magnification *Magnification, distribution map[string]int) {
	i.HandleDamageArea(magnification, distribution)
}

func (i *Initialization) HandleDamageArea(magnification *Magnification, distribution map[string]int) {
	switch magnification.DamageType {
	case common.Disorder:
		i.Output.BasicDamageArea = i.HandleDisorderArea(magnification)
	case common.Abnormal:
		i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
	default:
		i.handleDamageArea(i, magnification, distribution)
		i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
	}
}

func (i *Initialization) handleDamageArea(initializations *Initializations, magnification *Magnification, distribution map[string]int) {
	initializations.HandleBasicAttack(i, common.AttackPowerPercentage, distribution[common.AttackPowerPercentage], magnification.AttackInternalPercentage)
	i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
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
	//i.Output.BasicDamageArea = basicDamageArea
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
