package addon

// Provider is the implementation of addon sources
type Provider interface {
	Install(string)
	Update()
}
