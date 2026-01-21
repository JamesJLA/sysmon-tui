# Sysmon TUI - System Monitor

A lightweight system monitoring tool with an interactive terminal interface.

```
╭─────────────────────────────────────────────────────────────────────────────────────╮
│                              System Monitor TUI                                     │
│                                                                                     │
│  System   CPU   Memory   Disk   Network   Processes                                 │
│                                                                                     │
│  ╭─────────────────────────────────────────────────────────────────────╮            │
│  │ System Information                                                  │            │
│  │ Hostname: omarchy            Platform: linux                        │            │
│  │ OS: linux 6.5.0-27-generic   Uptime: 11h 0m                         │            │
│  │ Load Avg: 0.86 1.30 1.55                                            │            │
│  ╰─────────────────────────────────────────────────────────────────────╯            │
│                                                                                     │
│  Controls: ← → or Tab to switch | R to refresh | Q to quit                          │
╰─────────────────────────────────────────────────────────────────────────────────────╯
```

```
╭─────────────────────────────────────────────────────────────────────────────────────╮
│                              System Monitor TUI                                     │
│                                                                                     │
│  System   CPU   Memory   Disk   Network   Processes                                 │
│                                                                                     │
│  ╭─────────────────────────────────────────────────────────────────────╮            │
│  │ CPU Information                                                     │            │
│  │ Model: AMD Ryzen 7 7730U with Radeon Graphics                       │            │
│  │ Cores: 8                    Mhz: 2800                               │            │
│  │ Usage: 12.3%                 ████████████████████████████░░░░░░░░   │            │
│  ╰─────────────────────────────────────────────────────────────────────╯            │
│                                                                                     │
│  Controls: ← → or Tab to switch | R to refresh | Q to quit                          │
╰─────────────────────────────────────────────────────────────────────────────────────╯
```

```
╭─────────────────────────────────────────────────────────────────────────────────────╮
│                              System Monitor TUI                                     │
│                                                                                     │
│  System   CPU   Memory   Disk   Network   Processes                                 │
│                                                                                     │
│  ╭─────────────────────────────────────────────────────────────────────╮            │
│  │ Top Processes                                                       │            │ 
│  │ chrome [1234]               CPU: 45.2%    Mem: 856.2 MB             │            │
│  │ firefox [5678]               CPU: 23.1%    Mem: 642.1 MB            │            │
│  │ vscode [9012]                CPU: 12.8%    Mem: 432.7 MB            │            │
│  │ node [3456]                  CPU: 8.4%     Mem: 128.3 MB            │            │
│  │ docker [7890]                CPU: 6.2%     Mem: 256.8 MB            │            │
│  ╰─────────────────────────────────────────────────────────────────────╯            │
│                                                                                     │
│  Controls: ← → or Tab to switch | R to refresh | Q to quit                          │
╰─────────────────────────────────────────────────────────────────────────────────────╯
```
=======

## Features

- Interactive TUI with tabbed interface
- Real-time monitoring updates every 2 seconds
- System information (hostname, OS, uptime, load)
- CPU usage, model, cores, and frequency
- Memory usage with visual progress bars
- Disk usage and filesystem information
- Network traffic statistics
- Top processes by CPU and memory usage

## Installation

### From Source
```bash
git clone https://github.com/yourusername/sysmon-tui.git
cd sysmon-tui
go mod init sysmon-tui
go mod tidy
go build -o sysmon-tui main.go
cp sysmon-tui ~/.local/bin/
```

### Go Install
```bash
go install github.com/yourusername/sysmon-tui@latest
```

## Usage

```bash
# Launch interactive TUI
sysmon-tui

# Navigation
← → or Tab: Switch between tabs
R: Refresh data
Q or Ctrl+C: Quit
```

## Supported Platforms

- Linux
- macOS  
- Windows (partial support)

## Dependencies

- Go 1.21+
- gopsutil library for system information

## Contributing

1. Fork repository
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
