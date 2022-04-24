package function

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"net/http"
	"os"
)

var redisHost = os.Getenv("REDIS_HOST") // This should include the port which is most of the time 6379
var redisPassword = os.Getenv("REDIS_PASSWORD")

type GameSession struct{
	SessionId string
	Player string
}

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})

	nickname, _ := req.URL.Query()["nickname"]
	var g GameSession
	g.Player= nickname[0]
	g.SessionId = "game-" + uuid.NewString()

	gameSessionJson, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(g.SessionId, string(gameSessionJson), 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(res, string(gameSessionJson))


}
