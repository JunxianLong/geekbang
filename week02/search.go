import (
	"database/sql"
	"errors"
	pkgErrors "github.com/pkg/errors"
)


// GetUserNameByID 根据ID获取用户名
func (*seqDao) GetUserNameByID(userID int)(string ,error) {
	tx := DB.Begin()
	defer tx.Rollback()

	var userName string
	err := DB.Raw("SELECT user_name FROM user_info WHERE user_id = ?", userID).
		Row().Scan(&userName)
	if !errors.Is(err,sql.ErrNoRows){
		return "",pkgErrors.Wrap(err,"query error")
	}
	/*
		do something
	 */
	tx.Commit()
	return "",nil
}
