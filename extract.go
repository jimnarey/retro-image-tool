package main

import (
	"fmt"
	"github.com/gen2brain/go-unarr"
)

type archive interface {
	isValid() bool
	exttractTo(string) int
}

type unarrArchive struct {
	//archive unarr.Archive
	path string
}

func (u unarrArchive) isValid() bool {
	a, err := unarr.NewArchive(u.path)

	if err != nil {
		return false
	}
	defer a.Close()
	return true
}

func (u unarrArchive) extractTo(path string) error {
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





//func getUnarrArchive(path) (archive) {
//	a, err := unarr.NewArchive(archives[i])
//
//	if err != nil {
//		return "", err
//	}
//	return archive, nil
//}

//archive_getters := []func (string, error) {
//	getUnarrArchive
//}

func main() {

	//archives:= []string{"./fixtures/file.txt.zip", "./fixtures/file.txt.7z", "./fixtures/file.txt.rar", "./fixtures/file.txt.tar", "./fixtures/file.txt.gz"}
	archives:= []string{"./fixtures/file.txt.zip", "./fixtures/file.txt.7z", "./fixtures/file.txt.tar", "./fixtures/file.txt.gz"}

	for i:= 0; i < len(archives); i++ {
		func() {
			fmt.Println(archives[i])
			//a, err := unarr.NewArchive(archives[i])
			//if err != nil {
			//	panic(err)
			//}

			a := unarrArchive{path: archives[i]}
			fmt.Println(a.isValid())
		}()

	}
	//defer a.Close()
	//a, err := unarr.NewArchive("test.7z")
	//if err != nil {
	//	panic(err)
	//}
	//defer a.Close()
}
