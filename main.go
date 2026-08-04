package main

import (
	"fmt"
)

type Log struct {
	Service string
	Level   string
	Message string
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
func FilterByService(logs []Log, service string) []Log {
	serviceFilter := []Log{}
	for _, log := range logs {
		if log.Service == service {
			serviceFilter = append(serviceFilter, log)
		}
	}
	return serviceFilter
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

}
