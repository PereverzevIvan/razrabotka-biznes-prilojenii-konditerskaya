package params_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

type ToolGetAllParams struct {
	Sort *EToolGetAllSort `json:"sort"`

	ToolType     *int                  `json:"tool_type"`
	DegreeOfWear *models.EDegreeOfWear `json:"degree_of_wear"`
	SupplierID   *int                  `json:"supplier_id"`
	Name         *string               `json:"name"`
	// PurchaseDateFrom *time.Time            `json:"purchase_date"`
	// PurchaseDateTo   *time.Time            `json:"purchase_date_to"`
}

type EToolGetAllSort string

const (
	KDefaultSort        EToolGetAllSort = KSortByPurchaseDate
	KSortByToolType     EToolGetAllSort = "tool_type"
	KSortByDegreeOfWear EToolGetAllSort = "degree_of_wear_id"
	KSortByName         EToolGetAllSort = "name"
	KSortByPurchaseDate EToolGetAllSort = "purchase_date"
)
