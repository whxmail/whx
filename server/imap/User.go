// User.go
package main

//	"fmt"

type User struct {
	Username string
	Password string
	CMD      []string
	Flag     bool
}

/*
func (usr User) Set(src string, dst interface{}) {
	switch src {
	case "Username", "username":
		if value, ok := dst.(string); ok {
			usr.Username = value
		}

	case "Password", "password":
		if value, ok := dst.(string); ok {
			usr.Password = value
		}

	default:
		fmt.Println("Error para!")
	}
}
*/
