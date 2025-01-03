package services_product

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v3"
)

func (service *productService) SaveImage(
	ctx fiber.Ctx,
	id int,
	image *multipart.FileHeader,
) (string, error) {
	return service.productRepo.SaveImage(ctx, id, image)
}
