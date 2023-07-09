package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/irvantaufik28/go-hotel-project/internal/config"
	httpcontroler "github.com/irvantaufik28/go-hotel-project/internal/controller/http/v1"
	"github.com/irvantaufik28/go-hotel-project/internal/repository"
	"github.com/irvantaufik28/go-hotel-project/internal/service"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	port := os.Getenv("PORT")

	roomRepository := repository.NewRoomRepository(config.DB)
	roomService := service.NewRoomService(roomRepository)
	roomController := httpcontroler.NewRoomController(roomService)

	roomController.Route(r)
	runWithPort := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(runWithPort)
}
