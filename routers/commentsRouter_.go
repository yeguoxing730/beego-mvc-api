package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego-mvc-api/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
