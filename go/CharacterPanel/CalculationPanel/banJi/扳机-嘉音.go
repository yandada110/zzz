package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 安比01扳机01嘉音00() *Initializations {
	name := "安比01扳机01嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.SuoHunYingMou(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

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

func 安比01扳机0拘缚者嘉音00() *Initializations {
	name := "安比01扳机0拘缚者嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.JuFuZhe(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

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

func 安比01扳机0德玛拉嘉音00() *Initializations {
	name := "安比01扳机0德玛拉嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.DeMaLaDianChi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

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

func 安比01扳机0裁纸刀嘉音00() *Initializations {
	name := "安比01扳机0裁纸刀嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.CaiZhiDao(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.JiaYin00()})

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
