package main
import(
	"bufio"
	"fmt"
	"os"
	"net/http"
	"strings"
)

var host string
var statusCode string

func main(){
	userInput()
	words(host)
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
		enumerate(url, scanner.Text())
	}
	fmt.Println("Done...")
}

func enumerate(url, word string){
	conn, err := http.Get(url + "/" + word)
	if err != nil {
		fmt.Println(err)
	}
	if strings.Contains(conn.Status, statusCode){
		fmt.Printf("%s%s%s%s%s%s",url, "/", word, " : ", statusCode,"\n")
	}
}


