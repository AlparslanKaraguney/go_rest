package user

// Controller is the main entry point for the user service.
import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetAll(*fiber.Ctx) error
	GetByID(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	// Update(id uint, model Model) (Model, error)
	// Delete(id uint) error
}

type handler struct {
	service Service
}

// Compiled time check for interface implementation
var _ Handler = handler{}

func NewHandler(service Service) Handler {
	return handler{service: service}
}

type Response struct {
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func (h handler) GetAll(ctx *fiber.Ctx) error {
	models, err := h.service.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(Response{Error: err.Error(), Success: false})
	}
	return ctx.JSON(Response{Data: models, Success: true})
}

func (h handler) GetByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(500).JSON(Response{Error: err.Error(), Success: false})
	}
	model, err := h.service.GetByID(uint(id))
	if err != nil {
		return ctx.Status(500).JSON(Response{Error: err.Error(), Success: false})
	}
	return ctx.JSON(Response{Data: model, Success: true})
}

func (h handler) Create(ctx *fiber.Ctx) error {
	model := Model{}
	err := ctx.BodyParser(&model)
	if err != nil {
		return ctx.Status(500).JSON(Response{Error: err.Error(), Success: false})
	}
	id, err := h.service.Create(model)
	if err != nil {
		return ctx.Status(500).JSON(Response{Error: err.Error(), Success: false})
	}
	return ctx.JSON(Response{Data: id, Success: true})
}
