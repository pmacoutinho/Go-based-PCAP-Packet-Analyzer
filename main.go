package main

import (
    "flag"
    "fmt"
    "log"

    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

func main() {
    pcapFile := flag.String("pcap", "", "Path to the pcap file")
    flag.Parse()

    if *pcapFile == "" {
        log.Fatal("Please specify a pcap file with -pcap")
    }

    fmt.Printf("Reading %s...\n\n", *pcapFile)

    handle, err := pcap.OpenOffline(*pcapFile)
    if err != nil {
        log.Fatalf("Error opening pcap file: %v", err)
    }
    defer handle.Close()

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    count := 0
    for packet := range packetSource.Packets() {
        count++
        fmt.Printf("--- Packet #%d ---\n", count)

        // Print detected layers
        for _, layer := range packet.Layers() {
            fmt.Println("Layer:", layer.LayerType())
        }

        // Get raw bytes of the whole frame
        raw := packet.Data()
        fmt.Printf("Raw frame length: %d\n", len(raw))

        // Print hex dump (Wireshark style)
        printHexDump(raw)
        fmt.Println()
    }

    fmt.Printf("Done. Total packets: %d\n", count)
}

func printHexDump(data []byte) {
    for i := 0; i < len(data); i += 16 {
        end := i + 16
        if end > len(data) {
            end = len(data)
        }
        // Offset (like Wireshark’s left column)
        fmt.Printf("%04x  ", i)

        // Print bytes in hex with a space every byte
        for j := i; j < end; j++ {
            fmt.Printf("%02x ", data[j])
        }
        fmt.Println()
    }
}
