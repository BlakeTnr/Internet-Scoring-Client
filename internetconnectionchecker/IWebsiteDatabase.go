package internetchecker

type IWebsiteDatabase interface {
	GetDomainByIndex(index int) string
	GetNumberOfWebsites() int
}
