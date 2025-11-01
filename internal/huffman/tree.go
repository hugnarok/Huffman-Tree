package huffman

import "container/heap"

// Node representa um nó na árvore de Huffman
type Node struct {
	Word      string // palavra (vazio para nós internos)
	Frequency int    // frequência da palavra ou soma das frequências dos filhos
	Left      *Node  // filho esquerdo (código 0)
	Right     *Node  // filho direito (código 1)
}

// IsLeaf verifica se o nó é uma folha (contém uma palavra)
func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

// BuildHuffmanTree constrói a árvore de Huffman a partir de um mapa de frequências
// Usa um min-heap para construção bottom-up eficiente
func BuildHuffmanTree(frequencies map[string]int) *Node {
	// Caso especial: se houver apenas uma palavra, criar árvore mínima
	if len(frequencies) == 1 {
		for word, freq := range frequencies {
			// Cria uma árvore com raiz e uma folha
			root := &Node{
				Word:      "",
				Frequency: freq,
				Left: &Node{
					Word:      word,
					Frequency: freq,
					Left:      nil,
					Right:     nil,
				},
				Right: nil,
			}
			return root
		}
	}

	// Cria o heap inicial com todos os nós folha
	h := NewHeap(frequencies)

	// Constrói a árvore combinando os dois nós de menor frequência
	for h.Len() > 1 {
		// Remove os dois nós com menor frequência
		left := heap.Pop(h).(*Node)
		right := heap.Pop(h).(*Node)

		// Cria um novo nó interno com a soma das frequências
		parent := &Node{
			Word:      "", // nós internos não têm palavra
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}

		// Adiciona o novo nó de volta ao heap
		heap.Push(h, parent)
	}

	// O último nó restante é a raiz da árvore
	return heap.Pop(h).(*Node)
}
