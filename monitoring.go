package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 10
const delay = 5

func main() {

	introduction()

	for {

		showMenu()
		comando := getComand()

		// Utilizando IF

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 3 {
		// 	fmt.Println("Saindo...")
		// } else {
		// 	fmt.Println("Não conheço este comando")
		// }

		// Utilizando SWITCH

		switch comando {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)

		}
	}

}

func introduction() {

	nome := "Jonathan"
	versao := 1.21

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão:", versao)

}

func getComand() int {

	var comando int

	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)
	fmt.Println("")
	// fmt.Println("O endereço da variavel comando é: ", &comando)

	return comando
}

func showMenu() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func startMonitoring() {

	fmt.Println("Monitorando...")

	// sites := []string{"https://www.alura.com.br", "https://www.google.com.br", "https://instagram.com/"}

	sites := readSitesArchive()

	// Usando range (foreach, key & value)

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {

			fmt.Println("Testando site:", site)

			testSite(site)
		}

		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}

	fmt.Println("")

	// Tester for me

	// for i := 0; i < len(sites); i++ {

	// 	resp, _ := http.Get(sites[i])

	// 	if resp.StatusCode == 200 {
	// 		fmt.Println("Site:", sites[i], "foi carregado com sucesso!")
	// 	} else {
	// 		fmt.Println("Site:", sites[i], "está com problemas.", resp.StatusCode)
	// 	}

	// }

}

func testSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao testar site:", site, "Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas.", resp.StatusCode)
		registerLog(site, false)
	}

}

func readSitesArchive() []string {

	var sites []string

	arquivo, err := os.Open("archives/sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler arquivo:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registerLog(site string, status bool) {

	arquivo, err := os.OpenFile("archives/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func printLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao imprimir logs:", err)
	}

	fmt.Println(string(arquivo))
}
