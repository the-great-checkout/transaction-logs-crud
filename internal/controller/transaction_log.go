package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/the-great-checkout/transaction-logs-crud/internal/dto"
)

type TransactionLogService interface {
	Create(log dto.TransactionLog) error
	FindAllByTransactionID(transactionID uuid.UUID) ([]dto.TransactionLog, error)
}

// TransactionLogController handles HTTP requests for transaction logs
type TransactionLogController struct {
	service TransactionLogService
}

// NewTransactionLogController creates a new TransactionLogController
func NewTransactionLogController(service TransactionLogService) *TransactionLogController {
	return &TransactionLogController{service: service}
}

// CreateHandler godoc
// @Summary Create Transaction Log
// @Description Create a new transaction log
// @Accept json
// @Produce json
// @Param log body dto.TransactionLog true "Transaction Log"
// @Success 201 {object} dto.TransactionLog
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transaction-logs [post]
func (ctrl *TransactionLogController) CreateHandler(c echo.Context) error {
	var input dto.TransactionLog
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := ctrl.service.Create(input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, input)
}

// FindAllByTransactionIDHandler godoc
// @Summary Get Transaction Logs by Transaction ID
// @Description Retrieve all transaction logs associated with a specific transaction ID
// @Param id path string true "Transaction ID"
// @Produce json
// @Success 200 {array} dto.TransactionLog
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transaction-logs/{id} [get]
func (ctrl *TransactionLogController) FindAllByTransactionIDHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	logs, err := ctrl.service.FindAllByTransactionID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, logs)
}
