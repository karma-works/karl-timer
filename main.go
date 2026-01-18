package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type timerState struct {
	running   bool
	start     time.Time
	elapsed   time.Duration
	lastSaved time.Duration
}

func formatDuration(d time.Duration) string {
	s := int(d.Seconds()) % 60
	m := int(d.Minutes()) % 60
	h := int(d.Hours())
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func main() {
	myApp := app.NewWithID("com.karl.timer")
	myWindow := myApp.NewWindow("Karl Timer")

	state := &timerState{}

	// Create the large timer label using a widget for better thread safety
	label := widget.NewLabel("00:00:00")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}

	// Wrapper to handle mouse click
	content := container.NewStack(label)

	// Update loop
	go func() {
		ticker := time.NewTicker(time.Millisecond * 100)
		for range ticker.C {
			if state.running {
				currentElapsed := state.lastSaved + time.Since(state.start)
				newText := formatDuration(currentElapsed)
				if label.Text != newText {
					label.SetText(newText)
				}
			}
		}
	}()

	// Click to start/stop
	// We use a custom widget or simply capture clicks on the main container
	// In Fyne, we can use a Transparent button or a custom event handler.
	// We'll wrap the content in a Tappable container.

	// Create a transparent overlay or use the window's canvas directly
	myWindow.SetContent(container.NewCenter(content))

	toggleTimer := func() {
		if state.running {
			state.lastSaved += time.Since(state.start)
			state.running = false
		} else {
			state.start = time.Now()
			state.running = true
		}
	}

	resetTimer := func() {
		state.running = false
		state.lastSaved = 0
		state.elapsed = 0
		label.SetText("00:00:00")
	}

	// Handle Mouse Click
	// We can't easily put click handler on a Stack without creating a widget,
	// but we can use the window's Canvas.
	myWindow.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyF5:
			resetTimer()
		case fyne.KeyEscape:
			myWindow.SetFullScreen(!myWindow.FullScreen())
		}
	})

	// Capture all shortcuts/keys
	if des, ok := myWindow.Canvas().(desktop.Canvas); ok {
		des.SetOnKeyDown(func(k *fyne.KeyEvent) {
			// This covers more keys if needed
		})
	}

	// To handle clicks on the whole window, we can use a custom background button
	bgButton := newTappableButton(toggleTimer)

	finalContent := container.NewStack(bgButton, container.NewCenter(content))
	myWindow.SetContent(finalContent)

	// Apply custom theme to increase text size
	myApp.Settings().SetTheme(&customTheme{Theme: theme.DefaultTheme()})

	myWindow.SetFullScreen(true)
	myWindow.ShowAndRun()
}

type customTheme struct {
	fyne.Theme
}

func (m *customTheme) Size(n fyne.ThemeSizeName) float32 {
	if n == theme.SizeNameText {
		return 240
	}
	return m.Theme.Size(n)
}

// Custom widget to capture taps on the entire window area
type tappableButton struct {
	widget.BaseWidget
	OnTapped func()
}

func newTappableButton(tapped func()) *tappableButton {
	b := &tappableButton{OnTapped: tapped}
	b.ExtendBaseWidget(b)
	return b
}

func (t *tappableButton) CreateRenderer() fyne.WidgetRenderer {
	return &tappableRenderer{canvas.NewRectangle(theme.BackgroundColor())}
}

func (t *tappableButton) Tapped(_ *fyne.PointEvent) {
	if t.OnTapped != nil {
		t.OnTapped()
	}
}

type tappableRenderer struct {
	rect *canvas.Rectangle
}

func (r *tappableRenderer) Layout(size fyne.Size) {
	r.rect.Resize(size)
}

func (r *tappableRenderer) MinSize() fyne.Size {
	return r.rect.MinSize()
}

func (r *tappableRenderer) Refresh() {
	r.rect.FillColor = theme.BackgroundColor()
	r.rect.Refresh()
}

func (r *tappableRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect}
}

func (r *tappableRenderer) Destroy() {}
