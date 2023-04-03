package main

import (
	"classic"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hangman-web/utils"
	"net/http"
	"strings"
	"text/template"
)

type HangManData struct {
	ID                     string
	Score                  int
	ScoreAdd               int
	WordToSearchUpper      string
	WordToSearch           string
	WordAffiche            string
	WordListLettre         []string
	WordListLettreToSearch []string
	Attemp                 int
	LetterTested           string
	LettreAlreadyTest      string
}

const port = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	UserName := HangManData{
		ID: strings.ToUpper(r.FormValue("ID")),
	}
	renderTemplateHome(w, "page_start", UserName, r)
}

func level(w http.ResponseWriter, r *http.Request) {
	RenderTemplateLevel(w, "page_level", r)
}
func game(w http.ResponseWriter, r *http.Request) {
	Lettre := HangManData{
		LetterTested: r.FormValue("choice"),
	}
	RenderTemplateGame(w, "page_template", r, Lettre)
}

func Loose(w http.ResponseWriter, r *http.Request) {
	RenderTemplateLooseWin(w, "Loose", r)
}

func Win(w http.ResponseWriter, r *http.Request) {
	RenderTemplateLooseWin(w, "Win", r)
}

func RenderTemplateLooseWin(w http.ResponseWriter, tmpl string, r *http.Request) {
	t, _ := template.ParseFiles("./template/" + tmpl + ".html")
	cook, _ := r.Cookie("DATAofGAMEcook")
	data := DecodeCook(cook)
	if r.FormValue("EndGame") == "Back" {
		data.ID = ""
		result := encodeCook(data)
		http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: result}))
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	if r.FormValue("EndGame") == "Replay" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

	t.Execute(w, data)
}

func RenderTemplateGame(w http.ResponseWriter, tmpl string, r *http.Request, lettre HangManData) {

	t, _ := template.ParseFiles("./template/" + tmpl + ".html")
	lettre.LetterTested = classic.Lower(lettre.LetterTested)
	cook, _ := r.Cookie("DATAofGAMEcook")
	dataToChange := DecodeCook(cook)

	if r.FormValue("game") == "Back" {
		dataToChange.ID = ""
		result := encodeCook(dataToChange)
		http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: result}))
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	if r.FormValue("game") == "Replay" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

	print(dataToChange.WordToSearch, "\n")
	if dataToChange.Attemp <= 0 {
		http.Redirect(w, r, "/Loose", http.StatusSeeOther)
		t.Execute(w, dataToChange)
	}
	if len(lettre.LetterTested) == 1 {
		test := utils.VerifLetterIsIn(lettre.LetterTested, dataToChange.WordListLettre, dataToChange.LettreAlreadyTest, dataToChange.WordListLettreToSearch)
		if test == true {
			dataToChange.LettreAlreadyTest += lettre.LetterTested
			dataToChange.WordListLettreToSearch = utils.ChangeLettre(lettre.LetterTested, dataToChange.WordListLettre, dataToChange.WordListLettreToSearch)
			dataToChange.WordAffiche = utils.VisualWord(dataToChange.WordListLettreToSearch)
			DataForCookies := encodeCook(dataToChange)
			if utils.VerifWords(dataToChange.WordAffiche) == true {
				dataToChange.ScoreAdd = 31456 * dataToChange.Attemp
				dataToChange.Score += dataToChange.ScoreAdd
				DataForCookies = encodeCook(dataToChange)
				http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: DataForCookies}))
				http.Redirect(w, r, "/Win", http.StatusSeeOther)
				t.Execute(w, dataToChange)
			}
			http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: DataForCookies}))
			t.Execute(w, dataToChange)
			return
		} else {
			if utils.LettreIsUse(lettre.LetterTested, dataToChange.LettreAlreadyTest) == true {
				DataForCookies := encodeCook(dataToChange)
				http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: DataForCookies}))
				t.Execute(w, dataToChange)
				return
			}
			dataToChange.LettreAlreadyTest += lettre.LetterTested
			dataToChange.Attemp--
			if dataToChange.Attemp <= 0 {
				http.Redirect(w, r, "/Loose", http.StatusSeeOther)
				t.Execute(w, dataToChange)
			}
			DataForCookies := encodeCook(dataToChange)
			http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: DataForCookies}))
			t.Execute(w, dataToChange)
			return
		}
	} else if lettre.LetterTested != "" {
		if lettre.LetterTested == dataToChange.WordToSearch {
			dataToChange.ScoreAdd = 31456 * dataToChange.Attemp
			dataToChange.Score += dataToChange.ScoreAdd
			DataForCookies := encodeCook(dataToChange)
			http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: DataForCookies}))
			print("Le score Final Attendu :", dataToChange.Score)
			http.Redirect(w, r, "/Win", http.StatusSeeOther)
			t.Execute(w, dataToChange)
		}
		dataToChange.Attemp -= 2
		if dataToChange.Attemp <= 0 {
			http.Redirect(w, r, "/Loose", http.StatusSeeOther)
			t.Execute(w, dataToChange)
		}
		DataForCookies := encodeCook(dataToChange)
		http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: DataForCookies}))
		t.Execute(w, dataToChange)
		return
	}

	DataForCookies := encodeCook(dataToChange)
	http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: DataForCookies}))

	CookiesToRead, err := r.Cookie("DATAofGAMEcook")
	if err != nil {
		data := CreateCookies(w)
		t.Execute(w, data)
		return
	}
	data := DecodeCook(CookiesToRead)
	t.Execute(w, data)
}

