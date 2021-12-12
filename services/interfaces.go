package services

type ProtocolService interface {
	NewProtocol() (string, error)
}
