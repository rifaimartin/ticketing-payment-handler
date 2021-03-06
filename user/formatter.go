package user

//UserFormatter struct
type UserFormatter struct {
	ID         		int    `json:"id"`
	Name       	    string `json:"name"`
	Gender 		    string `json:"gender"`
	Email	      	string `json:"email"`
	PhoneNumber     string `json:"phoneNumber"`
	Token   	    string `json:"token"`
}

//FormatUser function for fomating data
func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         	user.ID,
		Name:       	user.Name,
		Gender:	 		user.Gender,
		Email:      	user.Email,
		PhoneNumber:    user.PhoneNumber,
		Token:      token,

	}

	return formatter
}
