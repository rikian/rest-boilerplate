package common

type LogRequest struct {
	ID, Method, Path, Ip, Ua string
}

type RequestType struct {
	ID   string
	Code int
}

var RequestKey RequestType
