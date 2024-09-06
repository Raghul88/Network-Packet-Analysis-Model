package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Open the pcap file
	handle, err := pcap.OpenOffline("network_traffic.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Initialize counters for traffic analysis
	var totalPackets int
	var totalBytes int

	for packet := range packetSource.Packets() {
		totalPackets++
		totalBytes += len(packet.Data())

		// Print packet summary
		fmt.Printf("Packet #%d: %s\n", totalPackets, packet.String())
	}

	// Output total packets and bytes
	fmt.Printf("\nTotal Packets: %d\n", totalPackets)
	fmt.Printf("Total Bytes: %d\n", totalBytes)
}
