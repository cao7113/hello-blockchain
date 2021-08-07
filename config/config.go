package config

import "fmt"

//ACCOUNT ADDRESS 0xe597a285BE0Bc901cd4116BE73AbCCE49d478e30
//PRIVATE KEY 67babe83ed57edcb12033ccf61def161efb0d082500a6a5c557d1e34f387b68c
const DevPk = "67babe83ed57edcb12033ccf61def161efb0d082500a6a5c557d1e34f387b68c"

const (
	HostIP = "127.0.0.1"
	Port = "7545"
)

var NodeURL = fmt.Sprintf("http://%s:%s", HostIP, Port)