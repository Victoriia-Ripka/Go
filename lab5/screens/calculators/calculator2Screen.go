package calculators

import (
	"html/template"
	"net/http"
	"lab5/services"
)

func Calculator2Screen(w http.ResponseWriter, r *http.Request) {
	calculatorService := &service.CalculatorService{}

	result := calculatorService.CalculateLosses()

	tmpl, err := template.New("calculator2").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Calculator 2</title>
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
	<h1>Завдання 2</h1>
	<h3>Розрахунок збитків від перерв електропостачання</h3>
	<button class="button" onclick="window.location.href='/calculate2'">Розрахувати</button>
	<br>
	{{if .result}}
		<p>Математичне сподівання збитків від переривання електропостачання: {{.result}}</p>
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
		result string
	}{
		result: result,
	})
}