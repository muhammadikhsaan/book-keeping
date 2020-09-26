package services

import (
	"strings"
)

// type QueryResultInterface interface {
// 	Get(string) queryResult
// 	Post(string) queryResult
// 	Select(string, interface{}) *BuilderMapping
// 	Insert(string, interface{}) *BuilderMapping
// 	Where(string, string) *BuilderMapping
//	Join(string, string) *BuilderMapping
// }

type BuilderMapping struct {
	querySelect  []string
	objectSelect []interface{}

	queryInsert []string
	valueInsert []interface{}

	queryUpdate []string
	valueUpdate []interface{}

	queryWhere []string
	valueWhere []interface{}

	queryJoin []string
}

type queryResult struct {
	Query  string
	Value  []interface{}
	Object []interface{}
}

func (b *BuilderMapping) Get(tabel string) queryResult {
	var query string
	var value []interface{}

	query = "select " + strings.Join(b.querySelect, ", ") + " from " + tabel + " "

	if len(b.queryJoin) != 0 {
		query += strings.Join(b.queryJoin, " ")
	}

	if len(b.queryWhere) != 0 {
		query += " where " + strings.Join(b.queryWhere, " and ")
	}

	query += ";"

	for _, v := range b.valueWhere {
		value = append(value, v)
	}

	return queryResult{
		Query:  query,
		Value:  value,
		Object: b.objectSelect,
	}
}

func (b *BuilderMapping) Post(tabel string) queryResult {
	var query string
	var exec []string

	for range b.queryInsert {
		exec = append(exec, "?")
	}

	query += "insert into " + tabel + "(" + strings.Join(b.queryInsert, ", ") + ") values (" + strings.Join(exec, ",") + ");"

	return queryResult{
		Query: query,
		Value: b.valueInsert,
	}
}

func (b *BuilderMapping) Put(tabel string) queryResult {
	var query string
	var value []interface{}
	query = "update " + tabel

	if len(b.queryUpdate) != 0 {
		query += " set " + strings.Join(b.queryUpdate, ", ")
	}

	if len(b.queryWhere) != 0 {
		query += " where " + strings.Join(b.queryWhere, ", ")
	}

	query += ";"

	for _, v := range b.valueUpdate {
		value = append(value, v)
	}

	for _, v := range b.valueWhere {
		value = append(value, v)
	}

	return queryResult{
		Query: query,
		Value: value,
	}
}

func (b *BuilderMapping) Delete(tabel string) queryResult {
	var query string
	var value []interface{}

	query = "delete from " + tabel

	if len(b.queryWhere) != 0 {
		query += " where " + strings.Join(b.queryWhere, ", ")
	}

	for _, v := range b.valueWhere {
		value = append(value, v)
	}

	return queryResult{
		Query: query,
		Value: value,
	}
}

func (b *BuilderMapping) Select(key string, obj interface{}) {
	b.querySelect = append(b.querySelect, key)
	b.objectSelect = append(b.objectSelect, obj)
}

func (b *BuilderMapping) Insert(key string, obj interface{}) {
	b.queryInsert = append(b.queryInsert, key)
	b.valueInsert = append(b.valueInsert, obj)
}

func (b *BuilderMapping) Update(key string, obj interface{}) {
	b.queryUpdate = append(b.queryUpdate, key)
	b.valueUpdate = append(b.valueUpdate, obj)
}

func (b *BuilderMapping) Where(key string, value interface{}) {
	b.queryWhere = append(b.queryWhere, key)
	b.valueWhere = append(b.valueWhere, value)
}

func (b *BuilderMapping) Join(key string) {
	b.queryJoin = append(b.queryJoin, key)
}
