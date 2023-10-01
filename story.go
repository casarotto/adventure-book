package adventurebook

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("story").Parse(defaultHandlerTmpl))
}

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Story}}
    <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li><a href="{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </body>
</html>
`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chapter := r.URL.Path[1:]
	if chapter == "" {
		chapter = "intro"
	}

	err := tmpl.Execute(w, h.s[chapter])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func JsonStory(reader *os.File) (Story, error) {
	d := json.NewDecoder(reader)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Chapter
