// User.go
package types

type ID uint64

type User struct {
	Username string
	Password string
	Mail     string
	Id       ID
	CMD      []string
	Flag     String
}
