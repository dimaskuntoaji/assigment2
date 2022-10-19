package response

import (
	"time"
)

type Order struct {
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items"`
	ID           uint      `json:"orderId"`
	OrderedAt    time.Time `json:"orderedAt"`
}