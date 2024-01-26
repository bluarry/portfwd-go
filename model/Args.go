package model

//const (
//	Type = []string{"tcp", "udp"}
//)

type FwdArgs struct {
	Type           string
	SourceHostPort string
	DestHostPort   string
}
