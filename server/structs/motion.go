package structs

// Motion is the structure representing the motion table in the DB
type Motion struct {
	ID       int
	Motion   int
	Duration float64
	Start    float64
}
