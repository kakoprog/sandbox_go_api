package tasks

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/kakoprog/sandbox_go_api/server/util"
	m "github.com/kakoprog/sandbox_go_api/model/todo/task"
	"github.com/kakoprog/sandbox_go_api/repository"
)

func GetContent(c echo.Context) error {
	t, err := getContentByID(c)
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, t)
}

func GetContents(c echo.Context) error {
	t, err := m.Tasks()
	if err != nil {
		return c.NoContent(http.StatusNoContent)
  }

	return c.JSON(http.StatusOK, t)
}

func PostContent(c echo.Context) error {
	t := new(repository.Task)
	if err := c.Bind(t); err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest)
	}

	return saveContent(c, t, http.StatusCreated);
}

func PutContent(c echo.Context) error {
	t, err := getContentByID(c)
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	if err := c.Bind(t); err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest)
	}

	return saveContent(c, t, http.StatusOK);
}

func DeleteContent(c echo.Context) error {
	t, err := getContentByID(c)
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	if err := m.Delete(t); err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func getContentByID(c echo.Context) (*repository.Task, error) {
	id, err := util.ParamToInt(c.Param("id"), 64)
	if err != nil { return nil, err }

	t, err := m.TaskByID(id.(int64))
	if err != nil { return nil, err }

	return t, nil
}

func saveContent(c echo.Context, t *repository.Task, status int) error {
	res, err := m.Save(t)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(status, res)
}
