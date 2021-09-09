package finder

type Finder interface {
	ViewTrace(v bool) Finder
	FindPath(fileName string) (string, error)
}
