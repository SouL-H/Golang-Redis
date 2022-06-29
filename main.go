package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"creator"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkErr(err)
	defer conn.Close()
	_, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	checkErr(err)
	title, err := redis.String(conn.Do("HGET", "album:2", "title"))
	checkErr(err)
	fmt.Println(title)
	fee, err := redis.Float64(conn.Do("HGET", "album:2", "price"))
	checkErr(err)
	fmt.Println(fee)
	values, err := redis.StringMap(conn.Do("HGETALL", "album:2"))
	checkErr(err)
	for k, v := range values {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", v)
	}
	reply, err := redis.Values(conn.Do("HGETALL", "album:2"))
	checkErr(err)
	var album Album
	err = redis.ScanStruct(reply, &album)
	checkErr(err)
	fmt.Println(album)

}
