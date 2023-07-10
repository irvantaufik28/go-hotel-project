package httpcontroler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvantaufik28/go-hotel-project/domain"
	"gorm.io/gorm"
)

type RoomController struct {
	RoomService domain.RoomService
}

func NewRoomController(roomService domain.RoomService) RoomController {
	return RoomController{RoomService: roomService}
}

func (controller *RoomController) Route(app *gin.Engine) {
	route := app.Group("/api/v1")
	route.GET("/room", controller.FindAll)
	route.GET("/room/:id", controller.FindById)
	route.POST("/room", controller.Create)
	route.PUT("/room/:id", controller.Update)
	route.DELETE("/room/:id", controller.Delete)
}

func (controller *RoomController) FindAll(c *gin.Context) {
	res, err := controller.RoomService.GetAllRoom()

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "room not found"})
			return
		default:
			{
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (controller *RoomController) FindById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	res, err := controller.RoomService.GetById(idInt)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "room not found"})
			return
		default:
			{
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (controller *RoomController) Create(c *gin.Context) {
	var payload domain.RoomReq

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.RoomService.Create(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": payload})
}

func (controller *RoomController) Update(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var payload domain.RoomReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.RoomService.Update(idInt, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room updated successfully"})
}

func (controller *RoomController) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = controller.RoomService.Delete(idInt)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}
