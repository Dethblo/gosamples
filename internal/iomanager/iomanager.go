package iomanager

// IOManager a general interface for reading/writing the data used in the application
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
