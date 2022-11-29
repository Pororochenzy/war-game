package comm

import (
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib"
)

// 模块名定义处
const (
	Gate       core.M_Modules = lib.SM_GateModule //网关服务模块
	Http       core.M_Modules = lib.SM_HttpModule //http服务模块
	ModuleDemo core.M_Modules = "demo"            //演示模块
	Fighter    core.M_Modules = "FighterModule"   //战斗服务模块
)

const (
	// demo测试
	Rpc_ModuleDemoTest string = "Rpc_ModuleDemoTest" //测试

)
