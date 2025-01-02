package repos_rest_component

import (
	"fmt"
	"net/http"
	"strings"
)

func (repo *componentRepo) ConsumeProductComponents(product_id int) error {
	url := fmt.Sprintf(
		"%s/%s",
		repo.baseUrl,
		fmt.Sprintf(
			"api/products/%d/make",
			product_id,
		),
	)

	resp, err := http.Post(
		url,
		"application/json",
		strings.NewReader(""),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"status code not ok: %d, body: %s",
			resp.StatusCode,
			resp.Body,
		)
	}

	return nil
}
