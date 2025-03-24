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
	i.HandleBasicCritical(initialization, distribution[common.Critical])
	i.HandleBasicExplosiveInjury(initialization, common.ExplosiveInjury, distribution[common.ExplosiveInjury])
	i.HandleBasicIncreasedDamage(initialization, common.IncreasedDamage, distribution[common.IncreasedDamage])
	i.HandlePenetrateDamage(initialization, common.Penetrate, distribution[common.Penetrate])
}

// 根据攻击力词条增加攻击力
func (i *Initializations) HandleBasicAttackPercentage(initialization *Initialization, count int) {
	attackPowerPercentage := i.Gain.AttackPowerPercentage + initialization.Gain.AttackPowerPercentage
	attackPowerPercentage += 3 * float64(count)
	initialization.CurrentPanel.Attack = (i.Basic.BasicAttack*(1+attackPowerPercentage/100)+i.Gain.AttackValue)*(1+i.Gain.AttackInternalPercentage/100) + i.Gain.AttackValue2
}

func (i *Initializations) HandleBasicAttackValue(initialization *Initialization, count int) {
	i.Gain.AttackValue = i.Gain.AttackValue + initialization.Gain.AttackValue + 19*float64(count)
}

// HandleBasicCritical 根据暴击词条更新暴击率，并计算转换为爆伤的词条数
func (i *Initializations) HandleBasicCritical(initialization *Initialization, count int) {
	critical := i.Gain.Critical + initialization.Gain.Critical
	critical += 2.4 * float64(count)
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
func (i *Initializations) HandlePenetrateDamage(initialization *Initialization, key string, count int) {
	increasedDamage := i.Defense.Penetration
	if key == common.Penetrate {
		if count == 3 {
			increasedDamage += 8
		}
		if count == 10 {
			increasedDamage += 24
		}
		if count == 13 {
			increasedDamage += 32
		}
	}
	initialization.CurrentPanel.Penetration = i.Basic.Penetration + increasedDamage
	initialization.CurrentPanel.DefenseBreak = i.Defense.DefenseBreak
	initialization.CurrentPanel.PenetrationValue = i.Defense.PenetrationValue
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
		initialization.Output.ProficientArea
}

func (i *Initializations) InitializationArea(initialization *Initialization, magnification *Magnification) {
	initialization.BasicDamageArea(magnification)
	initialization.IncreasedDamageArea(magnification)
	initialization.ExplosiveInjuryArea(magnification)
	i.DefenseArea(initialization, magnification)
	initialization.ReductionResistanceArea(magnification)
	initialization.VulnerableArea()
	initialization.SpecialDamageArea()
	initialization.ProficientArea()
	initialization.GradeArea()
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

func (i *Initialization) ProficientArea() {
	i.Output.ProficientArea = i.CurrentPanel.Proficient / 100
}

func (i *Initialization) GradeArea() {
	x := 1 + 1.0/59*(60-1)
	factor := math.Pow(10, float64(4))
	result := math.Trunc(x*factor) / factor
	i.Output.GradeArea = result
}

func (i *Initializations) initializationCount() {
	// 暴击词条上限
	critical := i.Basic.BasicCritical + i.Gain.Critical
	remaining := CriticalCount - critical
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
	i.AttackPercentageCount = AttackPercentageEntriesLimit[i.NumberFour]
	i.ProficientCount = ProficientEntriesLimit[i.NumberFour]
	i.AttackValueCount = AttackValueEntriesLimit[i.NumberFour]
	i.PenetrationValueCount = PenetrationValueEntriesLimit[i.NumberFour]
}

func (i *Initializations) checkCondition(slots map[string]int) bool {
	pairStatus := false
	// 例如：校验（穿透+增伤）词条数量是否满足要求
	for _, pair := range common.AllowedGroupB {
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
		if slots[common.Critical]+slots[common.ExplosiveInjury]+slots[common.AttackPowerPercentage]+slots[common.Proficient] >= 13 {
			if slots[common.Critical] >= 3 || slots[common.ExplosiveInjury] >= 3 || slots[common.Proficient] >= 3 {
				// 2件套选择暴击或者爆伤或者精通的情况,5号位必须是攻击力
				if slots[common.AttackPowerPercentage] >= 10 {
					if (i.AttackPercentageCount+10 >= slots[common.AttackPowerPercentage]) && (i.ProficientCount+i.CriticalCount+i.ExplosiveInjuryCount+3 >= slots[common.Critical]+slots[common.ExplosiveInjury]+slots[common.Proficient]) {
						fiveStatus = true
					}
				}
			} else {
				// 2件套，5号位都是攻击力
				if slots[common.AttackPowerPercentage] >= 13 {
					if i.AttackPercentageCount+13 >= slots[common.AttackPowerPercentage] {
						fiveStatus = true
					}
				}
			}
		}
	}
	// 增伤+穿透 =3，说明2件套无法分配，攻击力必须>=10
	if (slots[common.IncreasedDamage]+slots[common.Penetrate] == 3) && slots[common.AttackPowerPercentage] >= 10 {
		if i.AttackPercentageCount+10 >= slots[common.AttackPowerPercentage] {
			fiveStatus = true
		}
	}
	// 增伤+穿透 =10，说明5号位无法分配，2件套可以是攻击，暴击，爆伤任意一个
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 10 {
		if slots[common.Critical] >= 3 || slots[common.ExplosiveInjury] >= 3 || slots[common.AttackPowerPercentage] >= 3 || slots[common.Proficient] >= 3 {
			if i.ProficientCount+i.AttackPercentageCount+i.CriticalCount+i.ExplosiveInjuryCount+3 >= slots[common.AttackPowerPercentage]+slots[common.Critical]+slots[common.ExplosiveInjury]+slots[common.Proficient] {
				fiveStatus = true
			}
		}
	}
	// 增伤+穿透 =13，说明承包了2件套和5号位选择，其他词条随意
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 13 {
		if i.AttackPercentageCount >= slots[common.AttackPowerPercentage] && i.CriticalCount >= slots[common.Critical] && i.ExplosiveInjuryCount >= slots[common.ExplosiveInjury] && i.ProficientCount >= slots[common.Proficient] {
			fiveStatus = true
		}
	}

	// 攻击力最少都有4个词条
	if slots[common.IncreasedDamage]+slots[common.Penetrate] >= 10 {
		if slots[common.AttackPowerPercentage] < 4 {
			return false
		}
	} else {
		if slots[common.AttackPowerPercentage] < 13 {
			return false
		}
	}
	if slots[common.AttackValue] < 6 {
		return false
	}
	if slots[common.PenetrateValue] < 6 {
		return false
	}
	////暴击爆伤可以不要
	//if i.NumberFour == common.Critical {
	//	if slots[common.Critical] < 5 {
	//		return false
	//	}
	//	if slots[common.ExplosiveInjury] < 6 {
	//		return false
	//	}
	//}
	//if i.NumberFour == common.ExplosiveInjury {
	//	if slots[common.Critical] < 6 {
	//		return false
	//	}
	//}
	//if i.NumberFour == common.ExplosiveInjury {
	//	if slots[common.ExplosiveInjury] < 5 {
	//		return false
	//	}
	//}
	if i.NumberFour == common.Proficient {
		if slots[common.Proficient] < 5 {
			return false
		}
	} else {
		if slots[common.Proficient] < 6 {
			return false
		}
	}
	return fiveStatus
}
