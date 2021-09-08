package algorithms

import "fmt"

type BFSFile interface {
	Visite()
	IsVisite() bool
	GetFiles() ([]BFSFile, error)
	Check() bool
}

type BFSQueue struct {
	Files []BFSFile
}

func (b *BFSQueue) EnQueue(f BFSFile) {
	f.Visite()
	b.Files = append(b.Files, f)
}

func (b *BFSQueue) DeQueue() BFSFile {
	current, files := b.Files[0], b.Files[1:]
	b.Files = files
	return current
}

func (b *BFSQueue) IsEmpty() bool {
	return len(b.Files) == 0
}

func NewBFSQueue() *BFSQueue {
	return &BFSQueue{}
}

func BFSSearch(bFiles []BFSFile) (BFSFile, error) {
	var current BFSFile
	queue := NewBFSQueue()
	for _, f := range bFiles {
		queue.EnQueue(f)
	}
	for {
		if queue.IsEmpty() {
			break
		}

		df := queue.DeQueue()
		if df.Check() {
			current = df
			break
		}

		files, err := df.GetFiles()
		if err != nil {
			return nil, fmt.Errorf("failed to get files: %v", err)
		}
		for _, file := range files {
			if !file.IsVisite() {
				queue.EnQueue(file)
			}
		}
	}
	return current, nil
}
