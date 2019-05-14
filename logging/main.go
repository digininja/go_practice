package main

import "fmt"
import "log"

//import log "github.com/sirupsen/logrus"

func main() {
	fmt.Println("Begin...")
	log.SetLevel(log.DebugLevel)
	/*
		log.SetFormatter(&log.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	*/
	log.SetReportCaller(true)
	log.Printf("Print level message: %d", 1)
	log.Infof("Info level message: %d", 1)
	log.Warningf("Print level message: %d", 1)
	log.Debugf("Debug level message: %d", 1)

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")

	request_id := 32
	user_ip := "1.2.3.4"

	requestLogger := log.WithFields(log.Fields{"request_id": request_id, "user_ip": user_ip})
	// will log request_id and user_ip
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")

	event := "Get Certificate"
	requestID := 123

	// Calls os.Exit(1) after logging
	log.WithFields(log.Fields{
		"event":      event,
		"request ID": requestID,
	}).Fatal("Failed to send event")

	// Calls panic() after logging
	log.Panic("I'm bailing.")
}
