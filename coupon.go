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
	"github.com/aizwellenstan/go_woocommerce/lib"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
			fmt.Println("fatal error config file: default \n", err)
			os.Exit(1)
	}

	url := viper.GetString("url")
	key := viper.GetString("key")
	secret := viper.GetString("secret")

	factory := client.Factory{}
	c := factory.NewClient(options.Basic{
		URL:    url,
		Key:    key,
		Secret: secret,
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v3",
		},
	})

	if r, err := c.Get("orders", nil); err != nil {
		log.Fatal(err)
	} else if r.StatusCode != http.StatusOK {
		log.Fatal("Unexpected StatusCode:", r)
	} else {
		defer r.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
			log.Fatal(err)
		} else {
			d1 := []byte(string(bodyBytes))
			err := ioutil.WriteFile("orders.json", d1, 0644)
			check(err)

			orders := lib.LoadJson(d1)
			inputs := orders
			outputlist := map[string][]string{}
			for _, inp := range inputs {
				for _, item := range inp.LineItems {
					if len(inp.CouponLines) > 0 {
						outputlist[inp.CouponLines[0].Code] = append(outputlist[inp.CouponLines[0].Code], item.Name,)
					}
				}
			}

			var usageList []Coupon
			for key, inp := range outputlist {
				duplicate := inp
				dup_map := dup_count(duplicate)

				fmt.Println(key)
				var coupon Coupon
				coupon.Code = key
				for k, v := range dup_map {
					fmt.Printf("Item : %s , Count : %d\n", k, v)
					if k == "GaoXu Matsutake" {
						coupon.M = v
					}
					if k == "Gao Xu Antrodia Cinnamomea" {
						coupon.C = v
					}
					if k == "gosh CBD Citrus Gummies" {
						coupon.G = v
					}
					if k == "gosh CBD Tincture 7200 mg – REJUVENATE" {
						coupon.T = v
					}
				}
				coupon.Total=coupon.M+coupon.C+coupon.G+coupon.T
				
				usageList = append(usageList, coupon,)
			}

			jsonStr := lib.ToJson(outputlist)
			fmt.Println(jsonStr)
			fmt.Println(usageList)
			jsonStr = lib.ToJson(usageList)
			fmt.Println(usageList)
		}
	}
}

type Coupon struct {
	Code string
	// Orders int
	// AmountDiscounted int
	M int
	C int
	G int
	T int
	Total int
}
