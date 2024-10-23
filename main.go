package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/hajimehoshi/oto/v2"
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

    mainContent := container.NewVBox(currentSong, albumArt)

    // Audio control variables
    var streamer beep.StreamSeekCloser
    var format beep.Format
    var context *oto.Context

    // Function to play music
    playMusic := func(url string) {
        if streamer != nil {
            streamer.Close() // Close any currently playing music
        }

        // Open the music file from a URL (you can also use a local file)
        res, err := os.Open(url) // Replace with actual music URL
        if err != nil {
            fmt.Println("Error opening audio file:", err)
            return
        }

        // Decode the music file
        streamer, format, err = mp3.Decode(res)
        if err != nil {
            fmt.Println("Error decoding audio file:", err)
            return
        }

        // Initialize audio context
        context, err = oto.NewContext(format.SampleRate, format.NumChannels, 2, 1024)
        if err != nil {
            fmt.Println("Error creating audio context:", err)
            return
        }

        // Play the music
        player := context.NewPlayer(streamer)
        player.Play()
    }

    // Function to pause music
    pauseMusic := func() {
        if context != nil {
            context.Close()
            context = nil
        }
    }

    // Bottom Controls: Media buttons
    playBtn := widget.NewButton("Play", func() {
        currentSong.SetText("Currently Playing: Song 1")
        playMusic("https://soundcloud.com/10_05/isolation-2021-2024-beat-tape?utm_source=clipboard&utm_medium=text&utm_campaign=social_sharing") // Replace with actual music URL
    })
    pauseBtn := widget.NewButton("Pause", func() {
        pauseMusic()
        currentSong.SetText("Currently Playing: None")
    })
    stopBtn := widget.NewButton("Stop", func() {
        pauseMusic()
        currentSong.SetText("Currently Playing: None")
    })

    bottomControls := container.NewHBox(playBtn, pauseBtn, stopBtn)

    // Combine Layout: Sidebar and Main Content
    splitLayout := container.NewHSplit(sidebar, mainContent)
    splitLayout.SetOffset(0.25)

    // Final Layout: Main content with bottom controls
    mainLayout := container.NewBorder(nil, bottomControls, nil, nil, splitLayout)

    // Set window content and run the app
    myWindow.SetContent(mainLayout)
    myWindow.Resize(fyne.NewSize(800, 600))
    myWindow.ShowAndRun()
}
