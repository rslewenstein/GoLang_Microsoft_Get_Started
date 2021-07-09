/*

ch := make(chan int) // declara um canal do tipo inteiro
ch <- x // envia (ou escreve) x através do canal ch
x = <-ch // x recebe (ou lê) dados enviados para o canal ch
<-ch // recebe dados, mas o resultado é descartado
close(ch) // fecha o canal

*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}
	fmt.Printf(<-ch)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

}
