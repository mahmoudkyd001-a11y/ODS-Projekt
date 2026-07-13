package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var asyncSpecFile string

func findAsyncAPIFile() (string, error) {
	if asyncSpecFile != "" {
		if _, err := os.Stat(asyncSpecFile); err == nil {
			return asyncSpecFile, nil
		}
		return "", fmt.Errorf("specified file not found: %s", asyncSpecFile)
	}

	candidates := []string{
		"asyncapi.yaml",
		"AsyncAPI.yaml",
		"asyncapi.yml",
		"AsyncAPI.yml",
	}

	for _, file := range candidates {
		if _, err := os.Stat(file); err == nil {
			return file, nil
		}
	}

	return "", fmt.Errorf("no asyncapi file found, use -a <file> to specify one")
}

func fullAsyncSpec() (string, error) {
	file, err := findAsyncAPIFile()
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func asyncSummary() (string, error) {
	file, err := findAsyncAPIFile()
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	var doc map[string]interface{}
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return "", fmt.Errorf("failed to parse AsyncAPI spec: %w", err)
	}

	out := ""

	if info, ok := doc["info"].(map[string]interface{}); ok {
		if title, ok := info["title"].(string); ok {
			out += "Title: " + title + "\n"
		}
		if version, ok := info["version"].(string); ok {
			out += "Version: " + version + "\n"
		}
		if desc, ok := info["description"].(string); ok {
			out += "Description: " + desc + "\n"
		}
	}

	if servers, ok := doc["servers"].(map[string]interface{}); ok {
		out += "\nServers:\n"
		for name, srv := range servers {
			if srvMap, ok := srv.(map[string]interface{}); ok {
				protocol := ""
				if p, ok := srvMap["protocol"].(string); ok {
					protocol = p
				}
				host := ""
				if h, ok := srvMap["host"].(string); ok {
					host = h
				}
				out += fmt.Sprintf("- %s: %s (%s)\n", name, host, protocol)
			}
		}
	}

	if channels, ok := doc["channels"].(map[string]interface{}); ok {
		out += "\nChannels:\n"
		for name, ch := range channels {
			out += "- " + name
			if chMap, ok := ch.(map[string]interface{}); ok {
				if addr, ok := chMap["address"].(string); ok {
					out += fmt.Sprintf(" (address: %s)", addr)
				}
				// v2: inline subscribe/publish
				if _, ok := chMap["subscribe"]; ok {
					out += " [subscribe]"
				}
				if _, ok := chMap["publish"]; ok {
					out += " [publish]"
				}
			}
			out += "\n"
		}
	}

	// v3 operations
	if operations, ok := doc["operations"].(map[string]interface{}); ok {
		out += "\nOperations:\n"
		for name, op := range operations {
			if opMap, ok := op.(map[string]interface{}); ok {
				action := ""
				if a, ok := opMap["action"].(string); ok {
					action = a
				}
				summary := ""
				if s, ok := opMap["summary"].(string); ok {
					summary = s
				}
				out += fmt.Sprintf("- %s (%s): %s\n", name, action, summary)
			}
		}
	}

	if components, ok := doc["components"].(map[string]interface{}); ok {
		if messages, ok := components["messages"].(map[string]interface{}); ok {
			out += "\nMessages:\n"
			for name, msg := range messages {
				if msgMap, ok := msg.(map[string]interface{}); ok {
					summary := ""
					if s, ok := msgMap["summary"].(string); ok {
						summary = s
					}
					if title, ok := msgMap["title"].(string); ok {
						if summary != "" {
							summary = title + " - " + summary
						} else {
							summary = title
						}
					}
					out += fmt.Sprintf("- %s: %s\n", name, summary)
				}
			}
		}
	}

	return out, nil
}
