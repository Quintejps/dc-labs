package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type infoPack struct {
	status   string
	namePack string
	Date     string
}

func main() {
	// Open the file.
	f, _ := os.Open("pacman.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Create a new array to save the txt file.
	var listTxt []string
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		listTxt = append(listTxt, line)
	}

	//var temp int = 0
	var a []infoPack
	installedPacks := make(map[string]bool)
	upgradedPacks := make(map[string]int)
	removedPacks := make(map[string]bool)

	for i := 0; i < len(listTxt); i++ {
		testArray := strings.Fields(listTxt[i])
		for _, v := range testArray {
			if v == "installed" || v == "upgraded" || v == "removed" {
				a = append(a, infoPack{status: testArray[3], namePack: testArray[4], Date: testArray[0] + " " + testArray[1]})
			}
		}
	}

	for i := 0; i < len(a); i++ {
		switch {
		case a[i].status == "installed":
			for j := 0; j < len(removedPacks); j++ {
				if removedPacks[a[i].namePack] {
					installedPacks[a[i].namePack] = true
					removedPacks[a[i].namePack] = false
					break
				}
			}
			installedPacks[a[i].namePack] = true

		case a[i].status == "upgraded":
			upgradedPacks[a[i].namePack] += 1

		case a[i].status == "removed":
			for j := 0; j < len(installedPacks); j++ {
				if installedPacks[a[i].namePack] {
					installedPacks[a[i].namePack] = false
					removedPacks[a[i].namePack] = true
				}
			}
		}
	}

	var removed, install int = 0, 0
	var isRemoved, isUpgraded bool = false, false
	var lastUpdate, lastInstall []string

	for _, v := range removedPacks {
		if v == true {
			removed++
		}
	}

	for _, v := range installedPacks {
		if v == true {
			install++
		}
	}

	f, err := os.Create("lines.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	var lines []string

	lines = append(lines, "")
	lines = append(lines, "Pacman Packages Report")
	lines = append(lines, "----------------------")
	lines = append(lines, "- Installed packages : "+strconv.Itoa(len(installedPacks)))
	lines = append(lines, "- Removed packages   : "+strconv.Itoa(removed))
	lines = append(lines, "- Upgraded packages  : "+strconv.Itoa(len(upgradedPacks)))
	lines = append(lines, "- Current installed  : "+strconv.Itoa(install))
	lines = append(lines, "")
	lines = append(lines, "List of packages")

	for i, _ := range installedPacks {
		lines = append(lines, "")
		lines = append(lines, "- Package Name: "+i)

		for k, _ := range a {
			if i == a[k].namePack {
				if a[k].status == "installed" {
					lastInstall = append(lastInstall, a[k].Date)
				}

			}
		}

		lines = append(lines, "- Install date: "+lastInstall[len(lastInstall)-1])

		for k, _ := range a {
			if i == a[k].namePack {
				if a[k].status == "upgraded" {
					lastUpdate = append(lastUpdate, a[k].Date)
					isUpgraded = true
				}

			}
		}

		if isUpgraded {
			lines = append(lines, "- Last update date: "+lastUpdate[len(lastUpdate)-1])
		}

		for j, s := range upgradedPacks {
			if i == j {
				lines = append(lines, "- How many updates: "+strconv.Itoa(s))
			}
		}

		if isUpgraded == false {
			lines = append(lines, "- Last update date: - ")
			lines = append(lines, "- How many updates: 0 ")
		}

		for k, _ := range a {
			if i == a[k].namePack {
				if a[k].status == "removed" {
					lines = append(lines, "- Removal date: "+a[k].Date)
					isRemoved = true
				}

			}
		}

		if isRemoved == false {
			lines = append(lines, "- Removal date: - ")
		}

		lines = append(lines, "")

		isRemoved = false
		isUpgraded = false
		lastUpdate = nil
	}

	for _, v := range lines {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
