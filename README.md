# Sysmon - System Monitor CLI

A lightweight, fast system monitoring CLI tool written in Go. Provides comprehensive system information in multiple output formats.

## Features

- 🖥️ System information (hostname, OS, uptime, load)
- 💻 CPU usage, model, and core information  
- 🧠 Memory usage statistics
- 💾 Disk usage and available space
- 🌐 Network traffic statistics
- 📋 Top processes by CPU usage
- 📊 Multiple output formats (text, JSON)
- 👀 Watch mode for continuous monitoring

## Installation

### From Source

```bash
git clone https://github.com/yourusername/sysmon.git
cd sysmon
go build -o sysmon ./cmd/sysmon
cp sysmon ~/.local/bin/
```

### Go Install

```bash
go install github.com/yourusername/sysmon@latest
```

## Usage

```bash
# Show system information
sysmon

# Output as JSON
sysmon json

# Continuous monitoring (updates every 2 seconds)
sysmon --watch

# Show help
sysmon --help
```

## Examples

```bash
# Quick system overview
sysmon

# Use in scripts
sysmon json | jq '.cpu.usage'

# Monitor system continuously  
sysmon --watch

# Get specific info
sysmon json | jq '.memory.used_percent'
```

## Output Examples

### Default Output
```
🖥️  System Monitor
===================
Hostname: omarchy | arch | Uptime: 11h 0m
Load: 0.86 1.30 1.55

💻 CPU: AMD Ryzen 7 7730U (8 cores) @ 2800 MHz
Usage: 12.3%

🧠 Memory: 15.0 GB total, 4.2 GB used (28.0%)
Available: 10.8 GB | Cached: 2.1 GB

💾 Disk (ext4): 476.8 GB total, 125.3 GB used (26.3%)
Free: 351.5 GB

🌐 Network (wlan0): Sent: 1.2 GB, Recv: 3.4 GB
📋 Top Processes
PID    NAME                 CPU%    MEMORY
1234   chrome              45.2    856.2 MB
5678   firefox             23.1    642.1 MB
```

### JSON Output
```json
{
  "host": {
    "hostname": "omarchy",
    "platform": "arch",
    "platform_version": "",
    "uptime": 39664
  },
  "cpu": {
    "model": "AMD Ryzen 7 7730U",
    "cores": 8,
    "mhz": 2800,
    "usage": 12.3
  },
  "memory": {
    "total": 16106127360,
    "available": 11652746752,
    "used": 4517800640,
    "used_percent": 28.0
  }
}
```

## Supported Platforms

- Linux
- macOS  
- Windows (partial support)

## Dependencies

- Go 1.21+
- gopsutil library for system information

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Similar Projects

- [htop](https://github.com/htop-dev/htop) - Interactive process viewer
- [glances](https://github.com/nicolargo/glances) - System monitoring tool
- [gotop](https://github.com/xxxserxxx/gotop) - Terminal based graphical activity monitor