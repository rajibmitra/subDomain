# Subdomains CLI Tool
Fetches subdomains for a given domain using the subdomain.center API.

# Installation

Prerequisites:
Ensure you have Go installed (version 1.16 or newer). If not, download and install it from Go's official website.

Steps:
Clone the Repository:

Copy code
git clone https://github.com/YOUR_USERNAME/subdomains-cli.git
cd subdomains-cli

Build the Tool:

go build -o subdomains .

This command compiles the code and produces a binary named subdomains.

(Optional) Add to PATH:

If you want to use the tool from any directory, move the subdomains binary to a directory in your PATH or add the current directory to your PATH.

Replace YOUR_USERNAME with your actual GitHub username (or the username of the repository owner if different). If the repository URL is different, make sure to update that as well.

subdomains [flags] [domain]
Flags:
-u, --url: Domain to fetch subdomains for.
-o, --output: Output format. Currently, only "txt" is supported, which writes the subdomains to a file named subdomains.txt.

Examples:
Fetch subdomains for google.com and print them in the console:


subdomains -u google.com
Fetch subdomains for google.com and write them to subdomains.txt:

subdomains -u google.com -o txt

Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

License
Apache 2 
