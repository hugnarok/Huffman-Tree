package main

import (
	"fmt"
	"huffman/internal/compressor"
	"huffman/internal/decoder"
	"huffman/internal/encoder"
	"huffman/internal/frequency"
	"huffman/internal/huffman"
	"huffman/internal/reader"
	"huffman/internal/writer"
	"log"
	"os"
)

const (
	inputFile  = "data/input.dat"
	outputFile = "data/output.dat"
)

func main() {
	// Verifica se o arquivo de entrada existe
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("Error: input file '%s' not found", inputFile)
	}

	// 1. Ler textos do arquivo de entrada
	texts, err := reader.ReadTexts(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	if len(texts) == 0 {
		log.Fatal("Error: no text found in input file")
	}

	// 2. Processar cada texto
	var results []writer.CompressedText
	var i int = 1
	for _, text := range texts {

		// Calcular frequências das palavras
		fmt.Printf("Sentence: %d\n", i)
		frequencies := frequency.CalculateFrequency(text)
		fmt.Printf("     -> %d unique word(s) found\n", len(frequencies))

		// Construir árvore de Huffman
		tree := huffman.BuildHuffmanTree(frequencies)

		// Gerar códigos binários
		codes := encoder.GenerateCodes(tree)
		fmt.Printf("     -> %d code(s) generated\n", len(codes))

		// Comprimir texto (agora preserva espaços e pontuação)
		compressedResult := compressor.CompressText(text, codes)
		fmt.Printf("     -> Text compressed into %d bits\n", len(compressedResult.CompressedBits))

		// Decodificar para validar
		decodedText := decoder.DecodeText(compressedResult)
		
		// Calcular estatísticas
		originalBits, compressedBits, compressionRate := compressor.CalculateCompressionStats(text, compressedResult.CompressedBits, codes)
		fmt.Printf("     -> Compression rate: %.2f%% (%d → %d bits)\n", compressionRate, originalBits, compressedBits)
		
		// Verificar se a decodificação está correta
		if decodedText == text {
			fmt.Printf("     -> ✓ Decoding validation: SUCCESS\n")
		} else {
			fmt.Printf("     -> ⚠ Decoding validation: WARNING (texts differ)\n")
		}

		// Armazenar resultado
		results = append(results, writer.CompressedText{
			Tree:           tree,
			Codes:          codes,
			CompressedData: compressedResult.CompressedBits,
			OriginalText:   text,
			CompressedResult: compressedResult,
			DecodedText:     decodedText,
		})

		fmt.Println()
		i++
	}

	err = writer.WriteOutput(outputFile, results)
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}
	fmt.Println("| Output file successfully generated |")
}
