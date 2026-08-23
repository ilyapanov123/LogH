package main

import (
	"fmt"
)

type Log struct {
	Service string
	Level   string
	Message string
}

type ServiceStats struct {
	Total   int
	Error   int
	Warning int
	Info    int
}

func (stats *ServiceStats) Update(level string) {
	stats.Total++

	if level == "ERROR" {
		stats.Error++
	}
	if level == "WARNING" {
		stats.Warning++
	}
	if level == "INFO" {
		stats.Info++
	}
}

func FilterByLevel(logs []Log, level string) []Log {
	var filteredLogs []Log

	for _, log := range logs {
		if log.Level == level {
			filteredLogs = append(filteredLogs, log)
		}
	}

	return filteredLogs
}

func PrintLogs(logs []Log) {
	for _, log := range logs {
		fmt.Println(
			log.Level,
			log.Service,
			log.Message,
		)
	}
}

func PrintStats(stats map[string]ServiceStats) {
	for service, stat := range stats {
		fmt.Println(service)
		fmt.Println("All:", stat.Total)
		fmt.Println("Error:", stat.Error)
		fmt.Println("Warning:", stat.Warning)
		fmt.Println("Info:", stat.Info)
	}
}
func FilterByService(logs []Log, service string) []Log {
	serviceFilter := []Log{}
	for _, log := range logs {
		if log.Service == service {
			serviceFilter = append(serviceFilter, log)
		}
	}
	return serviceFilter
}

func CountLogsByService(logs []Log) map[string]int {
	counts := make(map[string]int)
	for _, log := range logs {
		counts[log.Service]++
	}
	return counts
}

func BuildServiceStats(logs []Log) map[string]ServiceStats {
	stats := make(map[string]ServiceStats)
	for _, log := range logs {
		currentStats := stats[log.Service]

		currentStats.Update(log.Level)

		stats[log.Service] = currentStats

	}
	return stats
}

func (stats ServiceStats) HasErrors() bool {
	if stats.Error > 0 {
		return true
	} else {
		return false
	}
}

func (stats ServiceStats) ErrorPercent() float64 {
	if stats.Total == 0 {
		return 0
	}
	return float64(stats.Error) / float64(stats.Total) * 100
}
func (stats ServiceStats) IsProblematic() bool {
	if stats.ErrorPercent() > 30 {
		return true
	} else {
		return false
	}
}
func main() {
	logs := []Log{
		{
			Service: "Minecraft",
			Level:   "INFO",
			Message: "Server started successfully",
		},
		{
			Service: "Minecraft",
			Level:   "WARNING",
			Message: "Memory usage is high",
		},
		{
			Service: "Minecraft",
			Level:   "ERROR",
			Message: "Server crashed",
		},
		{
			Service: "Auth",
			Level:   "ERROR",
			Message: "Invalid access token",
		},
	}

	errorLogs := FilterByLevel(logs, "ERROR")
	warningLogs := FilterByLevel(logs, "WARNING")

	counters := CountLogsByService(logs)
	for service, count := range counters {
		fmt.Println(service, "->", count)
	}

	if len(errorLogs) == 0 {
		fmt.Println("Логи ERROR не найдены")
	} else {
		fmt.Println("Найденные логи ERROR:")
		PrintLogs(errorLogs)
	}

	if len(warningLogs) == 0 {
		fmt.Println("Логи WARNING не найдены")
	} else {
		fmt.Println("Найденные логи WARNING:")
		PrintLogs(warningLogs)
	}

	minecraftLogs := FilterByService(logs, "Minecraft")
	fmt.Println("Найденные логи по SERVICE:")
	PrintLogs(minecraftLogs)

	buildService := BuildServiceStats(logs)
	PrintStats(buildService)

	for service, stats := range buildService {
		fmt.Println(service)
		fmt.Println(stats.HasErrors())
		fmt.Println(stats.IsProblematic())
		fmt.Println(stats.ErrorPercent())
	}

}
