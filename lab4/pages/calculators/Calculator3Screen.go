package calculators

import (
	"fmt"

	"lab4/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Calculator3Screen(goBack func(), calculatorService *logic.CalculatorService) fyne.CanvasObject {
	resultLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Розрахувати", func() {
		resultArray := calculatorService.DeterminateSubstationCurrent()
		if len(resultArray) < 4 {
			resultLabel.SetText("Помилка розрахунку")
			return
		}
		i3LN, i2LN, i3LNmin, i2LNmin := resultArray[0], resultArray[1], resultArray[2], resultArray[3]
		resultText := fmt.Sprintf("Результат: I 3 = %.2f А, I 2 = %.2f А, I 3 min = %.2f А, I 2 min = %.2f А.\nАварійний режим не передбачено", i3LN, i2LN, i3LNmin, i2LNmin)
		resultLabel.SetText(resultText)
	})

	backButton := widget.NewButton("Повернутися до головного меню", goBack)

	return container.NewVBox(
		widget.NewLabel("Практична робота №4.\nРозробник: Новотка Вікторія"),
		calculateButton,
		resultLabel,
		backButton,
	)
}