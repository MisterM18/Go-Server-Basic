
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/your/package/repository"
)

type UserHandler struct {
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) GetAllUser(c echo.Context) error {
	users := h.Repo.GetAllUsers()

	return c.JSON(http.StatusOK, users)
}


