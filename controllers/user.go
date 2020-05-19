package controllers

import (
	"authJWT/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

//UserController ...
type UserController struct {
	beego.Controller
}

//Post ...
func (c *UserController) Post() {
	response := struct {
		UserID int64
	}{}

	u := models.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &u)
	if err != nil {
		log.Println(err)
		c.Abort("500")
	}

	id, err := models.GetUserStorage().NewUser(u)
	if err != nil {
		log.Println(err)
		c.Abort("500")
	}

	response.UserID = id
	c.Data["json"] = response
	c.ServeJSON()
}
