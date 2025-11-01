# Huffman Compressor - AED2 Assignment

Implementation of Huffman's algorithm for text compression based on **word** frequency (not characters).

## 📋 Description

This project implements a text compressor using Huffman coding as a statistical encoding method. The program:

- Reads texts from an input file (`data/input.dat`)
- Calculates the frequency of each word
- Builds the Huffman binary tree
- Generates optimized binary codes (0=left, 1=right)
- Compresses texts by replacing words with codes
- Generates an output file (`data/output.dat`) with complete compression documentation

## 📁 Project Structure

```
Trabalho Huffman/
├── main.go                     # Main program
├── go.mod                      # Go module
├── README.md                   # This file
├── .gitignore                  # Files ignored by git
├── data/
│   ├── input.dat              # Input file (texts to compress)
│   └── output.dat             # Output file (compression result)
├── internal/
│   ├── reader/
│   │   └── reader.go          # File reader
│   ├── frequency/
│   │   └── frequency.go       # Frequency analyzer
│   ├── huffman/
│   │   ├── tree.go            # Tree structure and construction
│   │   └── heap.go            # Min-heap for efficient construction
│   ├── encoder/
│   │   └── encoder.go         # Binary code generator
│   ├── compressor/
│   │   └── compressor.go      # Text compressor
│   ├── serializer/
│   │   └── serializer.go      # Tree serializer
│   └── writer/
│       └── writer.go          # Output file writer
└── examples/
    └── sample_input.dat       # Input file example
```

## 🚀 How to Use

### Prerequisites

- Go 1.20 or higher installed
- Linux system (Ubuntu 24.04) or compatible

### Compilation

```bash
git clone <URL>
cd "Trabalho Huffman"
go build -o huffman main.go
```

### Execution

```bash
./huffman
```

The program will:
1. Read texts from `data/input.dat` file
2. Process each text individually
3. Generate `data/output.dat` file with results

### Input File Format

The `data/input.dat` file must contain texts separated by **blank lines**:

```
Text 1: first sentence of the first text.

Text 2: first sentence of the second text.

Text 3: first sentence of the third text.
```

## 📊 Output Example

For each processed text, `output.dat` contains:

### 1. Huffman Tree (hierarchical format)
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

### 2. Tree Visualization
```
└── (*) freq=12
    ├── (*) freq=8
    │   ├── (*) freq=4
    │   │   ├── [em] freq=1
    │   │   └── [com] freq=1
```

### 3. Compact Representation (for decoding)
```
I(12,I(4,I(2,N(1,alta),N(1,velocidade)),...))
```

### 4. Generated Codes
```
alta                : 000
com                 : 1110
computador          : 011
dados               : 1001
```

### 5. Original Text
```
O computador executa instruções em alta velocidade...
```

### 6. Compressed Text (binary)
```
10100111101111100000110001100100111101010
```

### 7. Compression Statistics
```
Original size:         672 bits (84 bytes)
Compressed size:       41 bits (5 bytes + overhead)
Compression rate:      93.90%
Unique words:          12
Average code length:   3.67 bits
```

## 🔍 How Huffman's Algorithm Works

1. **Frequency Analysis**: Counts how many times each word appears in the text
2. **Tree Construction**: 
   - Creates leaf nodes for each word with its frequency
   - Uses a min-heap to combine the two nodes with lowest frequency
   - Repeats until only one node remains (root)
3. **Code Generation**:
   - Traverses the tree from root to each leaf
   - Left child = 0, right child = 1
   - The path forms the word's code
4. **Compression**: Replaces each word in the text with its binary code

### Practical Example

Text: "a casa a casa"
- Frequencies: {a: 2, casa: 2}
- Tree: `a` and `casa` have same frequency, receive 1-bit codes
- Codes: a=0, casa=1
- Compressed: 0 1 0 1 = "0101"
- Original: 48 bits (6 bytes) → Compressed: 4 bits!

## 📝 Implementation Details

### Main Structures

```go
type Node struct {
    Word      string    // word (empty for internal nodes)
    Frequency int       // frequency
    Left      *Node     // left child (code 0)
    Right     *Node     // right child (code 1)
}
```

### Tokenization

- Converts text to lowercase
- Removes punctuation
- Splits by spaces and special characters
- Keeps only letters and numbers

### Min-Heap

Implements Go's `heap.Interface` for efficient tree construction with O(n log n) complexity.
