package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"miras210/user"
	"os"
	"strings"
)

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seenBrowsers := make(map[string]bool, 100)
	var isAndroid bool
	var isMSIE bool
	var email string

	sc := bufio.NewScanner(file)

	newUser := &user.User{}
	fmt.Fprintln(out, "found users:")
	for i := 0; sc.Scan(); i++ {
		line := sc.Bytes()

		if !(bytes.Contains(line, []byte("Android")) || bytes.Contains(line, []byte("MSIE"))) {
			continue
		}

		err = newUser.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%v %v\n", err, line)

		isAndroid = false
		isMSIE = false

		for _, browser := range newUser.Browsers {
			switch {
			case strings.Contains(browser, "Android"):
				isAndroid = true
			case strings.Contains(browser, "MSIE"):
				isMSIE = true
			default:
				continue
			}
			seenBrowsers[browser] = true
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email = strings.Replace(newUser.Email, "@", " [at] ", -1)
		fmt.Fprintf(out, "[%d] %s <%s>\n", i, newUser.Name, email)
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}
