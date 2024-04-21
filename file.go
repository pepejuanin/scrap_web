package main

 

import (

            "fmt"

            "github.com/gocolly/colly"

)

 

type Fact struct {

            ID          int    `json:"id"`

            Description string `json:"description"`

}

 

func main() {

            fmt.Println("Hello, 世界")

            allFacts := make([]Fact, 0)

            collector := colly.NewCollector(

                        colly.AllowedDomains("factretriever.com", www.factretriever.com),

            )

            fmt.Println("Hello, 世界")

}