func RenderTemplateLevel(w http.ResponseWriter, tmpl string, r *http.Request) {
	_, err2 := r.Cookie("DATAofGAMEcook")
	if err2 != nil {
		CreateCookies(w)
	}
	reset_game(w, r)
	t, _ := template.ParseFiles("./template/" + tmpl + ".html")

	if r.FormValue("Level") == "Easy" {
		ChoiceLevel(w, r, "")
	}
	if r.FormValue("Level") == "Medium" {
		ChoiceLevel(w, r, "2")
	}
	if r.FormValue("Level") == "Hard" {
		ChoiceLevel(w, r, "3")
	}
	if r.FormValue("Level") == "Extrem" {
		ChoiceLevel(w, r, "4")
	}
	if r.FormValue("Level") == "Back" {
		cookiesForId, _ := r.Cookie("DATAofGAMEcook")
		data := DecodeCook(cookiesForId)
		data.ID = ""
		result := encodeCook(data)
		http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: result}))
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	CookiesToRead, err := r.Cookie("DATAofGAMEcook")
	if err != nil {
		data := CreateCookies(w)
		t.Execute(w, data)
		return
	}

	dataFinal := DecodeCook(CookiesToRead)
	t.Execute(w, dataFinal)
}

func renderTemplateHome(w http.ResponseWriter, tmpl string, dataIN HangManData, r *http.Request) {
	t, _ := template.ParseFiles("./template/" + tmpl + ".html")

	cookiesForId, err := r.Cookie("DATAofGAMEcook")
	if err != nil {
		data := CreateCookies(w)
		t.Execute(w, data)
		return
	}
	data := DecodeCook(cookiesForId)

	if data.ID == "" {
		if dataIN.ID == "" {
			t.Execute(w, data)
			return
		}
		data.ID = dataIN.ID
		data.Score = 0
		CookiesToSave := encodeCook(data)
		http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: CookiesToSave}))
		http.Redirect(w, r, "/level", http.StatusSeeOther)
		t.Execute(w, data)
		return
	}
	http.Redirect(w, r, "/level", http.StatusSeeOther)
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/level", level)
	http.HandleFunc("/game", game)
	http.HandleFunc("/Loose", Loose)
	http.HandleFunc("/Win", Win)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("(http://localhost:8080/home) - Server started on port", port)
	http.ListenAndServe("localhost:8080", nil)
}

func encodeCook(data HangManData) string {
	datajson, _ := json.Marshal(data)
	return base64.URLEncoding.EncodeToString(datajson)
}

func DecodeCook(Cookies *http.Cookie) HangManData {

	FinalStruct := HangManData{}
	StructJson, _ := base64.URLEncoding.DecodeString(Cookies.Value)
	_ = json.Unmarshal(StructJson, &FinalStruct)
	return FinalStruct
}

func ChoiceLevel(w http.ResponseWriter, r *http.Request, lvl string) {
	word := utils.Word_result(lvl)
	Cook, err := r.Cookie("DATAofGAMEcook")
	if err != nil {
		CreateCookies(w)
	}
	data := DecodeCook(Cook)
	data.WordToSearch = word
	data.WordToSearchUpper = strings.ToUpper(data.WordToSearch)
	data.WordListLettre = utils.CutWord(word)
	data.WordListLettreToSearch = classic.Word_choice(word)
	data.WordAffiche = utils.VisualWord(data.WordListLettreToSearch)
	DataForCookies := encodeCook(data)
	http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: DataForCookies}))
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func CreateCookies(w http.ResponseWriter) HangManData {
	data := HangManData{
		ID:                     "",
		Score:                  0,
		WordToSearch:           "",
		WordListLettre:         []string{},
		WordListLettreToSearch: []string{},
		Attemp:                 10,
		LettreAlreadyTest:      "",
		WordAffiche:            "",
	}
	Cookies := encodeCook(data)

	http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: Cookies}))
	return data
}

func reset_game(w http.ResponseWriter, r *http.Request) {
	cookies, err := r.Cookie("DATAofGAMEcook")
	var data HangManData
	if err != nil {
		data = CreateCookies(w)
		data.ID = "Gest"
	} else {
		data = DecodeCook(cookies)
	}
	data.Attemp = 10
	data.LettreAlreadyTest = ""
	Cookies := encodeCook(data)

	http.SetCookie(w, utils.CreateCookie(utils.SCookieFunc{Name: "DATAofGAMEcook", Value: Cookies}))
}
