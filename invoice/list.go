package invoice

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

var path = ".exportedlist"

func loadList() []string {
    var _, err = os.Stat(path)
    if err != nil {
        fmt.Println(err)
    }

	// create file if not exists
	if os.IsNotExist(err) {
		var file, fileErr = os.Create(path)
		if fileErr != nil {
            fmt.Println(err)
            fmt.Println(fileErr)
        }
		defer file.Close()
	}

    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    return lines
}

func writeList(list []string) {
    file, err := os.Create(path)
    if err != nil {
        fmt.Println(err)
    }

    file.WriteString(strings.Join(list, "\n"))
}
