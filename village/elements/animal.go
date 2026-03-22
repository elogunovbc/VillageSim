package elements

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

func (a *Animal) AddEvent(str string, args ...any) {
	a.Events = append(a.Events, fmt.Sprintf(str, args...))
}

func (a *Animal) IncreaseAge() {
	a.Age++
}

func (a *Animal) Die() {
	if !a.Alive {
		return
	}

	a.AddEvent("Умер в %d лет. Какая досада.", a.Age)
	a.Alive = false
}

func (a *Animal) Update() {
	if !a.Alive {
		return
	}

	a.IncreaseAge()

	for i := 0; i < 12; i++ {
		if rand.IntN(1000) < 33 {
			a.AddEvent("Почесался.")
		}
		if rand.IntN(1000) < 15 {
			a.AddEvent("Покусал прохожего.")
		}
		if rand.IntN(1000) < 5 {
			a.AddEvent("Сломал лапу.")
		}
		if rand.IntN(1000) < a.Age {
			a.Die()
			return
		}
	}
}

func (a *Animal) FlushInfo() string {
	var builder strings.Builder

	var status string
	if a.Alive {
		status = "жив"
	} else {
		status = "мертв"
	}
	fmt.Fprintf(&builder, "Животное %s (%s, возраст: %d), статус: %s.\n", a.Name, a.Type, a.Age, status)

	fmt.Fprint(&builder, "События:")
	if len(a.Events) > 0 {
		fmt.Fprint(&builder, "\n")
		for _, event := range a.Events {
			fmt.Fprintf(&builder, "%s\n", event)
		}
	} else {
		fmt.Fprint(&builder, "Нет\n")
	}

	a.Events = []string{}
	return builder.String()
}
