package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 普通用户
	PermissionVisitor     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 封号
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {

	switch s {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "普通用户"
	case PermissionVisitor:
		return "游客"
	case PermissionDisableUser:
		return "封禁用户"
	default:
		return "其他"
	}
}
