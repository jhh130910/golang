// Channels are the pipes that connect concurrent goroutines
// send values into channels from one goroutine and receive those values into another goroutine

package main

import "fmt"

func main(){
    // create channel
    messages := make( chan string )

    go func(){
        // send value to channel ( one goroutine )
        // channels <- 
        messages <- "xx ..."
    }()

    // <- channels ( receive a value from channels ) 
    msg := <- messages
    fmt.Println(msg)

}
