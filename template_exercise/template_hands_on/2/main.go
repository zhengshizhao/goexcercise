package main 

import ("text/template"
		"log"
		"os"
	)

type hotel struct {
	Name string
	Address string
	City string
	Zip string
	Region string
}

var tpl *template.Template

func init (){
	tpl = template.Must(template.ParseFiles("default.html"))
}

func main(){
    californiaHotels := []hotel {
    	hotel{
    		"Residence Inn by Marriott Pleasant Hill Concord",
    		"700 Ellinwood Way",
    		"Pleasant Hill",
    		"94523",
    		"Central",
    	},
    	hotel{
    		"Hotel California",
    		"2431 Ash St",
    		"Palo Alto",
    		"94306",
    		"Central",
    	},

   }
    err := tpl.Execute(os.Stdout, californiaHotels)
    if err != nil {
    	log.Fatalln(err)
    }


}