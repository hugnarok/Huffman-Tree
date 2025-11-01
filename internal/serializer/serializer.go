package serializer

import (
	"fmt"
	"huffman/internal/huffman"
	"strings"
)

// SerializeTree serializa a árvore de Huffman em formato textual
// Usa travessia pré-ordem: raiz, esquerda, direita
// Formato: (frequencia:palavra) para folhas, (frequencia:*) para nós internos
func SerializeTree(root *huffman.Node) string {
	if root == nil {
		return ""
	}

	var builder strings.Builder
	serializeRecursive(root, &builder, 0)
	return builder.String()
}

// serializeRecursive serializa a árvore recursivamente com indentação
func serializeRecursive(node *huffman.Node, builder *strings.Builder, depth int) {
	if node == nil {
		return
	}

	// Adiciona indentação para visualização hierárquica
	indent := strings.Repeat("  ", depth)

	if node.IsLeaf() {
		// Nó folha: (frequencia:palavra)
		builder.WriteString(fmt.Sprintf("%s(%d:%s)\n", indent, node.Frequency, node.Word))
	} else {
		// Nó interno: (frequencia:*)
		builder.WriteString(fmt.Sprintf("%s(%d:*)\n", indent, node.Frequency))

		// Serializa filhos
		if node.Left != nil {
			builder.WriteString(fmt.Sprintf("%s  L:\n", indent))
			serializeRecursive(node.Left, builder, depth+2)
		}
		if node.Right != nil {
			builder.WriteString(fmt.Sprintf("%s  R:\n", indent))
			serializeRecursive(node.Right, builder, depth+2)
		}
	}
}

// SerializeTreeCompact cria uma representação compacta da árvore em uma linha
// Útil para armazenamento mais eficiente
func SerializeTreeCompact(root *huffman.Node) string {
	if root == nil {
		return ""
	}

	var builder strings.Builder
	serializeCompactRecursive(root, &builder)
	return builder.String()
}

// serializeCompactRecursive serializa a árvore em formato compacto
// Formato: N(freq,word) para folha, I(freq) para interno, seguido de L(...) R(...)
func serializeCompactRecursive(node *huffman.Node, builder *strings.Builder) {
	if node == nil {
		builder.WriteString("X")
		return
	}

	if node.IsLeaf() {
		builder.WriteString(fmt.Sprintf("N(%d,%s)", node.Frequency, node.Word))
	} else {
		builder.WriteString(fmt.Sprintf("I(%d,", node.Frequency))
		serializeCompactRecursive(node.Left, builder)
		builder.WriteString(",")
		serializeCompactRecursive(node.Right, builder)
		builder.WriteString(")")
	}
}

// VisualizeTree cria uma representação visual da árvore com os caminhos
func VisualizeTree(root *huffman.Node) string {
	if root == nil {
		return "Árvore vazia"
	}

	var builder strings.Builder
	builder.WriteString("Estrutura da Árvore de Huffman:\n")
	builder.WriteString(strings.Repeat("=", 50) + "\n")
	visualizeRecursive(root, &builder, "", true)
	return builder.String()
}

// visualizeRecursive cria visualização ASCII da árvore
func visualizeRecursive(node *huffman.Node, builder *strings.Builder, prefix string, isTail bool) {
	if node == nil {
		return
	}

	// Desenha o nó atual
	connector := "└── "
	if !isTail {
		connector = "├── "
	}

	if node.IsLeaf() {
		builder.WriteString(fmt.Sprintf("%s%s[%s] freq=%d\n", prefix, connector, node.Word, node.Frequency))
	} else {
		builder.WriteString(fmt.Sprintf("%s%s(*) freq=%d\n", prefix, connector, node.Frequency))

		// Prepara o prefixo para os filhos
		extension := "    "
		if !isTail {
			extension = "│   "
		}
		newPrefix := prefix + extension

		// Desenha os filhos
		if node.Right != nil {
			visualizeRecursive(node.Right, builder, newPrefix, false)
		}
		if node.Left != nil {
			visualizeRecursive(node.Left, builder, newPrefix, true)
		}
	}
}
