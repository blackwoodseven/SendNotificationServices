package notificationrequest

import (
	"time"

	"github.com/gofrs/uuid"

	"gorm.io/gorm"
)

// A CoffeeDateRegistry is an registry for a user, of its availability to participate on a CoffeeDate
// given week of a given year.
type NotificationRequest struct {
	RequestID     *uuid.UUID     `gorm:"column:requestid;type:UUID;primary_key"`
	ComponentID   string         `gorm:"column:componentid;type:varchar(255)"`
	ComponentName string         `gorm:"column:componentname;type:varchar(255);"`
	AuthToken     string         `gorm:"column:authtoken;type:varchar(1000);"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (n *NotificationRequest) BeforeCreate(tx *gorm.DB) error {
	if n.RequestID == nil {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		n.RequestID = &uuid
	}
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()

	return nil
}
