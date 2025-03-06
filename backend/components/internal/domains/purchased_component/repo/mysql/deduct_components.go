package purchased_component_repo_mysql

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
)

func (r *PurchasedComponentRepo) DeductComponents(component_id int, deduct_count int) error {

	result := r.db.
		Debug().
		Exec(`
WITH cte AS (
    SELECT 
        id,
    	shelf_life,
        SUM(quantity) OVER (
            PARTITION BY component_id
            ORDER BY shelf_life
            ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW
        ) AS running_sum
    FROM purchased_components
    WHERE 
        quantity > 0 
        AND component_id = @component_id
   	ORDER BY shelf_life
),
cte_begin AS (
	SELECT * 
    FROM cte
    WHERE running_sum < @deduct_count
),
cte_end AS (
	SELECT * 
    FROM cte
    WHERE running_sum >= @deduct_count
    LIMIT 1
)


-- Update statement
UPDATE purchased_components
SET quantity = CASE 
	WHEN id IN (SELECT id FROM cte_begin) THEN 0
	WHEN id IN (SELECT id FROM cte_end) THEN quantity - (@deduct_count - (SELECT IFNULL(SUM(quantity), 0) FROM cte_begin))
END
WHERE id IN (
    SELECT id FROM cte_begin 
    UNION ALL
    SELECT id FROM cte_end 
)
AND EXISTS (SELECT 1 FROM cte_end);
		`,
			map[string]interface{}{
				"component_id": component_id,
				"deduct_count": deduct_count,
			},
		)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return logic_errors.ErrNotEnoughPurchasedComponentsToDeduct
	}

	return nil
}
