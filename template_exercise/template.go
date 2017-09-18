package main

import( "fmt"
		"os"
		"log"
		"io"
		"strings"
)


func main() {
	name := "Lucia ZZ"
	// name := os.Args[1]
	// fmt.Println("Args0:", os.Args[0],"Args1:", os.Args[1])

	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Luis's page</title>
	</head>
	<body>
	<h1> Hi, my name is `+ name + `</h1>
	</body>
	</html>
	`)
    //https://golang.org/pkg/os/

	newfile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer newfile.Close()
	if _, err := io.Copy(newfile, strings.NewReader(str)); err !=nil {
		log.Fatal("error copying file", err)
	}

}
