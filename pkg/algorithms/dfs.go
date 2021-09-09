package algorithms

import "fmt"

type DFSStack struct {
	Files []SearchFile
}

func NewDFSStack() *DFSStack {
	return &DFSStack{}
}

func (b *DFSStack) Push(f SearchFile) {
	f.Visite()
	b.Files = append(b.Files, f)
}

func (b *DFSStack) Pop() SearchFile {
	files := b.Files
	length := len(files)
	if length == 1 {
		b.Files = []SearchFile{}
		return files[0]
	}

	last, files := b.Files[length-1], b.Files[:length-1]
	b.Files = files
	return last
}

func (b *DFSStack) IsEmpty() bool {
	return len(b.Files) == 0
}

func DFSSearch(dFile SearchFile) (SearchFile, error) {
	stack := NewDFSStack()
	stack.Push(dFile)
	var searchFile SearchFile
	for {
		if stack.IsEmpty() {
			break
		}

		popFile := stack.Pop()
		if popFile.Check() {
			searchFile = popFile
			break
		}

		files, err := popFile.GetFiles()
		if err != nil {
			return nil, fmt.Errorf("failed to get files: %v", err)
		}
		for i := len(files) - 1; i >= 0; i-- {
			if !files[i].IsVisite() {
				stack.Push(files[i])
			}
		}
	}
	return searchFile, nil
}
