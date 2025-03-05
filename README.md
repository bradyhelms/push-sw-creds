# push-sw-creds

## Overview
`push-sw-creds` is a Go-based script that automates the provisioning of new user accounts across multiple Cisco switches. The script collects administrator credentials and applies the new user credentials to a list of target switches.

## Features
- Securely collects admin username and password
- Uses a predefined list of IP addresses for provisioning
- Automates the creation of new user accounts across Cisco switches
- Written in Go for efficiency and portability

## Prerequisites
Ensure you have the following installed:
- Network access to the target Cisco switches
- Valid administrator credentials for authentication

## Installation
Clone the repository:
```sh
git clone https://github.com/bradyhelms/push-sw-creds.git
cd push-sw-creds
```

Usage:
```sh
push-sw-creds
```

Build from source:
```sh
go build -o push-sw-creds src/main.go
```

### Input File Format (`config/ips.txt`)
The file should contain a list of IP addresses, one per line:
```
192.168.1.1
192.168.1.2
192.168.1.3
```

## Contributions
Contributions are welcome! Please submit a pull request or open an issue if you have any suggestions or improvements.

## Author
[Brady Helms](https://github.com/bradyhelms)


