package handlers

import (
	"net/http"

	"github.com/dfreire/df0001/middleware"
	"github.com/dfreire/df0001/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// http http://localhost:3500/get-wine-comments-by-customer-id?customerId=customer-1
func GetWineCommentsByCustomerId(c echo.Context) error {
	db := c.Get(middleware.DB).(*gorm.DB)

	customerId := c.QueryParam("customerId")

	comments := []model.WineComment{}
	err := db.Where("customer_id = ?", customerId).Preload("Customer").Preload("Customer.Role").Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: comments})

}
