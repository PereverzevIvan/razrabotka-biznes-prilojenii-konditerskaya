package repos_mysql_supplier_component

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"
)

func (repo *repoSupplierComponent) FastestDeliveryComponentPrice(component_id int) (float64, error) {
	var price float64

	rows, err := repo.conn.
		Raw(`
SELECT price as price
FROM supplier_components 
INNER JOIN suppliers ON suppliers.id = supplier_components.supplier_id 
WHERE component_id = @component_id
ORDER BY suppliers.delivery_time_minutes ASC
LIMIT 1`,
			map[string]interface{}{
				"component_id": component_id,
			},
		).
		Rows()

	if err != nil {
		return 0, err
	}

	if !rows.Next() {
		return 0, logic_errors.ErrNoSupplierForComponent
	}

	err = rows.Scan(&price)
	if err != nil {
		return 0, err
	}

	return price, nil
}
