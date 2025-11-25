package boilerplate

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const TemplatesPath = "internal/boilerplate/templates"

type InitTmplData struct {
	MdlName string
}

type ModuleTmplData struct {
	MdlName    string
	MdlNameCap string
}

type ModuleBoilerplate interface {
	Create() error
}

func NewModuleBoilerplate(name, template, modulesPath string) ModuleBoilerplate {
	switch template {
	// Add more templates as needed
	default:
		return &SimpleModuleBoilerplate{name, modulesPath}
	}
}

type SimpleModuleBoilerplate struct {
	Name        string
	ModulesPath string
}

func (b *SimpleModuleBoilerplate) Create() error {
	modulesTemplatesPath := path.Join(TemplatesPath, "module")
	simpleModTemplatesPath := path.Join(modulesTemplatesPath, "simple")

	// Create init module file
	initPath := path.Join(b.ModulesPath, fmt.Sprintf("%s.go", b.Name))
	initFile, err := os.Create(initPath)
	if err != nil {
		return err
	}
	defer initFile.Close()

	initTmplPath := path.Join(modulesTemplatesPath, "init.go.tmpl")

	initTmpl, err := template.ParseFiles(initTmplPath)
	if err != nil {
		return err
	}

	var initBuf bytes.Buffer

	err = initTmpl.Execute(&initBuf, InitTmplData{MdlName: b.Name})
	if err != nil {
		return err
	}

	_, err = initFile.WriteString(initBuf.String())
	if err != nil {
		return err
	}

	// Create module directory
	modulePath := path.Join(b.ModulesPath, b.Name)
	if _, err := os.Stat(modulePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(modulePath, 0754); err != nil {
			return err
		}
	}

	// Create module file
	moduleFilePath := path.Join(modulePath, fmt.Sprintf("%s.go", b.Name))
	moduleFile, err := os.Create(moduleFilePath)
	if err != nil {
		return err
	}
	defer moduleFile.Close()

	moduleTmplPath := path.Join(simpleModTemplatesPath, "module.go.tmpl")
	moduleTmpl, err := template.ParseFiles(moduleTmplPath)
	if err != nil {
		return err
	}

	var moduleBuf bytes.Buffer

	capitalizedName := cases.Title(language.English, cases.Compact).String(b.Name)
	err = moduleTmpl.Execute(&moduleBuf, ModuleTmplData{
		MdlName:    b.Name,
		MdlNameCap: capitalizedName,
	})
	if err != nil {
		return err
	}

	_, err = moduleFile.WriteString(moduleBuf.String())
	if err != nil {
		return err
	}

	return nil
}
