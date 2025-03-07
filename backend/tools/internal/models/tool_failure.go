package models

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func (*ToolFailure) TableName() string {
	return "tool_failures"
}

func ToolFailureToGRPC(tf *ToolFailure) *tool_failure.ToolFailure {
	if tf == nil {
		return nil
	}

	res := &tool_failure.ToolFailure{
		Id:              int32(tf.ID),
		ToolId:          int32(tf.ToolID),
		MasterId:        int32(tf.MasterID),
		FailureReasonId: int32(tf.FailureReasonID),
		FailureAt:       timestamppb.New(tf.FailureAt),
		FixedAt:         nil,
	}

	if tf.FixedAt != nil {
		res.FixedAt = timestamppb.New(*tf.FixedAt)
	}

	return res
}

func GRPCToToolFailure(tf *tool_failure.ToolFailure) *ToolFailure {
	res := &ToolFailure{
		ID:              int(tf.Id),
		ToolID:          int(tf.ToolId),
		MasterID:        int(tf.MasterId),
		FailureReasonID: int(tf.FailureReasonId),
		FailureAt:       tf.FailureAt.AsTime(),
		FixedAt:         nil,
	}
	if tf.FixedAt != nil {
		fixedAt := tf.FixedAt.AsTime()
		res.FixedAt = &fixedAt
	}

	return res
}
