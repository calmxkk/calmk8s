package k8sin

type Pod struct {
	BaseModel
	MetaData
	PodSpec
}

type PodSpec struct {
	Containers    []Container `json:"containers"`
	RestartPolicy string      `json:"restartPolicy"`
}

type GetPodInp struct {
	ClusterName string `json:"cluster_name"`
	NameSpace   string `json:"name_space"`
	PodName     string `json:"pod_name"`
}

type GetPodRes struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Status    string `json:"status"`
	Ready     string `json:"ready"`
	IP        string `json:"IP"`
	Node      string `json:"node"`
	Age       string `json:"age"`
}
