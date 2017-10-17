package main

import ("text/template"
		"log"
		"os"
)

type item struct {
	Name string
	Descr string
	Price float64
	Vege bool
	Spicy int
}
func (p item) IsVege(vege bool) string {
	if vege {
		return "Vegeterian"
	}
	return ""
}

type items struct {
	Breakfast []item
	Lunch []item
	Dinner []item
}
type restaurant struct {
	Name string
	Menu items
}
var tpl *template.Template

func init (){
	tpl = template.Must(template.ParseFiles("default.html"))
}

func main(){
	lucias := items{
		[]item{item{"Egg party","two eggs,bacon,onions",5.00,false,0,},item{"Puncakes","3 punccakes",6.00,true,0,},},
		[]item{item{"Chicken","toasted chicken served with muchrooms",12.00,false,1,}},
		[]item{item{"Cheess platter","swiss, chedar, goat",11.00,false,0}},
	}

	err := tpl.Execute(os.Stdout,lucias)
    if err != nil {
    	log.Fatalln(err)
    }

}
