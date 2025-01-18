package delivery

import (
	"assyarif-backend-web-go/domain"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OrderUC domain.OrderUseCase
}

func NewOrderHandler(c *fiber.App, das domain.OrderUseCase) {
	handler := &OrderHandler{
		OrderUC: das,
	}
	api := c.Group("/order")

	stuff := api.Group("/stuff")
	stuff.Get("/", handler.ShowOrders)
	stuff.Get("/:id", handler.ShowOrderById)
	stuff.Post("/", handler.AddOrder)
	stuff.Put("/:id", handler.EditOrderById)
	stuff.Delete("/:id", handler.DeleteOrderById)

	stuff.Post("/multiple", handler.AddOrders)
	stuff.Delete("/multiple/delete", handler.DeleteOrdersById)

	private := api.Group("/private")
	private.Get("/outlet/:id", handler.ShowOrderByOutletId)
}

func (t *OrderHandler) ShowOrders(c *fiber.Ctx) error {
	res, er := t.OrderUC.ShowOrders(c.Context())
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

func (t *OrderHandler) ShowOrderById(c *fiber.Ctx) error {
	id := c.Params("id")
	res, er := t.OrderUC.ShowOrderById(c.Context(), id)
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

func (t *OrderHandler) AddOrder(c *fiber.Ctx) error {
	var in domain.Order
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	res, er := t.OrderUC.AddOrder(c.Context(), in)
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

func (t *OrderHandler) EditOrderById(c *fiber.Ctx) error {
	var in domain.Order
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	res, er := t.OrderUC.EditOrderById(c.Context(), in)
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

func (t *OrderHandler) DeleteOrderById(c *fiber.Ctx) error {
	id := c.Params("id")
	er := t.OrderUC.DeleteOrderById(c.Context(), id)
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

func (t *OrderHandler) ShowOrderByOutletId(c *fiber.Ctx) error {
	id := c.Params("id")
	res, er := t.OrderUC.ShowOrderByOutletId(c.Context(), id)
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

func (t *OrderHandler) AddOrders(c *fiber.Ctx) error {
	var in []domain.Order
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	res, er := t.OrderUC.AddOrders(c.Context(), in)
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

func (t *OrderHandler) DeleteOrdersById(c *fiber.Ctx) error {
	var in []uint
	fmt.Println("Before DeleteOrdersById Handler", in)
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}
	fmt.Println("After DeleteOrdersById Handler", in)
	er := t.OrderUC.DeleteOrdersById(c.Context(), in)
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
