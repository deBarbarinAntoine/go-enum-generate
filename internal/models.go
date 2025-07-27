package internal

import "time"

type Enum struct {
	Date     time.Time `json:"-" yaml:"-"`
	Name     string    `json:"name" yaml:"name"`
	Plural   string    `json:"plural,omitempty" yaml:"plural,omitempty"`
	EnumType string    `json:"-" yaml:"-"`
	EnumVar  string    `json:"-" yaml:"-"`
	Values   []struct {
		Key   string `json:"key" yaml:"key"`
		Value string `json:"value,omitempty" yaml:"value,omitempty"`
	} `json:"values" yaml:"values"`
}
