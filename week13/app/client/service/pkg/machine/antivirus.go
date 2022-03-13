package machine

import "github.com/StackExchange/wmi"

type AntiVirusProduct struct {
	DisplayName              string
	InstanceGuid             string
	PathToSignedProductExe   string
	PathToSignedReportingExe string
	ProductState             int
}

func GetAntiVirusProducts() ([]*AntiVirusProduct, error) {
	var result []*AntiVirusProduct
	if err := wmi.QueryNamespace(
		"Select * from AntiVirusProduct", &result, "root\\SecurityCenter2"); err != nil {
		return nil, err
	}
	return result, nil
}
