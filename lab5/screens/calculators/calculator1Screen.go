package calculators

import (
	"html/template"
	"net/http"
	"lab5/services"
)

func calculator1Screen(w http.ResponseWriter, r *http.Request) {
	calculatorService := &service.CalculatorService{}

	result := calculatorService.CompareReliabilitySystems()
	wDk := result[0]
	wDs := result[1]

	tmpl, err := template.New("calculator").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Calculator 1</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			text-align: center;
			padding: 20px;
		}
		.button {
			font-size: 18px;
			font-weight: bold;
			padding: 10px 20px;
			margin: 10px;
			cursor: pointer;
			border: 2px solid #007BFF;
			background-color: #007BFF;
			color: white;
			border-radius: 5px;
		}
		.button:hover {
			background-color: #0056b3;
		}
	</style>
</head>
<body>
	<h1>Завдання 1</h1>
	<h3>Порівняння надійності одноколової та двоколової систем електропередач</h3>
	<button class="button" onclick="window.location.href='/calculate'">Розрахувати</button>
	<br>
	{{if .wDk}}
		<p>Результат: w_dk = {{.wDk}}, w_ds = {{.wDs}} <br>Надійність двоколової системи є вищою</p>
	{{else}}
		<p>Натисніть "Розрахувати" для отримання результатів.</p>
	{{end}}
	<br>
	<button class="button" onclick="window.location.href='/'">Повернутися назад</button>
</body>
</html>
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, struct {
		wDk float64
		wDs float64
	}{
		wDk: wDk,
		wDs: wDs,
	})
}