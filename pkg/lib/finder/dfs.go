package finder

import (
	pkgAlg "file-finder/pkg/algorithms"
	pkgCommon "file-finder/pkg/lib/common"
	"fmt"
	"path/filepath"
)

type DFSFinder struct {
	root  string
	trace bool
}

func NewDFSFinder(root string) *DFSFinder {
	return &DFSFinder{
		root:  root,
		trace: false,
	}
}

func (f *DFSFinder) ViewTrace(v bool) Finder {
	f.trace = v
	return f
}

func (f *DFSFinder) FindPath(fileName string) (string, error) {
	root := f.root
	fileInfo, err := pkgCommon.GetFileInfo(root)
	if err != nil {
		return "", fmt.Errorf("failed to get file's information: %v", err)
	}
	searchFile := NewBFSFile(fileInfo).SetPath(root).SetMatch(fileName).SetDebug(f.trace).SetRoot(true)
	file, err := pkgAlg.DFSSearch(searchFile)
	if err != nil {
		return "", fmt.Errorf("failed to search the file: %v", err)
	}
	if file == nil {
		return "", nil
	}
	bf := file.(*SearchFile)
	return filepath.Join(bf.path, bf.file.Name()), nil
}
