package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func bruteForce(domain string, txtPath string) {
	originDomain := domain

	file, err := os.Open(txtPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Printf("\n----Brute force is starting---\n\n")

	for scanner.Scan() {
		line := scanner.Text()
		domain = line + "." + domain

		if checkIP(domain) != "IP ERROR" {
			fmt.Print("SUBDDOMAIN: "+domain + "  " + "IP: "+checkIP(domain) + "  "+"RESPONSE CODE: ")
			fmt.Print(checkResponse(domain))
			fmt.Println()
		}
		domain = originDomain
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Brute force is finished!")
}

func checkIP(domain string) string {

	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return "IP ERROR"
}

func checkResponse(domain string) string {
    resp, err := http.Get("https://" + domain)
    if err != nil {
        return "STATUS CODE NOT FOUND"
    }
    defer resp.Body.Close()
    return strconv.Itoa(resp.StatusCode)
}

func main() {

	fmt.Println("\033[36m"+`
 ___  __    ________  _____ ______   ___  ___  __    ________  ________  _______      
|\  \|\  \ |\   __  \|\   _ \  _   \|\  \|\  \|\  \ |\   __  \|\_____  \|\  ___ \     
\ \  \/  /|\ \  \|\  \ \  \\\__\ \  \ \  \ \  \/  /|\ \  \|\  \\|___/  /\ \   __/|    
 \ \   ___  \ \   __  \ \  \\|__| \  \ \  \ \   ___  \ \   __  \   /  / /\ \  \_|/__  
  \ \  \\ \  \ \  \ \  \ \  \    \ \  \ \  \ \  \\ \  \ \  \ \  \ /  /_/__\ \  \_|\ \ 
   \ \__\\ \__\ \__\ \__\ \__\    \ \__\ \__\ \__\\ \__\ \__\ \__\\________\ \_______\
    \|__| \|__|\|__|\|__|\|__|     \|__|\|__|\|__| \|__|\|__|\|__|\|_______|\|_______|                                                     
   `+"\033[0m")
   
	
	// Komut satırı argümanlarını kontrol et
    if len(os.Args) < 2 {
        fmt.Println("--Kamikaze--\nTo choose domain -d <EXAMPLE.COM>\n"+
					"To Choose wordlist -wl <WORDLIST PATH>\n" +
					"NOTE: You can leave wordlist space or write default for default wordlist")
        return
    }

    var domain string
    var wordlist string

    // Argümanları kontrol et ve değerlere ata
    for i := 1; i < len(os.Args); i++ {
        switch os.Args[i] {
        case "-d":
            if i+1 < len(os.Args) {
                domain = os.Args[i+1]
                i++ // Bir sonraki değeri atla
            }
        case "-wl":
            if i+1 < len(os.Args) {
                wordlist = os.Args[i+1]
                i++ // Bir sonraki değeri atla
            }
		case "--help":
			fmt.Println("\n--THIS IS HELP PAGE--\n==>To choose domain -d <EXAMPLE.COM>\n"+
					"==>To Choose wordlist -wl <WORDLIST PATH>\n" +
					"NOTE: You can leave wordlist space or write default for default wordlist")
			return
        }
    }


    // Varsayılan değerler
    if wordlist == "" || strings.ToLower(wordlist) == "default"{
        wordlist = "subdomain_wordlist.txt"
    }

	if domain != ""{
		fmt.Println(domain,wordlist)
		bruteForce(domain,wordlist)
	}else{
		fmt.Println("Wrong argument. Please enter a domain or --help to take help.")
	}
}
