
package models

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var users = []User{
    {ID: 1, Name: "Alice", Age: 30},
    {ID: 2, Name: "Bob", Age: 25},
}

func GetAllUsers() []User {
    return users
}

func AddUser(user User) {
    user.ID = len(users) + 1
    users = append(users, user)
}
