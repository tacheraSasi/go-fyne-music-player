package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	albumArt := canvas.NewImageFromFile("./assets/music-image-bg.jpg") // Ensure this path is correct
	albumArt.FillMode = canvas.ImageFillContain // Keeping aspect ratio

	// Create a container for album art and current song
	mainContent := container.NewVBox(currentSong, albumArt)

	// Bottom Controls: Media buttons
	playBtn := widget.NewButton("Play", func() {
		// Logic to play music will go here
		fmt.Println("Play button clicked")
	})
	pauseBtn := widget.NewButton("Pause", func() {
		// Logic to pause music will go here
		fmt.Println("Pause button clicked")
	})
	stopBtn := widget.NewButton("Stop", func() {
		// Logic to stop music will go here
		fmt.Println("Stop button clicked")
	})

	// Current song details
	currentSongName := canvas.NewText("Current Track: tachBeat 1", color.White)
	currentSongName.TextStyle = fyne.TextStyle{Bold: true} // Making it bold for emphasis

	// Progress bar
	progressBar := widget.NewProgressBar()
	progressBar.Min = 0
	progressBar.Max = 100

	// Simulate progress for demonstration purposes
	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progressBar.SetValue(i)
		}
	}()

	// Create a container for the bottom controls
	bottomControls := container.NewHBox(currentSongName, playBtn, pauseBtn, stopBtn, progressBar)
	bottomControls.Resize(fyne.NewSize(800, 50)) // Setting a fixed height for the control bar

	// Combine Layout: Sidebar and Main Content
	splitLayout := container.NewHSplit(sidebar, mainContent)
	splitLayout.SetOffset(0.25) // Adjusting sidebar width

	// Final Layout: Main content with bottom controls
	mainLayout := container.NewBorder(nil, bottomControls, nil, nil, splitLayout)

	// Set window content and run the app
	myWindow.SetContent(mainLayout)
	myWindow.Resize(fyne.NewSize(800, 400)) // Adjusting the window size
	myWindow.ShowAndRun()
}
