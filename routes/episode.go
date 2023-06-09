package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func EpisodeRoutes(e *echo.Group) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(EpisodeRepository)

	e.GET("/episodes", h.FindEpisode)
	e.GET("/film/:id/episode", h.FindEpisodeByFilm)
	e.GET("/film/:id/episode/:ide", h.GetEpisodeByFilm)
	e.GET("/episode/:id", h.GetEpisode)
	e.POST("/episode", middleware.Auth(middleware.UploadFile(h.CreateEpisode)))
	e.DELETE("/episode/:id", middleware.Auth(h.DeleteEpisode))
	e.PATCH("/episode/:id", middleware.Auth(h.UpdateEpisode))
	// }
}
