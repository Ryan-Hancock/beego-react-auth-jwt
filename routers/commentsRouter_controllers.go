package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["authJWT/controllers:AuthController"] = append(beego.GlobalControllerRouter["authJWT/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/auth/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["authJWT/controllers:AuthController"] = append(beego.GlobalControllerRouter["authJWT/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Refresh",
            Router: `/auth/refresh`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["authJWT/controllers:AuthController"] = append(beego.GlobalControllerRouter["authJWT/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Validate",
            Router: `/auth/validate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
