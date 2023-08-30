package SampleController

import (
	SampleDTO "go_echo_example/src/sample/dtos"
	SampleModel "go_echo_example/src/sample/repository/models"
	SampleSupabaseService "go_echo_example/src/sample/service"
	"net/http"

	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	body := new(SampleModel.Sample)

	if err := c.Bind(body); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	data := &SampleDTO.SampleDTO{
		Title: body.Title,
	}

	resut := SampleSupabaseService.Insert(data)

	response := map[string]interface{}{
		"data": resut,
	}
	return c.JSON(http.StatusOK, response)
}
