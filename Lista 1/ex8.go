package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Question string `json:"question,omitempty"`
	Yes      *Node  `json:"yes,omitempty"`
	No       *Node  `json:"no,omitempty"`
	Action   string `json:"action,omitempty"`
}

func (n *Node) IsLeaf() bool {
	return n != nil && n.Yes == nil && n.No == nil && n.Action != ""
}

func Evaluate(root *Node, answers []bool) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for _, ans := range answers {
		if cur.IsLeaf() {
			return cur
		}
		if ans {
			cur = cur.Yes
		} else {
			cur = cur.No
		}
		if cur == nil {
			return nil
		}
	}
	return cur
}

func RunCLI(root *Node) {
	if root == nil {
		fmt.Println("Árvore vazia.")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	cur := root
	for {
		if cur.IsLeaf() {
			fmt.Printf("\nResultado encontrado: %s\n", cur.Action)
			return
		}
		fmt.Printf("%s (y/n): ", cur.Question)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro de leitura:", err)
			return
		}
		text = strings.TrimSpace(strings.ToLower(text))
		if text == "y" || text == "s" {
			if cur.Yes == nil {
				fmt.Println("Caminho 'sim' não definido. Árvore incompleta.")
				return
			}
			cur = cur.Yes
			continue
		}
		if text == "n" {
			if cur.No == nil {
				fmt.Println("Caminho 'não' não definido. Árvore incompleta.")
				return
			}
			cur = cur.No
			continue
		}
		fmt.Println("Resposta inválida. Use 'y' (ou 's') para sim e 'n' para não.")
	}
}

func BuildExampleTree() *Node {
	// Folhas
	dog := &Node{Action: "É um cachorro"}
	cat := &Node{Action: "É um gato"}
	bird := &Node{Action: "É um pássaro"}
	fish := &Node{Action: "É um peixe"}

	// Nós intermediários
	flies := &Node{
		Question: "O animal voa?",
		Yes:      bird,
		No:       nil,
	}

	livesInWater := &Node{
		Question: "Vive na água?",
		Yes:      fish,
		No:       nil,
	}

	hasFur := &Node{
		Question: "Tem pelos?",
		Yes:      cat,
		No:       dog,
	}

	// Ajuste dos nós para ter estrutura binária completa onde necessário
	// Pergunta principal
	root := &Node{
		Question: "O animal tem quatro patas?",
		Yes:      hasFur,
		No: &Node{
			Question: "Tem penas?",
			Yes:      bird,
			No:       livesInWater,
		},
	}

	// Torna "voa" parte do ramo das penas (apenas para demonstrar)
	root.No.Yes = bird
	root.No.No = livesInWater

	// Se você quiser que "voa" seja um nível acima:
	_ = flies // mantém a variável caso extenda a árvore

	return root
}

// Example of serializing and deserializing the tree to/from JSON.
func exampleJSON(root *Node) {
	b, err := json.Marsh
