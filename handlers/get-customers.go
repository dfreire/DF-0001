package handlers

import (
	"net/http"

	"github.com/dfreire/df0001/middleware"
	"github.com/dfreire/df0001/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// http http://localhost:3500/get-customers
func GetCustomers(c echo.Context) error {
	db := c.Get(middleware.DB).(*gorm.DB)

	customers := []model.Customer{}
	err := db.Preload("Role").Find(&customers).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: customers})
}
