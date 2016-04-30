package commands

import (
	"github.com/dfreire/df0001/model"
	"github.com/jinzhu/gorm"
	"labix.org/v2/mgo/bson"
)

type SignupCustomerWithWineCommentRequestData struct {
	Name         string        `json:"name,omitempty"`
	Email        string        `json:"email"`
	RoleId       string        `json:"roleId"`
	WineComments []WineComment `json:"wineComments"`
}

type WineComment struct {
	WineId   string `json:"wineId"`
	WineYear int    `json:"wineYear"`
	Comment  string `json:"comment"`
}

func SignupCustomerWithWineComment(db *gorm.DB, reqData SignupCustomerWithWineCommentRequestData) error {
	customerToFind := model.Customer{
		Email: reqData.Email,
	}

	customerId := bson.NewObjectId().Hex()

	customerToCreate := model.Customer{
		ID:    customerId,
		Email: reqData.Email,
	}

	err := db.Where(customerToFind).FirstOrCreate(&customerToCreate).Error
	if err != nil {
		return err
	}

	customerToUpdate := model.Customer{
		Name:   reqData.Name,
		RoleId: reqData.RoleId,
	}

	err = db.Model(&customerToUpdate).Updates(customerToUpdate).Error
	if err != nil {
		return err
	}

	for _, comment := range reqData.WineComments {
		commentToFind := model.WineComment{
			CustomerId: customerId,
			WineId:     comment.WineId,
			WineYear:   comment.WineYear,
		}

		commentToCreate := model.WineComment{
			ID: bson.NewObjectId().Hex(),
		}

		err := db.Where(commentToFind).FirstOrCreate(&commentToCreate).Error
		if err != nil {
			return err
		}

		commentToUpdate := model.WineComment{
			Comment: comment.Comment,
		}
		err = db.Model(&commentToUpdate).Updates(commentToUpdate).Error
		if err != nil {
			return err
		}
	}

	return nil
}
