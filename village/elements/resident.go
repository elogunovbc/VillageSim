package elements

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

func (r *Resident) AddEvent(str string, args ...any) {
	r.Events = append(r.Events, fmt.Sprintf(str, args...))
}

func (r *Resident) IncreaseAge() {
	r.Age++
}

func (r *Resident) Die() {
	if !r.Alive {
		return
	}

	r.AddEvent("Умер в %d лет. Какая досада.", r.Age)
	r.Alive = false
}

func (r *Resident) ChangeMaritalStatus(married bool) {
	if r.Married == married {
		return
	}

	switch {
	case !r.Married && married:
		r.AddEvent("Наконец-то, найден спутник в жизни!!!")
	case r.Married && !married:
		r.AddEvent("Развод, больше я не в браке.")
	}

	r.Married = married
}

func (r *Resident) Update() {
	if !r.Alive {
		return
	}

	r.IncreaseAge()

	for i := 0; i < 10; i++ {
		if rand.IntN(10000) < 1000 && !r.Married {
			r.ChangeMaritalStatus(true)
		}
		if rand.IntN(10000) < 250 {
			r.AddEvent("Устроился на новую работу.")
		}
		if rand.IntN(10000) < 100 && r.Married {
			r.AddEvent("Поругался с супругой/ом.")
		}
		if rand.IntN(10000) < 50 && r.Married {
			r.ChangeMaritalStatus(false)
		}
		if rand.IntN(10000) < r.Age {
			r.Die()
			return
		}
	}

}

func (r *Resident) FlushInfo() string {
	var builder strings.Builder

	var status string
	if r.Alive {
		if r.Married {
			status = "в браке"
		} else {
			status = "холост"
		}
	} else {
		status = "мертв"
	}
	fmt.Fprintf(&builder, "Житель %s (возраст: %d), статус: %s.\n", r.Name, r.Age, status)

	fmt.Fprint(&builder, "События:")
	if len(r.Events) > 0 {
		fmt.Fprint(&builder, "\n")
		for _, event := range r.Events {
			fmt.Fprintf(&builder, "%s\n", event)
		}
	} else {
		fmt.Fprint(&builder, "Нет\n")
	}

	r.Events = []string{}
	return builder.String()
}
