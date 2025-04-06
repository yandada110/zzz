package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 悠真硫磺石青衣01妮可65() *Initializations {
	name := "悠真硫磺石青衣01妮可65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.YouZhen(), arms.LiuHuangShi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.ShiHengNiKe65()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase如影专武6命(),
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

func 悠真硫磺石青衣01丽娜11() *Initializations {
	name := "悠真硫磺石青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.YouZhen(), arms.LiuHuangShi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.ShiHengNiKe65()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase如影专武6命(),
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

func 悠真硫磺石青衣00丽娜11() *Initializations {
	name := "悠真硫磺石青衣00丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.YouZhen(), arms.LiuHuangShi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi00(), Role.LiNa11()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase如影专武6命(),
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

func 悠真硫磺石青衣00丽娜10() *Initializations {
	name := "悠真硫磺石青衣00丽娜10"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.YouZhen(), arms.LiuHuangShi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi00(), Role.LiNa10()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase如影专武6命(),
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

func 悠真硫磺石青衣01丽娜10() *Initializations {
	name := "悠真硫磺石青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.YouZhen(), arms.LiuHuangShi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LiNa10()})

	init.Initializations = []*Initialization{
		&Initialization{
			Magnifications: MagnificationBase如影专武6命(),
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

func 悠真硫磺石青衣01露西65() *Initializations {
	name := "悠真硫磺石青衣01露西65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.YouZhen(), arms.LiuHuangShi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LuXi65YaoBai()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase如影专武6命(),
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
