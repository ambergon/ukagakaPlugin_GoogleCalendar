package main

import (
    //"fmt"
    "time"
    "os"
    "context"
    "regexp"

    "google.golang.org/api/calendar/v3"
    "google.golang.org/api/option"
)


func Calendar() {
    credentialFilePath  := Config.FilePath
    Gmail               := Config.GmailAddress
    TimeZone            := Config.TimeZone
    Sep                 := Config.Sep



    //認証ファイルチェック
    _, err := os.Stat( credentialFilePath )
    if err != nil {
        CalendarText = "認証用ファイルがないよ。" 
        return
    }
    ctx := context.Background()
    calendarService, err := calendar.NewService(ctx, option.WithCredentialsFile( credentialFilePath ))
    if err != nil { 
        CalendarText = "認証ファイルの読み込みに失敗したよ。" 
        return 
    }


    var from string = ""
    var util string = ""
    from = time.Date(  time.Now().Year() , time.Now().Month() , time.Now().Day() , TimeZone , 0 , 0 , 0 , time.UTC ).AddDate( 0, 0, Config.From ).Format( time.RFC3339 )
    util = time.Date(  time.Now().Year() , time.Now().Month() , time.Now().Day() , TimeZone , 0 , 0 , 0 , time.UTC ).AddDate( 0, 0, Config.Util ).Format( time.RFC3339 )

    events, err := calendarService.Events.
        List(           Gmail       ).
        TimeMin(        from         ).
        TimeMax(        util         ).
        OrderBy(        "startTime" ).
        SingleEvents(   true        ).
        Do()
    if err != nil {
        //fmt.Printf( "check email : %v" , err ) 
        return
    }


    //予定が見つからなかった場合。
    if len( events.Items ) == 0 {
        //fmt.Println( "予定なし。" )
        return
    }
    res := ""
    for _,item := range events.Items {

        title := item.Summary
        if title == "" {
            title = "No Title"
        }

        //終日
        timeText := ""
        startTime := item.Start.Date 
        if startTime != "" {
            x := regexp.MustCompile( "^........" )
            startTime = x.ReplaceAllString( startTime , "" )
            timeText = startTime + "日 : " + title

        //指定時間あり。
        } else {
            startTime = item.Start.DateTime
            x := regexp.MustCompile( "^........")
            day := x.ReplaceAllString( startTime , "" )
            x = regexp.MustCompile( "^..")
            dayx := x.FindAllStringSubmatch( day , 1 )

            x = regexp.MustCompile( "^..........." )
            hhmm := x.ReplaceAllString( startTime , "" )
            x = regexp.MustCompile( `:00\+.*?$` )
            hhmm = x.ReplaceAllString( hhmm , "" )
            //timeText = "日 " + hhmm + " : " + title
            timeText = dayx[0][0] + "日 " + hhmm + " : " + title
        }
        res = res + timeText + Sep
    }
    //fmt.Println( res )
    CalendarText = res
}





















