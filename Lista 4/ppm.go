package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Nome dos arquivos de entrada e saída
	inputFile := "imagem.txt"
	outputFile := "imagem.ppm"

	// Lê o conteúdo do arquivo de texto
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de entrada: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	// Processa a primeira linha para obter a largura e altura
	dims := strings.Fields(lines[1])
	width, err := strconv.Atoi(dims[0])
	if err != nil {
		log.Fatalf("Erro ao converter largura: %v", err)
	}
	height, err := strconv.Atoi(dims[1])
	if err != nil {
		log.Fatalf("Erro ao converter altura: %v", err)
	}

	// Cria o arquivo de saída PPM
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo de saída: %v", err)
	}
	defer file.Close()

	// Escreve o cabeçalho do arquivo PPM
	fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)

	// Itera sobre as linhas de dados de pixel
	for i := 2; i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		if len(fields) == 3 {
			r := fields[0]
			g := fields[1]
			b := fields[2]
			fmt.Fprintf(file, "%s %s %s\n", r, g, b)
		}
	}

	fmt.Printf("Arquivo %s gerado com sucesso para %s\n", outputFile, inputFile)
}
