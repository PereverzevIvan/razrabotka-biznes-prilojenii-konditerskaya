package results_purchased_component

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

type GetAllResults struct {
	TotalRows int64                       `json:"total_rows"`
	Data      []models.PurchasedComponent `json:"Data"`

	TotalPrice int64 `json:"total_price"`
	TotalCount int64 `json:"total_count"`
}
