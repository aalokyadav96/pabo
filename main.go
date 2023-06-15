package main
import (
	"html/template"
	"net/http"
  	"log"
	"fmt"
	 
    "github.com/julienschmidt/httprouter"
)

const PORT = "localhost:4000"
var tmpl = template.Must(template.ParseGlob("templates/*.html"))


func main() {
    //Creating sub-domain
    server1 := http.NewServeMux()
    server1.HandleFunc("/", Translate)

//Running First Server
    go func() {
        log.Println("Server started on: http://localhost:9001")
        http.ListenAndServe("localhost:9001", server1)
    }()

	HandleRoutes()

}


func HandleRoutes() {

	router := httprouter.New()
	router.GET("/", Index)
	/*
	router.GET("/new", NewPhotoGet)
	router.POST("/upload", NewPhotoPost)
	router.GET("/post/:postid", ShowPost)
	router.GET("/delete/:postid", DeletePost)
	router.GET("/user/:name", UserProfile)	
*/	

//--------Search--------//
//	router.GET("/search", Search)

//--------Translate--------//
//	router.GET("/translate", Translate)

//--------Weather--------//
//	router.GET("/weather", Weather)

//--------News--------//
//	router.GET("/news", News)

//--------Chats--------//
//	router.GET("/chats", Chat)
//	router.GET("/chat/:userId", ShowPost)

//--------Shopping--------//
/*	router.GET("/buy", Buy)
	router.POST("/buy", Buy)
	router.GET("/sell", Sell)
	router.POST("/sell", Sell)
	router.GET("/deliver", Deliver)
	router.POST("/deliver", Deliver)*/


//--------CRUD--------//
/*	router.GET("/custom", CustomPage)
	router.POST("/custom", CustomPage)*/

//--------QnA--------//
	router.GET("/qna", QnA)
	router.GET("/qna/new", QnA_New)
	router.POST("/qna/new", QnA_New)
	router.GET("/qna/edit/:postid", QnA_Edit)
	router.POST("/qna/edit/:postid", QnA_Edit)
	router.POST("/qna/answer/:postid", QnA_Answer)
	router.GET("/qna/view/:postid", QnA_View)
	router.POST("/qna/delete/:postid", QnA_Delete)

//--------Blog--------//
	router.GET("/blog", Blog)
	router.GET("/blog/new", CreateBlogPost)
	router.POST("/blog/new", CreateBlogPost)
	router.GET("/blog/edit/:postid", EditBlogPost)
	router.POST("/blog/edit/:postid", EditBlogPost)
	router.GET("/blog/view/:postid", ViewBlogPost)
	router.POST("/blog/delete/:postid", DeleteBlogPost)

//--------Login--------//
	router.GET("/register", Register)
	router.POST("/register", Register)
	router.GET("/login", loginHandler)
	router.POST("/login", loginHandler)
	router.POST("/logout", logoutHandler)

//--------FileServer--------//
	router.NotFound = http.FileServer(http.Dir(""))
	router.ServeFiles("/img/*filepath", http.Dir("uploads"))
	router.ServeFiles("/static/*filepath", http.Dir("static"))

//--------Server--------//
	log.Println("Starting erver on ", PORT)
	err := http.ListenAndServe(PORT, router)
//err := http.ListenAndServe(GetPort(), router)
 	if err != nil {
		log.Fatal("error starting http server : ", router)
 	}

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method == "GET" {
		if isLoggedIn(r) {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
			tmpl.ExecuteTemplate(w, "index.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		} else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}
/*
func CustomPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method == "GET" {
		if isLoggedIn(r) {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
			tmpl.ExecuteTemplate(w, "custompage.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		} else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}*/