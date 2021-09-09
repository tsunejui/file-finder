package algorithms

import "fmt"

type BFSQueue struct {
	Files []SearchFile
}

func NewBFSQueue() *BFSQueue {
	return &BFSQueue{}
}

func (b *BFSQueue) EnQueue(f SearchFile) {
	f.Visite()
	b.Files = append(b.Files, f)
}

func (b *BFSQueue) DeQueue() SearchFile {
	files := b.Files
	length := len(files)
	if length == 1 {
		b.Files = []SearchFile{}
		return files[0]
	}

	current, files := b.Files[0], b.Files[1:]
	b.Files = files
	return current
}

func (b *BFSQueue) IsEmpty() bool {
	return len(b.Files) == 0
}

func BFSSearch(bFile SearchFile) (SearchFile, error) {
	queue := NewBFSQueue()
	queue.EnQueue(bFile)
	var current SearchFile
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
