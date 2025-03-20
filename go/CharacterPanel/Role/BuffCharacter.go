package Role

/**
辅助角色增益
*/

func JiaYin00() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              1200,
		AttackInternalPercentage: 12,
		Critical:                 0,
		ExplosiveInjury:          25,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          20 + 24,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func BanJi00() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          0,
		Vulnerable:               35,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}
