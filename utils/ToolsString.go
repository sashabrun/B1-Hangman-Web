package utils

import "classic"

func VisualWord(word []string) string {
	finalWord := ""
	for i := 0; i < len(word); i++ {
		finalWord += word[i]
		finalWord += " "
	}
	return finalWord
}

func Word_result(level string) string {
	word := classic.Random_word("words" + level + ".txt")
	return word
}

func VerifLetterIsIn(lettre string /*lettre*/, chaine []string /*mot chercher*/, lettrealready string /*lst letrte utiliser*/, lstlettre []string /*mot en crypt*/) bool {
	for i := 0; i < len(lettrealready); i++ {
		if lettre[0] == lettrealready[0] {
			return false
		}
	}
	for i := 0; i < len(chaine); i++ {
		if chaine[i] == lettre && chaine[i][0] != lstlettre[i][0] {
			return true
		}
	}
	return false
}

func ChangeLettre(lettre string, chaine []string, final []string) []string {
	for i := 0; i < len(chaine); i++ {
		if chaine[i] == lettre {
			final[i] = chaine[i]
		}
	}
	return final
}

func CutWord(word string) []string {
	var lstWord []string
	for i := 0; i < len(word); i++ {
		lstWord = append(lstWord, string(word[i]))
	}
	return lstWord
}

func VerifWords(Word string) bool {
	for i := 0; i < len(Word); i++ {
		if Word[i] == '_' {
			return false
		}
	}
	return true
}

func LettreIsUse(lettre string, LettreAlreadyUse string) bool {
	for i := 0; i < len(LettreAlreadyUse); i++ {
		if lettre[0] == LettreAlreadyUse[i] {
			return true
		}
	}
	return false
}
