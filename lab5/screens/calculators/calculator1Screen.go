package calculators

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"lab5/services"
)

func Calculator1Screen(w http.ResponseWriter, r *http.Request) {
	tmplPath, err := filepath.Abs("screens/html/screen1.html")
	if err != nil {
		http.Error(w, "Помилка отримання шляху до шаблону", http.StatusInternalServerError)
		fmt.Println("Помилка отримання шляху:", err)
		return
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Помилка завантаження шаблону", http.StatusInternalServerError)
		fmt.Println("Помилка парсингу шаблону:", err)
		return
	}

	if r.Method != http.MethodPost {
		err = tmpl.Execute(w, struct {
			Success bool
			wDk     float64
			wDs     float64
		}{Success: false})
		if err != nil {
			fmt.Println("Помилка рендеру шаблону:", err)
		}
		return
	}

	calculatorService := &service.CalculatorService{}
	result := calculatorService.CompareReliabilitySystems()
	if len(result) < 2 {
		http.Error(w, "Помилка розрахунку", http.StatusInternalServerError)
		fmt.Println("Помилка: функція повернула недостатньо значень")
		return
	}

	err = tmpl.Execute(w, struct {
		Success bool
		WDk     float64
		WDs     float64
	}{
		Success: true,
		WDk:     result[0],
		WDs:     result[1],
	})
	if err != nil {
		fmt.Println("Помилка рендеру шаблону:", err)
	}
}