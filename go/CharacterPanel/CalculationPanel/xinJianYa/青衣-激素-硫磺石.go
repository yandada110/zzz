package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 悠真_硫磺石_朋克青衣01丽娜11() *Initializations {
	name := "悠真_硫磺石_朋克青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase朋克(Role.YouZhen(), arms.LiuHuangShi(false))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LiNa11()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase硫磺石(),
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

func 悠真_硫磺石_如影青衣01丽娜11() *Initializations {
	name := "悠真_硫磺石_如影青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase如影(Role.YouZhen(), arms.LiuHuangShi(false))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LiNa11()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase硫磺石(),
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
