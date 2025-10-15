# 🌐 OrchardNet API

A high-performance network stress testing framework built in Go, featuring multiple attack vectors and an intuitive web dashboard.

![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-Educational-red)
![Status](https://img.shields.io/badge/status-Active-success)

## ⚠️ LEGAL DISCLAIMER

**THIS SOFTWARE IS FOR EDUCATIONAL AND AUTHORIZED TESTING PURPOSES ONLY.**

- ✅ Use ONLY on systems you own or have explicit written permission to test
- ❌ Unauthorized network attacks are **ILLEGAL** and punishable by law
- ❌ The authors are NOT responsible for misuse or damage caused by this software
- ❌ Deploying this against systems without permission may result in criminal prosecution

**BY USING THIS SOFTWARE, YOU AGREE TO:**
- Take full responsibility for your actions
- Comply with all applicable local, state, and federal laws
- Only use this tool in authorized penetration testing or research environments

## 🚀 Features

- **Multiple Attack Vectors**
  - 🔥 HTTP Flood - Application layer attacks
  - ⚡ SYN Flood - TCP handshake exhaustion
  - 💣 UDP Amplification - Bandwidth saturation

- **Advanced Evasion**
  - IP Spoofing with geo-targeting
  - Randomized user agents
  - Custom packet crafting

- **Modern Interface**
  - Beautiful web dashboard
  - Auto-opens browser on startup
  - Real-time attack monitoring
  - RESTful API endpoints

- **High Performance**
  - Multi-threaded worker system
  - Raw socket implementation
  - Efficient packet generation

## 📋 Prerequisites

- **Go 1.24+** installed
- **Root/Administrator privileges** (required for raw sockets)
- **Linux** (recommended) or compatible Unix system
- Network interface with raw socket support

## 🔧 Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/orchardnet-api.git
cd orchardnet-api

# Install dependencies
go mod download

# Build the binary
go build -o orchardnet-api

# Run with sudo (required for raw sockets)
sudo ./orchardnet-api
```

The dashboard will automatically open in your browser at `http://localhost:8080`

## 💻 Usage

### Web Dashboard (Recommended)

1. Start the server:
   ```bash
   sudo ./orchardnet-api
   ```

2. Browser automatically opens to the dashboard

3. Configure your test:
   - Enter target domain/IP
   - Select attack type (HTTP/SYN/UDP)
   - Adjust workers (10-500)
   - Set duration (5-120 seconds)

4. Click "🚀 Launch Attack"

### API Endpoints

#### Launch Attack
```bash
POST /api/v1/attack
Content-Type: application/json

{
  "target": "example.com",
  "type": "http",
  "workers": 100,
  "duration": 10,
  "api_key": "th0rn3-0rch@rd-k3y-2024"
}
```

**Response:**
```json
{
  "job_id": "job_20251015074620_1",
  "status": "launched"
}
```

#### Check Status
```bash
GET /api/v1/status/{job_id}
```

### Attack Types

| Type | Description | Use Case |
|------|-------------|----------|
| `http` | HTTP flood attack | Web server stress testing |
| `syn` | TCP SYN flood | Network infrastructure testing |
| `udp` | UDP amplification | Bandwidth saturation tests |

## 🏗️ Project Structure

```
orchardnet-api/
├── core/
│   ├── attack_engine.go    # Core attack logic
│   └── scheduler.go         # Job scheduling system
├── modules/
│   ├── httpflood/          # HTTP attack module
│   ├── synflood/           # SYN flood module
│   └── udpamp/             # UDP amplification module
├── evasion/
│   ├── ipspoof/            # IP spoofing utilities
│   └── proxyrotator/       # Proxy rotation (unused)
├── utils/
│   ├── packet/             # Packet building utilities
│   └── logger.go           # Logging utilities
└── main.go                 # Entry point + dashboard
```

## 🔒 Security Features

- **API Key Authentication** - Prevents unauthorized access
- **IP Spoofing** - Geo-targeted source IP generation
- **Rate Limiting** - Built-in request throttling
- **Evasion Delays** - Timing randomization

## 🛠️ Configuration

### Change API Key
Edit `main.go`:
```go
if req.APIKey != "your-new-api-key-here" {
    http.Error(w, "Forbidden", http.StatusForbidden)
    return
}
```

### Change Port
Edit `main.go`:
```go
port := ":8080"  // Change to your desired port
```

### Adjust Worker Limits
Edit dashboard HTML in `main.go`:
```html
<input type="range" id="workers" min="10" max="1000" value="100" step="10">
```

## 🐛 Troubleshooting

### "Permission denied" errors
```bash
# Raw sockets require root privileges
sudo ./orchardnet-api
```

### "Cannot bind to port"
```bash
# Port 8080 is in use, kill the process:
sudo lsof -ti:8080 | xargs kill -9
```

### Browser doesn't open automatically
```bash
# Manually open:
# Linux: xdg-open http://localhost:8080
# macOS: open http://localhost:8080
# Windows: start http://localhost:8080
```

### Module errors
```bash
# Clean and rebuild dependencies
go clean -modcache
go mod tidy
go build
```

## 📊 Performance Tips

1. **Optimize Workers**: More workers ≠ better performance
   - Start with 100-200 workers
   - Monitor CPU/network usage
   - Adjust based on target capacity

2. **Use SYN Floods for Infrastructure**: Best for router/firewall testing

3. **Use HTTP Floods for Applications**: Best for web server testing

4. **Monitor Your Network**: High packet rates may trigger ISP throttling

## 🤝 Contributing

Contributions for educational improvements are welcome:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/improvement`)
3. Commit your changes (`git commit -am 'Add improvement'`)
4. Push to the branch (`git push origin feature/improvement`)
5. Open a Pull Request

## 📜 License

This project is provided for **EDUCATIONAL PURPOSES ONLY**.

**No warranty is provided. Use at your own risk.**

## 📧 Contact

For authorized security research inquiries only.

---

**Remember: With great power comes great responsibility. Use this tool ethically and legally.**

## 🙏 Acknowledgments

- Built with [Go](https://golang.org/)
- Packet manipulation via [golang.org/x/net](https://pkg.go.dev/golang.org/x/net)
- Dashboard styled with modern CSS

---

⚠️ **FINAL WARNING**: Unauthorized use of this software against systems you don't own or have permission to test is ILLEGAL and can result in severe criminal penalties including imprisonment and substantial fines.