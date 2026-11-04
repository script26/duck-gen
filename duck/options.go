package duck

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var logFile string = ConfigPath() + "/aliases_history.txt"

func UserHome() (home string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}

func ConfigPath() (configPath string) {
	return UserHome() + "/.local/share/duck-gen"
}

func ShowHistory() {
	file, err := os.Open(logFile)

	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("No aliases history found.")
	}
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		start := strings.Index(line, " - ") + 3
		if start == 2 {
			start = 0
		}

		removeTime := line[start:]
		fmt.Println(removeTime)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func ShowVerboseHistory() {
	data, err := os.ReadFile(logFile)

	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("No aliases history found.")
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func GenerateAllAddresses(token string) {
	write := []byte("\nGenerated 30 Aliases\n")
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	file.WriteString(string(write))

	for i := 0; i < 30; i++ {
		address, err := GenerateAddress(token)
		if err != nil {
			log.Fatal("Could not generate Private Duck Address: ", err)
		}
		fmt.Println(address)
	}

	file.WriteString("")
}

func ShowMenu(token string) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	var args []string = os.Args

	var choice string = ""
	if len(args) != 2 {
		fmt.Println("Please choose an option:")
		fmt.Println("1 - Generate email alias")
		fmt.Println("2 - Show aliases history")
		fmt.Println("3 - Show verbose aliases history")
		fmt.Println("4 - Generate all aliases")
		fmt.Print("Enter your choice (1 - 4): ")
		fmt.Scanln(&choice)
	} else {
		choice = args[1]
	}

	switch choice {
	case "generate", "1":
		address, err := GenerateAddress(token)
		if err != nil {
			log.Fatal("Could not generate Private Duck Address: ", err)
		}
		fmt.Println(address)

	case "history", "2":
		ShowHistory()

	case "verbose_history", "3":
		ShowVerboseHistory()

	case "generate_all", "4":
		GenerateAllAddresses(token)

	default:
		if len(args) != 2 {
			fmt.Println("Invalid choice. Please run the script " +
				"again and select a valid option.")
		} else {
			fmt.Println("Usage: " + args[0] + " [generate|history|" +
				"verbose_history|generate_all]")
		}
	}
}
