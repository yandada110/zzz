package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 柏妮思01嘉音00() *Initializations {
	name := "柏妮思01嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Proficient,
		//NumberSix:  common.AttackPowerPercentage,
	}
	// 初始化基础数值
	init.InitializationBase0命(Role.BoNiSi(), arms.ZhuoXinYaoHu(true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase1(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
	}
	return init
}

func 柳01薇薇安01嘉音00() *Initializations {
	name := "柳01薇薇安01嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Proficient,
		//NumberSix:  common.AttackPowerPercentage,
	}
	// 初始化基础数值
	init.InitializationBase0命(Role.WeiWeiAn(), arms.FeiNiaoXingMeng(true, 6))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase2(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
	}
	return init
}

func 简01薇薇安1双生嘉音00() *Initializations {
	name := "简01薇薇安1双生嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Proficient,
		//NumberSix:  common.AttackPowerPercentage,
	}
	// 初始化基础数值
	init.InitializationBase0命(Role.WeiWeiAn(), arms.ShuangShengQiXing(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase双生1命(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
	}
	return init
}

func 简01薇薇安11嘉音00() *Initializations {
	name := "简01薇薇安11嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Proficient,
		//NumberSix:  common.AttackPowerPercentage,
	}
	// 初始化基础数值
	init.InitializationBase0命(Role.WeiWeiAn(), arms.FeiNiaoXingMeng(true, 6))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase1薇薇安11(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
	}
	return init
}

func 简01薇薇安21嘉音00() *Initializations {
	name := "简01薇薇安21嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Proficient,
		//NumberSix:  common.AttackPowerPercentage,
	}
	// 初始化基础数值
	init.InitializationBase0命(Role.WeiWeiAn(), arms.FeiNiaoXingMeng(true, 6))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase1薇薇安21(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
	}
	return init
}
