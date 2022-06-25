package reponi

import "fmt"

func (r *repoPermission) DeleteId(id int64) error {
	sql := "delete from permission where id = ?"

	effect, err := r.db.Exec(sql, id)
	if err != nil {
		return err
	}

	num, err := effect.RowsAffected()
	if err != nil {
		return err
	}

	if num == 0 {
		return fmt.Errorf("row effect 0")
	}
	return err
}
