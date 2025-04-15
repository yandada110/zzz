package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 简01赛斯65嘉音00() *Initializations {
	name := "简01赛斯65嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Proficient,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.JianDu(), arms.CuiFengQianCi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.SaiSi(), Role.JiaYin00()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase1(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
		//&Initialization{
		//	Magnifications: MagnificationBase2(),
		//	Gain: &Gain{
		//		Vulnerable: 25,
		//	},
		//	Name:         name + "-失衡",
		//	Output:       &Output{},
		//	CurrentPanel: &CurrentPanel{},
		//},
	}
	return init
}
