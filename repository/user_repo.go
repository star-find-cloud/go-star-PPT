package repository

//type User struct {
//}
//
//func (u User) NameIsExists(name string) bool {
//	db := mysql.GetDB()
//	defer db.Close()
//
//	var i int
//
//	var d = logger.MySQLLogger{}
//	db_logger := d.GetLogger()
//
//	var err = db.Get(&i, "select 1 from business_test.user where id = ? limit 1", name)
//	if err != nil {
//		db_logger.Errorf("Sqlx.mysql get rows faild: %s", err)
//	}
//
//	if i == 1 {
//		return true
//	}
//
//	return false
//}
