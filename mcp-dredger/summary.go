package main

func summary() (string, error) {
	spec, err := loadSpec()
	if err != nil {
		return "", err
	}

	out := ""

	out += "Schemas:\n"

	for name := range spec.Components.Schemas {
		out += "- " + name + "\n"
	}

	out += "\nEndpoints:\n"

	for path, item := range spec.Paths.Map() {

		if item.Get != nil {
			out += "GET " + path + "\n"
		}

		if item.Post != nil {
			out += "POST " + path + "\n"
		}

		if item.Put != nil {
			out += "PUT " + path + "\n"
		}

		if item.Delete != nil {
			out += "DELETE " + path + "\n"
		}
	}

	return out, nil
}