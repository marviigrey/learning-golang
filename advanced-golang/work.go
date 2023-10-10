package main
import (
    "path/filepath"
    "fmt"
    "os"
    )
    func main() {
        path := filepath.Join("/dir1/", "/dir2/", "text.txt")
        fmt.Println(path)
        fmt.Println(filepath.IsAbs("dir1"))
        fmt.Println(filepath.IsAbs(path))
        fileInfo, err := os.Stat("/home/ec2-user/environment/temp.txt")
        if err != nil {
            fmt.Println(fileInfo.Name())
            fmt.Println(fileInfo.Size())
            fmt.Println(fileInfo.Mode())
            fmt.Println(fileInfo.IsDir())
            
        }
    }