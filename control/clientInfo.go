package control

import (
	"cip/service"
	"cip/utils"
	"github.com/kataras/iris"
)

func RemoteIpInfo(ctx iris.Context)  {
	//clientInfo:= model.ClientINfo{}
	clientInfo,err := service.RemoteIpInfo(ctx)
	if err!=nil{
		ctx.JSON(utils.ResultUtil.Failure(err.Error()))
	}else {
		ctx.JSON(utils.ResultUtil.Success(clientInfo))
	}

}
