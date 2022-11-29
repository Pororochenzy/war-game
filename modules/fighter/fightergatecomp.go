package fighter

import (
	"fmt"
	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib/m_comps"
	"war-game/comm"
	"war-game/pb"
)

type Fighter_GateComp struct {
	m_comps.MComp_GateComp
	service base.IClusterService
	Fighter *Fighter
}

func (this *Fighter_GateComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	this.MaxGoroutine = 100 //runtime.NumCPU()
	err = this.MComp_GateComp.Init(service, module, comp, options)
	this.service = service.(base.IClusterService)
	this.Fighter = module.(*Fighter)
	this.ComId = comm.FighterComId
	this.IsLog = true

	this.RegisterHandle(comm.FighterReq, &pb.DemoTestReq{}, this.HelloReq)
	return
}
func (this *Fighter_GateComp) Start() (err error) {
	err = this.MComp_GateComp.Start()
	return
}
func (this *Fighter_GateComp) Destroy() (err error) {
	return
}

func (this *Fighter_GateComp) HelloReq(session core.IUserSession, _msg interface{}) {
	msg := _msg.(*pb.DemoTestReq)
	rmsg := &pb.DemoTestResp{}
	//log.Warnf("reqMsg:%s", msg.Name)
	fmt.Printf("reqMsg:%s", msg.Name)
	rmsg.Name = "hello res"
	defer func() {
		session.SendMsg(this.ComId, comm.FighterRes, rmsg) //session 用的是  LocalSession
	}()

}
