package main

import (
	"math"
	"zzz/CharacterPanel/common"
)

// CalculateExternalPanel 根据当前模型和词条分配计算局外面板
//
//	攻击力 = BasicAttack * (1 + (AttackPowerPercentage + 攻击力词条数*3)/100) + AttackValue
//	暴击率 = BasicCritical + 暴击词条数*2.4
//	爆伤   = BasicExplosiveInjury + (爆伤词条数 + 暴击转换爆伤词条数量)*4.8
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
	i.HandleBasicAttackValue(initialization, distribution[common.AttackValue])
	i.HandleBasicAttackPercentage(initialization, distribution[common.AttackPowerPercentage])
	i.HandleBasicIncreasedDamage(initialization, distribution[common.IncreasedDamage])
	i.HandlePenetrate(initialization, distribution[common.Penetrate])
	i.HandlePenetrationValue(initialization, distribution[common.PenetrationValue])
	i.HandleBasicProficient(initialization, distribution[common.Proficient])
}

// 根据攻击力词条增加攻击力
func (i *Initializations) HandleBasicAttackPercentage(initialization *Initialization, count int) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage + initialization.Gain.AttackPowerPercentage
	attackPowerPercentage += 3 * float64(count)
	initialization.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100)+i.Gain.AttackValue)*(1+i.Gain.AttackInternalPercentage/100) + i.Gain.AttackValue2 + initialization.CurrentPanel.Attack
}

func (i *Initializations) HandleBasicAttackValue(initialization *Initialization, count int) {
	initialization.CurrentPanel.Attack = i.Gain.AttackValue + initialization.Gain.AttackValue + 19*float64(count)
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

// 根据穿透词条更新穿透率
func (i *Initializations) HandlePenetrate(initialization *Initialization, count int) {
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
}

// 根据穿透词条更新穿透率
func (i *Initializations) HandlePenetrationValue(initialization *Initialization, count int) {
	penetrationValue := i.Defense.PenetrationValue
	penetrationValue += 9 * float64(count)
	initialization.CurrentPanel.PenetrationValue = i.Defense.PenetrationValue + penetrationValue
}

// 处理递增精通
func (i *Initializations) HandleBasicProficient(initialization *Initialization, count int) {
	critical := i.Gain.Proficient + initialization.Gain.Proficient
	critical += 9 * float64(count)
	initialization.CurrentPanel.Proficient = i.Basic.BasicProficient + critical
}

// ===== 以下各函数计算各分区加成 =====
func (i *Initializations) CalculatingTotalDamage(initialization *Initialization) float64 {
	totalDamage := 0.0
	for _, mag := range initialization.Magnifications {
		i.InitializationArea(initialization, mag)
		switch mag.DamageType {
		case common.Disorder:

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
		initialization.Output.ProficientArea *
		initialization.Output.ExplosiveInjuryArea
}

func (i *Initializations) DisorderDamage(initialization *Initialization) float64 {
	return initialization.Output.BasicDamageArea *
		initialization.Output.IncreasedDamageArea *
		initialization.Output.DefenseArea *
		initialization.Output.ReductionResistanceArea *
		initialization.Output.VulnerableArea *
		initialization.Output.SpecialDamageArea *
		initialization.Output.GradeArea *
		initialization.Output.ProficientArea *
		initialization.Output.ExplosiveInjuryArea
}

func (i *Initializations) InitializationArea(initialization *Initialization, magnification *Magnification) {
	initialization.BasicDamageArea(magnification)
	initialization.IncreasedDamageArea(magnification)
	initialization.ExplosiveInjuryArea(magnification)
	i.DefenseArea(initialization, magnification)
	initialization.ReductionResistanceArea(initialization, magnification)
	initialization.VulnerableArea(initialization)
	initialization.SpecialDamageArea(initialization, magnification)
	initialization.ProficientArea(initialization)
	initialization.GradeArea(initialization)
}

func (i *Initialization) BasicDamageArea(magnification *Magnification) {
	i.HandleDamageArea(magnification)
}

func (i *Initialization) HandleDamageArea(magnification *Magnification) {
	switch magnification.DamageType {
	case common.Disorder:
		i.Output.BasicDamageArea = i.HandleDisorderArea(magnification)
	case common.Abnormal:
		i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
	default:
		i.Output.BasicDamageArea = i.CurrentPanel.Attack * magnification.MagnificationValue / 100 * magnification.TriggerTimes
	}
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

func (i *Initialization) ReductionResistanceArea(initialization *Initialization, magnification *Magnification) {
	initialization.Output.ReductionResistanceArea = 1 + (magnification.ReductionResistance+i.CurrentPanel.ReductionResistance)/100
}

func (i *Initialization) VulnerableArea(initialization *Initialization) {
	initialization.Output.VulnerableArea = 1 + (i.CurrentPanel.Vulnerable)/100
}

func (i *Initialization) SpecialDamageArea(initialization *Initialization, magnification *Magnification) {
	initialization.Output.SpecialDamageArea = 1 + (i.CurrentPanel.SpecialDamage+magnification.SpecialDamage)/100
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
