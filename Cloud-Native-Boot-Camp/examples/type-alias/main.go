package main

type ServiceType string

const (
	ServiceTypeClusterIPServiceType    = "ClusterIP"
	ServiceTypeNodePortServiceType     = "NodePort"
	serviceTypeLoadBalancerServiceType = "LoadBalancer"
	serviceTypeExternalNameServiceType = "ExternalName"
)

func main() {

}
