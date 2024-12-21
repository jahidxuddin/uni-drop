package networkscanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

func scanIP(ip string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	pinger, err := ping.NewPinger(ip)
	if err != nil {
		fmt.Printf("Fehler beim Erstellen eines Pingers für %s: %v\n", ip, err)
		return
	}

	pinger.SetPrivileged(true) // Ermöglicht Pings ohne Root-Rechte
	pinger.Count = 1
	pinger.Timeout = 1 * time.Second
	err = pinger.Run()
	if err != nil {
		fmt.Printf("Fehler beim Pingen von %s: %v\n", ip, err)
		return
	}

	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		results <- ip
	}
}

func getHostname(ip string) string {
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		return "Unbekannt"
	}
	return strings.TrimSuffix(names[0], ".")
}

func getLocalSubnet() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			if strings.HasPrefix(ipNet.IP.String(), "169.254.") {
				continue
			}
			return ipNet.String(), nil
		}
	}

	return "", fmt.Errorf("keine gültige lokale Subnetzadresse gefunden")
}

func RunNetworkScan() (map[string]string, error) {
	subnet, err := getLocalSubnet()
	if err != nil {
		return nil, fmt.Errorf("fehler beim Abrufen des lokalen Subnetzes: %w", err)
	}

	ipRange := strings.Split(subnet, "/")[0]
	ipParts := strings.Split(ipRange, ".")
	if len(ipParts) != 4 {
		return nil, fmt.Errorf("ungültiges Subnetzformat")
	}

	var wg sync.WaitGroup
	results := make(chan string, 256)
	devices := make(map[string]string) // Speichert IP-Adressen und Hostnamen

	for i := 1; i < 255; i++ {
		ip := fmt.Sprintf("%s.%s.%s.%d", ipParts[0], ipParts[1], ipParts[2], i)
		wg.Add(1)
		go scanIP(ip, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for ip := range results {
		hostname := getHostname(ip)
		devices[ip] = hostname
	}

	return devices, nil
}
