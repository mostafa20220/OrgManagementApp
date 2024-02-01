package models

type User struct {
	ID       string `bson:"_id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	// Other fields as needed
}

