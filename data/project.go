package data

import (
	"archeboot/utils"
	// "encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Executable interface {
	Exec() error
}

type HasAbsPath interface {
	GetAbsPath() string
}

type Floder struct {
	Name    string    `yaml:"name"`
	Floder  *[]Floder `yaml:"dir"`
	AbsPath string
	Parent  *Floder
}

type Project struct {
	Name      string             `yaml:"name"`
	Version   string             `yaml:"version"`
	Kind      string             `yaml:"kind"`
	Variables *map[string]string `yaml:"variables"`
	Files     *[]interface{}     `yaml:"files"`
	AbsPath   string
	Floder    *[]Floder `yaml:"dir"`
}

func (p *Project) Read(config *Config) (*Project, error) {
	yamlFile, err := ioutil.ReadFile(config.Info.DefaultBootFileName)
	if err != nil {
		utils.LogError(err)
	}
	err = yaml.Unmarshal(yamlFile, p)
	if err != nil {
		utils.LogError(err)
	}
	return p, err
}

func execFloder(floders *[]Floder) error {
	if floders != nil {
		for _, f := range *floders {
			err := f.Exec()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func initFloder(floder *Floder, parent *Floder, p *Project) {
	if parent != nil {
		floder.Parent = parent
		floder.AbsPath = parent.GetAbsPath() + "/" + floder.Name
	} else {
		floder.AbsPath = p.GetAbsPath() + "/" + floder.Name
	}
	initFloderList(floder.Floder, floder, p)
}

func initFloderList(floders *[]Floder, parent *Floder, p *Project) {
	if floders != nil {
		fs := *floders
		for i := 0; i < len(fs); i++ {
			f := &fs[i]
			initFloder(f, parent, p)
		}
	}
}

func (p *Project) Init() {
	fmt.Println("Init Project")
	floders := p.Floder
	p.AbsPath = utils.GetCurrentDirectory()
	initFloderList(floders, nil, p)
}

func (p *Project) GetAbsPath() string {
	return p.AbsPath
}

func (f *Floder) GetAbsPath() string {
	return f.AbsPath
}

func (p *Project) Exec() error {
	fmt.Println("Project Name:", p.Name)
	fmt.Println("Project Version:", p.Version)
	fmt.Println("Current Path:", p.GetAbsPath())
	fmt.Println("Variables:", p.Variables)
	fmt.Println("Files:", p.Files)
	return execFloder(p.Floder)
}

func (f *Floder) Exec() error {
	err := utils.Mkdir(f.GetAbsPath())
	if err != nil {
		return err
	}
	return execFloder(f.Floder)
}

func (p *Project) Print() {
	// data, _ := json.Marshal(p)
	// fmt.Println(string(data))
}
