package model

type Permission struct {
	Id             int64  `json:"id"`
	PermissionName string `json:"permission_name"`
	PAlliass       string `json:"permission_alias"`
	PDes           string `json:"permission_description"`
	PisActive      string `json:"is_active"`
	PCreatedBy     string `json:"created_by"`
	PUpdatedBy     string `json:"updated_by"`
}

//id	permission_name	permission_alias	permission_description	is_active	created_by	updated_by
