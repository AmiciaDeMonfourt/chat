package domain

type User struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	// username
	// email ???
}

func NewUser(id int64, firstname, secondname string) User {
	return User{
		ID:         id,
		FirstName:  firstname,
		SecondName: secondname,
	}
}
