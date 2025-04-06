package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 雨果01莱特01嘉音00() *Initializations {
	name := "雨果01莱特01嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.YuGuo(), arms.QianMianRiYun(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.LaiTe01(), Role.BanJi01()})

	init.Initializations = []*Initialization{
		//&Initialization{
		//	Magnifications: MagnificationBase1(),
		//	Name:           name + "-站场",
		//	Gain:           &Gain{},
		//	Output:         &Output{},
		//	CurrentPanel:   &CurrentPanel{},
		//},
		&Initialization{
			Magnifications: MagnificationBase2(),
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
