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

func JiaYin01() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              1200,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          25,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          20 + 20 + 24,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func JiaYin20() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              1600,
		AttackInternalPercentage: 12,
		Critical:                 0,
		ExplosiveInjury:          25,
		Proficient:               0,
		ReductionResistance:      15,
		IncreasedDamage:          20 + 24,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func JiaYin21() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              1600,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          25,
		Proficient:               0,
		ReductionResistance:      15,
		IncreasedDamage:          20 + 20 + 24,
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

func BanJi10() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          0,
		Vulnerable:               35 + 20,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func BanJi01() *BuffCharacter {
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
		DefenseBreak:             25,
	}
}

func BanJi11() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          0,
		Vulnerable:               35 + 20,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             25,
	}
}

func BanJi21() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          24,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          0,
		Vulnerable:               35 + 20,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             25,
	}
}

func LiNa10() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          0,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              37.5,
		DefenseBreak:             0,
	}
}

func LiNa11() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          20.2,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              38.5,
		DefenseBreak:             0,
	}
}

func KaiSa00() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              1000,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          25 + 15,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func KaiSa01() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              1000,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          25 + 15 + 18,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func KaiSa21() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              1500,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      15,
		IncreasedDamage:          25 + 15 + 18,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func BoKeNa65() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          30,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func SaiSi() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               100,
		ReductionResistance:      0,
		IncreasedDamage:          15,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func LaiTe01() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          30,
		Proficient:               0,
		ReductionResistance:      15,
		IncreasedDamage:          75,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func NiKe65() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 5,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          10,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             20,
	}
}

/*
*
朱鸢+妮可配置 4摇摆
*/
func NiKe65_ZhuYuan() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 15,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          24 + 25 + 15,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             40,
	}
}

func LuXi65YaoBai() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              600,
		AttackInternalPercentage: 12,
		Critical:                 0,
		ExplosiveInjury:          10,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          15,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func LuXi65JiaYin() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              600,
		AttackInternalPercentage: 12,
		Critical:                 0,
		ExplosiveInjury:          10,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          15,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func QingYi01() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          20,
		Vulnerable:               80,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func QingYi00() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 0,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          00,
		Vulnerable:               80,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             0,
	}
}

func ShiHengNiKe65() *BuffCharacter {
	return &BuffCharacter{
		AttackValue:              0,
		AttackInternalPercentage: 0,
		Critical:                 15,
		ExplosiveInjury:          0,
		Proficient:               0,
		ReductionResistance:      0,
		IncreasedDamage:          20 + 15,
		Vulnerable:               0,
		SpecialDamage:            0,
		Penetration:              0,
		DefenseBreak:             40,
	}
}
