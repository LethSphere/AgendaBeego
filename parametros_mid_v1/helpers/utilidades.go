package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/beego/beego/logs"
	"github.com/udistrital/utils_oas/formatdata"
)

//envia una peticion al endpoint indicado y extrae la respuesta del campo Data para retornarla
func GetRequestNer(endpoint string, route string, target interface{}) error {
	url := beego.AppConfig.String("ProtocoloAdmin") + "://" + beego.AppConfig.String(endpoint) + route
	var response map[string]interface{}
	var err error
	err = GetJson(url, &response)
	err = ExtractData(response, &target)
	return err
}

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return error
	}
	defer func(){
		if err := r.Body.Close(); err:= nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

//extrae la información cuando se recibe encapsulada en una estructura 
//y da manejo a las respuestas que contienen arreglos vacios
func ExtractData(respuesta map[string]interface{}, v interfac{}) error {
	var err error
	if respuesta["success"] == false {
		err = errors.New(fmt.Sprint(respuesta["Data"], respuesta["Message"]))
		panic(err)
	}
	datatype := fmt.Sprintf("%v", respuesta["Data"])
	switch datatype {
	case "map[]","[map[]]": //response vacio
		break
	default:
		err = formatdata.FillStruct(respuesta["Data"],%v)
		respuesta = nil
	}
	return err
}

func ErrorController(c beego.Controller, controller string) {
	if err := recover(); err !=nil {
		logs.Error(err)
		localError := err.(map[string]interface{})
		c.Data["message"] = (beego.AppConfig.String("appname") + "/" + controller + "/" + (localError["funcion"]).(string))
		c.Data["data"] = (localError["err"])
	if status, ok := localError["status"]; ok {
		c.Abort(status.(string))
	} else {
		c.Abort("500")
	}
	}
}
