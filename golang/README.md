## Phase 1: Learn Go + Networking

### 1. Concurrent Port Scanner

Build:

* TCP connect scan
* CIDR support
* Concurrent workers
* Service banner grabbing

Learn:

* Goroutines
* Channels
* Sockets
* Timeouts

---

### 2. HTTP Recon Tool

Like a miniature `httpx`.

Features:

* Status code
* Page title
* Server header
* TLS certificate info
* Redirect detection

Learn:

* HTTP
* TLS
* Parsing

---

### 3. DNS Recon Toolkit

Features:

* A record lookup
* MX records
* TXT records
* Reverse DNS
* Subdomain brute forcing

Learn:

* DNS internals
* Resolvers
* Concurrency

---

## Phase 2: Pentesting Tools

### 4. Web Directory Brute Forcer

Mini Gobuster.

Features:

* Wordlists
* Extension support
* Filtering by response size
* Rate limiting

Learn:

* Efficient HTTP requests
* Thread pools

---

### 5. Nuclei-style Vulnerability Scanner

One of the best learning projects.

Features:

* YAML templates
* HTTP requests
* Matchers
* Reporting

Example concepts:

```yaml
id: exposed-git

request:
  method: GET
  path: /.git/config

matchers:
  - status: 200
```

Learn:

* Scanner architecture
* Rule engines
* Automation

---

### 6. HTTP Proxy

A basic Burp-like proxy.

Features:

* Intercept requests
* Modify requests
* Logging
* Response inspection

Learn:

* HTTP internals
* MITM concepts
* Traffic analysis

---

## Phase 3: Active Directory & Internal Networks

### 7. SMB Enumerator

Features:

* Detect SMB shares
* Enumerate versions
* Authentication testing against systems you own or are authorized to assess

Learn:

* SMB protocol basics
* Windows networking

---

### 8. LDAP Enumeration Tool

Features:

* Query LDAP
* Enumerate users/groups in lab environments
* Export results

Learn:

* Enterprise network protocols
* Directory services

---

### 9. BloodHound Data Parser

Instead of collecting data, start by parsing existing BloodHound JSON exports.

Learn:

* Graphs
* AD attack path concepts
* Large datasets

---

## Phase 4: WiFi Security

WiFi work is harder in Go because most serious tooling relies on raw packet access and monitor mode support.

### 10. 802.11 Packet Analyzer

Using:

* `gopacket`

Features:

* Parse PCAP files
* Detect:

  * SSIDs
  * BSSIDs
  * Clients
  * Encryption types

Learn:

* WiFi frame structure
* Packet analysis

---

### 11. Wireless Recon Dashboard

Read PCAP captures and produce:

* AP inventory
* Client inventory
* Security settings
* Channel usage

Learn:

* Data processing
* Reporting

---

### 12. EAPOL Handshake Analyzer

Features:

* Detect WPA/WPA2 handshakes in captures
* Validate completeness

Learn:

* WPA authentication flow
* Wireless security concepts

---

## Phase 5: Advanced Offensive Security

### 13. Asset Discovery Platform

Combine:

* DNS
* HTTP
* Port scanning
* TLS inspection

Output:

```json
{
  "host": "example.com",
  "ports": [80,443],
  "title": "Example",
  "technologies": ["nginx"]
}
```

This becomes a real recon framework.

---

### 14. Attack Surface Mapper

Input:

* Domain

Output:

* Subdomains
* Open ports
* Technologies
* Certificates

Basically a lightweight Attack Surface Management tool.

---

### 15. Network Traffic Analyzer

Features:

* Parse PCAPs
* Extract:

  * DNS requests
  * HTTP metadata
  * TLS metadata

Learn:

* Incident response fundamentals
* Protocol analysis

---

## Go Libraries Worth Learning

```text
net
net/http
crypto/tls
context
sync
bufio
encoding/json
```

Then:

```text
github.com/google/gopacket
github.com/spf13/cobra
github.com/spf13/viper
gopkg.in/yaml.v3
```

If I were preparing for a professional pentesting or red-team role today, I'd build these in order:

1. Concurrent Port Scanner
2. HTTP Recon Tool
3. DNS Recon Toolkit
4. Directory Brute Forcer
5. Nuclei-style Scanner
6. HTTP Proxy
7. SMB Enumerator
8. 802.11 Packet Analyzer
9. Attack Surface Mapper


