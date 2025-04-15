package pkg

type PodInterface interface {
	GetAllPodsInfo() ([]PodInfo, error) // 获取所有的pod信息
}

type PodInfo struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Status    string `json:"status"`
	Age       string `json:"age"`
}
