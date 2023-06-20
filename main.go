package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
	"github.com/sirupsen/logrus"
)

func main() {
	opsGenieApiKey := flag.String("api-key", os.Getenv("OPSGENIE_API_KEY"), "OpsGenie API Key, or specify via the OPSGENIE_API_KEY environment variable")
	opsGenieHeartbeat := flag.String("heartbeat", os.Getenv("OPSGENIE_HEARTBEAT"), "OpsGenie Heartbeat Name, or OPSGENIE_HEARTBEAT")
	flag.Parse()

	if opsGenieApiKey == nil || opsGenieHeartbeat == nil ||
		*opsGenieApiKey == "" || *opsGenieHeartbeat == "" {
		flag.Usage()
		log.Fatal("OpsGenie API key and heartbeat name are required")
	}

	// create opsgenie client
	heartbeatClient, err := heartbeat.NewClient(&client.Config{
		ApiKey:   *opsGenieApiKey,
		LogLevel: logrus.WarnLevel,
	})
	if err != nil {
		log.Fatalf("could not create OpsGenie client: %+v", err)
	}

	// ping
	_, err = heartbeatClient.Ping(context.Background(), *opsGenieHeartbeat)
	if err != nil {
		log.Fatalf("could not ping heartbeat: %+v", err)
	}
}
