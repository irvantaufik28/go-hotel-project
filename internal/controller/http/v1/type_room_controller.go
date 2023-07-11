package httpcontroler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irvantaufik28/go-hotel-project/domain"
	"gorm.io/gorm"
)

type TypeRoomController struct {
	TypeRoomService domain.TypeRoomService
}

func NewTypeRoomController(typeRoomService domain.TypeRoomService) TypeRoomController {
	return TypeRoomController{TypeRoomService: typeRoomService}
}

func (controller *TypeRoomController) Route(app *gin.Engine) {
	route := app.Group("/api/v1")
	route.GET("/type", controller.FindAllTypeRoom)

}

func (controller *TypeRoomController) FindAllTypeRoom(c *gin.Context) {
	res, err := controller.TypeRoomService.GetAllTypeRoom()

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "room not found"})
			return
		default:
			{
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error()})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
