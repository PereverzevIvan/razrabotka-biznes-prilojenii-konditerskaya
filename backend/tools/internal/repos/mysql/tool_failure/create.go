package repos_mysql_tool_failure

import (
	"fmt"
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	params_tool_failure "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models/params/tool_failure"
)

func (repo *toolFailureRepo) Create(
	params *params_tool_failure.CreateParams,
) (*models.ToolFailure, error) {

	var failure_reason_id int
	if params.FailureReasonID != nil {
		failure_reason_id = *params.FailureReasonID
	}

	if failure_reason_id == 0 && params.FailureReason != "" {
		err := repo.conn.
			Debug().
			Exec(`
				INSERT INTO `+models.ToolFailureReason{}.TableName()+` (name)
				VALUES ( ? ) ON DUPLICATE KEY UPDATE name=name;
			`, params.FailureReason).
			Error

		if err != nil {
			return nil, err
		}

		err = repo.conn.
			Debug().
			Raw(`
				SELECT id 
				FROM `+models.ToolFailureReason{}.TableName()+`
				WHERE name = ? ;
			`, params.FailureReason).
			Row().
			Scan(&failure_reason_id)

		if err != nil {
			return nil, err
		}
	}

	if failure_reason_id == 0 {
		return nil, fmt.Errorf("failure reason not found")
	}

	tool_failure := &models.ToolFailure{
		ToolID:          params.ToolID,
		MasterID:        params.MasterID,
		FailureReasonID: failure_reason_id,
		FailureAt:       time.Now(),
	}

	if params.FailureAt != nil {
		tool_failure.FailureAt = *params.FailureAt
	}

	err := repo.conn.
		Create(&tool_failure).
		Error

	if err != nil {
		return nil, err
	}

	return tool_failure, nil
}
