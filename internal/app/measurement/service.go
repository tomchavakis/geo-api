package measurement

// Service ...
type Service interface {
	GetDistance() (*float64, error)
}
