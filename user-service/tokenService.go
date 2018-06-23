package main

type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

type TokenService struct {
	repo Repository
}

func (s *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}

func (s *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}
