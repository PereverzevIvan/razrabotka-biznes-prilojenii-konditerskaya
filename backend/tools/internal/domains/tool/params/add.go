package tool_params

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
)

type ToolAddParams struct {
	ToolType       int       `json:"tool_type"`
	DegreeOfWearID int       `json:"degree_of_wear"`
	SupplierID     int       `json:"supplier_id"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	PurchaseDate   time.Time `json:"purchase_date"`
}

func (params *ToolAddParams) Validate() []string {
	var errors []string
	if params.ToolType <= 0 {
		errors = append(errors, "tool_type is required")
	}
	if params.DegreeOfWearID == 0 {
		errors = append(errors, "degree_of_wear is required")
	}
	if params.SupplierID == 0 {
		errors = append(errors, "supplier_id is required")
	}
	if params.Name == "" {
		errors = append(errors, "name is required")
	}
	if params.PurchaseDate.IsZero() {
		errors = append(errors, "purchase_date is required")
	}
	return errors
}

func ToolAddParamsFromGRPC(req *tool.ToolAddRequest) *ToolAddParams {
	return &ToolAddParams{
		ToolType:       int(req.ToolType),
		DegreeOfWearID: int(req.DegreeOfWear),
		SupplierID:     int(req.SupplierId),
		Name:           req.Name,
		Description:    req.Description,
		PurchaseDate:   req.PurchaseDate.AsTime(),
	}
}
