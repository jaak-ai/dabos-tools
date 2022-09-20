package transclude

import (
	"github.com/gofiber/fiber/v2"
)

func PopulateFromRequest(model interface{}, ctx *fiber.Ctx) error {
	err := ctx.BodyParser(model)
	if err != nil {
		return err
	}

	return nil
}
