package k8sin

type Cluster struct {
	Name           string         `json:"name"`
	CaCertificate  Certificate    `json:"caCertificate"`
	PrivateKey     []byte         `json:"privateKey"`
	Status         Status         `json:"status"`
	Labels         []string       `json:"labels"`
	Namespace      string         `json:"namespace"`
	Authentication Authentication `json:"authentication"`
}

type Connect struct {
	Direction string  `json:"direction"`
	Forward   Forward `json:"forward"`
}

type Forward struct {
	ApiServer string `json:"apiServer"`
	Proxy     Proxy  `json:"proxy"  `
}

type Proxy struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authentication struct {
	Mode              string      `json:"mode"`
	BearerToken       string      `json:"bearerToken"`
	Certificate       Certificate `json:"certificate"`
	ConfigFileContent []byte      `json:"configFileContent"`
}

type Certificate struct {
	KeyData  []byte `json:"keyData"`
	CertData []byte `json:"certData"`
}

type Status struct {
	Version string `json:"version"`
	Phase   string `json:"phase"`
	Message string `json:"message"`
}
