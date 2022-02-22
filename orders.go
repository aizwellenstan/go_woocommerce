package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/tgglv/wc-api-go/client"
	"github.com/tgglv/wc-api-go/options"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"github.com/aizwellenstan/go_woocommerce/lib"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
			fmt.Println("fatal error config file: default \n", err)
			os.Exit(1)
	}

	urlString := viper.GetString("url")
	key := viper.GetString("key")
	secret := viper.GetString("secret")

	factory := client.Factory{}
	c := factory.NewClient(options.Basic{
		URL:    urlString,
		Key:    key,
		Secret: secret,
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v3",
		},
	})

	v := url.Values{}
	v.Set("orderby", "popularity")
	v.Add("order", "desc")

	if r, err := c.Get("products", v); err != nil {
		log.Fatal(err)
	} else if r.StatusCode != http.StatusOK {
		log.Fatal("Unexpected StatusCode:", r)
	} else {
		defer r.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
			log.Fatal(err)
		} else {
			d1 := []byte(string(bodyBytes))
			inputs := lib.LoadProductsJson(d1)
			popularList := make([]int, 0)
			for _, inp := range inputs {
				popularList = append(popularList, inp.ID)
				// }
			}
			fmt.Println(popularList)
		}
	}
}
