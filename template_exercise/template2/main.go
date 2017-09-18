package main
import (
        "log"
		"os"
		"text/template"
	)


func errorLog (err error) {
	log.Fatalln(err)
}
var tpl *template.Template

func init (){

	tpl = template.Must(template.ParseGlob("./*"))
}

func main() {
	// tpl, err := template.ParseFiles("index.html"); 
	// if err != nil {
	// 	errorLog(err)
	// }

    // tpl, err := template.ParseGlob("./*"); 
	// if err != nil {
	// 	errorLog(err)
	// }
    err := tpl.Execute(os.Stdout, "Monkey")
	if err != nil {
		errorLog (err)
	}

	// tpl, err = tpl.ParseFiles("index2.gogo");
	// if err != nil {
	// 	errorLog (err)
	// }
	
	err = tpl.ExecuteTemplate(os.Stdout, "index.html", "Lucia")
	if err != nil {
		errorLog (err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "index2.gogo",nil)
	if err != nil {
		errorLog (err)
	}

} 