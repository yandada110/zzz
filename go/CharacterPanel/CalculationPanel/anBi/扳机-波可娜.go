package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 安比01扳机00波可娜65() *Initializations {
	name := "安比01扳机00波可娜65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.BoKeNa65()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase1(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
		&Initialization{
			Magnifications: MagnificationBase2(),
			Gain: &Gain{
				Vulnerable: 25,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 安比01扳机01波可娜65() *Initializations {
	name := "安比01扳机01波可娜65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi01(), Role.BoKeNa65()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase1(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
		&Initialization{
			Magnifications: MagnificationBase2(),
			Gain: &Gain{
				Vulnerable: 25,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 安比21扳机00波可娜65() *Initializations {
	name := "安比21扳机00波可娜65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase2(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.BoKeNa65()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase3(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
		&Initialization{
			Magnifications: MagnificationBase4(),
			Gain: &Gain{
				Vulnerable: 25,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 安比21扳机01波可娜65() *Initializations {
	name := "安比21扳机01波可娜65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase2(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi01(), Role.BoKeNa65()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase3(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
		&Initialization{
			Magnifications: MagnificationBase4(),
			Gain: &Gain{
				Vulnerable: 25,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}
