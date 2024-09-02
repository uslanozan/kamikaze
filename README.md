# Kamikaze 🚀
Unleash the power of subdomain discovery!

![Ekran görüntüsü 2024-08-29 125343](https://github.com/user-attachments/assets/abe5b4fb-8342-47de-94af-f4058b355294)


## Features 🔍
Kamikaze comes packed with powerful features to supercharge your subdomain reconnaissance:

- 🌐 Subdomain discovery
- 🖥️ IP address resolution for found subdomains
- 🚦 HTTP response code checking

## Installation 💻
Get Kamikaze up and running in no time:

1. Clone the repository:
   ```bash
   git clone https://github.com/uslanozan/kamikaze.git
   
2. Navigate to the project directory:
   ```bash
   cd kamikaze

3. Build the project:
   ```bash
   go build
   
## Usage 🛠️
Start your subdomain exploration journey:

1. To discover subdomains:
   ```bash
   ./kamikaze -d example.com -wl wordlist.txt
   
2. Use following options:
- -d : Specify the target domain.
- -wl: Provide the path to the wordlist file (Optional).
  
## Example 📋

See Kamikaze in action:
1. Bruteforce with default txt:
   ```bash
   ./kamikaze -d github.com -wl default

2. Bruteforce with spesific txt:
   ```bash
   ./kamikaze -d github.com -wl /user/bin/rockyou.txt

## Contributing 🤝
Join the Kamikaze community and make it even better:

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer ⚠️
Use responsibly and ethically:

This tool is for educational and ethical testing purposes only. Always ensure you have permission before scanning domains you do not own or have explicit permission to test.

## Contact 📬
Get in touch with the me:

- GitHub: [@uslanozan](https://github.com/uslanozan)
- Linkedin: [@uslanozan](https://www.linkedin.com/in/uslanozan/)
