package ent

import (
	"context"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

func (c *Client) Driver() dialect.Driver {
	return c.driver
}

/*
usage

type Row struct {
	Id       int
	Name     string
	NickName string
}

func hoge() {
	var args []interface{}
	var rows []Row

	sql := "SELECT users.id, users.name, pets.name as nickName FROM users INNER JOIN pets on users.id = pets.user_pets"
	if err := client.DB.RawQuery(ctx, sql, args, &rows); err != nil {
		return err
	}

	fmt.Println(rows)
}
*/
func (c *Client) RawQuery(ctx context.Context, query string, args, mapResult interface{}) error {
	rows := &sql.Rows{}
	if err := c.Driver().Query(ctx, query, args, rows); err != nil {
		return err
	}

	defer rows.Close()
	if err := sql.ScanSlice(rows, mapResult); err != nil {
		return err
	}

	return rows.Err()
}

/*
usage

func hoge() {
	var args []interface{}
	args = append(args, "hoge")
	result, err := client.DB.RawExec(ctx, "INSERT INTO users (name, created_at) VALUES (?, NOW())", args)
	if err != nil {
		return nil, err
	}
}
*/
func (c *Client) RawExec(ctx context.Context, query string, args interface{}) (sql.Result, error) {
	var result sql.Result
	if err := c.Driver().Exec(ctx, query, args, &result); err != nil {
		return nil, err
	}

	return result, nil
}
