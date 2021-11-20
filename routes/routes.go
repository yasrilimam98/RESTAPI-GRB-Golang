package routes

// Import dengan parameter agar bisa banyak import pkg
import (
	"net/http"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()
	// c echo context untuk mengakses request or respon dan error untuk notif ketika jadi kesalahan
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Word!!!")
	})
	return e
}
