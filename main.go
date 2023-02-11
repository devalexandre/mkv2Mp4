package main

import (
	"fmt"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("MKV 2 MP4")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(500, 400))

	filePathIcon := widget.NewFileIcon(nil)
	filePathUri := widget.NewLabel("Select File")
	filePathField := widget.NewButton("Select File", func() {
		dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {

			if err != nil {
				d := dialog.NewError(err, w)
				d.Show()
			}

			if file != nil {
				filePathUri.SetText(file.URI().String())
				file.Close()

			}
		}, w)
	})

	filePath := container.NewHBox(filePathIcon, filePathUri, filePathField)

	message := widget.NewLabel("")
	infinitbar := widget.NewProgressBarInfinite()
	infinitbar.Hide()

	//escolher onde salvar
	outputFolder := widget.NewLabel("")
	btnOutput := widget.NewButton("Select output folder", func() {
		output := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {

			outputFolder.SetText(lu.String())
		}, w)

		output.Show()
	})

	converter := widget.NewButton("MKV to MP4", func() {
		infinitbar.Show()

		path := filePathUri.Text
		path = path[7:]
		path = strings.ReplaceAll(path, " ", "\\ ")

		command := fmt.Sprintf("ffmpeg -i %s -y -s hd1080 -c copy ~/Downloads/output.mp4", path)
		_, err := exec.Command("bash", "-c", command).Output()

		if err != nil {
			d := dialog.NewError(err, w)
			d.Show()
			message.SetText(err.Error())
		}
		infinitbar.Hide()

		dialog.ShowInformation("Success", "File converted", w)

	})

	box := container.NewVBox(filePath, btnOutput, outputFolder, message, converter, infinitbar)

	w.SetContent(box)

	w.ShowAndRun()
}
