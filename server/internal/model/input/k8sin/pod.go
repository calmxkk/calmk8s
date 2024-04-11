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
