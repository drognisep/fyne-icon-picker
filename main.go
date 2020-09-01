package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--regenerate" {
		params := TemplateParams{}
		bytes, err := params.GenerateSourceText()
		if err != nil {
			fmt.Printf("Error occurred: '%v'", err)
		}
		err = ioutil.WriteFile("generated.go", bytes, 0644)
		if err != nil {
			fmt.Printf("Error occurred writing to file: '%v", err)
		} else {
			fmt.Println("File generated successfully")
		}
	} else {
		myApp := app.New()
		win := myApp.NewWindow("Icon Picker")
		win.Resize(fyne.NewSize(300, 500))
		win.SetIcon(ResourceIconPng)
		win.SetMaster()

		win.SetContent(getContents(win))
		win.CenterOnScreen()
		win.ShowAndRun()
	}
}
