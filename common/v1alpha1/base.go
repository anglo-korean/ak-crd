package v1alpha1

const (
	CRDVersion string = "v1alpha1"
)

// Base contains the options and keys all anko
// CRDs include, including logic for templating
// bits and bobs, along with defaults where necessary.
//
// Base is embedded into Cronjob, HTTP, and GRPC
type Base struct {
	Name            string      `json:"name"`
	Container       string      `json:"container"`
	Env             string      `json:"env"`
	Version         string      `json:"version"`
	Uniqueish       bool        `json:"uniqueish"`
	Singleton       bool        `json:"singleton"`
	Contact         string      `json:"contact"`
	Repo            string      `json:"repo"`
	ArbitraryConfig interface{} `json:"arbitrary_config"`
	ResourceFamily  string      `json:"resource_family"`

	// The following are optional values, if they don't exist/ are
	// null then we ignore them
	Storage     Storage `json:"storage,omitempty"`
	Preinstall  string  `json:"preinstall,omitempty"`
	Postinstall string  `json:"postinstall,omitempty"`
}

// Storage represents additional config necessary for accessing StorageClaims
type Storage struct {
	Claim      string `json:"claim"`
	MountPoint string `json:"mount_point"`
}
