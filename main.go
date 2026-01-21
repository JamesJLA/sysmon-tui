package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type model struct {
	tabs       []string
	activeTab  int
	cpuInfo    []cpu.InfoStat
	cpuPercent float64
	memInfo    *mem.VirtualMemoryStat
	diskInfo   []*disk.UsageStat
	netInfo    []net.IOCountersStat
	loadStat   *load.AvgStat
	hostInfo   *host.InfoStat
	processes  []*process.Process
	quitting   bool
}

type tabMsg struct {
	cpuInfo    []cpu.InfoStat
	cpuPercent float64
	memInfo    *mem.VirtualMemoryStat
	diskInfo   []*disk.UsageStat
	netInfo    []net.IOCountersStat
	loadStat   *load.AvgStat
	processes  []*process.Process
}

func initialModel() model {
	hostInfo, _ := host.Info()
	tabs := []string{"System", "CPU", "Memory", "Disk", "Network", "Processes"}

	return model{
		tabs:      tabs,
		activeTab: 0,
		hostInfo:  hostInfo,
		quitting:  false,
	}
}

func (m model) Init() tea.Cmd {
	return fetchData()
}

func fetchData() tea.Cmd {
	return func() tea.Msg {
		cpuInfo, _ := cpu.Info()
		cpuPercent, _ := cpu.Percent(time.Second, false)
		memInfo, _ := mem.VirtualMemory()
		diskInfo, _ := disk.Usage("/")
		netInfo, _ := net.IOCounters(true)
		loadStat, _ := load.Avg()
		processes, _ := process.Processes()

		return tabMsg{
			cpuInfo:    cpuInfo,
			cpuPercent: cpuPercent[0],
			memInfo:    memInfo,
			diskInfo:   []*disk.UsageStat{diskInfo},
			netInfo:    netInfo,
			loadStat:   loadStat,
			processes:  processes,
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "right", "l", "tab":
			m.activeTab = (m.activeTab + 1) % len(m.tabs)
			return m, fetchData()
		case "left", "h", "shift+tab":
			m.activeTab = (m.activeTab - 1 + len(m.tabs)) % len(m.tabs)
			return m, fetchData()
		case "r":
			return m, fetchData()
		}
	case tabMsg:
		m.cpuInfo = msg.cpuInfo
		m.cpuPercent = msg.cpuPercent
		m.memInfo = msg.memInfo
		m.diskInfo = msg.diskInfo
		m.netInfo = msg.netInfo
		m.loadStat = msg.loadStat
		m.processes = msg.processes
		return m, tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
			return fetchData()()
		})
	case tea.WindowSizeMsg:
		return m, nil
	}
	return m, nil
}

var (
	// Fixed width for consistent layout
	fixedWidth = 80

	// Consistent border styling
	mainBorder    = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 1).Width(fixedWidth)
	contentBorder = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 1).Width(fixedWidth - 4)

	// Tab styling
	activeTab   = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true).Background(lipgloss.Color("236"))
	tabInactive = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	// Color themes for different metrics
	headerStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	valueStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	cpuStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	memStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	diskStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
	netStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("33"))
	processStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("141"))
)

func (m model) View() string {
	if m.quitting {
		return "Thanks for using sysmon-tui!\n"
	}

	// Render tabs
	var tabs []string
	for i, t := range m.tabs {
		style := tabInactive
		if i == m.activeTab {
			style = activeTab
		}
		tabs = append(tabs, style.Render(t))
	}
	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)

	// Render content with consistent styling
	content := m.renderTabContent()

	return mainBorder.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			headerStyle.Render("System Monitor TUI"),
			"",
			tabRow,
			"",
			content,
			"",
			tabInactive.Render("Controls: ← → or Tab to switch | R to refresh | Q to quit"),
		),
	)
}

func (m model) renderTabContent() string {
	switch m.activeTab {
	case 0:
		return m.renderSystemTab()
	case 1:
		return m.renderCPUTab()
	case 2:
		return m.renderMemoryTab()
	case 3:
		return m.renderDiskTab()
	case 4:
		return m.renderNetworkTab()
	case 5:
		return m.renderProcessesTab()
	default:
		return ""
	}
}

func (m model) renderSystemTab() string {
	content := lipgloss.JoinVertical(lipgloss.Left,
		headerStyle.Render("System Information"),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Hostname: %s", m.hostInfo.Hostname)),
			"  ",
			valueStyle.Render(fmt.Sprintf("Platform: %s", m.hostInfo.Platform)),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("OS: %s %s", m.hostInfo.Platform, m.hostInfo.PlatformVersion)),
			"  ",
			valueStyle.Render(fmt.Sprintf("Uptime: %s", formatUptime(m.hostInfo.Uptime))),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			func() string {
				if m.loadStat != nil {
					return valueStyle.Render(fmt.Sprintf("Load Avg: %.2f %.2f %.2f", m.loadStat.Load1, m.loadStat.Load5, m.loadStat.Load15))
				}
				return valueStyle.Render("Load Avg: Loading...")
			}(),
		),
	)
	return contentBorder.Render(content)
}

