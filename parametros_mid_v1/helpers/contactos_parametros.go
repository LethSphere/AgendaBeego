package helpers

import(
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/agenda_v1/parametros_mid_v1/models"
)

func ListarContactosParametros() (contactosParametros models.ContactosParametros, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "ListarContactosParametros", "err": err, "status": "500"}
			panic(err)
		}
	}()

	var contactos []models.Contacto
	var parametros []models.Parametro

	url := "parametro?query=TipoParametroId__Nombre:Cargos&limit=0"
	if err := GetRequestNew("UrlCrudParametros", url, &parametros); err != nil {
		logs.Error(err)
		panic(err.Error())
	}

	url = "contacto/?limit=0"
	if err := GetRequestNew("UrlCrudContactos", url, &contactos); err != nil {
		logs.Error(err)
		panic(err.Error())
	}

	contactosParametros.Parametros = parametros
	contactosParametros.Contactos = contactos

	return contactosParametros, outputError
}