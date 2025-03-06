package main

import (
	"fmt"

	"lab4/logic"
	"lab4/pages/calculators"
	"lab4/pages/entry"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const (
	MainScreen  = "Main Screen"
	Calculator1 = "Calculator 1"
	Calculator2 = "Calculator 2"
	Calculator3 = "Calculator 3"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Lab 4")

	calculatorService, err := logic.NewCalculatorService("assets/cable_data.json")
	if err != nil {
		fmt.Println("Error loading calculator service:", err)
		return
	}

	tabs := container.NewAppTabs()

	tabs.Append(container.NewTabItem(MainScreen, entry.EntryScreen(
		func() { tabs.SelectIndex(1) },
		func() { tabs.SelectIndex(2) },
		func() { tabs.SelectIndex(3) },
	)))

	tabs.Append(container.NewTabItem(Calculator1, calculators.Calculator1Screen(
		func() { tabs.SelectIndex(0) },
	)))

	tabs.Append(container.NewTabItem(Calculator2, calculators.Calculator2Screen(
		func() { tabs.SelectIndex(0) }, calculatorService,
	)))

	tabs.Append(container.NewTabItem(Calculator3, calculators.Calculator3Screen(
		func() { tabs.SelectIndex(0) }, calculatorService,
	)))

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
