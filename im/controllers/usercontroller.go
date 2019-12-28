package controllers

import "im/models"

// UserController handles WebSocket requests.
type UserController struct {
	baseController
}

// Get method handles GET requests for UserController.
func (this *UserController) Login() {
	// Safe check.
	//uname := this.GetString("uname")

	this.Data["json"] = models.Result{}
	this.ServeJSON()

}
