package compressor

import (
	"huffman/internal/frequency"
	"strings"
	"unicode"
)

// Token representa um elemento do texto: palavra ou separador (espaço, pontuação)
type Token struct {
	Type      string // "word" ou "separator"
	Value     string // palavra ou caractere separador
	IsWord    bool   // true se é palavra, false se é separador
}

// CompressedResult contém o texto comprimido e informações para decodificação
type CompressedResult struct {
	CompressedBits string   // sequência de bits comprimida
	Tokens         []Token  // sequência de tokens (palavras e separadores)
	Codes          map[string]string // códigos de Huffman
}

// CompressText comprime um texto preservando espaços e pontuação para decodificação
func CompressText(text string, codes map[string]string) CompressedResult {
	var compressedBits strings.Builder
	var tokens []Token
	
	// Tokeniza preservando separadores
	tokens = tokenizeWithSeparators(text)
	
	// Comprime cada token
	for _, token := range tokens {
		if token.IsWord {
			// É uma palavra: substitui pelo código
			if code, exists := codes[token.Value]; exists {
				compressedBits.WriteString(code)
			}
		} else {
			// É um separador: preserva como está (não comprime)
			// Na decodificação, os separadores serão restaurados diretamente
		}
	}
	
	return CompressedResult{
		CompressedBits: compressedBits.String(),
		Tokens:         tokens,
		Codes:          codes,
	}
}

// tokenizeWithSeparators divide o texto em tokens preservando separadores
func tokenizeWithSeparators(text string) []Token {
	var tokens []Token
	var currentWord strings.Builder
	
	// Verifica frequências para normalizar palavras
	freq := frequency.CalculateFrequency(text)
	
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			// Adiciona caractere à palavra atual
			currentWord.WriteRune(r)
		} else {
			// Fim de uma palavra (se existir)
			if currentWord.Len() > 0 {
				word := strings.ToLower(currentWord.String())
				// Só adiciona se a palavra existe no dicionário (foi contada)
				if _, exists := freq[word]; exists {
					tokens = append(tokens, Token{
						Type:   "word",
						Value:  word,
						IsWord: true,
					})
				}
				currentWord.Reset()
			}
			
			// Adiciona o separador (espaço, pontuação, etc.)
			tokens = append(tokens, Token{
				Type:   "separator",
				Value:  string(r),
				IsWord: false,
			})
		}
	}
	
	// Adiciona a última palavra se existir
	if currentWord.Len() > 0 {
		word := strings.ToLower(currentWord.String())
		if _, exists := freq[word]; exists {
			tokens = append(tokens, Token{
				Type:   "word",
				Value:  word,
				IsWord: true,
			})
		}
	}
	
	return tokens
}

// tokenize divide o texto em palavras (mesma lógica do frequency package)
// Mantida para compatibilidade
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
