package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func printAll() {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error Q")
	} else {
		s := string(b)
		fmt.Println(s)

	}
}
func appendLine(line string) {
	file, err := os.ReadFile("data.txt")

	if err != nil {
		fmt.Println("err", err)
	}

	if len(file) == 0 {
		os.WriteFile("data.txt", []byte(line), 0644)
	} else {
		newFile := ""
		if strings.HasSuffix(string(file), "\n") {
			newFile = string(file) + line

		} else {
			newFile = string(file) + "\n" + line

		}

		os.WriteFile("data.txt", []byte(newFile), 0644)
	}
}
func prependLine(line string) {
	file, err := os.ReadFile("data.txt")

	if err != nil {
		fmt.Println("err", err)
	}

	if len(file) == 0 {
		os.WriteFile("data.txt", []byte(line), 0644)
	} else {
		newFile := ""
		if strings.HasPrefix(string(file), "\n") {
			newFile = line + string(file)

		} else {
			newFile = line + "\n" + string(file)

		}

		os.WriteFile("data.txt", []byte(newFile), 0644)
	}
}

func readLines() []string {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error Q")
		return []string{}
	} else {
		s := string(b)
		return strings.Split(s, "\n")
	}
}

func removeFirst(s []string) error {
	if len(s) == 0 {
		return errors.New("err")
	}

	newSlice := append([]string{}, s[1:]...)
	err := os.WriteFile("data.txt", []byte(strings.Join(newSlice, "\n")), 0644)

	if err != nil {
		return err
	}

	return nil
}
func removeLast(s []string) error {
	if len(s) == 0 {
		return errors.New("err")
	}

	newSlice := append([]string{}, s[:len(s)-1]...)
	err := os.WriteFile("data.txt", []byte(strings.Join(newSlice, "\n")), 0644)

	if err != nil {
		return err
	}

	return nil
}

func removeAt(s []string, n int) error {
	if len(s) == 0 {
		return errors.New("err")
	}

	if n < 0 || n >= len(s) {
		return errors.New("индекс за пределами диапазона")
	}

	newSlice := append(s[:n], s[n+1:]...)
	err := os.WriteFile("data.txt", []byte(strings.Join(newSlice, "\n")), 0644)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	printAll()
	// appendLine("yaya")
	// prependLine("prependLine")
	fmt.Println(readLines())
	// removeFirst(readLines())
	// removeLast(readLines())

	fmt.Println(removeAt(readLines(), 1))

}
