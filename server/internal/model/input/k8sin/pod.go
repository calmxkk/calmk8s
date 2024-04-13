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