func (m model) renderCPUTab() string {
	if len(m.cpuInfo) == 0 {
		return contentBorder.Render(valueStyle.Render("Loading CPU information..."))
	}

	cpu := m.cpuInfo[0]
	cpuDetails := lipgloss.JoinVertical(lipgloss.Left,
		headerStyle.Render("CPU Information"),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Model: %s", cpu.ModelName)),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Cores: %d", cpu.Cores)),
			"  ",
			valueStyle.Render(fmt.Sprintf("Mhz: %.0f", cpu.Mhz)),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Usage: %.1f%%", m.cpuPercent)),
			"  ",
			cpuStyle.Render(m.renderCPUBar(m.cpuPercent)),
		),
	)
	return contentBorder.Render(cpuDetails)
}

func (m model) renderMemoryTab() string {
	if m.memInfo == nil {
		return contentBorder.Render(valueStyle.Render("Loading memory information..."))
	}

	memContent := lipgloss.JoinVertical(lipgloss.Left,
		headerStyle.Render("Memory Information"),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Total: %s", formatBytes(m.memInfo.Total))),
			"  ",
			valueStyle.Render(fmt.Sprintf("Available: %s", formatBytes(m.memInfo.Available))),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Used: %s (%.1f%%)", formatBytes(m.memInfo.Used), m.memInfo.UsedPercent)),
		),
		"",
		memStyle.Render(m.renderMemBar(m.memInfo.UsedPercent)),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Cached: %s", formatBytes(m.memInfo.Cached))),
			"  ",
			valueStyle.Render(fmt.Sprintf("Buffers: %s", formatBytes(m.memInfo.Buffers))),
		),
	)
	return contentBorder.Render(memContent)
}

func (m model) renderDiskTab() string {
	if len(m.diskInfo) == 0 {
		return contentBorder.Render(valueStyle.Render("Loading disk information..."))
	}

	disk := m.diskInfo[0]
	diskContent := lipgloss.JoinVertical(lipgloss.Left,
		headerStyle.Render("Disk Information"),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Mountpoint: %s", disk.Path)),
			"  ",
			valueStyle.Render(fmt.Sprintf("Filesystem: %s", disk.Fstype)),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Total: %s", formatBytes(disk.Total))),
			"  ",
			valueStyle.Render(fmt.Sprintf("Free: %s", formatBytes(disk.Free))),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Used: %s (%.1f%%)", formatBytes(disk.Used), disk.UsedPercent)),
		),
		"",
		diskStyle.Render(m.renderDiskBar(disk.UsedPercent)),
	)
	return contentBorder.Render(diskContent)
}

func (m model) renderNetworkTab() string {
	if len(m.netInfo) == 0 {
		return contentBorder.Render(valueStyle.Render("Loading network information..."))
	}

	net := m.netInfo[0]
	netContent := lipgloss.JoinVertical(lipgloss.Left,
		headerStyle.Render("Network Information"),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Interface: %s", net.Name)),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Bytes Sent: %s", formatBytes(net.BytesSent))),
			"  ",
			valueStyle.Render(fmt.Sprintf("Bytes Recv: %s", formatBytes(net.BytesRecv))),
		),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top,
			valueStyle.Render(fmt.Sprintf("Packets Sent: %d", net.PacketsSent)),
			"  ",
			valueStyle.Render(fmt.Sprintf("Packets Recv: %d", net.PacketsRecv)),
		),
	)
	return contentBorder.Render(netContent)
}

func (m model) renderProcessesTab() string {
	if len(m.processes) == 0 {
		return contentBorder.Render(valueStyle.Render("Loading process information..."))
	}

	procContent := lipgloss.JoinVertical(lipgloss.Left,
		headerStyle.Render(fmt.Sprintf("Top %d Processes", len(m.processes))),
		"",
	)

	for i, proc := range m.processes {
		if i >= 10 { // Show top 10
			break
		}
		name, _ := proc.Name()
		pid := proc.Pid
		cpuPercent, _ := proc.CPUPercent()
		memInfo, _ := proc.MemoryInfo()

		procContent = lipgloss.JoinVertical(lipgloss.Left, procContent,
			lipgloss.JoinHorizontal(lipgloss.Top,
				processStyle.Render(fmt.Sprintf("%s [%d]", name, pid)),
				"  ",
				cpuStyle.Render(fmt.Sprintf("CPU: %.1f%%", cpuPercent)),
				"  ",
				memStyle.Render(fmt.Sprintf("Mem: %s", formatBytes(memInfo.RSS))),
			),
		)
	}
	return contentBorder.Render(procContent)
}

func (m model) renderCPUBar(percent float64) string {
	barWidth := 20
	filled := int(percent * float64(barWidth) / 100)
	bar := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}

func (m model) renderMemBar(percent float64) string {
	barWidth := 20
	filled := int(percent * float64(barWidth) / 100)
	bar := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}

func (m model) renderDiskBar(percent float64) string {
	barWidth := 20
	filled := int(percent * float64(barWidth) / 100)
	bar := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatUptime(seconds uint64) string {
	hours := int(seconds / 3600)
	minutes := int((seconds % 3600) / 60)
	return fmt.Sprintf("%dh %dm", hours, minutes)
}

func main() {
	program := tea.NewProgram(initialModel())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
