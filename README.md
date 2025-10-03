# Go-based-PCAP-Packet-Analyzer
A lightweight Go-based CLI tool for parsing and inspecting `.pcap` files.
It prints detected packet layers and displays a Wireshark-style hex dump for each captured frame.
---
## Features
- **Offline PCAP Analysis** – Reads packets from a `.pcap` file.
- **Layer Inspection** – Lists all protocol layers detected in each packet.
- **Hex Dump** – Displays raw packet data in a Wireshark-style hexadecimal format.
- **Simple CLI Interface** – Easy to use, minimal dependencies.
---
## Installation

1. **Clone the repository**
 ```bash
 git clone https://github.com/<your-username>/gopcap-inspector.git
 cd gopcap-inspector
 ```
2. **Install dependencies**
```bash
go mod tidy
```
3. **Build the executable**
```bash
Build the executable
```
---
## Usage

Run the tool with the -pcap flag to specify the path to your .pcap file:
```bash
./gopcap-inspector -pcap path/to/file.pcap
```

Example Output:
```
Reading sample.pcap...

--- Packet #1 ---
Layer: Ethernet
Layer: IPv4
Layer: TCP
Raw frame length: 66
0000  00 1c c0 a3 12 5b 00 0a 95 9d 68 16 08 00 45 00
0010  00 34 a6 f2 40 00 40 06 6b 6c c0 a8 01 64 c0 a8
0020  01 c8 00 50 1f 90 4a 3b 7e 25 00 00 00 00 50 02
0030  20 00 91 7c 00 00 00 00 00 00

--- Packet #2 ---
Layer: Ethernet
Layer: IPv4
Layer: UDP
Raw frame length: 74
...

Done. Total packets: 25
```
