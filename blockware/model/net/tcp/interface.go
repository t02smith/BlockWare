package tcp

/*
An abstraction over any maintained connection with another node.
Can be managed under the TCPServer or as one of the TCPClients.
*/
type TCPConnection interface {
	Send(command []byte) error
	SendString(command string) error
	SendStringf(command string, args ...any) error
	Info() string
}
