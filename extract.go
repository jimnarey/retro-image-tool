package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gen2brain/go-unarr"
	"github.com/mholt/archiver"
)

const fixturesPath string = "./fixtures"

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

type archive interface {
	extractAllTo(string) error
	basename() string
}

type fileBase struct {
	path string
}

func (f fileBase) basename() string {
	basename := path.Base(f.path)
	return strings.TrimSuffix(basename, path.Ext(basename))
}

type unarrArchive struct {
	fileBase
}

func (u unarrArchive) extractAllTo(path string) error {
	a, err := unarr.NewArchive(u.fileBase.path)
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

func (u unarrArchive) basename() string {
	return u.fileBase.basename()
}

type archiverArchive struct {
	fileBase
}

func (a archiverArchive) extractAllTo(path string) error {
	err := archiver.Unarchive(a.fileBase.path, path)
	if err != nil {
		fmt.Println("----Error:")
		fmt.Println(err)
		return err
	}
	return nil
}

func (a archiverArchive) basename() string {
	return a.fileBase.basename()
}

func newUnarrArchive(path string) (archive, error) {
	a, err := unarr.NewArchive(path)
	if err != nil {
		return nil, err
	}
	defer a.Close()
	return unarrArchive{fileBase: fileBase{path: path}}, nil
}

func newArchiverArchive(path string) (archive, error) {
	_, err := archiver.ByExtension(path)
	if err != nil {
		return nil, err
	}
	//TODO Add check
	return archiverArchive{fileBase: fileBase{path: path}}, nil
}

func getArchive(path string, archiveGetters []func(string) (archive, error)) (archive, error) {
	for i := 0; i < len(archiveGetters); i++ {
		archive, err := archiveGetters[i](path)
		if err == nil {
			// fmt.Printf("%T", archive)
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
		fmt.Println("****************")
		fmt.Println(archives[i].Name())
		fullPath := path.Join(fixturesPath, archives[i].Name())
		isDir, _ := isDirectory(fullPath)
		if !isDir {
			a, err := getArchive(path.Join(fullPath), archiveGetters)
			if err != nil {
				log.Fatal(err)

			} else {

				fmt.Printf("%T\n", a)
				fmt.Println(a.basename())
				fmt.Println(i)
				extractPath := path.Join(fixturesPath, "out", a.basename()+"-"+strconv.Itoa(i))
				fmt.Println(extractPath)
				fmt.Println()

				a.extractAllTo(path.Join(extractPath))
			}
		}

	}

}
