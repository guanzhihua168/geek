package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

func main()  {
	err := Dao("select * from user where id=1")
	if err != nil {
		log.Fatal("get the error:"+err.Error())
	}

}

func Dao(query string) error {
	err := mockError()
	if err == sql.ErrNoRows {
		return errors.Wrapf(err, fmt.Sprintf("not found%s", query))
	} else if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("system error %s", query))
	}
	return nil
}

func mockError() error {
	return errors.New("mock error")
}