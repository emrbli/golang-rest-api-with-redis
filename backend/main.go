package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-redis/redis"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome Kemal Emre Ballı!")
	// } else {
	// 	http.Error(w, http.StatusText(404), 404)
	// }
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/keys", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		switch r.Method {
		case http.MethodPut:
			{
				enableCors(&w)
				keys, ok := r.URL.Query()["key"]
				values, ok := r.URL.Query()["value"]
				if !ok || len(keys[0]) < 1 {
					log.Println("Url Param 'key' is missing")
					return
				}
				key := keys[0]
				value := values[0]
				log.Println("Url Param 'key' is: " + string(key))
				log.Println("Url Param 'value' is: " + string(value))
				SetValue(string(key), string(value))
				fmt.Fprint(w, string(key), " - ", string(value), " added to Redis.")
			}
		case http.MethodGet:
			enableCors(&w)
			getAllValue(w, r)
		case http.MethodDelete:
			enableCors(&w)
			DeleteAll(w, r)
		case http.MethodOptions:
			enableCors(&w)
		default:
			enableCors(&w)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/keys/", func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			enableCors(&rw)
			name := strings.Replace(req.URL.Path, "/keys/", "", 1)
			if len(name) > 0 {
				if GetValue(name) != redis.Nil.Error() {
					GetValue(name)
					fmt.Fprint(rw, GetValue(name))
					rw.WriteHeader(http.StatusOK)
				} else {
					rw.WriteHeader(http.StatusNotFound)
				}
			} else {
				fmt.Fprint(rw, "Please enter a parameter.")
			}

		case http.MethodDelete:
			enableCors(&rw)
			rw.WriteHeader(http.StatusOK)

			name := strings.Replace(req.URL.Path, "/keys/", "", 1)
			if len(name) > 0 {
				fmt.Fprint(rw, name, " was deleted from Redis.")
				DeleteValue(name)
			} else {
				fmt.Fprint(rw, "Please enter a parameter.")
			}

		case http.MethodHead:
			enableCors(&rw)
			name := strings.Replace(req.URL.Path, "/keys/", "", 1)
			if CheckValue(name) == false {
				rw.WriteHeader(http.StatusNotFound)
			} else {
				rw.WriteHeader(http.StatusOK)
			}
		case http.MethodOptions:
			enableCors(&rw)
		default:
			enableCors(&rw)
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/", homePage)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8081"
	}

	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
func rClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func main() {
	handleRequests()

	client := rClient()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func GetValue(key string) string {
	client := rClient()
	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	} else if val != "redis: nil" {
		return val
	} else {
		return "Key Not Found."
	}
}

func SetValue(key string, value string) error {

	client := rClient()
	err := client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteValue(key string) {
	client := rClient()
	val, err := client.Del(key).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) bool {
	client := rClient()
	client.FlushAll().Result()
	fmt.Fprintln(w, "All keys deleted from Redis.")
	return true
}

func CheckValue(key string) bool {
	client := rClient()
	val2, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist.")
		return false
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(key + " => " + val2 + " is true.")
		return true
	}
	return true
}
func getAllValue(w http.ResponseWriter, r *http.Request) {
	client := rClient()

	val := client.Keys("*").Val()

	if len(val) != 0 {
		for i := 0; i < len(val); i++ {
			fmt.Println("'" + val[i] + "'") // for all keys
			fmt.Fprintln(w, (val[i]))
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Redis does not contain a key.")
	}

}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	(*w).Header().Set("Content-Type", "application/x-www-form-urlencoded")
}
