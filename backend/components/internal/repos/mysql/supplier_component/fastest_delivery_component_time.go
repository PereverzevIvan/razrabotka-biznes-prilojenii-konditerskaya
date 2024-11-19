package repos_mysql_supplier_component

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/logic_errors"

func (repo *repoSupplierComponent) FastestDeliveryComponentTime(component_id int) (int, error) {
	var delivery_time_minutes int

	rows, err := repo.conn.
		Raw(`
SELECT IFNULL(MIN(delivery_time_minutes), -1) as delivery_time_minutes
FROM supplier_components 
INNER JOIN suppliers ON suppliers.id = supplier_components.supplier_id 
WHERE component_id = @component_id`,
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

	err = rows.Scan(&delivery_time_minutes)
	if err != nil {
		return 0, err
	}

	if delivery_time_minutes == -1 {
		return 0, logic_errors.ErrNoSupplierForComponent
	}

	return delivery_time_minutes, nil

}
