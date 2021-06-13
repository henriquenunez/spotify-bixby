package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"sync"
	"os"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

func searchSong(w http.ResponseWriter, r *http.Request) {

        keys, ok := r.URL.Query()["song"]

        if !ok || len(keys[0]) < 1 {
            log.Println("Url Param 'key' is missing")
            return
        }

        // Query()["key"] will return an array of items,
        // we only want the single item.
        key := keys[0]

        fmt.Println("Url Param 'key' is: " + string(key))

	results, err := client.Search(key, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	var item_uri spotify.URI
	// handle tracks results
	if results.Tracks != nil {
		item_uri = results.Tracks.Tracks[0].URI
		//for _, item := range results.Tracks.Tracks {
		//	fmt.Println("   ", item.URI)
		//}
	}
	fmt.Println("Tracks:", item_uri)

	w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintln(`{"uri": "`+item_uri+`"}`)))
}


var client spotify.Client //Global client so handler can access.

func refreshToken(m *sync.Mutex) {
	// Request token
	for {
		config := &clientcredentials.Config{
			ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			TokenURL:     spotify.TokenURL,
		}
		token, err := config.Token(context.Background())
		if err != nil {
			log.Fatalf("couldn't get token: %v", err)
		}
		fmt.Println("Refresing token")

		m.Lock()
		client = spotify.Authenticator{}.NewClient(token)
		m.Unlock()

		time.Sleep(time.Second * 3600)
	}
}

func main() {
	var m sync.Mutex
	go refreshToken(&m)
	http.HandleFunc("/search", searchSong)
	fmt.Println("Listening at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

