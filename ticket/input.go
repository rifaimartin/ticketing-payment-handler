package ticket

//GetTicketDetailInput this is struct
type GetTicketsDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

//CreateTicketInput this is struct
type CreateTicketInput struct {
	Name             string `json:"name" binding:"required"`
}
