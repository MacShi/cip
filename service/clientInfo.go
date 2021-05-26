package service

import (
	"cip/config"
	"cip/model"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
)

func RemoteIpInfo(ctx iris.Context) (model.ClientINfo,error)  {
	clientInfo := model.ClientINfo{}
	xRealIP := ctx.GetHeader("X-Real-IP")
	if len(xRealIP)>0{
		clientInfo.IpAddress = xRealIP
		//location,_:= requestIpLocation("117.158.210.253")
		location,err:= requestIpLocation(xRealIP)
		if err!=nil{
			return model.ClientINfo{}, err
		}else {
			clientInfo.Location.City = location.City
			clientInfo.Location.Province = location.Province
		}

	}else {
		clientInfo.IpAddress = ctx.RemoteAddr()
	}
	userAgent := ctx.GetHeader("User-Agent")
	if len(userAgent)>0{
		clientInfo.UserAgent = userAgent
	}else {
		clientInfo.UserAgent = "Null"
	}
	return clientInfo,nil

}
func requestIpLocation(xRealIP string) (*model.Location,error)  {
	url := config.Conf.Get("api.ipUrl").(string)
	key := config.Conf.Get("api.ipUrlKey").(string)
	url = fmt.Sprintf("%s?key=%s&type=4&ip=%s",url,key,xRealIP)
	resp, err := http.Get(url)
	if err!=nil{
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil, err
	}else {
		location :=model.Location{}
		bodyJson := map[string]string{}
		json.Unmarshal([]byte(string(body)),&bodyJson)
		fmt.Println(bodyJson["info"])
		location.Province = bodyJson["province"]
		location.City = bodyJson["city"]
		return &location,nil
	}
}

