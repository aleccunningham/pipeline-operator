package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Agent is a duke resource defining a CI lifecycle
type AutoScaler struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the ddesired behaviour of the pod terminator.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec AutoScalerSpec `json:"spec"`
}

// PipelineSpec is the spec for a Agent resource
type AutoScalerSpec struct {
	// Selector is how the target will be selected
	selector map[string]string `json:"selector,omitempty"`
	// sourceRepository is the location of a Deployments code
	config Config `json:"config,omitempty"`
}

// Config stores the configuration settings.
type Config struct {
	License   string        `json:"license,omitempty"`
	Interval  time.Duration `default:"5m" json:"interval,omitempty"`
	slack     Slack
	logs      Logs
	pool      Pool
	server    Server
	agent     Agent
	http      HTTP
	tls       TLS
	promtheus Prometheus
	database  Database
}

type Slack struct {
	Webhook string `json:"webhook,omitempty"`
}

type Logs struct {
	Color  bool `json:"color,omitempty"`
	Debug  bool `default:"true" json:"debug,omitempty"`
	Pretty bool `json:"pretty,omitempty"`
}

type Pool struct {
	Min    int           `default:"2" json:"min,omitempty"`
	Max    int           `default:"4" json:"max,omitempty"`
	MinAge time.Duration `default:"55m" split_words:"true" json:"minAge,omitempty"`
}

type Server struct {
	Host  string `json:"host,omitempty"`
	Proto string `json:"proto,omitempty"`
	Token string `json:"token,omitempty"`
}

type Agent struct {
	Host        string `json:"host,omitempty"`
	Token       string `json:"token,omitempty"`
	Image       string `default:"drone/agent:0.8" json:"image,omitempty"`
	Concurrency int    `default:"2" json:"concurrency,omitempty"`
}

type HTTP struct {
	Host string `json:"host,omitempty"`
	Port string `default:":8080" json:"port,omitempty"`
	Root string `default:"/" json:"root,omitempty"`
}

type TLS struct {
	Autocert bool   `json:"autocert,omitempty"`
	Cert     string `json:"cert,omitempty"`
	Key      string `json:"key,omitempty"`
}

type Prometheus struct {
	Token string `json:"token,omitempty"`
}

type Database struct {
	Driver     string `default:"sqlite3" json:"driver,omitempty"`
	Datasource string `default:"database.sqlite?cache=shared&mode=rwc&_busy_timeout=9999999" json:"datasource,omitempty"`
}

/*
	Amazon struct {
		Image         string            `json:"image,omitempty"`
		Instance      string            `json:"instance,omitempty"`
		PrivateIP     bool              `split_words:"true" json:"privateIP,omitempty"`
		Region        string            `json:"region,omitempty"`
		Retries       int               `json:"retries,omitempty"`
		SSHKey        string            `json:"SSHKey,omitempty"`
		SubnetID      string            `split_words:"true" json:"subnetID,omitempty"`
		SecurityGroup []string          `split_words:"true" json:"securityGroup,omitempty"`
		Tags          map[string]string `json:"tags,omitempty"`
		UserData      string            `envconfig:"DRONE_AMAZON_USERDATA" json:"userData,omitempty"`
		UserDataFile  string            `envconfig:"DRONE_AMAZON_USERDATA_FILE" json:"userDataFile,omitempty"`
	}

		DigitalOcean struct {
			Token        string   `json:"token,omitempty"`
			Image        string   `json:"image,omitempty"`
			Region       string   `json:"region,omitempty"`
			SSHKey       string   `json:"sshKey,omitempty"`
			Size         string   `json:"size,omitempty"`
			Tags         []string `json:"tags,omitempty"`
			UserData     string   `envconfig:"DRONE_DIGITALOCEAN_USERDATA" json:"userData,omitempty"`
			UserDataFile string   `envconfig:"DRONE_DIGITALOCEAN_USERDATA_FILE" json:"userDataFile,omitempty"`
		}

		Google struct {
			MachineType  string            `envconfig:"DRONE_GOOGLE_MACHINE_TYPE"`
			MachineImage string            `envconfig:"DRONE_GOOGLE_MACHINE_IMAGE"`
			Network      string            `envconfig:"DRONE_GOOGLE_NETWORK"`
			Labels       map[string]string `envconfig:"DRONE_GOOGLE_LABELS"`
			Scopes       string            `envconfig:"DRONE_GOOGLE_SCOPES"`
			DiskSize     int64             `envconfig:"DRONE_GOOGLE_DISK_SIZE"`
			DiskType     string            `envconfig:"DRONE_GOOGLE_DISK_TYPE"`
			Project      string            `envconfig:"DRONE_GOOGLE_PROJECT"`
			Tags         []string          `envconfig:"DRONE_GOOGLE_TAGS"`
			UserData     string            `envconfig:"DRONE_GOOGLE_USERDATA"`
			UserDataFile string            `envconfig:"DRONE_GOOGLE_USERDATA_FILE"`
			Zone         string            `envconfig:"DRONE_GOOGLE_ZONE"`
		}

		HetznerCloud struct {
			Datacenter   string `json:"datacenter,omitempty"`
			Image        string `json:"image,omitempty"`
			SSHKey       int    `json:"sshKey,omitempty"`
			Token        string `json:"token,omitempty"`
			Type         string `json:"type,omitempty"`
			UserData     string `envconfig:"DRONE_HETZNERCLOUD_USERDATA" json:"userData,omitempty"`
			UserDataFile string `envconfig:"DRONE_HETZNERCLOUD_USERDATA_FILE" json:"userDataFile,omitempty"`
		}
*/
