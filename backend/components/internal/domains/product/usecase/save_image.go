package product_usecase

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v3"
)

func (u *ProductUsecase) SaveImage(
	ctx fiber.Ctx,
	id int,
	image *multipart.FileHeader,
) (string, error) {
	return u.productRepo.SaveImage(ctx, id, image)
}
