package models

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
)

type Tool struct {
	ID int `json:"id"`

	ToolTypeID int       `json:"tool_type_id"`
	ToolType   *ToolType `json:"tool_type"`

	SupplierID int   `json:"supplier_id"`
	Supplier   *User `json:"supplier"`

	DegreeOfWearID int           `json:"degree_of_wear_id"`
	DegreeOfWear   *DegreeOfWear `json:"degree_of_wear"`

	Name        string  `json:"name"`
	Description *string `json:"description"`

	PurchaseDate time.Time `json:"purchase_date"`
}

func GRPCToTool(t *tool.Tool) *Tool {
	return &Tool{
		ID:             int(t.Id),
		ToolTypeID:     int(t.ToolTypeId),
		SupplierID:     int(t.SupplierId),
		DegreeOfWearID: int(t.DegreeOfWearId),
		Name:           t.Name,
		Description:    t.Description,
		PurchaseDate:   t.PurchaseDate.AsTime(),
	}
}

func ToolFromGRPC(t *tool.Tool) *Tool {
	return &Tool{
		ID:             int(t.Id),
		ToolTypeID:     int(t.ToolTypeId),
		SupplierID:     int(t.SupplierId),
		DegreeOfWearID: int(t.DegreeOfWearId),
		Name:           t.Name,
		Description:    t.Description,
		PurchaseDate:   t.PurchaseDate.AsTime(),
	}
}
