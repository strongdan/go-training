package main

import (
	"fmt"

	"log"

	"time"
)

type Person struct {
	FirstName string

	LastName string

	Age int
}

type FirePerson struct {
	Person

	NumFires int

	IsLit bool

	clockInTime time.Time
}

func (fp FirePerson) Job() string {

	return "fire person"

}

func (fp *FirePerson) ClockIn(t time.Time) error {

	if !fp.clockInTime.IsZero() {

		return fmt.Errorf("already working")

	}

	fp.clockInTime = t

	return nil

}

func (fp *FirePerson) ClockOut(t time.Time) error {

	if fp.clockInTime.IsZero() {

		return fmt.Errorf("not clocked in")

	}

	fmt.Printf("worked for %v\n", t.Sub(fp.clockInTime))

	fp.clockInTime = time.Time{}

	return nil

}

func (fp *FirePerson) DoWork() error {

	fp.NumFires++

	fp.IsLit = false

	return nil

}

type Worker interface {
	Job() string

	ClockIn(t time.Time) error

	DoWork() error

	ClockOut(t time.Time) error
}

func main() {

	var wk Worker = &FirePerson{Person: Person{FirstName: "bob"}}

	fmt.Printf("I'm a %v and I'm OK\n", wk.Job())

	if err := wk.ClockIn(time.Now()); err != nil {

		log.Fatal(err)

	}

	if err := wk.DoWork(); err != nil {

		log.Fatal(err)

	}

	if err := wk.ClockOut(time.Now().Add(time.Hour)); err != nil {

		log.Fatal(err)

	}

	firePerson := wk.(*FirePerson)

	fmt.Println(firePerson)

}

func (p Person) String() string {

	return fmt.Sprintf(

		"Person{Name: %v %v, Age: %v}",

		p.FirstName, p.LastName, p.Age)

}

func (p *Person) GrowOlder(years int) error {

	p.Age += years

	if p.Age > 99 {

		return fmt.Errorf("you ded")

	}

	return nil

}
