package notificationrequest

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(n *NotificationRequest) error
	GetAll() ([]*NotificationRequest, error)
	UpdateHashkey(n *NotificationRequest) error
	ComponentExists(id string, name string) bool
}

type service struct {
	db *gorm.DB
}

func InitDB(connection *gorm.DB) (Repository, error) {
	err := connection.AutoMigrate(&NotificationRequest{})
	if err != nil {
		return nil, err
	}

	return &service{
		db: connection,
	}, nil
}

func (r *service) Create(n *NotificationRequest) error {
	return r.db.Create(&n).Error
}

func (r *service) GetAll() ([]*NotificationRequest, error) {
	var notificationrequest []*NotificationRequest
	err := r.db.
		Preload(clause.Associations).
		Find(&notificationrequest).Error
	if err != nil {
		return nil, err
	}
	return notificationrequest, nil
}

func (r *service) UpdateHashkey(n *NotificationRequest) error {
	return r.db.Save(&n).Error
}

func (r *service) ComponentExists(componentid string, componentname string) bool {
	var count int64
	err := r.db.
		Model(&NotificationRequest{}).
		Where("componentid = ?", componentid).
		Where("componentname = ?", componentname).
		Count(&count).Error
	fmt.Println("Compid : ", componentid, " compname : ", componentname, " count : ", count)
	if err != nil {
		return false
	}

	return count > 0
}
