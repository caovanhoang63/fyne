package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	i := binding.NewInt()
	uri, _ := storage.ParseURI("https://upload.wikimedia.org/wikipedia/commons/4/47/PNG_transparency_demonstration_1.png")

	hello := widget.NewLabel("Hello Fyne!")
	image := canvas.NewImageFromURI(uri)
	image.FillMode = canvas.ImageFillOriginal
	countLabel := widget.NewLabelWithData(binding.IntToStringWithFormat(i, "Count %d"))
	w.SetContent(container.NewVBox(
		hello,
		countLabel,
		widget.NewButton("Inc", func() {
			current, _ := i.Get()
			_ = i.Set(current + 1)
		}),
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		widget.NewButton("Home", func() {
			w.SetContent(aboutContent(w))
		}),
		widget.NewButton("Popup", func() {
			popup := a.NewWindow("Popup Window")
			popup.SetContent(widget.NewLabel("Hello from Popup!"))
			popup.Resize(fyne.NewSize(200, 100))
			popup.Show()
		}),
		widget.NewButton("Dialog", func() {
			dialog.ShowInformation("Title", "This is a dialog box!", w)
		}),
		widget.NewCheck("Hello checked", func(checked bool) {
			if checked {
				hello.SetText("Welcome :)")
			}
		}),
		image,
	))

	w.ShowAndRun()
}

func aboutContent(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("About Page"),
		widget.NewButton("Back to Home", func() {
			w.SetContent(container.NewVBox(
				widget.NewLabel("Home Page"),
				widget.NewButton("Go to About", func() {
					w.SetContent(aboutContent(w))
				}),
			))
		}),
	)
}
