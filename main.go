package main
import (
	"cip/config"
	"cip/control"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func router(app  *iris.Application)  {
	app.Get("/",control.RemoteIpInfo)

}
func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Logger().SetLevel("info")
	router(app)
	listenInfo :=fmt.Sprintf("%s:%s",config.Conf.Get("listen.hostAddress").(string),config.Conf.Get("listen.port").(string))
	app.Run(iris.Addr(listenInfo), iris.WithoutInterruptHandler)
	//fmt.Println("aa",config.Conf.Get("app.host").(string))
	//fmt.Println("bb",config.Conf.Get("api.ipUrl").(string))


}

