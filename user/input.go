package user

//RegisterUserInput is struct
type RegisterUserInput struct {
	Name        string `json:"name" binding:"required"`
	Gender  	string `json:"gender" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Alamat   	string `json:"alamat" binding:"required"`
}

