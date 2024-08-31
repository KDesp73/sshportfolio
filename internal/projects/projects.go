package projects

import (
	"fmt"
	"os"
	"sshapp/internal/files"
	"gopkg.in/yaml.v2"
)

const projectsPath = "./docs/projects"

type Project struct {
	Title string `yaml:"title"`
	License string `yaml:"license"`
	Description string `yaml:"description"`
	Link string `yaml:"link"`
	Language string `yaml:"language"`
	Content string `yaml:"content"`
}

type ProjectPool struct {
	Items []*Project
	TitleMap map[string]int
}

func LoadFromYAML(filename string) (*Project, error) {
    p := &Project{}

    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("error reading file: %w", err)
    }

    if err := yaml.Unmarshal(data, p); err != nil {
        return nil, fmt.Errorf("error unmarshalling YAML: %w", err)
    }

    return p, nil
}

func LoadProjects() (*ProjectPool, error) {
	pool := &ProjectPool{}

	config := files.ListConfig {
		Extensions: []string {"yml", "yaml"},
		Paths: []string {projectsPath},
	}
	files, err := files.GrabFiles(config)


	if err != nil {
		return nil, err
	}

	pool.TitleMap = make(map[string]int)

	for _, file := range files {
		p, err := LoadFromYAML(file.Path)
		if err != nil {
			return nil, err
		}

		pool.TitleMap[p.Title] = len(pool.Items)
		pool.Items = append(pool.Items, p)
	}

	return pool, nil
}
