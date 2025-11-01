package frequency

import (
	"strings"
	"unicode"
)

// CalculateFrequency calcula a frequência de cada palavra em um texto
// Normaliza para lowercase e remove pontuação
func CalculateFrequency(text string) map[string]int {
	frequencies := make(map[string]int)

	// Tokeniza o texto em palavras
	words := tokenize(text)

	// Conta a frequência de cada palavra
	for _, word := range words {
		if word != "" {
			frequencies[word]++
		}
	}

	return frequencies
}

// tokenize divide o texto em palavras, removendo pontuação e convertendo para lowercase
func tokenize(text string) []string {
	var words []string
	var currentWord strings.Builder

	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			// Adiciona caractere à palavra atual (em lowercase)
			currentWord.WriteRune(unicode.ToLower(r))
		} else if currentWord.Len() > 0 {
			// Fim de uma palavra
			words = append(words, currentWord.String())
			currentWord.Reset()
		}
	}

	// Adiciona a última palavra se existir
	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words
}

// GetWords retorna uma lista ordenada de palavras únicas do texto
func GetWords(text string) []string {
	freq := CalculateFrequency(text)
	words := make([]string, 0, len(freq))

	for word := range freq {
		words = append(words, word)
	}

	return words
}
