package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 安比01扳机0拘缚者妮可65() *Initializations {
	name := "安比01扳机0拘缚者妮可65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase2(Role.BanJi(), arms.JuFuZhe(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.NiKe65()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase5(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
		&Initialization{
			Magnifications: MagnificationBase6(),
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

func 安比01扳机01妮可65() *Initializations {
	name := "安比01扳机01妮可65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.SuoHunYingMou(false))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi01(), Role.NiKe65()})

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
				Vulnerable: 50,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 安比01扳机1拘缚者妮可65() *Initializations {
	name := "安比01扳机1拘缚者妮可65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase2(Role.BanJi(), arms.JuFuZhe(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi10(), Role.NiKe65()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase7(),
			Name:           name + "-站场",
			Gain:           &Gain{},
			Output:         &Output{},
			CurrentPanel:   &CurrentPanel{},
		},
		&Initialization{
			Magnifications: MagnificationBase8(),
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

func 安比01扳机11妮可65() *Initializations {
	name := "安比01扳机11妮可65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.SuoHunYingMou(false))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi11(), Role.NiKe65()})

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
				Vulnerable: 50,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 安比01扳机21妮可65() *Initializations {
	name := "安比01扳机21妮可65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.SuoHunYingMou(false))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi21(), Role.NiKe65()})

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
				Vulnerable: 50,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 安比01扳机0德玛拉妮可65() *Initializations {
	name := "安比01扳机0德玛拉妮可65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.DeMaLaDianChi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.NiKe65()})

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

func 安比01扳机0裁纸刀妮可65() *Initializations {
	name := "安比01扳机0裁纸刀妮可65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.BanJi(), arms.CaiZhiDao(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.NiKe65()})

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
