package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/gen2brain/go-unarr"
	"github.com/mholt/archiver"
)

const fixturesPath string = "./fixtures"

type archive interface {
	extractAllTo(string) error
}

type unarrArchive struct {
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

type archiverArchive struct {
	path string
}

func (a archiverArchive) extractAllTo(path string) error {
	err := archiver.Unarchive(a.path, path)
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

func newArchiverArchive(path string) (archive, error) {
	fmt.Println()
	a, err := archiver.ByExtension(path)
	if err != nil {
		return nil, err
	}
	fmt.Printf("a: %v\n", a)
	if err != nil {
		return nil, err
	}
	return archiverArchive{path: path}, nil
}

func getArchive(path string, archiveGetters []func(string) (archive, error)) (archive, error) {
	for i := 0; i < len(archiveGetters); i++ {
		archive, err := archiveGetters[i](path)
		if err == nil {
			fmt.Printf("%T", archive)
			return archive, nil
		}
		fmt.Println(err)
	}
	return nil, errors.New("No compatible unarchiver found")
}

func main() {

	archiveGetters := []func(string) (archive, error){newArchiverArchive, newUnarrArchive}

	archives, err := ioutil.ReadDir(fixturesPath)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(archives); i++ {
		fmt.Println("**************")
		fmt.Println(archives[i])
		a, err := getArchive(path.Join(fixturesPath, archives[i].Name()), archiveGetters)
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Printf("%T\n", a)
		}
	}

}
