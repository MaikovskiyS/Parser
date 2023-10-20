package service

import (
	"bytes"
	"os"
)

type service struct {
	parser Parser
}

func New(p Parser) *service {

	return &service{
		parser: p,
	}
}
func (s *service) Parse() error {
	arr, err := s.parser.GetData()
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	for _, s := range arr {
		_, err := buf.WriteString("\n" + s)
		if err != nil {
			return err
		}
	}
	f, err := os.OpenFile("data.csv", 1, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(buf.Bytes())
	if err != nil {
		return err
	}
	//parsing rules
	return nil
}
