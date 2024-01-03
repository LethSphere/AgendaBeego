package controllers

import(
	"github.com/astaxie/beego"
	"github.com/udistrital/agenda_v1/parametros_mid_v1/helpers"
)

type AgendaParametrosController struct {
	beego.controller
}

func (c *AgendaParametrosController) URLMapping(){

}

//GetAll
//@Tittle GetAll
//@Description get AgendaParametros
//@Success 200 {object} []models.AgendaParametros
//@Failure 400 bad request
//@Failure 500 Internal server error
//@router / [get]
func (c *AgendaParametrosController) GetAll() {
	defer helpers.ErrorController(c.Controller, "AgendaParametrosController")

	if v, err := helpers.ListarAgendaParametros(); err == nil{
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true,"Status":200, "Message": "Listado consultado con exito". "Data": v}
	} else {
		panic(err)
	}

	c.serveJSON()
}