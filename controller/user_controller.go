package user

import (
	"github.com/cihub/seelog"
	user "github.com/empenguin1186/gin-gorm-tutorial/service"
	"github.com/gin-gonic/gin"
)

// Controller is user controller
type Controller struct {
	service user.Service
}

// Index action: GET /users
func (pc Controller) Index(c *gin.Context) {
	// var s user.Service
	p, err := pc.service.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		seelog.Error("Exception occurred in Getting Users.\n")
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (pc Controller) Create(c *gin.Context) {
	p, err := pc.service.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		seelog.Error("Exception occurred in Createing User.\n")
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
	p, err := pc.service.GetByID(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(400)
		seelog.Error("Exception occurred in Getting User.\n")
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(c *gin.Context) {
	p, err := pc.service.UpdateByID(c.Params.ByName("id"), c)
	if err != nil {
		c.AbortWithStatus(400)
		seelog.Error("Exception occurred in Updating User.\n")
	} else {
		c.JSON(200, p)
	}
}

// Delete  action: DELETE /users/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	err := pc.service.DeleteByID(id)
	if err != nil {
		c.AbortWithStatus(403)
		seelog.Error("Exception occurred in Deleting User.\n")
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
