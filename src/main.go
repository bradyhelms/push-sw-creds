package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
)

func main() {
	// Load the switch IPs
	ips := getIPs()

	admin_user, admin_pass := getAdminCreds()
	userUsername, userPass := getUserCreds()

	clientConfig := &ssh.ClientConfig{
		User: admin_user,
		Auth: []ssh.AuthMethod{
			ssh.Password(admin_pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	for _, ip := range ips {
		ipWithPort := fmt.Sprintf("%s:22", ip)

		// Connect to switch
		conn, err := ssh.Dial("tcp", ipWithPort, clientConfig)
		if err != nil {
			log.Fatalf("Failed to dial: %s", err)
		}
		defer conn.Close()

		// Create a new session
		session, err := conn.NewSession()
		if err != nil {
			log.Fatalf("Failed to create session: %s", err)
		}
		defer session.Close()

		// Run the command to add the SSH user
		command := fmt.Sprintf("username %s privelege 15 secret 5 %s", userUsername, userPass)
		output, err := session.CombinedOutput(command)
		if err != nil {
			log.Fatalf("Failed to run command: %s", err)
		}

		fmt.Printf("Command output: %s\n", output)
	}

}

func getAdminCreds() (admin_user, admin_pass string) {
	fmt.Print("Enter admin username: ")
	fmt.Scanln(&admin_user)

	fmt.Print("Enter admin password: ")
	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Failed to read password: %s", err)
	}

	admin_pass = string(pass)
	return admin_user, admin_pass
}

func getUserCreds() (userUsername, userPass string) {
	fmt.Print("Enter new username: ")
	fmt.Scanln(&userUsername)
	fmt.Print("Enter new user password: ")
	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Failed to read password: %s", err)
	}

	userPass = string(pass)
	return userUsername, userPass
}

func getIPs() []string {
	file, err := os.Open("config/ips.txt")
	if err != nil {
		log.Fatalf("Failed to open file; %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ips []string

	for scanner.Scan() {
		ip := scanner.Text()
		ips = append(ips, ip)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	return ips
}
