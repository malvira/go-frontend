package frontend

import (
	"fmt"
	"os"
	"os/exec"
)

type App struct {
	SourcePath string
	DeployLocation string
}

func (a *App) BuildAndServe() error {
	err := a.build(); if err != nil { return err }
	return nil
}

func (a *App) build() error {
	fmt.Printf("building app at %v with gopherjs\n", a.SourcePath)
	wd, _ := os.Getwd()
	if a.DeployLocation == "" {
		a.DeployLocation = "build/app.js"
	}
	fmt.Printf("app deployed at %v\n", a.DeployLocation)
	cmd := exec.Command("gopherjs", "build", fmt.Sprintf("-o../%v", a.DeployLocation))
	os.Chdir(a.SourcePath)
	out, err := cmd.CombinedOutput(); if err != nil { fmt.Println(err, string(out)); return err }
	os.Chdir(wd)
	return nil
}
