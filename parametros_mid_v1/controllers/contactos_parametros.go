package controllers

import(
	"github.com/astaxie/beego"
	"github.com/udistrital/agenda_v1/parametros_mid_v1/helpers"
)

type ContactosParametrosController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContactosParametrosController) URLMapping() {
	//c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}

// GetAll ...
// @Title Get All
// @Description Obtener los Contactos y Parametros
// @Success 200 {object} models.
// @Failure 403
// @router / [get]
func (c *ContactosParametrosController) GetAll() {
	defer helpers.ErrorController(c.Controller, "ContactoController")

	if v, err := helpers.ListarContactosParametros(); err == nil{
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Listado consultado con éxito", "Data": v}
	} else {
		panic(err)
	}

	c.ServeJSON()
}