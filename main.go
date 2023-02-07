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
) // Importing Fyne in my project
func main() {

	fmt.Println("Test Fyne...")
	// Start with go mod init myapp to create a package
	// we will create Our First Fyne Project

	// Our first line of code will creating a new app

	a := app.New()

	// Now we will create a new window

	w := a.NewWindow("MKV to MP4")   // You can you any title of your app
	w.SetFixedSize(true)             // This will fix the size of the window
	w.Resize(fyne.NewSize(500, 400)) // This will resize the window
	//create button for hello world

	filePathIcon := widget.NewFileIcon(nil)
	filePathUri := widget.NewLabel("Select File")
	filePathField := widget.NewButton("Select File", func() {
		fmt.Println("Select File")
		dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
			if file != nil {

				filePathUri.SetText(file.URI().String())

				file.Close()
				fmt.Printf("Selected file: %s", file.URI().String())

			}
		}, w)

	})

	filePath := container.NewHBox(filePathIcon, filePathField)

	message := widget.NewLabel("")
	converter := widget.NewButton("Mkv To Mp4", func() {
		/* remove file:// from uri */
		path := filePathUri.Text
		path = path[7:]
		//remove all spaces
		path = strings.ReplaceAll(path, " ", "\\ ")

		fmt.Println(path)
		//use ffmpeg to convert mkv to mp4 with exec
		command := fmt.Sprintf("ffmpeg -i %s -y -s hd1080 -c copy /home/alexandre/Videos/output.mp4", path)
		_, err := exec.Command("bash", "-c", command).Output()
		if err != nil {
			message.SetText(err.Error())
			return
		}
		dialog.ShowInformation("Success", "File converted", w)

	})

	box := container.NewVBox(filePath, filePathUri, converter) // grid com container layout for button
	// grid com container layout for button

	w.SetContent(box) // This will set the content of the window
	w.ShowAndRun()    // Finally Running our App
}
