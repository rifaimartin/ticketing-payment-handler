package ticket

//GetTicketDetailInput this is struct
type GetTicketsDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

//CreateTicketInput this is struct
type CreateTicketInput struct {
	Name             		string `json:"name" binding:"required"`
	ShortDescription        string `json:"short_description" binding:"required"`
	ImageUrl        		string `json:"image_url" binding:"required"`
	Qty             		int    `json:"qty" binding:"required"`
	Price           	    int    `json:"price" binding:"required"`
}
