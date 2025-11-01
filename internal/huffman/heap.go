package huffman

import "container/heap"

// NodeHeap implementa heap.Interface para criar um min-heap de nós de Huffman
// Nós com menor frequência têm prioridade
type NodeHeap []*Node

// Len retorna o tamanho do heap
func (h NodeHeap) Len() int {
	return len(h)
}

// Less compara dois nós pela frequência (menor frequência tem prioridade)
func (h NodeHeap) Less(i, j int) bool {
	return h[i].Frequency < h[j].Frequency
}

// Swap troca dois elementos no heap
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adiciona um elemento ao heap
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

// Pop remove e retorna o elemento de menor prioridade (menor frequência)
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NewHeap cria e inicializa um novo heap a partir de um mapa de frequências
func NewHeap(frequencies map[string]int) *NodeHeap {
	h := &NodeHeap{}
	heap.Init(h)

	// Cria um nó folha para cada palavra e adiciona ao heap
	for word, freq := range frequencies {
		node := &Node{
			Word:      word,
			Frequency: freq,
			Left:      nil,
			Right:     nil,
		}
		heap.Push(h, node)
	}

	return h
}
