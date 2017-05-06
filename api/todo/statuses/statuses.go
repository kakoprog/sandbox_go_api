package statuses

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/kakoprog/sandbox_go_api/server/util"
	m "github.com/kakoprog/sandbox_go_api/model/todo/status"
)

func GetContent(c echo.Context) error {
	id, err := util.ParamToInt(c.Param("id"), 64)
	if err != nil { return err }

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
