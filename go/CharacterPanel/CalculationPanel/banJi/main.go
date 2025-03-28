package main

import (
	"fmt"
	"strconv"
	"zzz/CharacterPanel/common"
)

// ------------------------ 主函数 ------------------------
func main() {
	// 初始化各套队伍（示例，具体初始化函数需自行实现）
	initializations := []*Initializations{
		安比01扳机0拘缚者丽娜11(),
		安比01扳机1拘缚者丽娜11(),
		安比01扳机01丽娜11(),
		安比01扳机11丽娜11(),
		安比01扳机21丽娜11(),
	}
	// 针对每套队伍进行计算
	for idx, initialization := range initializations {
		fmt.Printf("====== 队伍组合 %d: %s ======\n", idx+1, initialization.Name)
		bestSim, bestDistribution, total, efficientTotal := initialization.FindOptimalDistribution()
		// 输出整体最佳方案
		fmt.Println("【整体最佳方案】计算次数：", strconv.Itoa(total), "it有效计算次数：", strconv.Itoa(efficientTotal))
		fmt.Println("最佳词条分配方案:")
		fmt.Printf("  攻击力词条: %d, 暴击词条: %d, 爆伤词条: %d, 增伤词条: %d, 穿透词条: %d\n",
			bestDistribution[common.AttackPowerPercentage],
			bestDistribution[common.Critical],
			bestDistribution[common.ExplosiveInjury],
			bestDistribution[common.IncreasedDamage],
			bestDistribution[common.Penetrate],
		)
		fmt.Println("--------------------------------------------------")
		bestSim.OutputResult(bestDistribution)
	}
}

func (i *Initializations) OutputResult(bestDistribution map[string]int) {
	var status bool
	// 输出各模型（不同计算方法）的局内、局外面板及【单模型】的技能伤害明细
	for _, model := range i.Initializations { // 注意这里使用的是 Initialization 集合
		if !status {
			internalPanel := model.CurrentPanel
			fmt.Println("局内面板:")
			var penetration float64 = 0
			if bestDistribution[common.Penetrate] == 3 {
				penetration = 8
			}
			if bestDistribution[common.Penetrate] == 13 {
				penetration = 32
			}
			if bestDistribution[common.Penetrate] == 10 {
				penetration = 24
			}
			fmt.Printf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%, 增伤: %.2f%%, 穿透: %.2f%%，破防: %.2f%%\n",
				internalPanel.Attack,
				internalPanel.Critical,
				internalPanel.ExplosiveInjury,
				internalPanel.IncreasedDamage,
				i.Defense.Penetration+penetration,
				internalPanel.DefenseBreak,
			)
			fmt.Println("--------------------------------------------------")
			fmt.Println("局外面板:")
			if bestDistribution[common.Penetrate] == 3 {
				penetration = 8
			}
			if bestDistribution[common.Penetrate] == 13 {
				penetration = 32
			}
			if bestDistribution[common.Penetrate] == 10 {
				penetration = 24
			}
			attack := float64(bestDistribution[common.AttackPowerPercentage])*3 + i.Gain.AttackPowerPercentage
			fmt.Printf("  攻击力: %.2f, 暴击: %.2f%%, 爆伤: %.2f%%,穿透: %.2f%%\n",
				i.Basic.BasicAttack*(1+attack/100)+i.Gain.AttackValue,
				i.Basic.BasicCritical+float64(bestDistribution[common.Critical])*2.4,
				i.Basic.BasicExplosiveInjury+float64(bestDistribution[common.ExplosiveInjury])*4.8,
				penetration,
			)
			status = true
		}
		// ---------------- 新增部分 ----------------
		// 根据当前模型的最终参数，计算并输出各技能的伤害明细
		fmt.Println("--------------------------------------------------")
		fmt.Println(model.Name, "-最终伤害:")
		// 对每个技能分别调用 InitializationArea 更新输出数据后计算伤害
		totalModelSkillDamage := i.CalculatingTotalDamage(model)
		fmt.Printf("  技能总伤害: %.6f\n", totalModelSkillDamage)
		fmt.Println("--------------------------------------------------")
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
func (i *Initializations) FindOptimalDistribution() (bestSim *Initializations, bestDistribution map[string]int, total int, efficientTotal int) {
	distributions := generateDistributions(GlobalMainArticle, GlobalMainArticleTypeCount)
	var bestDamage = -1.0
	bestDistribution = make(map[string]int)
	//初始化属性词条的上限
	i.initializationCount()
	// 遍历所有分配方案
	for _, dist := range distributions {
		total++
		distribution := map[string]int{
			common.AttackPowerPercentage: dist[0],
			common.Critical:              dist[1],
			common.ExplosiveInjury:       dist[2],
			common.IncreasedDamage:       dist[3],
			common.Penetrate:             dist[4],
		}
		var damage = 0.0
		var lastSim []*Initialization
		if !i.checkCondition(distribution) {
			continue
		}
		efficientTotal++
		// 根据本分配方案，各模型计算伤害，并在计算前做必要的条件校验
		for _, initialization := range i.Initializations {
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
			lastSim = append(lastSim, initialization.DeepCopy())
		}
		if lastSim == nil {
			continue
		}
		if damage > bestDamage {
			bestDamage = damage
			bestDistribution = copyMap(distribution)
			bestSim = i.DeepCopyData(lastSim)
		}
	}
	if bestSim == nil {
		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")
		fmt.Println("出现错误，并没有获得最佳的面板")
		fmt.Println("--------------------------------------------------")
		fmt.Println("--------------------------------------------------")
		return bestSim, bestDistribution, total, efficientTotal
	}

	return bestSim, bestDistribution, total, efficientTotal
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
