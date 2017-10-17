//template.Funcs
// register functions 
// var fm = template.FuncMap{
// 	"uc": strings.ToUpper,
// 	"ft": firstThree,
// }

// func init() {
// 	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
// }

// func firstThree(s string) string {
// 	s = strings.TrimSpace(s)
// 	if len(s) >= 3 {
// 		s = s[:3]
// 	}
// 	return s
// }

//global func: https://godoc.org/text/template#hdr-Functions

package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	xs1 := []string{"zero", "one", "two", "three", "four", "five"}
    
    data := struct { 
    	AllUsers []user
    	Xs []string

    }{ users,
    	xs1,
    }

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}