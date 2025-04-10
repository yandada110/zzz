package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 朱鸢_强音_朋克青衣01丽娜11() *Initializations {
	name := "朱鸢_强音_朋克青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase朋克(Role.ZhuYuan(), arms.QiangYinReWang(true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.NiKe65_ZhuYuan()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase强音(),
			Gain: &Gain{
				Vulnerable: 50,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 朱鸢_强音_雷暴青衣01丽娜11() *Initializations {
	name := "朱鸢_强音_雷暴青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase雷暴(Role.YouZhen(), arms.QiangYinReWang(true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.NiKe65_ZhuYuan()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase强音(),
			Gain: &Gain{
				Vulnerable: 50,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}
