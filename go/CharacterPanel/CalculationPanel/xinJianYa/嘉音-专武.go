package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 星见雅01月城柳00嘉音00() *Initializations {
	name := "星见雅01月城柳00嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
		NumberSix:  common.AttackPowerPercentage,
	}
	// 初始化基础数值
	init.InitializationBase0命(Role.XinJianYa(), arms.XiaLuoXingDian(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase月城柳(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
	}
	return init
}
