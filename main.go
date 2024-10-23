package main

import (
	// "fmt"
	// "os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "github.com/faiface/beep"
	// "github.com/faiface/beep/mp3"
	// "github.com/hajimehoshi/oto/v2"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("L!STEN")

	// Sidebar: Playlist
	songs := []string{"Song 1", "Song 2", "Song 3"}
	sidebar := widget.NewList(
		func() int { return len(songs) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(songs[id])
		},
	)

	// Main Content: Song info and album art
	currentSong := widget.NewLabel("Currently Playing: None")
	albumArt := canvas.NewImageFromFile("path/to/album_art.png") // Replace with your image path
	albumArt.FillMode = canvas.ImageFillContain // Keep aspect ratio

	mainContent := container.NewVBox(currentSong, albumArt)

	// Bottom Controls: Media buttons
	playBtn := widget.NewButton("Play", func() {
		// Logic to play music will go here
	})
	pauseBtn := widget.NewButton("Pause", func() {
		// Logic to pause music will go here
	})
	stopBtn := widget.NewButton("Stop", func() {
		// Logic to stop music will go here
	})

	bottomControls := container.NewHBox(playBtn, pauseBtn, stopBtn)

	// Combine Layout: Sidebar and Main Content
	splitLayout := container.NewHSplit(sidebar, mainContent)
	splitLayout.SetOffset(0.25) // Adjust sidebar width

	// Final Layout: Main content with bottom controls
	mainLayout := container.NewBorder(nil, bottomControls, nil, nil, splitLayout)

	// Set window content and run the app
	myWindow.SetContent(mainLayout)
	myWindow.Resize(fyne.NewSize(800, 600)) // Adjust window size
	myWindow.ShowAndRun()
}
