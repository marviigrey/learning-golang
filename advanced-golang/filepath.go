package main
import (
    "fmt"
    "path/filepath"
    "strings"
    )
    func main() {
        fmt.Println(filepath.Dir("/test/page/mind.js"))
        fmt.Println(filepath.Base("/table/scripts/run.sh"))
        s := []string {"apple", "pear", "mango"}
        fmt.Println(s)
        fmt.Println(strings.Join(s, ", "))
        fmt.Println(filepath.Ext("/test/page/mind.js"))
    }