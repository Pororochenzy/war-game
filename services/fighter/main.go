package main

import (
	"flag"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/lib/s_comps"
	"war-game/modules/fighter"
	"war-game/modules/gate"
	"war-game/services"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/core"
)

var (
	sID  = flag.String("sId", "gate_1", "获取需要启动的服务id,id不同,读取的配置文件也不同") //启动服务的Id
	conf = flag.String("conf", "./conf/gate_1.yaml", "获取需要启动的服务配置文件")
)

func main() {
	flag.Parse()
	s := NewService(
		cluster.SetConfPath(*conf),
		cluster.SetVersion(*sID),
	)
	s.OnInstallComp( //装备组件
		s_comps.NewGateRouteComp(),
	)
	lego.Run(s, //运行模块
		fighter.NewModule(),
		gate.NewModule(),
	)

}

func NewService(ops ...cluster.Option) core.IService {
	s := new(Demo1Service)
	s.Configure(ops...)
	return s
}

type Demo1Service struct {
	services.ServiceBase
}

func (this *Demo1Service) InitSys() {
	this.ServiceBase.InitSys()
}
