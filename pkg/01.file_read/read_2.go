package main 

import ( 
    "log" 
    "os" 
) 

func main() { 

    f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755) 

    if err != nil { 
        log.Fatal(err) 
    } 

    if err := f.Close() 

    err != nil { 
        log.Fatal(err) 
    } 

}


// fi, err := os.Open(path)
//     if err != nil {
//         panic(err)
//     }
//     defer fi.Close()
