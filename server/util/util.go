package util

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func ParamToInt(param string, bit int) (interface{}, error) {
	result, err := strconv.ParseInt(param, 10, bit)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest)
	}

	return result, nil
}
