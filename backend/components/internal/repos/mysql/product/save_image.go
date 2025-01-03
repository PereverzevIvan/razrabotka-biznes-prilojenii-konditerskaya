package repos_mysql_product

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
)

const base_path = "images/products"

func (repo *productRepo) SaveImage(
	ctx fiber.Ctx,
	id int,
	image *multipart.FileHeader,
) (string, error) {

	image_extension := filepath.Ext(image.Filename)
	image_save_name := fmt.Sprintf("%d%s", id, image_extension)

	full_path := fmt.Sprintf(
		"%s/%s",
		base_path,
		image_save_name,
	)

	err := ctx.SaveFile(image, full_path)
	if err != nil {
		return "", err
	}

	return full_path, nil
}