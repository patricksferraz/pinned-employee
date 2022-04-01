package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-employee/utils"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base  `json:",inline" valid:"-"`
	Name  *string `json:"name" gorm:"column:name;not null" valid:"required"`
	Token *string `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
}

func NewEmployee(name *string) (*Employee, error) {
	token := primitive.NewObjectID().Hex()
	e := Employee{
		Name:  name,
		Token: &token,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Employee) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

type SearchEmployees struct {
	Pagination `json:",inline" valid:"-"`
}

func NewSearchEmployees(pagination *Pagination) (*SearchEmployees, error) {
	e := SearchEmployees{}
	e.PageToken = pagination.PageToken
	e.PageSize = pagination.PageSize

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}
