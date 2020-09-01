package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"regexp"
	"text/template"
)

const sourceTemplate =
`// This file is generated, do not edit.
// Use 'fyne-icon-picker --regenerate' to generate this source file.

package {{.SourcePackage}}

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const IconPreviewSize = 96
var confirmCopyDialog dialog.Dialog

type iconPreviewLayout struct {
}

func (i iconPreviewLayout) Layout(objs []fyne.CanvasObject, parent fyne.Size) {
	objs[0].Resize(fyne.NewSize(IconPreviewSize, IconPreviewSize))
}
func (i iconPreviewLayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(IconPreviewSize, IconPreviewSize)
}

type tappableIcon struct {
	widget.Icon
	Text   string
	window fyne.Window
}

func newTappableIcon(resource fyne.Resource, text string, window fyne.Window) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(resource)
	icon.window = window
	icon.Text = text

	return icon
}

func (t *tappableIcon) Tapped(e *fyne.PointEvent) {
	previewIcon := widget.NewIcon(t.Resource)
	content := fyne.NewContainerWithLayout(
		iconPreviewLayout{},
		previewIcon,
	)
	copyOnConfirmCallback := func(accepted bool) {
		if accepted {
			t.window.Clipboard().SetContent("theme." + t.Text)
		}
	}
	dialog.NewCustomConfirm(t.Text, "Copy", "Cancel", content, copyOnConfirmCallback, t.window)
}

type tappableLabel struct {
	widget.Label
	window fyne.Window
}

func newTappableLabel(text string, window fyne.Window) *tappableLabel {
	lbl := &tappableLabel{}
	lbl.ExtendBaseWidget(lbl)
	lbl.SetText(text)
	lbl.window = window

	return lbl
}

func (t *tappableLabel) Tapped(e *fyne.PointEvent) {
	t.window.Clipboard().SetContent("theme." + t.Text)
	if confirmCopyDialog == nil {
		confirmCopyDialog = dialog.NewInformation("Copied", "Copied constant to clipboard", t.window)
	}
	confirmCopyDialog.Show()
	go func() {
		time.Sleep(time.Second)
		confirmCopyDialog.Hide()
	}()
}

func newIcon(resource fyne.Resource, name string, window fyne.Window) fyne.CanvasObject {
	icon := newTappableIcon(resource, name, window)
	label := newTappableLabel(name, window)
	label.window = window

	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		icon,
		label,
	)
}

func getContents(win fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithColumns(6),{{range $icon := .IconStrings}}
		newIcon(theme.{{$icon}}(), "{{$icon}}", win),{{end}}
	)
}
`

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getIconCallStrings() []string {
	usr, _ := user.Current()
	source := path.Join(usr.HomeDir, "go", "src", "fyne.io", "fyne", "theme", "icons.go")
	fmt.Printf("Path to icons source: '%s'\n", source)
	if fileExists(source) {
		pattern := regexp.MustCompile(`func\s+([A-Z]\w+)\(\)\s+fyne.Resource`)
		sourceText, err := ioutil.ReadFile(source)
		if err != nil {
			fmt.Println("Failed to read source file")
			os.Exit(-1)
		}
		matches := pattern.FindAllStringSubmatchIndex(string(sourceText), -1)
		matchesLen := len(matches)
		returnedMatches := make([]string, len(matches))
		fmt.Println("Matches:")
		for i := 0; i < matchesLen; i++ {
			match := matches[i]
			iconString := string(sourceText[match[2]:match[3]])
			returnedMatches[i] = iconString
			fmt.Printf(" - %s\n", iconString)
		}
		return returnedMatches
	} else {
		fmt.Println("Failed to find and read source file")
		os.Exit(-1)
	}
	return []string{}
}

type TemplateParams struct {
	SourcePackage string
	IconStrings   []string
}

func (t *TemplateParams) GenerateSourceText() ([]byte, error) {
	if t.SourcePackage == "" {
		t.SourcePackage = "main"
	}
	if len(t.IconStrings) == 0 {
		t.IconStrings = getIconCallStrings()
	}
	var buf bytes.Buffer
	tmpl := template.Must(template.New("source").Parse(sourceTemplate))
	err := tmpl.Execute(&buf, t)
	return buf.Bytes(), err
}
