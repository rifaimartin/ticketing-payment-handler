package main

import (
	"log"
	"ticketing-payment-handler/auth"
	"ticketing-payment-handler/handler"
	"ticketing-payment-handler/helper"
	"ticketing-payment-handler/ticket"
	"ticketing-payment-handler/transaction"
	"ticketing-payment-handler/user"

	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	ticketService := ticket.NewService(ticketRepository)
	authService := auth.NewService()
	transactionService := transaction.NewService(transactionRepository, ticketRepository)


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

	userHandler := handler.NewUserHandler(userService, authService)
	ticketHandler := handler.NewTicketHandler(ticketService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	api.GET("/tickets", ticketHandler.GetTickets)
	api.POST("/tickets", ticketHandler.CreateTicket)

	api.GET("transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("transactions/payment/notification", transactionHandler.GetNotification)


	router.Run()
}


func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			fmt.Println("Bearer empity")
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// split by space example Bearer = Token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// context currentUser
		c.Set("currentUser", user)
	}
}

