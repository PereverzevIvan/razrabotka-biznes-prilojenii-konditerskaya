package tool_params

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
)

type ToolEditParams struct {
	ToolType       *int       `json:"tool_type"`
	DegreeOfWearID *int       `json:"degree_of_wear"`
	SupplierID     *int       `json:"supplier_id"`
	Name           *string    `json:"name"`
	Description    *string    `json:"description"`
	PurchaseDate   *time.Time `json:"purchase_date"`
}

func ToolEditParamsFromGRPC(req *tool.ToolEditRequest) *ToolEditParams {
	params := &ToolEditParams{
		Name:        req.Name,
		Description: req.Description,
	}
	if req.ToolType != nil {
		toolType := int(*req.ToolType)
		params.ToolType = &toolType
	}
	if req.DegreeOfWear != nil {
		degreeOfWear := int(*req.DegreeOfWear)
		params.DegreeOfWearID = &degreeOfWear
	}
	if req.SupplierId != nil {
		supplierID := int(*req.SupplierId)
		params.SupplierID = &supplierID
	}
	if req.PurchaseDate != nil {
		purchaseDate := req.PurchaseDate.AsTime()
		params.PurchaseDate = &purchaseDate
	}

	return params
}
