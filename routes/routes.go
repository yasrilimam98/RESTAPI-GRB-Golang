package routes

// Import dengan parameter agar bisa banyak import pkg
import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yasrilimam98/grb-restapi/controllers"
)

func Init() *echo.Echo {

	e := echo.New()
	// c echo context untuk mengakses request or respon dan error untuk notif ketika jadi kesalahan
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Word!!!")
	})
	e.GET("/karyawan", controllers.FetchAllKaryawan)
	e.POST("/karyawan", controllers.StoreKaryawan)
	e.PUT("/karyawan", controllers.UpdateKaryawan)
	e.DELETE("/karyawan",controllers.DeleteKaryawan)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	
	return e
}
