package delivery

import (
	"assyarif-backend-web-go/domain"

	"github.com/gofiber/fiber/v2"
)

type InHandler struct {
	InUC domain.InUseCase
}

func NewInHandler(c *fiber.App, das domain.InUseCase) {
	handler := &InHandler{
		InUC: das,
	}
	api := c.Group("/stuff")
	in := api.Group("/in")
	in.Get("/", handler.ShowIns)
	in.Get("/:id", handler.ShowInById)
	in.Post("/", handler.AddIn)
	in.Put("/:id", handler.EditInById)
	in.Delete("/:id", handler.DeleteInById)

	last := api.Group("/last")
	inLast := last.Group("/in")
	inLast.Get("/", handler.GetLastInNumber)

	period := api.Group("/period/in")
	period.Get("/", handler.GetInsByPeriod)
}

func (t *InHandler) ShowIns(c *fiber.Ctx) error {
	res, er := t.InUC.ShowIns(c.Context())
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success get data",
	})
}

func (t *InHandler) ShowInById(c *fiber.Ctx) error {
	id := c.Params("id")
	res, er := t.InUC.ShowInById(c.Context(), id)
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success get data",
	})
}

func (t *InHandler) GetLastInNumber(c *fiber.Ctx) error {
	res, er := t.InUC.ShowInLastNumber(c.Context())
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success get data",
	})
}

func (t *InHandler) AddIn(c *fiber.Ctx) error {
	var in domain.In
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	res, er := t.InUC.AddIn(c.Context(), in)
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"success": true,
		"data":    res,
		"message": "Success create data",
	})
}

func (t *InHandler) EditInById(c *fiber.Ctx) error {
	var in domain.In
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	res, er := t.InUC.EditInById(c.Context(), in)
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success update data",
	})
}

func (t *InHandler) DeleteInById(c *fiber.Ctx) error {
	id := c.Params("id")
	er := t.InUC.DeleteInById(c.Context(), id)
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    nil,
		"message": "Success delete data",
	})
}

func (t *InHandler) GetInsByPeriod(c *fiber.Ctx) error {
	res, er := t.InUC.GetInsByPeriod(c.Context())
	if er != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": er.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Success get data",
	})
}
