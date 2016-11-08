// User.go
package types

type (
	ID int64
)

type User struct {
	Username string
	Password string
	Mail     Mail
	ID       ID
	CMD      CMD
	Flag     string
}
