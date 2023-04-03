package classic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	WordBase   string
	ShowWord   []string
	LetterFind string
	Attemps    int
	Position   int
	Asci       string
}

func main() {
	var varia HangManData
	if len(Verif(os.Args[1], "--startWith")) > 0 {
		content, _ := os.ReadFile("save.txt")
		err1 := json.Unmarshal(content, &varia)
		if err1 != nil {
			print(err1)
			return
		}
	} else {
		word := Random_word(os.Args[1])
		varia = HangManData{word, Word_choice(word), "", 10, -1, ""}
	}
	var letter rune
	var choice string
	fail := false

	println(varia.WordBase)
	println(len(varia.WordBase))
	println("Good Luck, you have ", varia.Attemps, " attempts.")
	ShowWord := Word_choice(varia.WordBase)
	for varia.Attemps > 0 {
		for i := 0; i < len(varia.WordBase); i++ {
			print(ShowWord[i] + " ")
		}
		print("\n" + "\n" + "Choose: ")
		fmt.Scanln(&choice)
		var listInd []int
		for i := 0; i < len(varia.WordBase); i++ {
			if choice[0] == varia.WordBase[i] {
				listInd = append(listInd, i)
			}
		}
		if len(choice) == 1 {
			letter = rune(choice[0])
			if len(Verif(varia.LetterFind, choice)) > 0 {
				varia.Attemps--
				varia.Position = Gallows(1, varia.Position)
				println("\nalready present in the word,", varia.Attemps, "attempts remaining\n")
				fail = true
			}
			varia.LetterFind += choice
			if len(Verif(varia.WordBase, choice)) >= 1 {
				index := Verif(varia.WordBase, choice)
				for i := 0; i < len(index); i++ {
					ShowWord[index[i]] = string(letter - 32)
				}
			} else {
				if !fail {
					varia.Attemps--
					varia.Position = Gallows(1, varia.Position)
					println("\nNot present in the word,", varia.Attemps, "attempts remaining\n")
					fail = false
				}
			}
		} else {
			if choice == varia.WordBase {
				println("\nCongrats !")
				return
			} else if choice == "stop" {
				b, _ := json.Marshal(varia)
				save, _ := os.Create("save.txt")
				save.Write(b)
				return
			} else {
				varia.Attemps -= 2
				varia.Position = Gallows(2, varia.Position)
				println("\nlie! is not the real word,", varia.Attemps, "attempts remaining\n")
			}
		}
		if len(Verif(Listtostring(ShowWord), "_")) == 0 {
			for i := 0; i < len(varia.WordBase); i++ {
				print(ShowWord[i] + " ")
			}
			println("\n\nCongrats !")
			return
		}
	}
	println("You are bad, it's was :", varia.WordBase)
}

func ReadFiles(files ...string) []string {
	var word_list []string
	for _, fileName := range files {
		file, _ := os.Open(fileName)
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			word_list = append(word_list, fileScanner.Text())
		}
	}
	return word_list
}

func Random_word_With_List(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func Verif(word, choice string) []int {
	var ListInd []int
	for i := 0; i < len(word); i++ {
		if choice[0] == word[i] {
			ListInd = append(ListInd, i)
		}
	}
	return ListInd
}

func Gallows(nbr, position int) int {
	jose, _ := os.ReadFile("hangman.txt")
	position += 79 * nbr
	if position >= 790 {
		position = 790
	}
	fmt.Print(string(jose[position-78 : position]))
	return position
}

func Listtostring(list []string) string {
	char := ""
	for i := 0; i < len(list); i++ {
		char += list[i]
	}
	return char
}

func Lower(choice string) string {
	choice3 := ""
	for h := 0; h < len(choice); h++ {
		if choice[h] >= 65 && choice[h] <= 90 {
			choice3 += string(choice[h] + 32)
		} else {
			choice3 += string(choice[h])
		}
	}
	return choice3
}

func attemps(choice string, letter rune, word string) {
	LetterFind := ""
	ShowWord := Word_choice(word)
	fail := false
	Attemps := 10
	if len(choice) == 1 {
		letter = rune(choice[0])
		if len(Verif(LetterFind, choice)) > 0 {
			Attemps--
			fail = true
		}
		LetterFind += choice
		if len(Verif(word, choice)) >= 1 {
			index := Verif(word, choice)
			for i := 0; i < len(index); i++ {
				ShowWord[index[i]] = string(letter - 32)
			}
		} else {
			if !fail {
				Attemps--
				fail = false
			}
		}
	} else {
		if choice == word {
			return
		} else {
			Attemps -= 2
		}
	}
}

func Letter(word string) {
	var choice string
	fmt.Scan(&choice)
	var listInd []int
	for i := 0; i < len(word); i++ {
		if choice[0] == word[i] {
			listInd = append(listInd, i)
		}
	}
}

func Word_choice(word string) []string {
	var show_word []string
	nbrletter := len(word)/2 - 1
	for i := 0; i < len(word); i++ {
		show_word = append(show_word, "_")
	}
	for x := 0; x < nbrletter; x++ {
		ind := rand.Intn(len(word))
		show_word[ind] = string(word[ind])
	}
	return show_word
}

func Random_word(word string) string {
	file, _ := os.Open(word)
	var word_list []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		word_list = append(word_list, fileScanner.Text())
	}
	rand.Seed(time.Now().UnixNano())
	return word_list[rand.Intn(len(word))]
}
