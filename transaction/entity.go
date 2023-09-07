package transaction

import (
	"time"

	"github.com/vctrthe/api-go/campaign"
	"github.com/vctrthe/api-go/user"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentUrl string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
