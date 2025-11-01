package compressor

import (
	"huffman/internal/frequency"
	"strings"
)

// CompressText comprime um texto substituindo cada palavra pelo seu código binário
func CompressText(text string, codes map[string]string) string {
	var compressed strings.Builder

	// Tokeniza o texto da mesma forma que o analisador de frequências
	words := tokenize(text)

	// Substitui cada palavra pelo código correspondente
	for _, word := range words {
		if code, exists := codes[word]; exists {
			compressed.WriteString(code)
		}
	}

	return compressed.String()
}

// tokenize divide o texto em palavras (mesma lógica do frequency package)
func tokenize(text string) []string {
	// Reutiliza a lógica de tokenização do package frequency
	// Para garantir consistência
	freq := frequency.CalculateFrequency(text)

	// Reconstrói a lista de palavras na ordem original
	var words []string
	var currentWord strings.Builder

	for _, r := range text {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			currentWord.WriteRune(r)
		} else if currentWord.Len() > 0 {
			word := strings.ToLower(currentWord.String())
			if _, exists := freq[word]; exists {
				words = append(words, word)
			}
			currentWord.Reset()
		}
	}

	// Adiciona a última palavra se existir
	if currentWord.Len() > 0 {
		word := strings.ToLower(currentWord.String())
		if _, exists := freq[word]; exists {
			words = append(words, word)
		}
	}

	return words
}

// CalculateCompressionStats calcula estatísticas de compressão
func CalculateCompressionStats(originalText string, compressedBits string, codes map[string]string) (originalBits, compressedBitsCount int, compressionRate float64) {
	// Tamanho original em bits (assumindo 8 bits por caractere)
	originalBits = len(originalText) * 8

	// Tamanho comprimido em bits
	compressedBitsCount = len(compressedBits)

	// Taxa de compressão
	if originalBits > 0 {
		compressionRate = (1.0 - float64(compressedBitsCount)/float64(originalBits)) * 100
	}

	return
}
