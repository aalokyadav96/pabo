package main
import (
	"net/http"
	"fmt"
	"time"
    "github.com/julienschmidt/httprouter"
)

type Tok struct {
	Token string
}


func Blog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method == "GET" {
		if isLoggedIn(r) {
			tmpl.ExecuteTemplate(w, "head.html",nil)
			tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
			tmpl.ExecuteTemplate(w, "blog.html",nil)
		} else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}


func CreateBlogPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
		t := time.Now()
        token := t.Format("20060102150405")
		tok := Tok{Token: token}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
				tmpl.ExecuteTemplate(w, "blog_newpost.html", tok)
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
			posttitle := r.FormValue("posttitle")
			postbody := r.FormValue("postbody")
			token := r.FormValue("token")
			println(rdxSet(token, "{ 'PostTitle': '"+ posttitle + "', 'PostBody': '" + postbody + "' }"))
			
			http.Redirect(w, r, "/blog/view/"+token, http.StatusSeeOther)
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

type BlogPost struct {
	PostTitle string
	PostBody string
	PostId string
	LoggedIn string
	LoggedOut string
}

type LoginStatus struct {
	LoggedIn string
	LoggedOut string
}
/*
type userStatus struct {
	LoggedIn bool
}*/


func ViewBlogPost(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				token := postid.ByName("postid")
				postDetails, _ := rdxGet(token)
				res := BlogPost{PostTitle: postDetails[:], PostBody: postDetails[:], PostId: token,LoggedIn: "true"}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
				tmpl.ExecuteTemplate(w, "blog_viewpost.html", res)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				token := postid.ByName("postid")
				postDetails, _ := rdxGet(token)
				res := BlogPost{PostTitle: postDetails[:], PostBody: postDetails[:], PostId: token,LoggedOut: "true"}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedOut: "true"})
				tmpl.ExecuteTemplate(w, "blog_viewpost.html", res)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

func EditBlogPost(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				token := postid.ByName("postid")
				postDetails, _ := rdxGet(token)
				res := BlogPost{PostTitle: postDetails[:], PostBody: postDetails[:], PostId: token,LoggedIn: "true"}
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
				tmpl.ExecuteTemplate(w, "blog_editpost.html", res)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		case "POST" : {
			fmt.Println("path", r.URL.Path)
			posttitle := r.FormValue("posttitle")
			postbody := r.FormValue("postbody")
			token := r.FormValue("token")
			println(rdxSet(token, "{ 'PostTitle': '"+ posttitle + "', 'PostBody': '" + postbody + "' }"))
			http.Redirect(w, r, "/blog/view/"+token, http.StatusSeeOther)
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}

func DeleteBlogPost(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method  == "DELETE" {
		if isLoggedIn(r) {
			println(rdxDel(postid.ByName("postid")))
			http.Redirect(w, r, "/blog", http.StatusSeeOther)
		} 	else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}