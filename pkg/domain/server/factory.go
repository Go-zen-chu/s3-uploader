package server

type Factory interface {
	Service() (Service, error)
}