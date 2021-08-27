package main

import "fmt"

// Inheritance is not possible to use in the golang as we usually used in someother language
// but it is possible to implement composition in golang
// in golang we just have to add one struct in other struct.

// Lets first start with author

type author struct {
	firstname string
	lastname  string
	bio       string
}
type blogpost struct {
	author
	title   string
	content string
}
type website struct{
    blogpost []blogpost
}
func(a author) fullname() string {
    return fmt.Sprintf("%s %s", a.firstname , a.lastname)
}
func (b blogpost)details() {
    fmt.Println("title:",b.title)
    fmt.Println("content:",b.content)
    fmt.Println("author:",b.fullname())
    fmt.Println("bio:",b.bio)
}
func (w website)contents(){
    fmt.Println("contents of the websites\n")
    for _ , v  := range w.blogpost{
        v.details()
        fmt.Println()
    }

}
func main(){
    a := author{
        "abdul",
        "aleem",
        "golang",
    }
    b1 := blogpost{
        a,
        "Inheritance in Go",
        "Go composition using instead of inheritance",
    }
    b2 := blogpost{
        a,
        "Inheritance in Go",
        "Go composition using instead of inheritance",
    }
    b3 := blogpost{
        a,
        "Inheritance in Go",
        "Go composition using instead of inheritance",
    }
    w := website{
        blogpost: []blogpost{b1,b2,b3},
    }
    w.contents()
}