package duck

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var API_KEY string = "YOUR_API_KEY_HERE"
var HISTORY_PATH string = os.Getenv("HOME") + "/.local/share/ddg"
var LOG_FILE string = HISTORY_PATH + "/aliases_history.txt"
var onlyOneRun bool = true

	func generate_alias() {
		client := &http.Client{}
		var data = strings.NewReader(`{}`)
		req, err := http.NewRequest("POST",
			"https://quack.duckduckgo.com/api/email/addresses", data)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+API_KEY)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "DuckDuckGo/5")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		rsp, err := io.ReadAll(resp.Body)
		RESPONSE := string(rsp)
		if err != nil {
			log.Fatal(err)
		}

		ALIAS := RESPONSE[12 : len(RESPONSE)-2]

		if RESPONSE == "{\"error\":\"invalid_token\"}" {
			ALIAS = ""
		}

		if ALIAS != "" {
			FULL_ALIAS := ALIAS + "@duck.com"
			TIMESTAMP := time.Now().Format("0000-00-00 00:00:00 AM")
			write := []byte("\n" + TIMESTAMP + " - " + FULL_ALIAS)
			if onlyOneRun {
				write = []byte("\n\n" + TIMESTAMP + " - " + FULL_ALIAS)
			}

			file, err := os.OpenFile(LOG_FILE,
				os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close()
			if _, err = file.WriteString(string(write)); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Email alias generated: " + FULL_ALIAS)
			return
		}

		fmt.Println("Failed to generate alias. Response: " + RESPONSE)
	}

	func show_history() {
		file, err := os.Open(LOG_FILE)

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

	func show_verbose_history() {
		data, err := os.ReadFile(LOG_FILE)

		if errors.Is(err, os.ErrNotExist) {
			log.Fatal("No aliases history found.")
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(data))
	}

	func generate_all_aliases() {
		write := []byte("\n\nGenerated 30 Aliases")
		file, err := os.OpenFile(LOG_FILE,
			os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		if _, err = file.WriteString(string(write)); err != nil {
			log.Fatal(err)
		}

		onlyOneRun = false

		for i := 0; i < 30; i++ {
			generate_alias()
		}
	}

	func show_menu() {
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
			generate_alias()

		case "history", "2":
			show_history()

		case "verbose_history", "3":
			show_verbose_history()

		case "generate_all", "4":
			generate_all_aliases()

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

