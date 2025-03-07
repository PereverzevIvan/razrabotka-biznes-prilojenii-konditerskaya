package tool_params

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
)

type GetAllParams struct {
	Sort *EGetAllSort `json:"sort"`

	ToolType        *int                  `json:"tool_type"`
	DegreeOfWear    *models.EDegreeOfWear `json:"degree_of_wear"`
	SupplierID      *int                  `json:"supplier_id"`
	Name            *string               `json:"name"`
	OnlyServiceable bool                  `json:"only_serviceable"`
	// PurchaseDateFrom *time.Time            `json:"purchase_date"`
	// PurchaseDateTo   *time.Time            `json:"purchase_date_to"`
}

func GetAllParamsFromGRPC(req *tool.ToolGetAllRequest) *GetAllParams {
	params := &GetAllParams{
		Sort: (*EGetAllSort)(req.Sort),

		Name: (*string)(req.Name),
	}

	if req.ToolType != nil {
		toolType := int(*req.ToolType)
		params.ToolType = &toolType
	}
	if req.DegreeOfWear != nil {
		degreeOfWear := models.EDegreeOfWear(*req.DegreeOfWear)
		params.DegreeOfWear = &degreeOfWear
	}
	if req.SupplierId != nil {
		supplierID := int(*req.SupplierId)
		params.SupplierID = &supplierID
	}
	if req.OnlyServiceable != nil {
		params.OnlyServiceable = *req.OnlyServiceable
	}

	return params
}

type EGetAllSort string

const (
	KDefaultSort        EGetAllSort = KSortByPurchaseDate
	KSortByToolType     EGetAllSort = "tool_type"
	KSortByDegreeOfWear EGetAllSort = "degree_of_wear_id"
	KSortByName         EGetAllSort = "name"
	KSortByPurchaseDate EGetAllSort = "purchase_date"
)
