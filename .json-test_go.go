package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "os"
)

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func main() {
        args := os.Args[1:]
        configPath := args[0]
        configData, err := ioutil.ReadFile(configPath)
        check(err)
        fmt.Println("loaded " + configPath + "\n")

        var labs []interface{}
        err = json.Unmarshal(configData, &labs)
        check(err)

        for l := range labs {
                lab := labs[l].(map[string]interface{})
                fmt.Println("---- " + lab["title"].(string) + " ----")
                start := int(lab["start"].(float64))
                end := int(lab["end"].(float64))

                for i := start; i <= end; i++ {
                        fmt.Printf("%s-%02d.***REMOVED***\n", lab["prefix"], i)
                }
        }
}
