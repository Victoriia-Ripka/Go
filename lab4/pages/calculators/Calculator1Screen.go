package calculators

import (
	"fmt"
	"strconv"

	"lab4/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Calculator1Screen(goBack func()) fyne.CanvasObject {
	var conductor, cableType string
	timeRange := 1000
	result := ""

	conductors := []string{
		"Неізольовані проводи та шини",
		"Кабелі з паперовою та проводи гумовою ізоляцією",
		"Кабелі з гумовою та пластмасовою ізоляцією",
	}
	cableTypes := []string{"Алюмінієві", "Мідні"}

	conductorSelect := widget.NewSelect(conductors, func(s string) {
		conductor = s
	})
	conductorSelect.PlaceHolder = "Оберіть проводник"

	cableTypeSelect := widget.NewSelect(cableTypes, func(s string) {
		cableType = s
	})
	cableTypeSelect.PlaceHolder = "Оберіть тип кабелю"

	timeRangeSelect := widget.NewSelectEntry(nil)
	timeRangeSelect.SetText(strconv.Itoa(timeRange))
	timeRangeSelect.OnChanged = func(input string) {
		val, err := strconv.Atoi(input)
		if err == nil && val >= 1000 && val <= 10000 {
			timeRange = val
		} else {
			timeRangeSelect.SetText(strconv.Itoa(timeRange))
		}
	}

	resultLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Розрахувати", func() {
		calculatorService, err := logic.NewCalculatorService("assets/cable_data.json")
		if err != nil {
			fmt.Println("Error loading calculator service:", err)
			return
		}
	
		res, err := calculatorService.CalculateCablesCompatibility(conductor, cableType, float64(timeRange))
		if err != nil {
			fmt.Println("Error calculating cables compatibility:", err)
			return
		}
	
		// Ensure res is a float64 before formatting
		result = fmt.Sprintf("Результат: потрібно змінити переріз жил до %.2f мм^2", float64(res))
		resultLabel.SetText(result)
	})

	backButton := widget.NewButton("Повернутися до головного меню", goBack)

	return container.NewVBox(
		widget.NewLabel("Практична робота №4.\nРозробник: Новотка Вікторія"),
		conductorSelect,
		cableTypeSelect,
		widget.NewLabel("Зазначте час (1000-10000):"),
		timeRangeSelect,
		calculateButton,
		resultLabel,
		backButton,
	)
}