package dao

import "modulename/basic_homework/homework_for_redrock/test/业务/model"

func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into messageboard.user(name, password) values (?,?)", u.UserName, u.Password)
	return
} //可在选择架构中把homework架构更改为message架构

func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from messageboard.user  where name=?")
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.UserName, &u.Password)
	return
}
