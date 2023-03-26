package demo

import (
	"fmt"
	"strconv"

	"github.com/tianzhaocn/skyscraper/app/http/middleware"
	"github.com/tianzhaocn/skyscraper/app/provider/demoservice"
	"github.com/tianzhaocn/skyscraper/framework"
	"github.com/tianzhaocn/skyscraper/framework/containerService/contract"
	"gorm.io/gorm"
)

type DemoApi struct {
	dbServie   *gorm.DB
	logService contract.Log
	myServiece demoservice.MyService
}

func Register(c *framework.Core) error {
	api := NewDemoApi(c)
	c.Use(middleware.Mydemo())
	c.Get("/", api.Hello)
	c.Get("/echo/:id", api.Echo)
	c.Get("/users", api.GetUsers)
	return nil
}

func NewDemoApi(c *framework.Core) *DemoApi {
	db, _ := c.MustMake(contract.ORMKey).(contract.ORMService).GetDB()
	return &DemoApi{
		dbServie:   db,
		logService: c.MustMake(contract.LogKey).(contract.Log),
		myServiece: c.MustMake(demoservice.DemoKey).(demoservice.MyService),
	}
}

func (api *DemoApi) Hello(c *framework.Context) error {
	fmt.Println("HELLO")
	c.HtmlAfile("skyscraper", "app/http/module/demo/skyscraper.html", map[string]string{
		"Version": "1.0.0",
	})
	return nil
}

func (api *DemoApi) Echo(c *framework.Context) error {
	api.myServiece.SayHello()
	id, _ := c.ParamInt("id", 0)
	c.SetOKStatus().Json("your id is :" + strconv.Itoa(id))
	api.myServiece.SayByeBye()
	return nil

}

func (api *DemoApi) GetUsers(c *framework.Context) error {
	users := []UserModel{}
	api.dbServie.Find(&users)
	c.SetOKStatus().Json(users)
	api.logService.Info(c, "get users", map[string]interface{}{"result": "ok"})
	return nil

}
