package main

import "fmt"

func summary() (string, error) {
	spec, err := loadSpec()
	if err != nil {
		return "", err
	}

	out := ""

	if spec.Info != nil {
		out += fmt.Sprintf("Title: %s\n", spec.Info.Title)
		out += fmt.Sprintf("Version: %s\n", spec.Info.Version)
		if spec.Info.Description != "" {
			out += fmt.Sprintf("Description: %s\n", spec.Info.Description)
		}
		out += "\n"
	}

	if spec.Components != nil && len(spec.Components.Schemas) > 0 {
		out += "Schemas:\n"
		for name, schema := range spec.Components.Schemas {
			out += "- " + name
			if schema.Value != nil && schema.Value.Description != "" {
				out += ": " + schema.Value.Description
			}
			out += "\n"
		}
		out += "\n"
	}

	if spec.Paths != nil {
		out += "Endpoints:\n"
		for path, item := range spec.Paths.Map() {
			if item.Get != nil {
				out += "GET " + path
				if item.Get.Summary != "" {
					out += " - " + item.Get.Summary
				}
				out += "\n"
			}
			if item.Post != nil {
				out += "POST " + path
				if item.Post.Summary != "" {
					out += " - " + item.Post.Summary
				}
				out += "\n"
			}
			if item.Put != nil {
				out += "PUT " + path
				if item.Put.Summary != "" {
					out += " - " + item.Put.Summary
				}
				out += "\n"
			}
			if item.Delete != nil {
				out += "DELETE " + path
				if item.Delete.Summary != "" {
					out += " - " + item.Delete.Summary
				}
				out += "\n"
			}
			if item.Patch != nil {
				out += "PATCH " + path
				if item.Patch.Summary != "" {
					out += " - " + item.Patch.Summary
				}
				out += "\n"
			}
		}
	}

	return out, nil
}
