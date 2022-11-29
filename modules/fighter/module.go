package fighter

import (
	"fmt"
	"war-game/comm"
	"war-game/modules/gate"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
)

var fighter *Fighter //内部指针

func NewModule() core.IModule {
	m := new(Fighter)
	fighter = m
	return m
}

type Fighter struct {
	cbase.ModuleBase
	service     core.IService
	gate        *gate.Gate
	fighterComp *Fighter_GateComp
	//userMgrComp *UserMgrComp
}

func (this *Fighter) GetType() core.M_Modules {
	return comm.Fighter
}

func (this *Fighter) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	err = this.ModuleBase.Init(service, module, options)
	this.service = service
	return
}

func (this *Fighter) Start() (err error) {
	if err = this.ModuleBase.Start(); err != nil {
		return
	}
	if m, err := this.service.GetModule(comm.Gate); err != nil {
		return fmt.Errorf("fighter 获取网关模块失败！")
	} else {
		this.gate = m.(*gate.Gate)
	}
	return
}

func (this *Fighter) OnInstallComp() {
	this.ModuleBase.OnInstallComp()
	this.fighterComp = this.RegisterComp(new(Fighter_GateComp)).(*Fighter_GateComp)
}
