1) Install Go onto your OS(depending on which type of OS you use)(Example: Windows, Linux, MacOS)
2) Download VS code onto the system
3) Download go extension in VS Code
4) Install Wireshark
5) Run the Wireshark for your desired connection and save it as network_traffic.pcap.
6) open terminal and go the the location of the .pcap file directory.
7)open terminal on the system and initialize go module:
go mod init network-traffic-analysis
8) Install all dependencies which are gopacket library:
go get github.com/google/gopacket
9) Install the main.go into your directory
10) in VS code run main.go it will give all the traffic analysis and overflow details.