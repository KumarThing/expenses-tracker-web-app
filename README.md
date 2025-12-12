# Expenses Tracker Web App (Go)

A simple web-based Expense Tracker built using Go, net/http, and HTML templates.
This project demonstrates how to build a small dynamic web application using native Go features—no external frameworks required.

Users can:
- Add expenses with a name and amount
- View all expenses on the add page (live list)
- Navigate between pages
- Store expenses in memory during runtime

___

## Features 

- Add expenses
- Show expenses


___

## 🛠️ Tech Stack

Language: Go (Golang)

Web server: net/http

Templating: html/template

Frontend: HTML5, CSS

Files served: Static CSS from /static

___
 
## Project Structure 

```
expense-tracker/
│
├── main.go
├── go.mod
│
├── template/
│   ├── home.html
│   ├── add.html
│   ├── showall.html  (optional future use)
│   └── status.html   (optional future use)
│
└── static/
    └── style.css
```






