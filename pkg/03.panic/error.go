// defer ( 捕获panic )  ; panic ; recover ( 等待捕获panic的defer )
package main

import (
    "fmt"
    "time"    
)

func GO(){
    fmt.Println("GO ...")
}

func Say( say string){
    for i:=1; i<10; i++{
        time.Sleep(1)
        fmt.Println(say)
    }
}

func Python(){
    // defer grep panic exception ;  panic ( exit ) unless recover 
    defer func(){
        if err := recover(); err != nil{
            fmt.Println("Finall grep warning ...",err)
        }
    }() // go:18:6: expression in defer must be function call

    panic(" err ...")
    fmt.Println("Python ...")
}

func main(){
    GO()
    Python()
    // concurrent  
    go Say("thisworld")
    Say("THISWORLD")
}

// __END__
