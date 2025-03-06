package entry

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func EntryScreen(
	onCalculator1Navigate func(),
	onCalculator2Navigate func(),
	onCalculator3Navigate func(),
) fyne.CanvasObject {
	title := widget.NewLabel("Практична робота №4.\nРозробник: Новотка Вікторія")
	title.Alignment = fyne.TextAlignCenter

	button1 := widget.NewButton("Калькулятор 1 (вибір кабелів для живлення)", onCalculator1Navigate)
	button2 := widget.NewButton("Калькулятор 2 (визначення струму КЗ на шинах)", onCalculator2Navigate)
	button3 := widget.NewButton("Калькулятор 3 (визначення струму КЗ для підстанції)", onCalculator3Navigate)

	button1.Importance = widget.HighImportance
	button2.Importance = widget.HighImportance
	button3.Importance = widget.HighImportance

	buttons := container.NewVBox(button1, button2, button3)

	content := container.NewVBox(
		container.NewCenter(title),
		buttons,
	)

	return content
}