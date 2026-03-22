package database

type Database interface {
	Get() ([]Data, error)
}
