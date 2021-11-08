package auth

import "golara/models"

type auth struct {
	user models.User
}

var authVariable auth

func LoggedInUser(user models.User) {
	authVariable = auth{
		user: user,
	}
}

func Check() bool {
	if authVariable.user.ID != 0 {
		return true
	}
	return false
}

func User() models.User {
	return authVariable.user
}
