package delivery

import (
	"assyarif-backend-web-go/domain"
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	StockUC domain.StockUseCase
}

func NewStockHandler(c *fiber.App, das domain.StockUseCase) {
	handler := &StockHandler{
		StockUC: das,
	}
	api := c.Group("/stock")

	_ = api.Group("/public")

	private := api.Group("/private")
	private.Get("/stuff", handler.GetAllStock)
	private.Get("/stuff/:id", handler.GetStockByID)
	private.Post("/stuff", handler.CreateStock)
	private.Put("/stuff/:id", handler.UpdateStock)
	private.Delete("/stuff/:id", handler.DeleteStock)

	private.Put("/stuff/increase-dashboard/multiple", handler.IncreaseStocks)
	private.Put("/stuff/decrease-dashboard/multiple", handler.DecreaseStocks)
	private.Put("/stuff/update-description/multiple", handler.UpdateDescription)

	period := api.Group("/period/stock")
	period.Get("/", handler.GetStocksByPeriod)
}

func (t *StockHandler) GetAllStock(c *fiber.Ctx) error {
	res, err := t.StockUC.FetchStocks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully get all user",
	})
}

func (t *StockHandler) GetStockByID(c *fiber.Ctx) error {
	id := c.Params("id")
	strId, erStr := strconv.Atoi(id)
	if erStr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse id",
			"error":   erStr.Error(),
		})
	}
	res, err := t.StockUC.FetchStockByID(c.Context(), uint(strId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully get user by id",
	})
}

func (t *StockHandler) CreateStock(c *fiber.Ctx) error {
	req := new(domain.Stock)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   er.Error(),
		})
	}
	res, err := t.StockUC.CreateStock(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"success": true,
		"data":    res,
		"message": "Successfully create user",
	})
}

func (t *StockHandler) UpdateStock(c *fiber.Ctx) error {
	req := new(domain.Stock)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	valRes, er := govalidator.ValidateStruct(req)
	if !valRes {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   er.Error(),
		})
	}
	res, err := t.StockUC.UpdateStock(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully update user",
	})
}

func (t *StockHandler) DeleteStock(c *fiber.Ctx) error {
	id := c.Params("id")
	strId, erStr := strconv.Atoi(id)
	if erStr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse id",
			"error":   erStr.Error(),
		})
	}
	err := t.StockUC.DeleteStock(c.Context(), uint(strId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"message": "Successfully delete user",
	})
}

func (t *StockHandler) IncreaseStocks(c *fiber.Ctx) error {
	req := new([]domain.Stock)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	res, err := t.StockUC.IncreaseStocks(c.Context(), *req)
	fmt.Println("res", res)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully increase stocks",
	})
}

func (t *StockHandler) DecreaseStocks(c *fiber.Ctx) error {
	req := new([]domain.Stock)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	res, err := t.StockUC.DecreaseStocks(c.Context(), *req)
	fmt.Println("res", res)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully decrease stocks",
	})
}

func (t *StockHandler) UpdateDescription(c *fiber.Ctx) error {
	req := new([]domain.UpdateDescriptionRequest)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}
	res, err := t.StockUC.UpdateDescription(c.Context(), *req)
	fmt.Println("res", res)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"message": err,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"success": true,
		"data":    res,
		"message": "Successfully update description",
	})
}

func (t *StockHandler) GetStocksByPeriod(c *fiber.Ctx) error {
	res, er := t.StockUC.GetStocksByPeriod(c.Context())
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
