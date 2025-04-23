package main

import "zzz/CharacterPanel/common"

type Condition struct {
	AttackPercentageMax int // 攻击力百分比词条基础上限
	ProficientMax       int // 精通词条基础上限
	CriticalMax         int // 暴击词条基础上限
	ExplosiveInjuryMax  int // 爆伤词条基础上限
	AttackValueMax      int // 攻击力值词条基础上限
	PenetrationValueMax int // 穿透值基础上限

	AttackPercentageMin int // 攻击力百分比词条基础下限
	ProficientMin       int // 精通词条基础下限
	CriticalMin         int // 暴击词条基础下限
	ExplosiveInjuryMin  int // 爆伤词条基础下限
	AttackValueMin      int // 攻击力值词条基础下限
	PenetrationValueMin int // 穿透值基础下限
}

func (i *Initializations) initializationCount() {
	i.Condition = &Condition{}
	// 暴击词条上限
	critical := i.Basic.BasicCritical + i.Gain.Critical
	remaining := 100 - critical
	// 计算最大可用词条数，向下取整
	i.Condition.CriticalMax = int(remaining / 2.4)

	// 计算最终暴击率
	finalCritical := critical + float64(i.Condition.CriticalMax)*2.4

	// 如果最终暴击率超过 100，则减少 1 个词条
	if finalCritical > 100 {
		i.Condition.CriticalMax--
	}

	// 爆伤词条
	i.Condition.ExplosiveInjuryMin = 5
	// 5个磁盘+4
	i.Condition.ExplosiveInjuryMax = i.Condition.ExplosiveInjuryMin + 25
	// 如果4号位不是爆伤，上限提升
	if i.NumberFour != common.ExplosiveInjury {
		i.Condition.ExplosiveInjuryMin++
		i.Condition.ExplosiveInjuryMax = i.Condition.ExplosiveInjuryMin + 30
	}

	// 百分比攻击力
	i.Condition.AttackPercentageMin = 3
	// 1,2,3号位
	i.Condition.AttackPercentageMax = i.Condition.AttackPercentageMin * 5
	if i.NumberFour != common.AttackPowerPercentage {
		i.Condition.AttackPercentageMin++
		i.Condition.AttackPercentageMax = i.Condition.AttackPercentageMin * 5
	}
	if i.NumberSix != common.AttackPowerPercentage {
		i.Condition.AttackPercentageMin++
		i.Condition.AttackPercentageMax = i.Condition.AttackPercentageMin * 5
	}

	// 精通上限
	i.Condition.ProficientMin = 5
	// 1,2,3,5,6号位
	i.Condition.ProficientMax = i.Condition.ProficientMin * 5
	if i.NumberFour != common.Proficient {
		i.Condition.ProficientMin++
		i.Condition.ProficientMax = i.Condition.ProficientMin * 5
	}

	// 攻击值-穿透值上限
	//i.Condition.AttackValueMax = 25
	//i.Condition.PenetrationValueMax = 30
	i.Condition.AttackValueMin = 5
	i.Condition.PenetrationValueMin = 6
}

func (i *Initializations) checkCondition(slots map[string]int) bool {
	status := false
	// 增伤-穿透率是固定值，不满足的直接过滤
	for _, pair := range common.AllowedGroupB {
		if slots[common.IncreasedDamage] == pair[0] && slots[common.Penetrate] == pair[1] {
			status = true
			break
		}
	}
	if !status {
		return false
	}
	condition := &Condition{
		AttackPercentageMax: i.Condition.AttackPercentageMax,
		ProficientMax:       i.Condition.ProficientMax,
		CriticalMax:         i.Condition.CriticalMax,
		ExplosiveInjuryMax:  i.Condition.ExplosiveInjuryMax,
		AttackValueMax:      i.Condition.AttackValueMax,
		PenetrationValueMax: i.Condition.PenetrationValueMax,

		AttackPercentageMin: i.Condition.AttackPercentageMin,
		ProficientMin:       i.Condition.ProficientMin,
		CriticalMin:         i.Condition.CriticalMin,
		ExplosiveInjuryMin:  i.Condition.ExplosiveInjuryMin,
		AttackValueMin:      i.Condition.AttackValueMin,
		PenetrationValueMin: i.Condition.PenetrationValueMin,
	}
	// 假设穿透率-增伤都是0
	if !i.handle穿透增伤0(condition, slots) {
		return false
	}
	// 假设穿透率-增伤都是3
	if !i.handle穿透增伤3(condition, slots) {
		return false
	}
	// 假设穿透率-增伤都是10
	if !i.handle穿透增伤10(condition, slots) {
		return false
	}
	// 精通不满足退出
	if slots[common.Proficient] < i.Condition.ProficientMin || slots[common.Proficient] > i.Condition.ProficientMax {
		return false
	}
	// 攻击不满足退出
	if slots[common.AttackPowerPercentage] < condition.AttackPercentageMin || slots[common.AttackPowerPercentage] > condition.AttackPercentageMax {
		return false
	}
	return true
}

func (i *Initializations) handle穿透增伤0(condition *Condition, slots map[string]int) bool {
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 0 {
		// 5号位必须是攻击力
		if slots[common.AttackPowerPercentage] < i.Condition.AttackPercentageMin+10 {
			return false
		}
		condition.AttackPercentageMax += 10
		// 假设，2件套是其他的情况下
		if slots[common.AttackPowerPercentage] < i.Condition.AttackPercentageMin+13 {
			// 必须要有一个是2件套
			if slots[common.Proficient] < i.Condition.ProficientMin+3 {
				return false
			}
			condition.ProficientMax += 3
		} else {
			condition.AttackPercentageMax += 3
		}
	}
	return true
}

func (i *Initializations) handle穿透增伤3(condition *Condition, slots map[string]int) bool {
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 3 {
		// 5号位必须是攻击力
		if slots[common.AttackPowerPercentage] < i.Condition.AttackPercentageMin+10 {
			return false
		}
		condition.AttackPercentageMax += 10
	}
	return true
}

func (i *Initializations) handle穿透增伤10(condition *Condition, slots map[string]int) bool {
	if slots[common.IncreasedDamage]+slots[common.Penetrate] == 10 {
		condition.AttackPercentageMin++
		// 必须要有一个是2件套
		if (slots[common.Proficient] < i.Condition.ProficientMin+3) && (slots[common.AttackPowerPercentage] < i.Condition.AttackPercentageMin+3) {
			return false
		}
		condition.ProficientMax += 3
		condition.AttackPercentageMax += 3
	}
	return true
}
