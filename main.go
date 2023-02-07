package main

import (
	"log"
	"os"
)

const website = `
<html>
  <head>
    <title>GitHub Actions</title>
  </head>
  <body>
    <p>Hello, World!</p>
  </body>
</html>
`

func main() {
	f, err := os.Create("public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(website)); err != nil {
		log.Fatal(err)
	}
}
