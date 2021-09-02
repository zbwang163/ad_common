package env

func IsDev() bool {
	return true
}

func GetPodIp() string {
	return "192.168.100.1"
}

func GetPodName() string {
	return "pod_123456"
}
