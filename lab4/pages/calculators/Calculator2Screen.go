package calculators

import (
	"fmt"
	"strconv"

	"lab4/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


func Calculator2Screen(goBack func(), calculatorService *logic.CalculatorService) fyne.CanvasObject {
	uValue := widget.NewEntry()
	uValue.SetPlaceHolder("Введіть напругу")

	skValue := widget.NewEntry()
	skValue.SetPlaceHolder("Введіть потужність КЗ")

	sNomtValue := widget.NewEntry()
	sNomtValue.SetPlaceHolder("Введіть номінальну потужність трансформатора")

	resultLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Розрахувати", func() {
		formattedU, err1 := strconv.ParseFloat(uValue.Text, 64)
		formattedSk, err2 := strconv.ParseFloat(skValue.Text, 64)
		formattedsNomt, err3 := strconv.ParseFloat(sNomtValue.Text, 64)

		if err1 != nil || err2 != nil || err3 != nil {
			resultLabel.SetText("Помилка введення! Переконайтесь, що всі значення числові.")
			return
		}

		result := calculatorService.DeterminateCurrent(formattedU, formattedSk, formattedsNomt)
		resultLabel.SetText(fmt.Sprintf("Результат: %.2f кА", result))
	})

	backButton := widget.NewButton("Повернутися до головного меню", goBack)

	return container.NewVBox(
		widget.NewLabel("Практична робота №4.\nРозробник: Новотка Вікторія"),
		uValue,
		skValue,
		sNomtValue,
		calculateButton,
		resultLabel,
		backButton,
	)
}