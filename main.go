package main
import(
	"bufio"
	"fmt"
	"os"
	"net/http"
	"strings"
	"sync"
)

var host string
var statusCode string
var wg sync.WaitGroup

func main(){
	userInput()
	//wg.Add(30)  // attempting to set up concurrency.
	words(host)
	//wg.Wait()
}

func userInput(){
	fmt.Println("Host (include http(s)://> ")
	fmt.Scan(&host)
	fmt.Println("Status code (i.e. 200)> ")
	fmt.Scan(&statusCode)
}

func words(url string){
	file, ferr := os.Open("dictionaries/common.txt")
	if ferr != nil{
		fmt.Println(ferr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // For debugging purposes...
		enumerate(url, scanner.Text())
	}

}

func enumerate(url, word string){
	conn, err := http.Get(url + "/" + word)
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(conn.Status, statusCode){
		fmt.Printf("%s%s%s%s",url, "/", word, " : Good Endpoint\n")
		discoveredURl := url + "/" + word
		enumerate(discoveredURl, word)
	}
	//wg.Done()
}


