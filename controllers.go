package main
import (
//	"net/http"
	"fmt"
	"log"
	"io"
	"os/exec"
	 
//    "github.com/julienschmidt/httprouter"
)

func testRedis() {
	println(rdxGet("ac"))
	println(rdxSet("ac", "{Name: 'ac', Age: 25}"))
}

func init() {
	go func() {
	cmd := exec.Command("redis-server")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	}()
	go testRedis()
	
}

/*
func CustomPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	switch r.Method {
		case "GET" : {
			if isLoggedIn(r) {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nav.html", nil)
				tmpl.ExecuteTemplate(w, "custom.html", "custom")
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			} 	else {
				tmpl.ExecuteTemplate(w, "head.html", nil)
				tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
				tmpl.ExecuteTemplate(w, "footer.html", nil)
			}
		}
		case "POST" : {
			fmt.Println("POST")
		}
		default : {
			fmt.Println("Method Not allowed")
		}
	}
}*/
/*

type Person struct {
First string
Last string
}
func init() {*/
    /* This will marshal the JSON into []bytes */
/*
    p1 := Person{"alice", "bob"}
    bs, _ := json.Marshal(p1)
    fmt.Println(string(bs))
*/
    /* This will unmarshal the JSON from []bytes */
/*
    var p2 Person
    bs = []byte(`{"First":"alice","Last":"bob"}`)
    json.Unmarshal(bs, &p2)
    fmt.Println(p2)

}*/