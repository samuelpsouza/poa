package domain

type AuthenticationTyper interface {

}

type Authentication struct {
	Id int
	Key string
	Method AuthenticationTyper
}

type Request struct {
	Id int
	Body string
	Headers []string
	Method string
}

type Response struct {
	Id int
	Key string
	Body string
	StatusCode int
}

type Endpoint struct {
	Id int
	Key string
	Auth AuthenticationTyper
	DefaultResponse Response 
	Requests []Request
	Origin string
	Received string
}
