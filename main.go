package main

import (
	"log"
	"os"

	"github.com/fatih/color"
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
	color.Red("Hello, World!")

	f, err := os.Create("public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte(website))
}
