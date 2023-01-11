package package1

import (
	"context"
	"errors"
	"flag"
	"log"

	luno "github.com/luno/luno-go"
)

var (
	apiKeyID     = flag.String("api_key_id", "j42v5mwx72evv", "Luno API key ID")
	apiKeySecret = flag.String("api_key_secret", "E1LlaWNjWco8HaEexBkHuTbtRwjuc6v0EQSFf4U-7lU", "Luno API key secret")
	debug        = flag.Bool("debug", false, "Enable debug mode")
)

func GetAccount(ID string) (string,error) {
	flag.Parse()

	cl := luno.NewClient()
	cl.SetDebug(*debug)
	cl.SetAuth(*apiKeyID, *apiKeySecret)

	ctx := context.Background()
	req := luno.GetBalancesRequest{}
	res, err := cl.GetBalances(ctx, &req)
	if err != nil {
		log.Println(err)
		return "",err
	}
	log.Printf("\nResult:%+v", res)
	for ind:=0;ind<len(res.Balance);ind++ {
		if res.Balance[ind].Asset == ID {
			return res.Balance[ind].AccountId,nil
		}
	}
//	log.Printf("\nMy Balance:%+v", res.Balance[1])
	if res != nil && len(res.Balance) > 0 {
		return res.Balance[1].AccountId,nil
	}
	return "",errors.New("No account ID available")

}