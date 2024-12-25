package delivery

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lipzy13/dakas-backend.git/internal/delivery/http"
	"github.com/lipzy13/dakas-backend.git/internal/service"
)

func SetupRoutes(router *httprouter.Router, service service.GerobakService) {
	gerobakHandler := http.NewGerobakHandler(service)

	router.POST("/gerobak", gerobakHandler.CreateGerobak)
	router.GET("/gerobak", gerobakHandler.GetAllGerobaks)
	router.GET("/gerobak/:id", gerobakHandler.GetGerobakById)
}
