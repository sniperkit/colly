package feedify

import (
	"github.com/astaxie/beego"
	"github.com/sniperkit/colly/plugins/data/aggregate/web/core/context"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) GetInput() *context.Input {
	return &context.Input{c.Ctx.Input}
}

func (c *Controller) GetCtx() *context.Context {
	return &context.Context{c.Ctx}
}

func (c *Controller) GetJsonData() interface{} {
	return c.Controller.Data["json"]
}
