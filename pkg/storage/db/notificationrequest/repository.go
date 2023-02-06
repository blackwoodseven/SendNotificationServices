package notificationrequest

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(n *NotificationRequest) error
	GetAll() ([]*NotificationRequest, error)
	UpdateHashkey(n *NotificationRequest) error
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
