package repos_mysql_purchased_component

func (repo *purchasedComponentRepo) CalcPriceOfRequiredCount(component_id int, required_count int) (float64, error) {
	var purchased_price float64

	err := repo.conn.
		Raw(`
WITH cte AS (
    SELECT 
        id,
        quantity,
        purchase_price,
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
),
cte_begin AS (
    SELECT * 
    FROM cte
    WHERE running_sum < @required_count
),
cte_end AS (
    SELECT * 
    FROM cte
    WHERE running_sum >= @required_count
    ORDER BY shelf_life -- Added ORDER BY to ensure we get the correct row
    LIMIT 1
)

SELECT 
    COALESCE(
        SUM(cte_begin.quantity * cte_begin.purchase_price) +
        (@required_count - (SELECT SUM(quantity) FROM cte_begin)) * (SELECT purchase_price FROM cte_end),
        @required_count * (SELECT purchase_price FROM cte_end),
        -1
    ) AS purchased_price
FROM cte_begin;`,
			map[string]interface{}{
				"component_id":   component_id,
				"required_count": required_count,
			},
		).
		Scan(&purchased_price).
		Error

	if err != nil {
		return 0, err
	}

	return purchased_price, nil
}
