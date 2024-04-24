package main

import (
    "net/http"
    "html/template"
    "log"
    "os"
	"math/rand"
	"strings"
)

var HangManeasy HangManData
var HangManhard HangManData
var HangMan HangManData
var playerChoice string

// Page représente les données nécessaires pour rendre une page HTML
type Page struct {
    Title                string
    Content              string
    Wordpage             string
	ToFindpage           string
	Attemptspage         int
	HangmanPositionspage string
}

//Type HangManData
type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions []string
}

// fonction pour utiliser les templates

func renderTemplate(w http.ResponseWriter, tmpl string, p Page) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// fonction pour ma page d'acceuil

func indexHandler(w http.ResponseWriter, r *http.Request) {

	page := Page{
		Title:              "THE" + "\n" + "HANGMAN",
		Content:            "Good Luck",
	}
	renderTemplate(w, "index", page)
}

func PageEasy(w http.ResponseWriter, r *http.Request, tmpl string) {
    r.ParseForm()
    playerChoice = r.Form.Get("message")
    handlePlayerChoice(w, r, playerChoice, tmpl)

    page := Page{
        Title:                "THE" + "\n" + "HANGMAN",
        ToFindpage:           HangMan.ToFind,
        Wordpage:             separateWithSpace(HangMan.Word),
        Attemptspage:         HangMan.Attempts,
        HangmanPositionspage: HangMan.HangmanPositions[len(HangMan.HangmanPositions) - 1 - HangMan.Attempts],
    }
    renderTemplate(w, tmpl, page)
}

func PageNormal(w http.ResponseWriter, r *http.Request, tmpl string) {
    r.ParseForm()
    playerChoice = r.Form.Get("message")
    handlePlayerChoice(w, r, playerChoice, tmpl)

    page := Page{
        Title:                "THE" + "\n" + "HANGMAN",
        ToFindpage:           HangMan.ToFind,
        Wordpage:             separateWithSpace(HangMan.Word),
        Attemptspage:         HangMan.Attempts,
        HangmanPositionspage: HangMan.HangmanPositions[len(HangMan.HangmanPositions) - 1 - HangMan.Attempts],
    }
    renderTemplate(w, tmpl, page)
}

func PageHard(w http.ResponseWriter, r *http.Request, tmpl string) {
    r.ParseForm()
    playerChoice = r.Form.Get("message")
    handlePlayerChoice(w, r, playerChoice, tmpl)

    page := Page{
        Title:                "THE" + "\n" + "HANGMAN",
        ToFindpage:           HangMan.ToFind,
        Wordpage:             separateWithSpace(HangMan.Word),
        Attemptspage:         HangMan.Attempts,
        HangmanPositionspage: HangMan.HangmanPositions[len(HangMan.HangmanPositions) - 1 - HangMan.Attempts],
    }
    renderTemplate(w, tmpl, page)
}

func page1Handler(w http.ResponseWriter, r *http.Request) {
    PageEasy(w, r, "page1")
}

func pagereseteasy(w http.ResponseWriter, r *http.Request) {
    resetHangman(&HangMan, 1)
    PageEasy(w, r, "page1")
}

func page2Handler(w http.ResponseWriter, r *http.Request) {
    PageNormal(w, r, "page2")
}

func pageresetnormal(w http.ResponseWriter, r *http.Request) {
    resetHangman(&HangMan, 2)
    PageNormal(w, r, "page2")
}

func page3Handler(w http.ResponseWriter, r *http.Request) {
    PageHard(w, r, "page3")
}

func pageresethard(w http.ResponseWriter, r *http.Request) {
    resetHangman(&HangMan, 3)
    PageHard(w, r, "page3")
}

func winHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title: 				"You Win",
	}
	renderTemplate(w, "win", page)
}

func lostHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title: 				"You Lost",
	}
	renderTemplate(w, "lost", page)
}

// verifie si la lettre est dans le mot ou non 

func handlePlayerChoice(w http.ResponseWriter, r *http.Request, playerChoice string, tmpl string) {
    playerChoice = strings.ToUpper(playerChoice)

    if HangMan.Attempts > 0 && len(playerChoice) > 0 {
        success := false
        for i, c := range HangMan.ToFind {
            if string(c) == playerChoice {
                HangMan.Word = HangMan.Word[:i] + string(c) + HangMan.Word[i+1:]
                success = true
            }
        }
        if !success {
            HangMan.Attempts--
        }
        if HangMan.Attempts == 0 {
            http.Redirect(w, r, "/lost", http.StatusSeeOther)
            return
        }
        if HangMan.ToFind == HangMan.Word {
            http.Redirect(w, r, "/win", http.StatusSeeOther)
            return
        }
    }
}


// fonction pour séparer mon fichier hangman.txt

func separateWithSpace(input string) string {
	result := strings.Join(strings.Split(input, ""), " ")
	return result
}

// fonciton pour reset le hangman

func resetHangman(HangMan *HangManData, mode int) {

	HangMan.Attempts=10

	//=== Ouverture du fichier hangman.txt et ajout des positions à notre type HangManData ===
	content, err := os.ReadFile("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	HangMan.HangmanPositions=[]string(strings.Split(string(content), ","))

	//=== Ouverture du fichier words.txt et choix du mot ===
	if mode == 1 {
		content, err = os.ReadFile("words.txt")
	} else if mode == 2 {
		content, err = os.ReadFile("words2.txt")
	} else {
		content, err = os.ReadFile("words3.txt")
	}

	if err != nil {
		log.Fatal(err)
	}
	words:=strings.Split(string(content), "\n")
	HangMan.ToFind=strings.ToUpper(words[rand.Intn(len(words))])

	//=== Creation du mot vide ===
	wordHidden:=""
	lastLetter:=string(HangMan.ToFind[len(HangMan.ToFind)-1])
	for i:=0; i<len(HangMan.ToFind)-1; i++{
		if string(HangMan.ToFind[i])==lastLetter{
			wordHidden+=lastLetter
		} else {
			wordHidden+="_"
		}
	}
	wordHidden+=string(HangMan.ToFind[len(HangMan.ToFind)-1])
	HangMan.Word=wordHidden
}

func main() {

    resetHangman(&HangManeasy, 1)
	resetHangman(&HangMan, 2)
	resetHangman(&HangManhard, 3)

	fs := http.FileServer(http.Dir("./static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/reseteasy", pagereseteasy)
	http.HandleFunc("/resetnormal", pageresetnormal)
	http.HandleFunc("/resethard", pageresethard)
	http.HandleFunc("/page1", page1Handler)
	http.HandleFunc("/page2", page2Handler)
    http.HandleFunc("/page3", page3Handler)
	http.HandleFunc("/win", winHandler)
	http.HandleFunc("/lost", lostHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
