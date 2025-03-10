package calculators

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"lab5/services"
)

func Calculator2Screen(w http.ResponseWriter, r *http.Request) {
	tmplPath, err := filepath.Abs("screens/html/screen2.html")
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
			result string
		}{Success: false})
		if err != nil {
			fmt.Println("Помилка рендеру шаблону:", err)
		}
		return
	}

	calculatorService := &service.CalculatorService{}
	result := calculatorService.CalculateLosses()

	err = tmpl.Execute(w, struct {
		Success bool
		Result string
	}{
		Success: true,
		Result: result,
	})
	if err != nil {
		fmt.Println("Помилка рендеру шаблону:", err)
	}
}