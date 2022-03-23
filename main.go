package main

import (
	"log"
	"ticketing-payment-handler/handler"
	"ticketing-payment-handler/ticket"
	"ticketing-payment-handler/user"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:rifaimartin123@tcp(127.0.0.1:3306)/ticketing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	ticketRepository := ticket.NewRepository(db)


	userService := user.NewService(userRepository)
	ticketService := ticket.NewService(ticketRepository)


	// user, _ := userService.GetUserByID(1)

	// transactionService.CreateTransaction(input)

	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0fQ.gPQZuzlzKZsJevcIpEX_M5rjWbfhw_ZdPIdjaHd6IKE")
	// if err != nil {
	// 	fmt.Println("ERROR")
	// }

	// if token.Valid {
	// 	fmt.Println("VALID")
	// } else {
	// 	fmt.Println("INVALID")
	// }

	userHandler := handler.NewUserHandler(userService)
	ticketHandler := handler.NewTicketHandler(ticketService)


	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	api.GET("/tickets", ticketHandler.GetTickets)

	router.Run()
}


