package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Addr   Address
}

type Employee struct {
	Name   string
	Salary float64
	Addr   Address
}

type Address struct {
	Street   string
	City     string
	State    string
	PostCode string
}
