package services

// ProtocolService Interface to manage the protocol
type ProtocolService interface {
	NewProtocol() (string, error)
}
