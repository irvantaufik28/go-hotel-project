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
