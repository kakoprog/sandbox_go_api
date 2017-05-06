package statuses

import (
	m "github.com/kakoprog/sandbox_go_api/model/todo/status"
	"github.com/kakoprog/sandbox_go_api/server/util"
	"github.com/labstack/echo"
	"net/http"
)

func GetContent(c echo.Context) error {
	id, err := util.ParamToInt(c.Param("id"), 64)
	if err != nil {
		return err
	}

	s, err := m.StatusByID(id.(int64))
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, s)
}

func GetContents(c echo.Context) error {
	s, err := m.Statuses()
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, s)
}
