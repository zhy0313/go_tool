/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 数据库功能
 *
 */
package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//插入数据
//Insert Data
func Insert(db *sql.DB, prestring string, parm []interface{}) (int64, error) {
	stmt, err := db.Prepare(prestring)
	if err != nil {
		//log.Println(err)
		return 0, err
	}
	R, err := stmt.Exec(parm...)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	num,err:=R.RowsAffected()
	return num,err

}

//打开数据库连接
//username:password@protocol(address)/dbname?param=value
func Open(username string, password string, ip string, dbname string) *sql.DB {
	db, err := sql.Open("mysql", username + ":" + password + "@tcp(" + ip + ":3306)/" + dbname + "?charset=utf8")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

//查询数据库
func Select(db *sql.DB, prestring string, parm []interface{}) (returnrows []map[string]interface{},err error) {
	returnrows = []map[string]interface{}{}
	returnrow := map[string]interface{}{}
	rows, err := db.Query(prestring, parm...)
	if err != nil {
		return
	}

	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			returnrow[columns[i]] = value
			//log.Println(columns[i], ": ", value)

		}
		returnrows = append(returnrows, returnrow)
		//log.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}