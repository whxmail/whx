// interface.go
package types

type info interface {
	getInfo()
}

type Handler interface {
	Handle()
}

type Selector interface {
	Select()
}
