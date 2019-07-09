package build

import (
	"bytes"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/juliengk/go-caddy/pkg/utils"
	"github.com/juliengk/go-utils/filedir"
)

var dir = "/tmp/caddy"

func init() {
	if err := filedir.CreateDirIfNotExist(dir, false, 0775); err != nil {
		panic(err)
	}
}

func New(importPaths []string, telemetry bool) *Build {
	data := Data{
		Plugins:   importPaths,
		Telemetry: telemetry,
	}

	build := Build{
		Data: data,
	}

	return &build
}

func (b *Build) GenerateGoFile() error {
	t := template.Must(template.New("main.go").Parse(mainGoTemplate))

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, b.Data); err != nil {
		return err
	}

	b.Result = buf.Bytes()

	return nil
}

func (b *Build) WriteGoFile() error {
	mainFile := path.Join(dir, "main.go")

	return ioutil.WriteFile(mainFile, b.Result, 0644)
}

func (b *Build) Run() error {
	var cmd []string

	outBin := path.Join(dir, "caddy")

	cmd = []string{"go", "mod", "init", "caddy"}
	if err := utils.RunCommand(cmd, dir); err != nil {
		return err
	}

	cmd = []string{"go", "get", "github.com/caddyserver/caddy"}
	if err := utils.RunCommand(cmd, dir); err != nil {
		return err
	}

	cmd = []string{"go", "build", "-o", outBin}
	if err := utils.RunCommand(cmd, dir); err != nil {
		return err
	}

	cmd = []string{"strip", "--strip-all", outBin}
	if err := utils.RunCommand(cmd, dir); err != nil {
		return err
	}

	return nil
}
