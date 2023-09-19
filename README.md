# Subdomains CLI Tool

A tool that fetches subdomains for a given domain using the subdomain.center API.

## Installation

### Prerequisites:

- **Go**: Ensure you have Go installed. If not, you can download and install it from [Go's official website](https://golang.org/dl/).

### Step-by-step Guide:

1. **Clone the Repository**:
   ```bash
   https://github.com/rajibmitra/subDomain.git
   cd subDomain

2. **Build the Tool**:
   
  go build -o subdomains .

This command compiles the code and produces a binary named subdomains.

(Optional) Add to PATH:

If you want to use the tool from any directory, move the subdomains binary to a directory in your PATH or add the current directory to your PATH.


subdomains [flags] [domain]
Flags:
-u, --url: Domain to fetch subdomains for.
-o, --output: Output format. Currently, only "txt" is supported, which writes the subdomains to a file named subdomains.txt.

Examples:
Fetch subdomains for xyz.com and print them in the console:


subdomains -u xyz.com
Fetch subdomains for xyz.com and write them to subdomains.txt:

subdomains -u xyz.com -o txt
Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

License
Licensed under the Apache 2 License.
