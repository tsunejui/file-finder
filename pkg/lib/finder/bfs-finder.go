package lib

import (
	pkgAlg "file-finder/pkg/algorithms"
	pkgCommon "file-finder/pkg/lib/common"
	"fmt"
	"io/fs"
	"path/filepath"
)

var visitedDirectories []string
var visitedFiles []string

type BFSFile struct {
	path  string
	file  fs.FileInfo
	match string
	debug bool
}

func NewBFSFile(f fs.FileInfo) *BFSFile {
	return &BFSFile{
		file:  f,
		debug: false,
	}
}

func (b *BFSFile) SetDebug(debug bool) *BFSFile {
	b.debug = debug
	return b
}

func (b *BFSFile) SetMatch(fileName string) *BFSFile {
	b.match = fileName
	return b
}

func (b *BFSFile) SetPath(path string) *BFSFile {
	b.path = path
	return b
}

func (b *BFSFile) Visite() {
	fileName := b.file.Name()
	if b.file.IsDir() {
		visitedDirectories = append(visitedDirectories, filepath.Join(b.path, fileName))
	} else {
		visitedFiles = append(visitedFiles, fileName)
	}
}

func (b *BFSFile) IsVisite() bool {
	var isVisite bool
	if b.file.IsDir() {
		isVisite = pkgCommon.InArray(b.file.Name(), visitedDirectories)
	} else {
		isVisite = pkgCommon.InArray(b.file.Name(), visitedFiles)
	}
	return isVisite
}

func (b *BFSFile) GetFiles() ([]pkgAlg.BFSFile, error) {
	var files []pkgAlg.BFSFile
	if b.file.IsDir() {
		path := filepath.Join(b.path, b.file.Name())
		readFiles, err := pkgCommon.GetFiles(path)
		if err != nil {
			return nil, fmt.Errorf("failed to get files from the directories: %v", err)
		}
		files = newBFSFiles(readFiles, path, b.match, b.debug)
	}
	return files, nil
}

func (b *BFSFile) Check() bool {
	fileName := b.file.Name()
	if b.debug {
		fmt.Println(filepath.Join(b.path, fileName))
	}
	if b.file.IsDir() {
		return false
	}
	return fileName == b.match
}

type Finder struct {
	root  string
	trace bool
}

func NewFinder(root string) *Finder {
	return &Finder{
		root:  root,
		trace: false,
	}
}

func (f *Finder) ViewTrace(v bool) *Finder {
	f.trace = v
	return f
}

func (f *Finder) FindPath(fileName string) (string, error) {
	root := f.root
	files, err := pkgCommon.GetFiles(root)
	if err != nil {
		return "", fmt.Errorf("failed to get files from the directories: %v", err)
	}
	bFs := newBFSFiles(files, root, fileName, f.trace)
	file, err := pkgAlg.BFSSearch(bFs)
	if err != nil {
		return "", fmt.Errorf("failed to search the file: %v", err)
	}
	if file == nil {
		return "", nil
	}
	bf := file.(*BFSFile)
	return filepath.Join(bf.path, bf.file.Name()), nil
}

func newBFSFiles(files []fs.FileInfo, path string, match string, debug bool) []pkgAlg.BFSFile {
	var bFiles []pkgAlg.BFSFile
	for _, f := range files {
		bFiles = append(bFiles, NewBFSFile(f).SetPath(path).SetMatch(match).SetDebug(debug))
	}
	return bFiles
}
