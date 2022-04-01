package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-employee/domain/service"
	"github.com/gofiber/fiber/v2"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateEmployee godoc
// @Summary create a new employee
// @ID createEmployee
// @Tags Employee
// @Description Router for create a new employee
// @Accept json
// @Produce json
// @Param body body CreateEmployeeRequest true "JSON body for create a new employee"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /employees [post]
func (t *RestService) CreateEmployee(c *fiber.Ctx) error {
	var req CreateEmployeeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	employeeID, err := t.Service.CreateEmployee(c.Context(), &req.Name)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *employeeID})
}

// FindEmployee godoc
// @Summary find a employee
// @ID findEmployee
// @Tags Employee
// @Description Router for find a employee
// @Accept json
// @Produce json
// @Param employee_id path string true "Employee ID"
// @Success 200 {object} Employee
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /employees/{employee_id} [get]
func (t *RestService) FindEmployee(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "employee_id is not a valid uuid",
		})
	}

	employee, err := t.Service.FindEmployee(c.Context(), &employeeID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(employee)
}

// SearchEmployees godoc
// @Summary search employees
// @ID searchEmployees
// @Tags Employee
// @Description Router for search employees
// @Accept json
// @Produce json
// @Param page_size query int false "page size"
// @Param page_token query string false "page token"
// @Success 200 {object} SearchEmployeesResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /employees [get]
func (t *RestService) SearchEmployees(c *fiber.Ctx) error {
	var req SearchEmployeesRequest

	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	employees, nextPageToken, err := t.Service.SearchEmployees(c.Context(), &req.PageToken, &req.PageSize)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"employees":       employees,
		"next_page_token": nextPageToken,
	})
}
