/*
 * Copyright (c) 2019.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Config struct {
	file    string
	version string

	timeout    int
	timeperiod int
}

func (config *Config) ParseCommandLine() {
	paramFilePtr := flag.String("file", "", "File path for ready check.")
	paramVersionPtr := flag.String("version", "", "Version information for check")

	paramTimeout := flag.Int("timeout", 200, "Timeout for waiting in seconds")
	paramTimeperiod := flag.Int("timeperiod", 20, "Time between checks in seconds")

	flag.Parse()

	config.file = *paramFilePtr
	config.version = *paramVersionPtr

	//time configuration
	config.timeout = *paramTimeout
	config.timeperiod = *paramTimeperiod

	if *paramFilePtr == "" {
		fmt.Fprintln(os.Stderr, "Parameter 'file' is empty.")
		flag.CommandLine.Usage()
		os.Exit(101)
	}

	if *paramTimeout <= *paramTimeperiod {
		fmt.Fprintln(os.Stderr, "Parameter timeperiod is bigger or equal to timeout. The timeout configuration must be bigger than the timeperiod.")
		flag.CommandLine.Usage()
		os.Exit(106)
	}
	if *paramTimeout <= 0 {
		fmt.Fprintln(os.Stderr, "Parameter timeout must be bigger than 0.")
		flag.CommandLine.Usage()
		os.Exit(107)
	}
	if *paramTimeperiod <= 0 {
		fmt.Fprintln(os.Stderr, "Parameter timeperiod must be bigger than 0.")
		flag.CommandLine.Usage()
		os.Exit(108)
	}
}

func main() {
	config := &Config{}
	config.ParseCommandLine()

	runTime := 0
	available := false

	for runTime < config.timeout {
		available = config.FileExists()

		if available && config.version != "" {
			readVersionBuf, err := ioutil.ReadFile(config.file)

			if err == nil {
				readVersion := string(readVersionBuf)

				if strings.TrimSpace(readVersion) == config.version {
					os.Exit(0)
				} else {
					log.Printf("Wait %d seconds for file content '%s' in '%s'. Will give up in %d seconds.", runTime, config.version, config.file, (config.timeout - runTime))
				}
			} else {
				log.Fatal(err)
			}
		}

		if available && config.version == "" {
			os.Exit(0)
		}

		if !available {
			log.Printf("Wait %d seconds for File '%s'. Will give up in %d seconds.", runTime, config.file, (config.timeout - runTime))
		}

		runTime += config.timeperiod
		time.Sleep(time.Duration(rand.Int31n(int32(config.timeperiod))) * time.Second)
	}

	if !available {
		log.Fatalf("No file '%s' found within %d seconds. Please check your configuration", config.file, runTime)
		os.Exit(10)
	} else {
		log.Fatalf("Content '%s' not found in file %s within %d seconds. Please check your configuration", config.file, config.version, runTime)
		os.Exit(20)
	}

}

func (config *Config) FileExists() bool {
	if _, err := os.Stat(config.file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
