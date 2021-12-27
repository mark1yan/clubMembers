package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//Member struct of club member
type Member struct {
	Index int
	Email string
	Name  string
	Date  string
}

//MembersByEmail Global variable contains Member
var MembersByEmail []Member

//emailIsExist function check email address if all ready in MembersByEmail
func emailIsExist(Email string, Members *[]Member) bool {
	if len(*Members) >= 1 {
		for _, mem := range *Members {
			if Email == mem.Email {
				return true
			}
		}
	}
	return false
}

//addMember function for add member
func addMember(Email string, Name string, Members *[]Member) bool {
	if emailIsExist(Email, Members) {
		return false
	}
	member := Member{
		Index: len(*Members) + 1,
		Name:  Name,
		Email: Email,
		Date:  time.Now().Format("02-01-2006"),
	}
	log.Println(member.Index)
	*Members = append(*Members, member)
	return true
}

//Web Server
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

// indexHandler handler of index.gohtml
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	log.Println("Get index")
}

//addHandler for adding members with POST request
func addHandler(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("name")
	femail := r.FormValue("email")
	if len(fname) != 0 && len(femail) != 0 {
		if addMember(femail, fname, &MembersByEmail) {
			log.Println("Add member")
		} else {
			log.Println("All-ready exist email.")
		}
	} else {
		log.Println("Empty name or email!")
	}
	tpl.ExecuteTemplate(w, "index.gohtml", MembersByEmail)
}

//main function
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":"+port, nil)
}
