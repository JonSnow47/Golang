/*
 * Revision History:
 *     Initial: 2018/05/16        Chen Yanchen
 */

package news

import (
	"fmt"
	"log"

	"github.com/kaelanb/newsapi-go"
)

func main()  {
	
}

func startServer()  {
	apikey := "apikeyhere"
	client := newsapi.New(apikey)

	query1 := []string{"country=ca"}
	newsResponse, err := client.GetTopHeadlines(query1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newsResponse)
}