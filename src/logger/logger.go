package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	logChannel = make(chan string, 100)
	wg         sync.WaitGroup
	logFile    *os.File
)

func init() {
	var err error
	// Open log file in append mode, creating it if it doesn't exist
	logFile, err = os.OpenFile("./logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Start the log writer goroutine
	wg.Add(1)
	go logWriter()
}

// Log formats a log message and sends it to the logChannel
func Log(format string, v ...any) {
	timestamp := time.Now().Format("2006-01-02 15:04:05") // Format timestamp
	message := fmt.Sprintf("[%s] %s", timestamp, fmt.Sprintf(format, v...))
	select {
	case logChannel <- message:
	default:
		// Fall back to standard logger if channel is full
		log.Println("Log channel is full, dropping log message:", message)
	}
	log.Println(message)
}

// logWriter reads from the logChannel and writes to the log file
func logWriter() {
	defer wg.Done()
	for logMessage := range logChannel {
		_, err := logFile.WriteString(logMessage + "\n")
		if err != nil {
			log.Printf("Failed to write log message: %s", err)
		}
	}
}

// Close should be called to flush logs and close the log file
func Close() {
	// Close the log channel and wait for the log writer to finish
	close(logChannel)
	wg.Wait()

	// Close the log file
	if err := logFile.Close(); err != nil {
		log.Printf("Failed to close log file: %s", err)
	}
}
