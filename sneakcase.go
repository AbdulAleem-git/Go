package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main(){
        var n int
        var str string
        fmt.Scanf("%d",&n)
        for i := 0 ; i < n ; i++{
            fmt.Scanf("%s",&str)
            str := getsnakeCase(str)
            fmt.Println(str)
        }
}
func getsnakeCase(str string) string{
    for i , v := range str {
        if i == 0 && unicode.IsUpper(v){
            temp := strings.ToLower(string(v))
            str = temp+str[1:]
            continue
        }
        if unicode.IsUpper(v){
            tempstr := ""
            if i+1 < len(str){
                 tempstr = "_"+strings.ToLower(string(v)) + str[i+1:]
            }
            str  = str[:i] + tempstr
        }
    }
    return str
}