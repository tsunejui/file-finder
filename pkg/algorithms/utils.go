package algorithms

type SearchFile interface {
	Visite()
	IsVisite() bool
	GetFiles() ([]SearchFile, error)
	Check() bool
}
