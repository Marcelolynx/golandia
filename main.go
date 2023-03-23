package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo criado com sucesso")

	time.Sleep(8 * time.Second)

	tamanho, err := f.WriteString("Hello World!")
	if err != nil {
		panic(err)
	}

	println(tamanho, "bytes written successfully")
	f.Close()

	// Lendo o arquivo
	arquivo, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	time.Sleep(8 * time.Second)

	// Lendo o arquivo linha a linha
	arquivo2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 4)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(string(buffer[:n]))
	}

	time.Sleep(8 * time.Second)

	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Arquivo removido com sucesso")
}
