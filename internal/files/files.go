package files

import (
	"fmt"
	"os"
	"path/filepath"
)


type File struct {
	Path string
	Name string
	Extension string
	Content string
}

func (file File) Log() {
	fmt.Printf("Path: %s\n", file.Path)
	fmt.Printf("Name: %s\n", file.Name)
	fmt.Printf("Extension: %s\n", file.Extension)
}

type ListConfig struct {
	Extensions []string
	Paths []string
	RecursiveDepth int
}

func GrabFiles(config ListConfig) ([]File, error) {
    var list []File
	for _, path := range config.Paths {
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			fullPath := filepath.Join(path, entry.Name())
			if entry.IsDir() {
				files, err := GrabFiles(config)
				if err != nil {
					return nil, err
				}
				list = append(list, files...)
				continue
			}

			ext := filepath.Ext(entry.Name())
			ext = ext[1:] // Remove the leading "."

			found := false
			for _, e := range config.Extensions {
				if ext == e {
					found = true
					break
				}
			}

			if found {
				content, err := os.ReadFile(fullPath)
				if err != nil {
					return nil, err
				}

				list = append(list, File{
					Path:      fullPath,
					Name:      entry.Name(),
					Extension: ext,
					Content:   string(content),
				})
			}
		}
	}

    return list, nil
}


func SrcDir(path string) string {
	return path + "/src"
}

func IncludeDir(path string) string {
	return path + "/include"
}

