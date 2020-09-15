package config

// Repository describes where we access data
type Repository interface {
	Name() string
	URL() string
	Version() string
	Provider() string
}

func Download(r *Repository, dest string) {
	// switch r.Provider() := p {
	// case "http":

	// case "git":

	// default:
	// 	log.Printf("Could not find provider %s \n", p)
	// }

	// if r.Provider() == "http" {
	// 	providers.HTTPDownloader()
	// }
}
