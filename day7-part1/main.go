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

	currentDirectoryPath := ""
	currentDirectorySlice := make([]string, 0)
	directoryMap := make(map[string][]string)
	directorySizeMap := make(map[string]int)
	// var unresolvedDirectory bool
	var previousLs bool

	count := 0

	for scanner.Scan() {
		currentString := scanner.Text()
		// log.Println(currentString)
		isCommand := strings.Contains(currentString, "$")

		if isCommand {
			if previousLs {
				directoryMap[currentDirectoryPath] = currentDirectorySlice
				ok, currentDirectorySize := checkDirectory(currentDirectoryPath, currentDirectorySlice, directorySizeMap)
				if ok {
					directorySizeMap[currentDirectoryPath] = currentDirectorySize
				}
				previousLs = false
			}
			commandSplit := strings.Split(currentString, "$ ")
			command := commandSplit[1]
			cd := strings.Contains(command, "cd ")
			if cd {
				cdSplit := strings.Split(command, "cd ")
				directory := cdSplit[1]
				isUpDirectory := strings.Contains(directory, "..")
				if isUpDirectory {
					removeLast := currentDirectoryPath[:len(currentDirectoryPath) - 1]
					lastIndex := strings.LastIndex(removeLast, "/")
					if lastIndex != 0 {
						newDirectoryPath := currentDirectoryPath[:lastIndex+1]
						currentDirectoryPath = newDirectoryPath
					} else {
						currentDirectoryPath = "/"
					}

					ok, currentDirectorySize := checkDirectory(currentDirectoryPath, directoryMap[currentDirectoryPath], directorySizeMap)
					if ok {
						directorySizeMap[currentDirectoryPath] = currentDirectorySize
					}

				} else {
					if len(currentDirectoryPath) != 0 {
					}
					if currentDirectoryPath == "" {
						currentDirectoryPath += directory
					} else {
						currentDirectoryPath += directory + "/"
					}
					currentDirectorySlice = make([]string, 0)

				}
			}
		} else {
			previousLs = true
			currentDirectorySlice = append(currentDirectorySlice, currentString)
			// log.Println(currentString)
		}
		count++
	}
	directoryMap[currentDirectoryPath] = currentDirectorySlice

	totalSize := 0
	for len(directoryMap) > 0 {
		for k, v := range(directoryMap){
			directoryFound := true
			if value, ok := directorySizeMap[k]; ok {
				if value < 100000 {
					totalSize += value
				} else {
				}
			} else {
				ok, currentDirectorySize := checkDirectory(k, v, directorySizeMap)
				if ok {
					directorySizeMap[k] = currentDirectorySize
					if currentDirectorySize < 100000 {
						totalSize += currentDirectorySize
					}
				} else {
					directoryFound = false
				}
			}
			if directoryFound {
				delete(directoryMap, k)
			}
		}

	}

	log.Println(totalSize)

}

func checkDirectory(currentDirectory string, directorySlice []string, directorySizeMap map[string]int) (bool, int){
	currentDirectorySize := 0
	for _, currentString := range(directorySlice) {
		isDir := strings.Contains(currentString, "dir ")
		spaceSplit := strings.Split(currentString, " ")
		if isDir {
			directoryToCheck := currentDirectory + spaceSplit[1] + "/"
			if value, ok := directorySizeMap[directoryToCheck]; ok {
				currentDirectorySize += value
			} else {
				return false, 0
			}
		} else {
			currentSize , _ := strconv.Atoi(spaceSplit[0])
			currentDirectorySize += currentSize
		}
	}
	return true, currentDirectorySize
}
