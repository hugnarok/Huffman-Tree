package decoder

import (
	"huffman/internal/compressor"
	"strings"
)

// DecodeText decodifica um texto comprimido usando os tokens e códigos
// Os tokens já preservam a ordem e os separadores, então só precisamos
// mapear os bits de volta para as palavras
func DecodeText(result compressor.CompressedResult) string {
	var decoded strings.Builder
	bitIndex := 0
	
	// Cria um mapa reverso: código -> palavra
	codeToWord := make(map[string]string)
	for word, code := range result.Codes {
		codeToWord[code] = word
	}
	
	// Percorre os tokens na ordem original
	for _, token := range result.Tokens {
		if token.IsWord {
			// É uma palavra: precisa decodificar do código binário
			// Procura o código correspondente nos bits restantes
			found := false
			maxCodeLen := 0
			for code := range codeToWord {
				if len(code) > maxCodeLen {
					maxCodeLen = len(code)
				}
			}
			
			// Tenta encontrar o código mais longo que corresponde
			for codeLen := 1; codeLen <= maxCodeLen && bitIndex+codeLen <= len(result.CompressedBits) && !found; codeLen++ {
				currentBits := result.CompressedBits[bitIndex : bitIndex+codeLen]
				if word, exists := codeToWord[currentBits]; exists {
					// Verifica se a palavra corresponde ao token esperado
					if word == token.Value {
						decoded.WriteString(word)
						bitIndex += codeLen
						found = true
					}
				}
			}
			
			if !found {
				// Erro: não conseguiu decodificar
				decoded.WriteString("[ERRO]")
				bitIndex++ // Avança um bit para não travar
			}
		} else {
			// É um separador: restaura diretamente
			decoded.WriteString(token.Value)
		}
	}
	
	return decoded.String()
}

// SerializeTokens serializa os tokens para armazenamento no output.dat
func SerializeTokens(tokens []compressor.Token) string {
	var builder strings.Builder
	for i, token := range tokens {
		if i > 0 {
			builder.WriteString("|")
		}
		if token.IsWord {
			builder.WriteString("W:" + token.Value)
		} else {
			// Escapa separadores especiais
			separator := token.Value
			if separator == "|" {
				separator = "\\|"
			} else if separator == "\n" {
				separator = "\\n"
			} else if separator == "\r" {
				separator = "\\r"
			} else if separator == "\t" {
				separator = "\\t"
			} else if separator == "\\" {
				separator = "\\\\"
			}
			builder.WriteString("S:" + separator)
		}
	}
	return builder.String()
}

// DeserializeTokens deserializa tokens do formato armazenado
func DeserializeTokens(data string) []compressor.Token {
	var tokens []compressor.Token
	parts := strings.Split(data, "|")
	
	for _, part := range parts {
		if strings.HasPrefix(part, "W:") {
			// É uma palavra
			tokens = append(tokens, compressor.Token{
				Type:   "word",
				Value:  part[2:],
				IsWord: true,
			})
		} else if strings.HasPrefix(part, "S:") {
			// É um separador
			separator := part[2:]
			// Desescapa caracteres especiais
			if separator == "\\|" {
				separator = "|"
			} else if separator == "\\n" {
				separator = "\n"
			} else if separator == "\\r" {
				separator = "\r"
			} else if separator == "\\t" {
				separator = "\t"
			} else if separator == "\\\\" {
				separator = "\\"
			}
			tokens = append(tokens, compressor.Token{
				Type:   "separator",
				Value:  separator,
				IsWord: false,
			})
		}
	}
	
	return tokens
}
