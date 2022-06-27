package main

import (
	"fmt"
	"github.com/gen2brain/go-unarr"
)

type archive interface {
	extractAllTo(string) error
}

type unarrArchive struct {
	//archive unarr.Archive
	path string
}

func (u unarrArchive) extractAllTo(path string) error {
	a, err := unarr.NewArchive(u.path)
	if err != nil {
		return err
	}
	defer a.Close()
	_, err = a.Extract(path)
	if err != nil {
		return err
	}
	return nil
}

func newUnarrArchive(path string) (archive, error) {
	a, err := unarr.NewArchive(path)
	if err != nil {
		return nil, err
	}
	defer a.Close()
	return unarrArchive{path: path}, nil
}


func main() {

	//archives:= []string{"./fixtures/file.txt.zip", "./fixtures/file.txt.7z", "./fixtures/file.txt.rar", "./fixtures/file.txt.tar", "./fixtures/file.txt.gz"}
	archives:= []string{"./fixtures/file.txt.zip", "./fixtures/file.txt.7z", "./fixtures/file.txt.tar", "./fixtures/file.txt.gz"}

	for i:= 0; i < len(archives); i++ {
		a, err := newUnarrArchive(archives[i])
		fmt.Println(archives[i])
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println(a)
		}
	}

}
