/*
 * Grabber main program.
 */
package main

import (
	"flag"
	"fmt"
	// "os"
)

func main() {
	var log_filename string
	var server_url string
	var api_id string
	var api_key string
	var config_filename string
	var daemonize bool
	flag.StringVar(&log_filename, "l", "", "Log filename. Default is empty, and log to stdout.")
	flag.StringVar(&server_url, "u", "http://localhost:8938/", "Server URL for report interface.")
	flag.StringVar(&api_id, "id", "", "API ID for server API access.")
	flag.StringVar(&api_key, "key", "", "API Key for server API access.")
	flag.StringVar(&config_filename, "c", "", "Configration file.")
	flag.Bool("d", false, "Run in deamon mode.")
	flag.Parse()

	if api_id == "" || api_key == "" || server_url == "" {
		flag.PrintDefaults()
	}

	if config_filename != "" {
		fmt.Printf("Using configuration %s ...", config_filename)
	}

	if log_filename != "" {
		fmt.Printf("Logging into %s ...", log_filename)
	}

}
