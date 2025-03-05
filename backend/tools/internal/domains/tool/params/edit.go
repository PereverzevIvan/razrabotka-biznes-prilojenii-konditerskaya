package tool_params

import "time"

type ToolEditParams struct {
	ToolType       *int       `json:"tool_type"`
	DegreeOfWearID *int       `json:"degree_of_wear"`
	SupplierID     *int       `json:"supplier_id"`
	Name           *string    `json:"name"`
	Description    *string    `json:"description"`
	PurchaseDate   *time.Time `json:"purchase_date"`
}
