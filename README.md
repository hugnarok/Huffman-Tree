# Compressor de Huffman - Trabalho de AED2

Implementação do algoritmo de Huffman para compressão de texto baseada em frequência de **palavras** (não caracteres).

## 📋 Descrição

Este projeto implementa um compressor de texto utilizando o código de Huffman como método de codificação estatística. O programa:

- Lê textos de um arquivo de entrada (`data/input.dat`)
- Calcula a frequência de cada palavra
- Constrói a árvore binária de Huffman
- Gera códigos binários otimizados (0=esquerda, 1=direita)
- Comprime os textos substituindo palavras por códigos
- Gera um arquivo de saída (`data/output.dat`) com toda a documentação da compressão

## 📁 Estrutura do Projeto

```
Trabalho Huffman/
├── main.go                     # Programa principal
├── go.mod                      # Módulo Go
├── README.md                   # Este arquivo
├── .gitignore                  # Arquivos ignorados pelo git
├── data/
│   ├── input.dat              # Arquivo de entrada (textos a comprimir)
│   └── output.dat             # Arquivo de saída (resultado da compressão)
├── internal/
│   ├── reader/
│   │   └── reader.go          # Leitor de arquivos
│   ├── frequency/
│   │   └── frequency.go       # Analisador de frequências
│   ├── huffman/
│   │   ├── tree.go            # Estrutura e construção da árvore
│   │   └── heap.go            # Min-heap para construção eficiente
│   ├── encoder/
│   │   └── encoder.go         # Gerador de códigos binários
│   ├── compressor/
│   │   └── compressor.go      # Compressor de texto
│   ├── serializer/
│   │   └── serializer.go      # Serializador da árvore
│   └── writer/
│       └── writer.go          # Escritor do arquivo de saída
└── examples/
    └── sample_input.dat       # Exemplo de arquivo de entrada
```

## 🚀 Como Usar

### Pré-requisitos

- Go 1.20 ou superior instalado
- Sistema Linux (Ubuntu 24.04) ou compatível

### Compilação

```bash
git clone <URL>
cd "Trabalho Huffman"
go build -o huffman main.go
```

### Execução

```bash
./huffman
```

O programa irá:
1. Ler os textos do arquivo `data/input.dat`
2. Processar cada texto individualmente
3. Gerar o arquivo `data/output.dat` com os resultados

### Formato do Arquivo de Entrada

O arquivo `data/input.dat` deve conter textos separados por **linhas em branco**:

```
Texto 1: primeira frase do primeiro texto.

Texto 2: primeira frase do segundo texto.

Texto 3: primeira frase do terceiro texto.
```

## 📊 Exemplo de Saída

Para cada texto processado, o `output.dat` contém:

### 1. Árvore de Huffman (formato hierárquico)
```
(12:*)
  L:
    (4:*)
      L:
        (2:*)
          L:
            (1:alta)
          R:
            (1:velocidade)
```

### 2. Visualização da Árvore
```
└── (*) freq=12
    ├── (*) freq=8
    │   ├── (*) freq=4
    │   │   ├── [em] freq=1
    │   │   └── [com] freq=1
```

### 3. Representação Compacta (para decodificação)
```
I(12,I(4,I(2,N(1,alta),N(1,velocidade)),...))
```

### 4. Códigos Gerados
```
alta                : 000
com                 : 1110
computador          : 011
dados               : 1001
```

### 5. Texto Original
```
O computador executa instruções em alta velocidade...
```

### 6. Texto Comprimido (binário)
```
10100111101111100000110001100100111101010
```

### 7. Estatísticas de Compressão
```
Tamanho original:      672 bits (84 bytes)
Tamanho comprimido:    41 bits (5 bytes + overhead)
Taxa de compressão:    93.90%
Número de palavras únicas: 12
Comprimento médio dos códigos: 3.67 bits
```

## 🔍 Como Funciona o Algoritmo de Huffman

1. **Análise de Frequências**: Conta quantas vezes cada palavra aparece no texto
2. **Construção da Árvore**: 
   - Cria nós folha para cada palavra com sua frequência
   - Usa um min-heap para combinar os dois nós de menor frequência
   - Repete até restar apenas um nó (raiz)
3. **Geração de Códigos**:
   - Percorre a árvore da raiz até cada folha
   - Filho esquerdo = 0, filho direito = 1
   - O caminho forma o código da palavra
4. **Compressão**: Substitui cada palavra do texto pelo seu código binário

### Exemplo Prático

Texto: "a casa a casa"
- Frequências: {a: 2, casa: 2}
- Árvore: `a` e `casa` têm mesma frequência, recebem códigos de 1 bit
- Códigos: a=0, casa=1
- Comprimido: 0 1 0 1 = "0101"
- Original: 48 bits (6 bytes) → Comprimido: 4 bits!

## 📝 Detalhes de Implementação

### Estruturas Principais

```go
type Node struct {
    Word      string    // palavra (vazio para nós internos)
    Frequency int       // frequência
    Left      *Node     // filho esquerdo (código 0)
    Right     *Node     // filho direito (código 1)
}
```

### Tokenização

- Converte texto para lowercase
- Remove pontuação
- Separa por espaços e caracteres especiais
- Mantém apenas letras e números

### Min-Heap

Implementa a interface `heap.Interface` do Go para construção eficiente da árvore com complexidade O(n log n).

