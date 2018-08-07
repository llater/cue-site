package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

var CueSiteTemplate = `
    {{define "site"}}
    <!DOCTYPE html>
    <html>
      <head>
        <meta http-equiv="content-type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>{{.Title}}</title>
        <style>
           html, body { 
            margin: 0; 
            padding: 0; 
      background-color: #f5f5f5; 
      font-family: 'Helvetica Neue', Arial, Helvetica, sans-serif;
            }
            #top-container {
              margin: 0 auto;
              width: 100vw;
              height: 8vh;
              background-color: #7f1ae5;
            }
    h1 { 
      text-align: center; 
    }
    #main-container { 
      margin: 0 auto;
      height: 108vh;
      min-height: 360px;
      width: 100vw;
      display: flex;
      flex-direction: column;
      justify-content: center;
      flex-wrap: wrap; 
      background-color: #0D0D0C;
    }
    .main-content {
      color: #f5f5f5;
      font-eight: 700;
      font-size: 10vh;
      display: flex;
      justify-content: top;
      text-align: center;
      padding-left: 12px;
      min-width: 360px;
    }
    #bottom-container {
      margin: 0 auto;
      height: 80vh;
      background-color: #7f1ae5;
    }
      </style>
      </head>
      <body>
        <section id="top-container">
          <div class="top-content"></div>
        </section>
        <section id="main-container">
          <div class="main-content">
            <h1>{{.Content}}</h1>
          </div>
        </section>
        <section id="bottom-container">
          <div class="bottom-content"></div>
       </section> 
      <script>
      </script>
      </body>
    </html>
    {{end}}
`

func init() {
	log.SetOutput(os.Stdout)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("site")
		te, _ := t.Parse(CueSiteTemplate)
		type siteVars struct {
			Title   string
			Content string
		}
		cueSite := &siteVars{
			"Cue Labs",
			"Cue Labs",
		}
		if err := te.Execute(w, cueSite); err != nil {
			log.Println("Error while templating in main()")
			log.Println(err)
		}
	})
	logR := handlers.CombinedLoggingHandler(os.Stdout, handlers.ProxyHeaders(r))
	log.Fatal(http.ListenAndServe(":8000", logR))
}
