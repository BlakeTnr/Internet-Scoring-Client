package database

type Database interface {
	GetDomainByIndex(index int) string
	GetNumberOfWebsites() int
}
