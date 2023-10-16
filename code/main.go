package main

/*
   #include <windows.h>
   #include <stdlib.h>
   #include <string.h>
*/
import "C"

import (
    "fmt"
    "unsafe"
    "strings"
    "regexp"
)

func main() {
    fmt.Println( "test" )
}


var Directory string
var References []string 
var CheckID         = regexp.MustCompile("^ID: ")
var CheckReference  = regexp.MustCompile("^Reference.+?: ")


type ResponseStruct struct {
    Shiori  string
    Sender  string
    Charset string
    Marker  string
    Value   string
}
func GetResponse( r *ResponseStruct ) string {
    V := ""
    if r.Value  != "" { V = "Value: "  + r.Value     + "\r\n" }
    res :=  r.Shiori    + "\r\n" + 
            r.Sender    + "\r\n" + 
            r.Charset   + "\r\n" + 
            V + "\r\n\r\n"
    return res
}

var CalendarText  string = ""
var weekText string = ""
var PastMin     int = 0



//export load
func load(h C.HGLOBAL, length C.long ) C.BOOL {
    fmt.Println( "load GoogleCalendar" )
    Directory = C.GoStringN(( *C.char )( unsafe.Pointer( h )), ( C.int )( length ))
    fmt.Println( Directory  )

    //設定読み込み。
    LoadJson()

    //カレンダー機能の実行 day & week
    if Config.StartMin > 0 {
        go Calendar()
    }


	C.GlobalFree( h )
	return C.TRUE
}


//export unload
func unload() bool {
    fmt.Println( "unload GoogleCalendar" )
	return true
}


//export request
func request( h C.HGLOBAL, length *C.long ) C.HGLOBAL {
	RequestText := C.GoStringN(( *C.char )( unsafe.Pointer( h )), ( C.int )( *length ))
	C.GlobalFree( h )


    Value           := ""
    Marker          := ""
    ID              := ""
    References      = []string{}
    //var NOTIFY bool = false

    Response := new( ResponseStruct )
    Response.Sender  = "Sender: GolangGoogleCalendar"
    Response.Charset = "Charset: UTF-8"

    //IDとReference
    //必要な情報を分解する。
    RequestLines := strings.Split( RequestText , "\r\n" )
    for _ , line := range RequestLines {
        if( line == "NOTIFY PLUGIN/2.0" ){
            //"GET PLUGIN/2.0";
            //NOTIFY = true

        } else if CheckID.MatchString( line )  {
            //fmt.Println( line )
            ID = CheckID.ReplaceAllString( line , "" )

        } else if CheckReference.MatchString( line )  {
            //fmt.Println( line )
            ref := CheckReference.ReplaceAllString( line , "" )
            References = append( References , ref )

        } else {
            //fmt.Println( line )
        }
    }





    //そもそも取得しないようにしておくか。
    //Config.StartMin != 0 

    //起動の時点で実行。
    //別スレッド





    //これの実行タイミングを60にしておくことで、一分間隔に変更
    //最終チェックは起動後三時間まで。
    if ID == "OnSecondChange" && 180 >= PastMin {
        PastMin++
        //0なら無効
        if CalendarText !=  "" && PastMin >= Config.StartMin {
            Value = CalendarText
            CalendarText = ""

        }


    //} else if ID == ""  {
    } else {
        //fmt.Println( "no touch :" + ID )
        //fmt.Print( "NOTIFY : " )
        //fmt.Println( NOTIFY )
        //fmt.Print( "References : " )
        //fmt.Println( References )
        //fmt.Println( "" )
    }


    if Value == "" {
        Response.Shiori  = "PLUGIN/2.0 204 No Content"
    } else {
        Response.Shiori = "PLUGIN/2.0 200 OK"
        Response.Value  = Value
    }
    if Marker != "" {
        Response.Marker  = Marker
    }

    res_buf := C.CString( GetResponse( Response ))
    defer C.free( unsafe.Pointer( res_buf ))

	res_size := C.strlen( res_buf )
	ret      := C.GlobalAlloc( C.GPTR , ( C.SIZE_T )( res_size ))
	C.memcpy(( unsafe.Pointer )( ret ) , ( unsafe.Pointer )( res_buf ) , res_size )
	*length = ( C.long )( res_size )
	return ret
}






















