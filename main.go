package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)


var (
	homeTmpl = template.Must(template.ParseFiles("template/home.html"))
	addTmpl = template.Must(template.ParseFiles("template/add.html"))
	showallTmpl = template.Must(template.ParseFiles("template/showall.html"))
	statusTmpl = template.Must(template.ParseFiles("template/status.html"))
)

type Expense struct {
	Name string
	Amount int
}

var Expenses [] Expense

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		homeTmpl.Execute(w, nil)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request){

		if r.Method == http.MethodPost{
			r.ParseForm()

			newName := r.FormValue("name")
			amountStr := r.FormValue("amount")

			amountInt, err := strconv.Atoi(amountStr)
			if err != nil {
				fmt.Println("Error converting string to int", err)
			}

			Expenses = append(Expenses, Expense{
				Name : newName,
				Amount: amountInt,
			})

		}
		addTmpl.Execute(w, Expenses)
	})


	fmt.Println("Your server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}