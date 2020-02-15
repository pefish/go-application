package go_application


type ApplicationClass struct {
	Debug bool
}

var Application = ApplicationClass{true}

var OnTerminated = make(chan interface{})
