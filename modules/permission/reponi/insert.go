package reponi 

import "002/modules/permission/model"

func (rPermission *repoPermission) InsertPermission(req model.Permission) error {
	stmt := `insert into permission(permission_name, permission_alias, permission_description, is_active, created_by, updated_by)
	values (?, ?, ?, ?, ?, ?)`
	result, err := rPermission.db.Exec(stmt, req.PermissionName, req.PAlliass, req.PDes, req.PisActive, req.PCreatedBy, req.PUpdatedBy)
	if err != nil {
		return err
	}
	if num, err := result.RowsAffected(); num == 0 || err != nil {
		return err
	}
	return nil
}
