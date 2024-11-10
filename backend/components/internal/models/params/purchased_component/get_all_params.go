package params_purchased_component

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models/params"
)

type GetAllParams struct {
	*params.PaginateParams

	ComponentName    *string `query:"component_name"`
	ComponentArticle *string `query:"component_article"`

	ShelfLifeFrom *time.Time `query:"shelf_life_from"`
	ShelfLifeTo   *time.Time `query:"shelf_life_to"`

	sort *EGetAllSort `query:"sort"`
}

func (p *GetAllParams) Validate() []string {
	errs := make([]string, 0)

	if p.PaginateParams == nil {
		p.PaginateParams = &params.PaginateParams{}
	}

	if p.ShelfLifeFrom != nil && p.ShelfLifeTo != nil {
		if p.ShelfLifeFrom.After(*p.ShelfLifeTo) {
			errs = append(errs, "shelf_life_from must be less than shelf_life_to")
		}
	}

	return errs
}

func (p *GetAllParams) GetSort() EGetAllSort {
	if p.sort == nil {
		return KGetAllSortDefault
	}

	switch *p.sort {
	case KGetAllSortByQuantity:
		return KGetAllSortByQuantity
	case KGetAllSortByShelfLife:
		return KGetAllSortByShelfLife
	}

	return KGetAllSortDefault
}

type EGetAllSort string

const (
	KGetAllSortDefault     EGetAllSort = KGetAllSortByShelfLife
	KGetAllSortByQuantity  EGetAllSort = "quantity"
	KGetAllSortByShelfLife EGetAllSort = "shelf_life"
)
