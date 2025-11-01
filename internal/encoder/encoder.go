package encoder

import (
	"huffman/internal/huffman"
)

// GenerateCodes gera os códigos binários de Huffman para cada palavra
// Percorre a árvore recursivamente: 0 para esquerda, 1 para direita
func GenerateCodes(root *huffman.Node) map[string]string {
	codes := make(map[string]string)

	if root == nil {
		return codes
	}

	// Caso especial: árvore com apenas uma palavra
	if root.IsLeaf() {
		codes[root.Word] = "0"
		return codes
	}

	// Percorre a árvore gerando códigos
	generateCodesRecursive(root, "", codes)

	return codes
}

// generateCodesRecursive é uma função auxiliar recursiva que percorre a árvore
// e constrói os códigos binários
func generateCodesRecursive(node *huffman.Node, code string, codes map[string]string) {
	if node == nil {
		return
	}

	// Se é uma folha, armazena o código
	if node.IsLeaf() {
		codes[node.Word] = code
		return
	}

	// Percorre o filho esquerdo com código 0
	if node.Left != nil {
		generateCodesRecursive(node.Left, code+"0", codes)
	}

	// Percorre o filho direito com código 1
	if node.Right != nil {
		generateCodesRecursive(node.Right, code+"1", codes)
	}
}
