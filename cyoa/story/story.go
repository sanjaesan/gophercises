package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

type storyHandler struct {
	story Story
	tpl   *template.Template
}

// NewStoryHandler --
func NewStoryHandler(s Story, t *template.Template) http.Handler {
	return storyHandler{s, t}
}

func (s storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == " " || path == "/" {
		path = "/intro"
	}
	//Trim trailing slash
	path = path[1:]

	storyArc, present := s.story[path]
	if present {
		err := s.tpl.Execute(w, storyArc)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong ...", http.StatusNotFound)
		}
		return
	}
	http.Error(w, "chapter not found", http.StatusNotFound)

}

//MakeTemplate -
func MakeTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}

// Story --
type Story map[string]StoryArc

// StoryArc --
type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option --
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

//JsonifyStory --- convert story file to json
func JsonifyStory(r io.Reader) (Story, error) {
	var storyMap Story
	err := json.NewDecoder(r).Decode(&storyMap)
	if err != nil {
		panic(err)
	}
	return storyMap, nil
}
