package entry

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	// "lab5/service"
)

func MainScreen(w http.ResponseWriter, r *http.Request) {
	tmplPath, err := filepath.Abs("entry.html")
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
		}{Success: false})
		if err != nil {
			fmt.Println("Помилка рендеру шаблону:", err)
		}
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

	// calculatorService := &service.CalculatorService{}
	// result := calculatorService...

	// err = tmpl.Execute(w, struct {
	// 	Success bool
	// }{
	// 	Success: true,
	// })
	// if err != nil {
	// 	fmt.Println("Помилка рендеру шаблону:", err)
	// }
}
