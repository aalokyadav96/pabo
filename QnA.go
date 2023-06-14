package main
import (
	"net/http"
	"fmt"
	"time"
    "github.com/julienschmidt/httprouter"
)

func QnA(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method == "GET" {
		if isLoggedIn(r) {
			tmpl.ExecuteTemplate(w, "head.html",nil)
			tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
			tmpl.ExecuteTemplate(w, "QnA_home.html",nil)
		} else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}


type QnAPost struct {
	Q string
	A string
	PostId string
	LoggedIn string
	LoggedOut string
}

func QnA_New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
		t := time.Now()
        token := t.Format("20060102150405")
		tok := Tok{Token: token}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
				tmpl.ExecuteTemplate(w, "QnA_new.html", tok)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedOut: "true"})
				tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		case "POST" : {
			fmt.Println("path", r.URL.Path)
			Q := r.FormValue("Q")
			A := r.FormValue("A")
			token := r.FormValue("token")
			println(rdxSet(token, "{ 'Q': '"+ Q + "', 'A': '" + A + "' }"))
			
			http.Redirect(w, r, "/qna/view/"+token, http.StatusSeeOther)
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

func QnA_View(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				token := postid.ByName("postid")
				postDetails, _ := rdxGet(token)
				res := QnAPost{Q: postDetails[:], A: postDetails[:], PostId: token,LoggedIn: "true"}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
				tmpl.ExecuteTemplate(w, "QnA_view.html", res)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				token := postid.ByName("postid")
				postDetails, _ := rdxGet(token)
				res := QnAPost{Q: postDetails[:], A: postDetails[:], PostId: token,LoggedOut: "true"}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedOut: "true"})
				tmpl.ExecuteTemplate(w, "QnA_view.html", res)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

func QnA_Edit(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				token := postid.ByName("postid")
				postDetails, _ := rdxGet(token)
				res := QnAPost{Q: postDetails[:], A: postDetails[:], PostId: token,LoggedIn: "true"}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
				tmpl.ExecuteTemplate(w, "QnA_edit.html", res)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		case "POST" : {
			fmt.Println("path", r.URL.Path)
			Q := r.FormValue("Q")
			A := r.FormValue("A")
			token := r.FormValue("token")
			println(rdxSet(token, "{ 'Q': '"+ Q + "', 'A': '" + A + "' }"))
			http.Redirect(w, r, "/qna/view/"+token, http.StatusSeeOther)
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

func QnA_Answer(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				token := postid.ByName("postid")
				postDetails, _ := rdxGet(token)
				res := QnAPost{Q: postDetails[:], A: postDetails[:], PostId: token}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
				tmpl.ExecuteTemplate(w, "QnA_edit.html", res)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		case "POST" : {
			fmt.Println("path", r.URL.Path)
			Q := r.FormValue("Q")
			A := r.FormValue("A")
			token := r.FormValue("token")
			println(rdxSet(token, "{ 'Q': '"+ Q + "', 'A': '" + A + "' }"))
			http.Redirect(w, r, "/qna/view/"+token, http.StatusSeeOther)
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

func QnA_Delete(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method  == "DELETE" {
		if isLoggedIn(r) {
			println(rdxDel(postid.ByName("postid")))
			http.Redirect(w, r, "/qna", http.StatusSeeOther)
		} 	else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}