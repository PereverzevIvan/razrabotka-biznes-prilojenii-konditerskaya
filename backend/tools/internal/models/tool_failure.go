package models

import "time"

type ToolFailure struct {
	ID int `json:"id" gorm:"Column:id"`

	ToolID int   `json:"tool_id" gorm:"Column:tool_id"`
	Tool   *Tool `json:"tool" gorm:"foreignKey:ToolID"`

	MasterID int   `json:"master_id" gorm:"Column:master_id"`
	Master   *User `json:"master" gorm:"foreignKey:MasterID"`

	FailureReasonID int                `json:"failure_type_id" gorm:"Column:failure_reason_id"`
	FailureReason   *ToolFailureReason `json:"failure_type" gorm:"foreignKey:FailureReasonID"`

	// IsFixed bool       `json:"is_fixed" gorm:"Column:is_fixed"`
	FailureAt time.Time  `json:"failure_at" gorm:"Column:failure_at"`
	FixedAt   *time.Time `json:"fixed_at" gorm:"Column:fixed_at"`
}
