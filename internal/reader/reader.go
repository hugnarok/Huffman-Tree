package reader

import (
	"bufio"
	"os"
	"strings"
)

// ReadTexts lê um arquivo e retorna uma lista de textos separados por linhas em branco
func ReadTexts(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var texts []string
	var currentText strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Se a linha está vazia, é um separador de textos
		if line == "" {
			if currentText.Len() > 0 {
				texts = append(texts, strings.TrimSpace(currentText.String()))
				currentText.Reset()
			}
		} else {
			// Adiciona a linha ao texto atual
			if currentText.Len() > 0 {
				currentText.WriteString(" ")
			}
			currentText.WriteString(line)
		}
	}

	// Adiciona o último texto se existir
	if currentText.Len() > 0 {
		texts = append(texts, strings.TrimSpace(currentText.String()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return texts, nil
}
