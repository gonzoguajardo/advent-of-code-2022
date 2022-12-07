package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	foundLs := false;
	directoryMap := make(map[string][]string)
	var currentDirectoryFiles []string
	var currentDirectory string
	for scanner.Scan() {
		currentString := scanner.Text()
		if currentString[0:1] == "$" {
			upDirectorySlice := strings.Split(currentString, "$ cd ")
			if len(upDirectorySlice) > 1 {
				if upDirectorySlice[0] == ".." {
				} else {
					if foundLs {
						directoryMap[currentDirectory] = currentDirectoryFiles
						foundLs = false
					}
					currentDirectory = upDirectorySlice[1]
					currentDirectoryFiles = make([]string, 0)
				}
				continue
			}
			dollarSplit := strings.Split(currentString, "$ ")
			if dollarSplit[1] == "ls" {
				foundLs = true
			} else {
				directoryMap[currentDirectory] = currentDirectoryFiles
				foundLs = false
			}
		} else if foundLs {
			currentDirectoryFiles = append(currentDirectoryFiles, currentString)
		}
	}
	directoryMap[currentDirectory] = currentDirectoryFiles

	directorySizeMap := make(map[string]int)

	for len(directorySizeMap) < len(directoryMap) {
	MapLoop:
		for k, v := range(directoryMap) {
			totalSize := 0
			for _, file := range(v) {
				dirSlice := strings.Split(file, "dir ")
				if file[0:4] == "dir "{
					if size, ok := directorySizeMap[dirSlice[1]]; ok {
						totalSize += size
					} else {
						continue MapLoop
					}
				} else {
					spaceSplit := strings.Split(file, " ")
					size, _ := strconv.Atoi(spaceSplit[0])
					totalSize += size
				}

			}
			directorySizeMap[k] = totalSize
		}
	}
	log.Println(directorySizeMap)
	total := 0
	for  k, value := range(directorySizeMap){
		// if k != "/" && value < 100000 {
		// 	total += value
		// }
		if k != "/" {
			total += value
		}
	}
	log.Println(total)
}
