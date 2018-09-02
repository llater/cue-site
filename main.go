package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

 var (
  CueSiteHTML = `
{{- define "site_html" -}}
<!DOCTYPE html>
    <html>
      <head>
        <meta http-equiv="content-type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>{{.Title}}</title>
        <style>{{template "site_css"}}</style>
      </head>
      <body>
        <section id="top-container">
          <div class="top-content"></div>
        </section>
        <section id="main-container">
          <div class="main-content">{{template "main_content" .}}</div>
        </section>
        <section id="bottom-container">
          <div class="bottom-content">
           <div class="gap"></div>
           <section id="app">
          <h1>What is Cue?</h1>
          <h2>Cue is a free app. Sign in with Spotify and play music with your friends!</h2>
 </section>
 </div>
 </section>
       <section id="base">Cue Labs 2018</section>
      <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
      </body>
    </html>
    {{end}}
    `
    MainContent = `
    {{define "main_content"}}
    <h1>{{.Header}}</h1>
    <h2>{{.Subheader}}</h2>
    <form action="/redirect-to-app" method="GET"><button type="submit">GO TO APP</button></form>
    {{end}}
    `
    CueSiteCSS = `
    {{define "site_css"}}
html, body { 
margin: 0; 
padding: 0; 
background-color: #f5f5f5; 
      font-family: 'Helvetica Neue', Arial, Helvetica, sans-serif;
            }
            #top-container {
              margin: 0 auto;
              width: 100vw;
              height: 5vh;
              background-color: #7f1ae5;
            }
    #main-container { 
      margin: 0 auto;
      height: 100vh;
      min-height: 360px;
      width: 100vw;
      display: flex;
      flex-direction: column;
      justify-content: center;
      flex-wrap: wrap; 
      background-color: #0D0D0C;
      min-width: 480px;
    }
.main-content {
      color: #f5f5f5;
      display: flex;
      flex-direction: column;
      justify-content: center;
      text-align: left;
      padding-left: 64px;
    }
.main-content h1 {
    font-size: calc(72px + (24 - 16) * (100vw - 360px)/(960-360));
    font-weight: 700;
}
.main-content h2 {
    font-size:  calc(24px + (24 - 16) * (100vw - 360px)/(960-360));
    padding-bottom: 24px;
}
.main-content button {
    font-size: calc(24px + (24 - 16) * (100vw - 360px)/(960-360));
    width: 20vw;
    height: 8vh;
    border-radius: 10px;
    border-width: 6px;
    border-color: #7f1ae5;
    font-weight: 700;
    text-align: center;
background-color: #f5f5f5;
    min-width: 190px;
}
    #bottom-container {
      margin: 0 auto;
      height: 100vh;
      width: 100vw;
      background-color: #7f1ae5;
      color: #0D0D0C;
    }

.bottom-content {
display: flex;
flex-direction: column;
justify-content: center;
text-align: left;
color: #0D0D0C;
width: 70vw;
}
.bottom-content h1, .bottom-content h2 {
  font-size: calc(48px + (24 - 16) * (100vw - 360px)/(960-360));
  text-shadow: 2px 2px #f5f5ff5;
padding-left: 64px;
}

.bottom-content h2 {
  font-size: calc(24px + (24 - 16) * (100vw - 360px)/(960-360));
}

.gap {
  width: 100vw;
  height: 28vh;
  margin: 0 auto;
}
    #base {
  margin: 0 auto;
  width: 100vw;
height: 8vh;
background-color: #0D0D0C;
  font-weight: 700;
  color: #7f1ae5;
  text-align: center;
  justify-content: flex-end;
}
    {{end}}
    `
    )

func init() {
	log.SetOutput(os.Stdout)

}

type mainContentVars struct {
	Title   string
	Header string
	Subheader string
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("site_html")
		te, _ := t.Parse(CueSiteHTML + CueSiteCSS + MainContent)

		cueSite := &mainContentVars{
			"Cue Labs",
			"Cue",
			"play music with friends",

		}
		if err := te.Execute(w, cueSite); err != nil {
			log.Println("Error while templating in main()")
			log.Println(err)
		}
	})
	http.HandleFunc("/redirect-to-app", func(w http.ResponseWriter, r *http.Request) {
		// Take browser fingerprint
		// Look for Spotify login
		log.Print("Taking a browser fingerprint...")
		log.Print("Redirecting...")
		http.Redirect(w, r, "https://app.cue.zone", 301)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
