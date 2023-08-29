package gormcmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuyan94zl/gotools/core/utils"
	"github.com/wuyan94zl/sql2gorm/parser"
	"os"
)

var VarStringSource string
var VarStringDir string
var VarTable string

type Command struct {
	wd        string
	nameSpace string
	dir       string

	tableName    string
	structName   string
	structData   string
	structImport string
}

func getAllTable(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var tables []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

func (c *Command) Run() error {
	c.wd, _ = os.Getwd()
	nameSpace, err := utils.GetPackage(c.wd)
	if err != nil {
		return err
	}
	c.nameSpace = nameSpace
	c.dir = VarStringDir + "/model"

	db, err := sql.Open("mysql", VarStringSource)

	if err != nil {
		panic(err)
	}
	tables, _ := getAllTable(db)
	for _, table := range tables {
		if VarTable != "" && VarTable != table {
			continue
		}
		rows, err := db.Query(fmt.Sprintf("SHOW CREATE TABLE %s", table))
		if err != nil {
			panic(err)
		}
		if rows.Next() {
			var t, s string
			err := rows.Scan(&t, &s)
			if err != nil {
				panic(err)
			}
			c.genModel(t, s)
		}
	}
	return nil
}

func (c *Command) genModel(tableName, code string) error {
	structData, err := parser.ParseSql(code, parser.WithNoNullType(), parser.WithGormType(), parser.WithIndex(), parser.WithJsonTag())
	if err != nil {
		return err
	}

	if len(structData.ImportPath) > 0 {
		for _, v := range structData.ImportPath {
			c.structImport = fmt.Sprintf("%s\"%s\"\n", c.structImport, v)
		}
	}

	c.structData = structData.StructCode[0].Code
	c.structName = structData.StructCode[0].Table
	c.tableName = tableName

	if err := c.setGormBaseModel(); err != nil {
		return err
	}

	if err := c.createTables(); err != nil {
		return err
	}

	if err := c.setGormCustomModel(); err != nil {
		return err
	}
	c.setMigrate()
	fmt.Println(c.tableName, "model gen Down.")
	return nil
}
