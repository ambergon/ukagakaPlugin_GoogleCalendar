package main

import (
    "fmt"
    "os"
    "encoding/json"
)


type CalendarConfig struct {
    FilePath        string
    GmailAddress    string
    TimeZone        int
    StartMin        int
    From            int
    Util            int
    Sep             string

}
var Config CalendarConfig


func LoadJson(){
	JsonChatGPTConfig, err := os.Open( Directory + "/Config.json")
	if err != nil {
        fmt.Println( err )
	}
	defer JsonChatGPTConfig.Close()
    decoder := json.NewDecoder( JsonChatGPTConfig )
    err     = decoder.Decode( &Config )
	if err != nil {
        fmt.Println(  err  )
    }
}






