package main

import "os"

func main() {
	f, err := os.Create("public/index.html")
	if err != nil {
		panic(err)
	}
	contents := `
		<html>
		<head>
			<title>GitHub Actions</title>
		</head>
		<body>
			<h1>Hello</h1>
			<p>Hello, World!</p>
		</body>
		</html>
	`
	f.Write([]byte(contents))
}
