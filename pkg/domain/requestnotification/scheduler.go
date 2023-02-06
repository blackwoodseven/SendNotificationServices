package requestnotification

import (
	"encoding/hex"
	"log"

	"github.com/jasonlvhit/gocron"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/utility"
)

var updatefrequencyinhours uint64 = 2

func (d *domain) UpdateHashkey() {
	models, err := d.NotificationRequest.GetAll()

	if err == nil && models != nil && len(models) > 0 {
		key := utility.GetRamdomKey()
		for _, model := range models {
			model.AuthToken = utility.Encrypt(hex.EncodeToString([]byte(key)), model.AuthToken)
			if err := d.NotificationRequest.UpdateHashkey(model); err != nil {
				log.Println("Error while trying to update:", err)
			}
		}
	}

}

func (d *domain) Scheduler() {
	s := gocron.NewScheduler()
	s.Every(updatefrequencyinhours).Hours().Do(d.UpdateHashkey)
	<-s.Start()
}
