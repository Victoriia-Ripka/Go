package main

import (
	// "lab4/logic"
	"lab4/pages/entry"
	"lab4/pages/calculators"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

// Routes Enum Equivalent
const (
	MainScreen     = "Main Screen"
	Calculator1    = "Calculator 1"
	// Calculator2    = "Calculator 2"
	// Calculator3    = "Calculator 3"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Lab 4")

	// calculatorService := logic.NewCalculatorService("/assets/cable_data.json")

	// Define tabs variable
	tabs := container.NewAppTabs()

	// Add the first tab (MainScreen) with EntryScreen content
	tabs.Append(container.NewTabItem(MainScreen, entry.EntryScreen(
		func() { tabs.SelectIndex(1) }, // Navigate to Calculator1 tab
		func() { tabs.SelectIndex(2) },
		func() { tabs.SelectIndex(3) },
	)))

	// Add the second tab (Calculator1) with Calculator1Screen content
	tabs.Append(container.NewTabItem(Calculator1, calculators.Calculator1Screen(
		func() { tabs.SelectIndex(0) }, // Go back to the main screen
	)))

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}