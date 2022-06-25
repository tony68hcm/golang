package reponi

import (
	"002/modules/permission/model"
	"fmt"
)

func (r *repoPermission) SelectAll() ([]model.Permission, error) {
	sql := "Select id, permission_name, permission_description, is_active from permission"
	all, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	permission := model.Permission{}
	listPermission := []model.Permission{}

	for all.Next() {
		err = all.Scan(&permission.Id, &permission.PermissionName, &permission.PDes, &permission.PisActive)
		if err != nil {
			return nil, err
		}
		fmt.Println(permission)
		listPermission = append(listPermission, permission)
	}

	return listPermission, nil
}
