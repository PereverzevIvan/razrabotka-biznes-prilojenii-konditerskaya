package models

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Tool struct {
	ID int `json:"id" gorm:"Column:id"`

	ToolTypeID int       `json:"tool_type_id" gorm:"Column:type_id"`
	ToolType   *ToolType `json:"tool_type" gorm:"foreignKey:ToolTypeID"`

	SupplierID int   `json:"supplier_id" gorm:"Column:supplier_id"`
	Supplier   *User `json:"supplier" gorm:"foreignKey:SupplierID"`

	DegreeOfWearID int           `json:"degree_of_wear_id" gorm:"Column:degree_of_wear_id"`
	DegreeOfWear   *DegreeOfWear `json:"degree_of_wear" gorm:"foreignKey:DegreeOfWearID"`

	Name        string  `json:"name" gorm:"Column:name"`
	Description *string `json:"description" gorm:"Column:description"`

	PurchaseDate time.Time `json:"purchase_date" gorm:"Column:purchase_date"`
}

func (Tool) TableName() string {
	return "tools"
}

func ToolToGRPC(model *Tool) *tool.Tool {
	return &tool.Tool{
		Id:         int32(model.ID),
		ToolTypeId: int32(model.ToolTypeID),
		// ToolType:       ToolTypeToGRPC(tool.ToolType),
		SupplierId: int32(model.SupplierID),
		// Supplier:       UserToGRPC(tool.Supplier),
		DegreeOfWearId: int32(model.DegreeOfWearID),
		// DegreeOfWear:   DegreeOfWearToGRPC(tool.DegreeOfWear),
		Name:         model.Name,
		Description:  model.Description,
		PurchaseDate: timestamppb.New(model.PurchaseDate),
	}
}
