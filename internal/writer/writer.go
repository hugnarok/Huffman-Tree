package writer

import (
	"fmt"
	"huffman/internal/compressor"
	"huffman/internal/huffman"
	"huffman/internal/serializer"
	"os"
	"sort"
	"strings"
)

// CompressedText contém todas as informações de um texto comprimido
type CompressedText struct {
	Tree           *huffman.Node
	Codes          map[string]string
	CompressedData string
	OriginalText   string
}

// WriteOutput escreve os resultados da compressão no arquivo output.dat
func WriteOutput(filename string, results []CompressedText) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escreve cada texto comprimido
	for i, result := range results {
		writeTextResult(file, i+1, result)

		// Adiciona separador entre textos (exceto no último)
		if i < len(results)-1 {
			file.WriteString("\n" + strings.Repeat("=", 80) + "\n\n")
		}
	}

	return nil
}

// writeTextResult escreve o resultado da compressão de um único texto
func writeTextResult(file *os.File, textNumber int, result CompressedText) {
	// Cabeçalho
	file.WriteString(fmt.Sprintf("%s TEXTO %d %s\n", strings.Repeat("=", 33), textNumber, strings.Repeat("=", 33)))
	file.WriteString("\n")

	// Seção: Árvore de Huffman
	file.WriteString("[ÁRVORE DE HUFFMAN]\n")
	file.WriteString(strings.Repeat("-", 80) + "\n")
	file.WriteString(serializer.SerializeTree(result.Tree))
	file.WriteString("\n")

	// Seção: Visualização da Árvore
	file.WriteString("[VISUALIZAÇÃO DA ÁRVORE]\n")
	file.WriteString(strings.Repeat("-", 80) + "\n")
	file.WriteString(serializer.VisualizeTree(result.Tree))
	file.WriteString("\n")

	// Seção: Representação Compacta da Árvore
	file.WriteString("[REPRESENTAÇÃO COMPACTA DA ÁRVORE (para decodificação)]\n")
	file.WriteString(strings.Repeat("-", 80) + "\n")
	file.WriteString(serializer.SerializeTreeCompact(result.Tree))
	file.WriteString("\n\n")

	// Seção: Códigos Gerados
	file.WriteString("[CÓDIGOS GERADOS]\n")
	file.WriteString(strings.Repeat("-", 80) + "\n")
	writeCodes(file, result.Codes)
	file.WriteString("\n")

	// Seção: Texto Original
	file.WriteString("[TEXTO ORIGINAL]\n")
	file.WriteString(strings.Repeat("-", 80) + "\n")
	file.WriteString(result.OriginalText)
	file.WriteString("\n\n")

	// Seção: Texto Comprimido
	file.WriteString("[TEXTO COMPRIMIDO (bits)]\n")
	file.WriteString(strings.Repeat("-", 80) + "\n")
	writeCompressedData(file, result.CompressedData)
	file.WriteString("\n\n")

	// Seção: Estatísticas
	file.WriteString("[ESTATÍSTICAS DE COMPRESSÃO]\n")
	file.WriteString(strings.Repeat("-", 80) + "\n")
	writeStats(file, result)
	file.WriteString("\n")
}

// writeCodes escreve os códigos gerados ordenados alfabeticamente
func writeCodes(file *os.File, codes map[string]string) {
	// Ordena as palavras alfabeticamente
	words := make([]string, 0, len(codes))
	for word := range codes {
		words = append(words, word)
	}
	sort.Strings(words)

	// Escreve cada código
	for _, word := range words {
		file.WriteString(fmt.Sprintf("%-20s: %s\n", word, codes[word]))
	}
}

// writeCompressedData escreve os dados comprimidos com quebras de linha para legibilidade
func writeCompressedData(file *os.File, compressedData string) {
	const lineWidth = 80

	for i := 0; i < len(compressedData); i += lineWidth {
		end := i + lineWidth
		if end > len(compressedData) {
			end = len(compressedData)
		}
		file.WriteString(compressedData[i:end] + "\n")
	}
}

// writeStats escreve as estatísticas de compressão
func writeStats(file *os.File, result CompressedText) {
	originalBits, compressedBits, compressionRate := compressor.CalculateCompressionStats(
		result.OriginalText,
		result.CompressedData,
		result.Codes,
	)

	file.WriteString(fmt.Sprintf("Tamanho original:      %d bits (%d bytes)\n", originalBits, originalBits/8))
	file.WriteString(fmt.Sprintf("Tamanho comprimido:    %d bits (%d bytes + overhead)\n", compressedBits, compressedBits/8))
	file.WriteString(fmt.Sprintf("Taxa de compressão:    %.2f%%\n", compressionRate))
	file.WriteString(fmt.Sprintf("Número de palavras únicas: %d\n", len(result.Codes)))

	// Calcula comprimento médio dos códigos
	totalCodeLength := 0
	for _, code := range result.Codes {
		totalCodeLength += len(code)
	}
	avgCodeLength := float64(totalCodeLength) / float64(len(result.Codes))
	file.WriteString(fmt.Sprintf("Comprimento médio dos códigos: %.2f bits\n", avgCodeLength))
}
