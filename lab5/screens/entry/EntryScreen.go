package entry

import (
	"html/template"
	"net/http"
)

func MainScreen(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("entry").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Entry Screen</title>
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
	<h1>Lab 5</h1>
	<button class="button" onclick="window.location.href='/calculator1'">Калькулятор 1</button>
	<br/>
	<button class="button" onclick="window.location.href='/calculator2'">Калькулятор 2</button>
</body>
</html>
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
