package api

type ApiConnection interface {
	Dial(username string, password string) (*ApiConnection, error)
}
