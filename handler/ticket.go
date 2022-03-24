package handler

import (
	"net/http"
	"strconv"
	"ticketing-payment-handler/helper"
	"ticketing-payment-handler/ticket"

	"github.com/gin-gonic/gin"
)

type ticketHandler struct {
	service ticket.Service
}

func NewTicketHandler(service ticket.Service) *ticketHandler {
	return &ticketHandler{service}
}

//api/v1/tickets
func (h *ticketHandler) GetTickets(c *gin.Context) {
	userdID, _ := strconv.Atoi(c.Query("user_id"))

	tickets, err := h.service.GetTickets(userdID)
	if err != nil {
		response := helper.APIResponse("Error to get tickets", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("List of tickets", http.StatusOK, "success", ticket.FormatTickets(tickets))
	c.JSON(http.StatusOK, response)
}


func (h *ticketHandler) CreateTicket(c *gin.Context) {
	var input ticket.CreateTicketInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("failed to create ticket", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTicket, err := h.service.CreateTicket(input)
	if err != nil {
		response := helper.APIResponse("Failed to create ticket", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create ticket", http.StatusOK, "success", ticket.FormatTicket(newTicket))
	c.JSON(http.StatusOK, response)
}