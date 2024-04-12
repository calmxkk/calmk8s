package k8sin

import "github.com/gogf/gf/v2/frame/g"

type YamlByte struct {
	data []byte
}

type BaseModel struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type MetaData struct {
	Name   string      `json:"name"`
	Labels g.MapStrStr `json:"lables"`
}

type Container struct {
	Name            string      `json:"name"`
	Image           string      `json:"image"`
	ImagePullPolicy string      `json:"imagePullPolicy"`
	Ports           []Port      `json:"ports"`
	Resources       Resources   `json:"resources"`
	Envs            []Env       `json:"envs"`
	VolumeMounts    g.MapStrStr `json:"env"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Port struct {
	ContainerPort int `json:"containerPort"`
}

type Resources struct {
	Requests Requests `json:"requestes"`
	Limits   Limits   `json:"limits"`
}

type Requests struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

type Limits struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}
