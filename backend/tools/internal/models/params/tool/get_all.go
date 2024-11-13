package params_tool

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
)

type GetAllParams struct {
	Sort *EGetAllSort `json:"sort"`

	ToolType     *int                  `json:"tool_type"`
	DegreeOfWear *models.EDegreeOfWear `json:"degree_of_wear"`
	SupplierID   *int                  `json:"supplier_id"`
	Name         *string               `json:"name"`
	// PurchaseDateFrom *time.Time            `json:"purchase_date"`
	// PurchaseDateTo   *time.Time            `json:"purchase_date_to"`
}

type EGetAllSort string

const (
	KDefaultSort        EGetAllSort = KSortByPurchaseDate
	KSortByToolType     EGetAllSort = "tool_type"
	KSortByDegreeOfWear EGetAllSort = "degree_of_wear_id"
	KSortByName         EGetAllSort = "name"
	KSortByPurchaseDate EGetAllSort = "purchase_date"
)
