package models

type User struct {
    ID   int    `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
    Age  int    `json:"age"`
    Password string `json:"password"`
}

var users = []User{
    {
    	ID: 1,
     	Username: "Alice",
      	Email: "test1@example.com",
      	Age: 30,
       	Password: "password123",
    },
    {
    	ID: 2,
     	Username: "Bob",
     	Email: "test2@example.com",
      	Age: 25,
       	Password: "password123",
    },
}

func GetAllUsers() []User {
    return users
}

func GetUserByName(username string) *User {
	for _, user := range users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

func AddUser(user User) {
    user.ID = len(users) + 1
    users = append(users, user)
}
