package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ClubsHandler struct {
	db                  *gorm.DB
	clubRepository      ClubRepository
	clubunionRepository ClubUnionRepository
	responseHandler     ResponseHandler
	requestHandler      RequestHandler
}

func InitClubsHandler(app *App) *ClubsHandler {
	h := &ClubsHandler{
		app.Db(),
		NewClubRepository(app.Db()),
		NewClubUnionRepository(app.Db()),
		app.ResponseHandler(),
		app.requestHandler,
	}

	authMiddleware := NewAuthMiddleware(app)

	v1 := app.engine.Group("/v1")
	{
		v1.Use(authMiddleware).GET("/clubs", h.List)
		v1.Use(authMiddleware).GET("/clubs/:id", h.Get)
	}

	return h
}
func (h *ClubsHandler) List(c *gin.Context) {
	user := c.MustGet("user").(User)

	clubs, err := h.clubRepository.FindByUserId(int(user.ID))
	if err == nil {
		manager, err := h.clubunionRepository.FindClubsByUserId(int(user.ID))
		if err == nil {
			clubs = append(clubs, manager...)
		}
		h.responseHandler.JSON(c, http.StatusOK, clubs)
	} else {
		h.responseHandler.InternalServerError(c)
		return
	}
}

func (h *ClubsHandler) Get(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	club, err := h.clubRepository.FindById(id)
	if err != nil {
		h.responseHandler.NotFound(c)
		return
	}
	var count int
	h.db.Table("club_users").Where("club_id = ? and flag = '3'", id).Count(&count)
	club.ManagerCount = count
	h.responseHandler.JSON(c, http.StatusOK, club)
}
