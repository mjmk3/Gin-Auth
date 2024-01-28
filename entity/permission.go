package entity

type Permission struct {
	ID             int64  `json:"id"`
	PermissionName string `json:"permissionName"`
	Role           []Role
}
