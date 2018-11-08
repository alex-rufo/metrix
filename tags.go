package metrix

import "fmt"

type Tags map[string]interface{}

func (t Tags) Add(tags Tags) {
	for name, value := range tags {
		t[name] = value
	}
}

func (t Tags) toStringMap() map[string]string {
	tags := map[string]string{}
	for name, value := range t {
		tags[name] = fmt.Sprintf("%v", value)
	}

	return tags
}
