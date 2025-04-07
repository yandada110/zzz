package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 悠真01朋克青衣01丽娜11() *Initializations {
	name := "悠真01朋克青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase朋克(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LiNa11()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase朋克专武(),
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

func 悠真01如影青衣01丽娜11() *Initializations {
	name := "悠真01如影青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase如影(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LiNa11()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase如影专武(),
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

func 悠真01雷暴青衣01丽娜11() *Initializations {
	name := "悠真01雷暴青衣01丽娜11"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase雷暴(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LiNa11()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase如雷专武(),
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

func 悠真01朋克青衣01露西65() *Initializations {
	name := "悠真01朋克青衣01露西65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase朋克(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LuXi65YaoBai()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase朋克专武(),
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

func 悠真01如影青衣01露西65() *Initializations {
	name := "悠真01如影青衣01露西65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase如影(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LuXi65YaoBai()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase如影专武(),
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

func 悠真01雷暴青衣01露西65() *Initializations {
	name := "悠真01雷暴青衣01露西65"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase雷暴(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.LuXi65YaoBai()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase如雷专武(),
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

func 悠真61朋克青衣01嘉音00() *Initializations {
	name := "悠真61朋克青衣01嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.JiaYin00()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase朋克专武6命(),
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

func 悠真61如影青衣01嘉音00() *Initializations {
	name := "悠真61如影青衣01嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase2(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.JiaYin00()})

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

func 悠真61雷暴青衣01嘉音00() *Initializations {
	name := "悠真61雷暴青衣01嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
	}
	// 初始化基础数值
	init.InitializationBase3(Role.YouZhen(), arms.CanXinQingNang(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.QingYi01(), Role.JiaYin00()})

	init.Initializations = []*Initialization{
		{
			Magnifications: MagnificationBase如雷专武6命(),
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
