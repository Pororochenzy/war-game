package comm

import "github.com/liwei1dao/lego/core"

const (
	ErrorCode_Success core.ErrorCode = 0 //成功

	ErrorCode_NoFoundData           core.ErrorCode = 9  //未查找到数据
	ErrorCode_NoFindService         core.ErrorCode = 10 //没有找到远程服务器
	ErrorCode_RpcFuncExecutionError core.ErrorCode = 11 //Rpc方法执行错误
	ErrorCode_CacheReadError        core.ErrorCode = 12 //缓存读取失败
	ErrorCode_SqlExecutionError     core.ErrorCode = 13 //数据库执行错误
	ErrorCode_ReqParameterError     core.ErrorCode = 14 //请求参数错误
	ErrorCode_SignError             core.ErrorCode = 15 //签名错误
	ErrorCode_UserIdInvalid         core.ErrorCode = 16 //目标用户Id无效
	ErrorCode_UserOffLine           core.ErrorCode = 17 //目标用户不在线
	ErrorCode_DeadlineExceeded      core.ErrorCode = 18 //调用net接口超时

	//fighter 战斗
	ErrorCode_NoLoginNoCanChat    core.ErrorCode = 3001 //未登录不能聊天
	ErrorCode_RedWalletStateError core.ErrorCode = 3002 //聊天红包状态错误
)
