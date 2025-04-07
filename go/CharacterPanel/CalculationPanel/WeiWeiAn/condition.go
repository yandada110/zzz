package main

import "zzz/CharacterPanel/common"

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
		if slots[common.AttackPowerPercentage]+slots[common.Proficient] >= 13 {
			// 2精通-5攻击力
			if slots[common.AttackPowerPercentage] >= 10 && slots[common.Proficient] >= 3 {
				if i.AttackPercentageCount+10 >= slots[common.AttackPowerPercentage] && i.ProficientCount+3 >= slots[common.Proficient] {
					fiveStatus = true
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
	// 增伤+穿透 =10，说明5号位无法分配，2件套可以是攻击，精通
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 10 {
		if slots[common.AttackPowerPercentage] >= 3 || slots[common.Proficient] >= 3 {
			if i.ProficientCount+i.AttackPercentageCount+3 >= slots[common.AttackPowerPercentage]+slots[common.Proficient] {
				fiveStatus = true
			}
		}
	}
	// 增伤+穿透 =13，说明承包了2件套和5号位选择，其他词条随意
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 13 {
		if i.AttackPercentageCount >= slots[common.AttackPowerPercentage] && i.ProficientCount >= slots[common.Proficient] && i.AttackValueCount >= slots[common.AttackValue] && i.PenetrationValueCount >= slots[common.PenetrationValue] {
			fiveStatus = true
		}
	}

	// 2件套增伤或者穿透 攻击力最少都有4个词条
	if slots[common.IncreasedDamage]+slots[common.Penetrate] >= 10 {
		if slots[common.AttackPowerPercentage] < 5 {
			return false
		}
		if slots[common.Proficient] < 5 {
			return false
		}
	} else {
		if slots[common.Proficient] >= 3 {
			if slots[common.AttackPowerPercentage] < 10 {
				return false
			}
		} else {
			if slots[common.AttackPowerPercentage] < 13 {
				return false
			}
		}
	}
	if slots[common.AttackValue] < 5 {
		return false
	}
	if slots[common.PenetrationValue] < 6 {
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
