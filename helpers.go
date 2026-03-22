package user_dashboard

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// GetOS returns the operating system name
func GetOS() string {
	return runtime.GOOS
}

// GetOSVersion returns the operating system version
func GetOSVersion() string {
	return runtime.Version()
}

// GetGoVersion returns the Go version
func GetGoVersion() string {
	return runtime.Version()
}

// GetCurrentTime returns the current time
func GetCurrentTime() time.Time {
	return time.Now()
}

// GetTimezone returns the current timezone
func GetTimezone() string {
	return time.Now().Location().String()
}

// GetMemoryUsage returns the memory usage in bytes
func GetMemoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Sys
}

// GetCPUCount returns the number of CPU cores
func GetCPUCount() int {
	return runtime.NumCPU()
}

// GetCPUUsage returns the current CPU usage
func GetCPUUsage() (int64, int64) {
	var (
		totalTime  int64
		idleTime   int64
		previousTime int64
	)

	start := time.Now()
	runtime.ReadMemStats(&totalTime)
	time.Sleep(100 * time.Millisecond)
	runtime.ReadMemStats(&idleTime)
	elapsed := time.Since(start).Nanoseconds() / 1e9
	previousTime = totalTime
	totalTime = idleTime
	return previousTime - totalTime, totalTime - idleTime
}

// GetHostname returns the hostname of the system
func GetHostname() string {
	return os.Hostname()
}

// GetUptime returns the system uptime in seconds
func GetUptime() float64 {
	u, _ := syscall.Uptime()
	return float64(u)
}

// GetProcesses returns the number of running processes
func GetProcesses() int {
	ps := syscall.Processes()
	return len(ps)
}

// GetOpenFiles returns the number of open files
func GetOpenFiles() int {
	files, _ := os.ReadDir("/")
	return len(files)
}

// GetSystemLoadAverage returns the system load average
func GetSystemLoadAverage() (int64, int64, int64) {
	load := os.Getloadavg()
	return load[0], load[1], load[2]
}

// GetDiskUsage returns the disk usage in bytes
func GetDiskUsage() (uint64, uint64, uint64) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err!= nil {
		return 0, 0, 0
	}
	return fs.Blocks * uint64(fs.Bsize), fs.Bavail * uint64(fs.Bsize), fs.Bfree * uint64(fs.Bsize)
}

// GetDiskUsagePercentage returns the disk usage percentage
func GetDiskUsagePercentage() float64 {
	_, used, free := GetDiskUsage()
	return float64(used) / float64(free) * 100
}

// GetDiskFreeSpace returns the free disk space in bytes
func GetDiskFreeSpace() uint64 {
	_, free, _ := GetDiskUsage()
	return free
}

// GetDiskTotalSpace returns the total disk space in bytes
func GetDiskTotalSpace() uint64 {
	_, _, total := GetDiskUsage()
	return total
}

// GetDiskUsedSpace returns the used disk space in bytes
func GetDiskUsedSpace() uint64 {
	_, used, _ := GetDiskUsage()
	return used
}

// GetDiskUsageString returns the disk usage as a string
func GetDiskUsageString() string {
	used, total, free := GetDiskUsage()
	return fmt.Sprintf("%d/%d (%.2f%%) free", free, total, float64(free)/float64(total)*100)
}

// GetNetworkInterfaces returns the network interfaces
func GetNetworkInterfaces() []string {
	ifaces, err := net.Interfaces()
	if err!= nil {
		return nil
	}
	var ifacesStr []string
	for _, i := range ifaces {
		ifacesStr = append(ifacesStr, i.Name)
	}
	return ifacesStr
}

// GetNetworkTraffic returns the network traffic in bytes
func GetNetworkTraffic() (uint64, uint64) {
	ifaces, err := net.Interfaces()
	if err!= nil {
		return 0, 0
	}
	var rx, tx uint64
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err!= nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil ||!ip.To4().HasAllZeros() {
				continue
			}
			stats, err := i.Statistics()
			if err!= nil {
				continue
			}
			rx += stats.RxBytes
			tx += stats.TxBytes
		}
	}
	return rx, tx
}

// GetNetworkTrafficString returns the network traffic as a string
func GetNetworkTrafficString() string {
	rx, tx := GetNetworkTraffic()
	return fmt.Sprintf("%d/%d (%.2f%%) rx, %d/%d (%.2f%%) tx", rx, GetDiskUsedSpace(), float64(rx)/float64(GetDiskUsedSpace())*100, tx, GetDiskUsedSpace(), float64(tx)/float64(GetDiskUsedSpace())*100)
}

// GetProcessList returns the process list
func GetProcessList() []string {
	ps := syscall.Processes()
	var psList []string
	for _, p := range ps {
		psList = append(psList, p.Name())
	}
	return psList
}

// GetProcessMemoryUsage returns the process memory usage in bytes
func GetProcessMemoryUsage(pid int) uint64 {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return 0
	}
	m, err := p.MemoryInfo()
	if err!= nil {
		return 0
	}
	return m.Rss
}

// GetProcessCPUTime returns the process CPU time
func GetProcessCPUTime(pid int) (int64, int64) {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return 0, 0
	}
	t := p.CPU()
	return t.User, t.System
}

// GetProcessStatus returns the process status
func GetProcessStatus(pid int) string {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return "Unknown"
	}
	s, err := p.Status()
	if err!= nil {
		return "Unknown"
	}
	return s.String()
}

// GetProcessOpenFiles returns the number of open files for a process
func GetProcessOpenFiles(pid int) int {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return 0
	}
	f, err := p.OpenFiles()
	if err!= nil {
		return 0
	}
	return len(f)
}

// GetProcessThreads returns the number of threads for a process
func GetProcessThreads(pid int) int {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return 0
	}
	t, err := p.NumThreads()
	if err!= nil {
		return 0
	}
	return t
}

// GetProcessName returns the process name
func GetProcessName(pid int) string {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return ""
	}
	return p.Name()
}

// GetProcessUsername returns the process username
func GetProcessUsername(pid int) string {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return ""
	}
	u, err := p.Username()
	if err!= nil {
		return ""
	}
	return u
}

// GetProcessCommand returns the process command
func GetProcessCommand(pid int) string {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return ""
	}
	c, err := p.Command()
	if err!= nil {
		return ""
	}
	return c
}

// GetProcessArguments returns the process arguments
func GetProcessArguments(pid int) []string {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return nil
	}
	a, err := p.Args()
	if err!= nil {
		return nil
	}
	return a
}

// GetProcessEnvironment returns the process environment
func GetProcessEnvironment(pid int) map[string]string {
	p, err := os.FindProcess(pid)
	if err!= nil {
		return nil
	}
	e, err := p.Environ()
	if err!= nil {
		return nil
	}
	env := make(map[string]string)
	for _, v := range e {
		kv := strings.SplitN(v, "=", 2)
		if len(kv) == 2 {
			env[kv[0]] = kv[1]
		}
	}
	return env
}