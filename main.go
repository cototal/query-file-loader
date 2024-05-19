package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage...")
		os.Exit(0)
	}

	fileName := args[0]
	file, err := os.Open(fileName)
	dir, _ := filepath.Abs(filepath.Dir(fileName))
	fmt.Println(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	queryMethods := make(map[string]string)

	scanner := bufio.NewScanner(file)

	var methodName string
	var queryBuilder strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "--") {
			if methodName != "" {
				queryMethods[methodName] = queryBuilder.String()
				queryBuilder.Reset()
			}
			methodName = strings.TrimSpace(strings.TrimPrefix(line, "--"))
			fmt.Println(methodName)
		} else {
			queryBuilder.WriteString(line + " ")
		}
	}
	queryMethods[methodName] = queryBuilder.String()

	var fileContentBuilder strings.Builder

	fileContentBuilder.WriteString("package queries")
	for method, query := range queryMethods {
		fileContentBuilder.WriteString("\n\n" + fmt.Sprintf(`func %s() string {
	return %q
}`, method, strings.TrimSpace(query)))
	}

	goFileName := filepath.Join(dir, "all.go")
	err = os.WriteFile(goFileName, []byte(fileContentBuilder.String()), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
