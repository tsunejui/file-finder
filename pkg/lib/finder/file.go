package finder

import (
	pkgAlg "file-finder/pkg/algorithms"
	pkgCommon "file-finder/pkg/lib/common"
	"fmt"
	"io/fs"
	"path/filepath"
)

var visitedDirectories []string
var visitedFiles []string

type SearchFile struct {
	path   string
	file   fs.FileInfo
	match  string
	debug  bool
	isRoot bool
}

func NewBFSFile(f fs.FileInfo) *SearchFile {
	return &SearchFile{
		file:  f,
		debug: false,
	}
}

func (b *SearchFile) SetDebug(debug bool) *SearchFile {
	b.debug = debug
	return b
}

func (b *SearchFile) SetMatch(fileName string) *SearchFile {
	b.match = fileName
	return b
}

func (b *SearchFile) SetPath(path string) *SearchFile {
	b.path = path
	return b
}

func (b *SearchFile) SetRoot(isRoot bool) *SearchFile {
	b.isRoot = isRoot
	return b
}

func (b *SearchFile) Visite() {
	fileName := b.file.Name()
	if b.file.IsDir() {
		visitedDirectories = append(visitedDirectories, filepath.Join(b.path, fileName))
	} else {
		visitedFiles = append(visitedFiles, fileName)
	}
}

func (b *SearchFile) IsVisite() bool {
	var isVisite bool
	if b.file.IsDir() {
		isVisite = pkgCommon.InArray(b.file.Name(), visitedDirectories)
	} else {
		isVisite = pkgCommon.InArray(b.file.Name(), visitedFiles)
	}
	return isVisite
}

func (b *SearchFile) GetFiles() ([]pkgAlg.SearchFile, error) {
	var files []pkgAlg.SearchFile
	if b.file.IsDir() {
		path := b.file.Name()
		if !b.isRoot {
			path = filepath.Join(b.path, b.file.Name())
		}
		readFiles, err := pkgCommon.GetFiles(path)
		if err != nil {
			return nil, fmt.Errorf("failed to get files from the directories: %v", err)
		}
		files = newSearchFiles(readFiles, path, b.match, b.debug)
	}
	return files, nil
}

func (b *SearchFile) Check() bool {
	fileName := b.file.Name()
	if b.debug {
		path := fileName
		if !b.isRoot {
			path = filepath.Join(b.path, fileName)
		}
		fmt.Println(path)
	}
	if b.file.IsDir() {
		return false
	}
	return fileName == b.match
}

func newSearchFiles(files []fs.FileInfo, path string, match string, debug bool) []pkgAlg.SearchFile {
	var bFiles []pkgAlg.SearchFile
	for _, f := range files {
		bFiles = append(bFiles, NewBFSFile(f).SetPath(path).SetMatch(match).SetDebug(debug))
	}
	return bFiles
}
