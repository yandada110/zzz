package main

import (
	"zzz/CharacterPanel/Role"
	"zzz/CharacterPanel/arms"
	"zzz/CharacterPanel/common"
)

func 安比01凯撒00嘉音00() *Initializations {
	name := "安比01凯撒00嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.KaiSa00(), Role.JiaYin00()})

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

func 安比01扳机00嘉音00() *Initializations {
	name := "安比01扳机00嘉音00"
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
				Vulnerable: 50,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 安比01扳机01嘉音00() *Initializations {
	name := "安比01扳机01嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi01(), Role.JiaYin00()})

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

func 安比01扳机10嘉音00() *Initializations {
	name := "安比01扳机10嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi10(), Role.JiaYin00()})

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

func 安比01扳机11嘉音00() *Initializations {
	name := "安比01扳机11嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi11(), Role.JiaYin00()})

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

func 安比01扳机21嘉音00() *Initializations {
	name := "安比01扳机21嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi21(), Role.JiaYin00()})

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

func 心弦夜响安比扳机00嘉音00() *Initializations {
	name := "硫磺石安比扳机00嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XinXianEeXiang(true))
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

func 防爆者安比扳机00嘉音00() *Initializations {
	name := "硫磺石安比扳机00嘉音00"
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

func 残心青囊安比扳机00嘉音00() *Initializations {
	name := "硫磺石安比扳机00嘉音00"
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
func 硫磺石安比扳机00嘉音00() *Initializations {
	name := "硫磺石安比扳机00嘉音00"
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

func 强音热望安比扳机00嘉音00() *Initializations {
	name := "强音热望安比扳机00嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.ExplosiveInjury,
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
				Vulnerable:               25,
				AttackInternalPercentage: 9.6,
			},
			Name:         name + "-失衡",
			Output:       &Output{},
			CurrentPanel: &CurrentPanel{},
		},
	}
	return init
}

func 硫磺石安比扳机00嘉音20() *Initializations {
	name := "硫磺石安比扳机00嘉音00"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.LiuHuangShi(true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin20()})

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

func 安比01扳机00嘉音01() *Initializations {
	name := "安比01扳机00嘉音01"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin01()})

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

func 安比01扳机01嘉音01() *Initializations {
	name := "安比01扳机01嘉音01"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi01(), Role.JiaYin01()})

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

func 安比01扳机00嘉音20() *Initializations {
	name := "安比01扳机00嘉音20"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin20()})

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

func 安比01扳机01嘉音20() *Initializations {
	name := "安比01扳机01嘉音20"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi01(), Role.JiaYin20()})

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

func 安比01扳机00嘉音21() *Initializations {
	name := "安比01扳机00嘉音21"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin21()})

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

func 安比01扳机01嘉音21() *Initializations {
	name := "安比01扳机01嘉音21"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase1(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi01(), Role.JiaYin21()})

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

func 安比21扳机00嘉音21() *Initializations {
	name := "安比21扳机00嘉音21"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase2(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi00(), Role.JiaYin21()})

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

func 安比21扳机01嘉音21() *Initializations {
	name := "安比21扳机01嘉音21"
	init := &Initializations{
		Name:       name,
		NumberFour: common.Critical,
	}
	// 初始化基础数值
	init.InitializationBase2(Role.AnBi(), arms.XiShengChunJie(true, true, true))
	// 初始化角色增益
	init.InitializationRole([]*Role.BuffCharacter{Role.BanJi01(), Role.JiaYin21()})

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
