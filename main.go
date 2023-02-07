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
			<h2>Hello</h2>
			<p>Hello, World!</p>
		</body>
		</html>
	`
	if _, err := f.Write([]byte(contents)); err != nil {
		panic(err)
	}
}
