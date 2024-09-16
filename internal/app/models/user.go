
package models

type User struct {
    ID   int    `json:"id"`
    Username string `json:"username"`
    Age  int    `json:"age"`
    Password string `json:"password"`
}

var users = []User{
    {ID: 1, Username: "Alice", Age: 30, Password: "password123"},
    {ID: 2, Username: "Bob", Age: 25, Password: "password123"},
}

func GetAllUsers() []User {
    return users
}

func AddUser(user User) {
    user.ID = len(users) + 1
    users = append(users, user)
}
