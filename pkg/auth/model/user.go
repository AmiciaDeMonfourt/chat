package model

type User struct {
	ID string `json:"id"`
	// Personal information
	UserProfile
	// Authentificate and verification information
	UserCredentials
}

type UserCredentials struct {
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"-"`
}

type UserProfile struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}
