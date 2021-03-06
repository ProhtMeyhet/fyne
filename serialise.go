package fyne

import "bytes"
import "fmt"
import "io/ioutil"

func fromFile(name string) Resource {
	data, err := ioutil.ReadFile(cachePath(name))

	if err != nil {
		return nil
	}

	return NewStaticResource(name, data)
}

func toFile(res *StaticResource) {
	ioutil.WriteFile(cachePath(res.StaticName), res.StaticContent, 0644)
}

// ToGo converts a Resource object to Go code.
// This is useful if serialising to a go file for compilation into a binary
func (r *StaticResource) ToGo() string {
	var buffer bytes.Buffer

	buffer.WriteString("&fyne.StaticResource{\n")
	buffer.WriteString("\tStaticName: \"" + r.StaticName + "\",\n")
	buffer.WriteString("\tStaticContent: []byte{\n\t\t")
	for i, v := range r.StaticContent {
		if i > 0 {
			buffer.WriteString(", ")
		}

		buffer.WriteString(fmt.Sprint(v))
	}
	buffer.WriteString("}}")

	return buffer.String()
}
