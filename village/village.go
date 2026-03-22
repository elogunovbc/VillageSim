package village

import (
	"fmt"
	"strings"
)

type VillageElement interface {
	Update()
	FlushInfo() string
}

type Village struct {
	Elements []VillageElement
}

func (v *Village) AddElement(element VillageElement) {
	v.Elements = append(v.Elements, element)
}

func (v *Village) UpdateAll() {
	for _, element := range v.Elements {
		element.Update()
	}
}

func (v Village) ShowAllInfo() string {
	var builder strings.Builder
	for _, element := range v.Elements {
		fmt.Fprintf(&builder, "%s\n", element.FlushInfo())
	}
	return builder.String()
}
