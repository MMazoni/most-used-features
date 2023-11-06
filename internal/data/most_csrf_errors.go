package data

type MostCsrfErrors struct {
	Path       string
	Method     string
	Env        string
	Controller string
	Action     string
	Error      int
}
