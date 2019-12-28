package controllers

import (
	"im/models"
	"time"
)

// UserController handles WebSocket requests.
type UserController struct {
	baseController
}

// Get method handles GET requests for UserController.
func (this *UserController) Login() {
	// Safe check.
	//uname := this.GetString("uname")
	result := models.Result{}
	result.Success = true

	data := models.LoginData{}
	data.Id = time.Now().Nanosecond()
	result.Data=data
	result.Data=data
	this.Data["json"] = result

	this.ServeJSON()

}
