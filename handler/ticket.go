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
