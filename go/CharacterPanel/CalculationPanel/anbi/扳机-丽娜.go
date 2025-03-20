package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 安比01扳机00丽娜10() *Initializations {
	name := "安比01扳机00丽娜10"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin00()})

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

func 心弦夜响安比扳机00丽娜10() *Initializations {
	name := "心弦夜响安比扳机00丽娜10"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XinXianEeXiang(true, false))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin00()})

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

func 防爆者安比扳机00丽娜10() *Initializations {
	name := "防爆者安比扳机00丽娜10"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.FangBaoZhe(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin00()})

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

func 残心青囊安比扳机00丽娜10() *Initializations {
	name := "残心青囊安比扳机00丽娜10"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin00()})

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

func 硫磺石安比扳机00丽娜10() *Initializations {
	name := "硫磺石安比扳机00丽娜10"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.LiuHuangShi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin00()})

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

func 强音热望安比扳机00丽娜10() *Initializations {
	name := "强音热望安比扳机00丽娜10"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.QiangYinReWang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin00()})

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
