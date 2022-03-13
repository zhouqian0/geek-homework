package dto

type Cert struct {
	Cert []byte `json:"cert"`
	Key  []byte `json:"key"`
}

// Host 机房信息
type Host struct {
	Name    string `json:"name"`    // 名称
	Manager string `json:"manager"` // 管理员
	Phone   string `json:"phone"`   // 联系方式
}
