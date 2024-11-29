package params_order

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

type CreateParams struct {
	AuthorId   int    `json:"-"`
	AuthorRole string `json:"-"`

	ProductID int `json:"product_id"`

	CustomerID int  `json:"customer_id"`
	ManagerID  *int `json:"-"`

	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Cost        *float64 `json:"cost"`

	PlannedCompletionAt *time.Time `json:"planned_completion_at"`

	// ExampleImages []string `json:"example_images"`
}

func (p *CreateParams) Validate() []string {
	errs := []string{}

	if p.ProductID == 0 {
		errs = append(errs, "product_id is required")
	}

	if p.Name == "" {
		errs = append(errs, "name is required")
	}

	if p.CustomerID == 0 {
		errs = append(errs, "customer_id is required")
	}

	if p.AuthorRole == models.KRoleClientManager.Name() && p.ManagerID == nil {
		errs = append(errs, "manager_id is not set for author client manager")
	}

	return errs
}
