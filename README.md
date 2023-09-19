# Subdomains CLI Tool

Subdomains play a crucial role in the structure and navigation of websites. They can be indicative of different services, regions, or purposes under the umbrella of the main domain. However, from a cybersecurity perspective, each subdomain can also be a potential entry point for attackers. As applications grow in complexity, managing and keeping track of subdomains becomes increasingly challenging.

The Subdomains CLI Tool is designed to simplify this process. By leveraging the power of the subdomain.center API, it provides a streamlined method to fetch and list subdomains associated with a primary domain. This is not only beneficial for developers and website administrators in managing their online assets but is also a vital tool for cybersecurity professionals. Here's why:

Vulnerability Scanning: Before launching a penetration test, security experts often enumerate subdomains to identify potential weak points. The more subdomains they're aware of, the more comprehensive their tests can be.

Phishing Detection: Malicious actors often create subdomains under legitimate-looking domains to deceive users. By regularly monitoring and understanding the subdomains associated with a primary domain, businesses can detect suspicious activity early.

Asset Management: Large organizations with multiple online services can use this tool to maintain an inventory of their web assets, ensuring no subdomain goes unnoticed or unchecked.

SEO and Brand Monitoring: Brands can monitor subdomains to detect potential infringements or misuse of their brand name in URLs.

This CLI tool provides a fast and reliable way to fetch subdomains, aiding both in developmental operations and security assessments. Whether you're a website administrator, a developer, or a security expert, this tool is designed to make your life easier and your operations more secure.

## Installation

### Prerequisites:

- **Go**: Ensure you have Go installed. If not, you can download and install it from [Go's official website](https://golang.org/dl/).

### Step-by-step Guide:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/rajibmitra/subDomain.git
   cd subDomain

2. **Build the Tool**:
   go build -o subdomains .

   This command compiles the code and produces a binary named subdomains.
      
   (Optional) Add to PATH:
   If you want to use the tool from any directory, move the subdomains binary to a directory in your PATH or add the current directory to your PATH.

4. **Download from release**
   If you are unable to build the tool - goto https://github.com/rajibmitra/subDomain/releases/tag/1.0.0
   Download the binary as per your need, and use it. 

   subdomains [flags] [domain]
   Flags:
   -u, --url: Domain to fetch subdomains for.
   -o, --output: Output format. Currently, only "txt" is supported, which writes the subdomains to a file named subdomains.txt.
   
   Examples:

   Fetch subdomains for xyz.com and print them in the console:
   subdomains -u xyz.com
   
   Fetch subdomains for xyz.com and write them to subdomains.txt:
   subdomains -u xyz.com -o txt

   **Contributing**

   Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
   
   **License**

   Licensed under the Apache 2 License.
