package service

type Parser interface {
	GetData() ([]string, error)
}
