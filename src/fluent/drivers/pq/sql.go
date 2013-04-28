package pq

import "fmt"
import _ "github.com/bmizerany/pq"
import nsql "fluent"
import gsql "database/sql"
import "fluent/drivers/utils"

/*============================================================================*
 * {{{ Public api
 *============================================================================*/

// Execute a query and return a rowset
func (self *sql) Execute(query *nsql.Query) (nsql.Rowset, error) {
	switch {
		case query.Type == nsql.INSERT:
			return self.Insert(query)
		case query.Type == nsql.SELECT:
			return self.Select(query)
	}
	return nil, utils.Fail(1, "Invalid query type; not supported")
}
	
// Execute a raw SQL query for anything not supported, eg. CREATE TABLE, etc.
func (self *sql) Raw(query string, values ...[]interface{}) (nsql.Rowset, error) {
	return self.exec(query, []string{}, values...)	
}

// Close the connection
func (self *sql) Close() {
	self.db.Close()
}

/*============================================================================*
 * }}}
 *============================================================================*/
 
/*============================================================================*
 * {{{ Implementation
 *============================================================================*/
 
type sql struct {
	db *gsql.DB
  factory *sqlFactory
}

func newSql(uri string) nsql.Sql {
	var db, err = gsql.Open("postgres", uri)
	if err != nil {
		fmt.Printf("Failed to open database connection: %s", err.Error())
		return nil
	}
	return &sql{
		db : db,
    factory : newSqlFactory(),
	}
}

// Execute an sql 'query' that returns 'rows'
func (self *sql) query(query string, values ...[]interface{}) (*rowset, error) {
	var rows *gsql.Rows = nil
	var err error = nil
	
	// Pass params over 
	if len(values) > 0 {
		rows, err = self.db.Query(query, values[0]...)
	} else {
		rows, err = self.db.Query(query)
	}
	if err != nil {
		return nil, utils.Fail(1, "Error perfroming explicit SQL query \"%s\": %s", query, err.Error())
	} 
	
	// Generate output
	var rtn = newRowset(rows).(*rowset)
	return rtn, err
}

// Execute an sql 'exec' that returns 'insert ids' (api is stupid)
func (self *sql) exec(query string, columns []string, values ...[]interface{}) (nsql.Rowset, error) {
	var output gsql.Result = nil
	var err error = nil
	
	// Pass params over 
	if len(values) > 0 {
		output, err = self.db.Exec(query, values[0]...)
	} else {
		output, err = self.db.Exec(query)
	}
	if err != nil {
		return nil, utils.Fail(1, "Error perfroming explicit SQL query \"%s\": %s", query, err.Error())
	}
	
	// Process output
	var rtn = newRowset(nil)
	var id, id_err = output.LastInsertId()
	if id_err == nil {
		var record = map[string]interface{} {}
		for _, c := range(columns) {
			record[c] = id
		}
		(rtn.(*rowset)).Values.PushBack(record)
	}
	
	return rtn, err
}

// Perform an insert operation
func (self *sql) Insert(query *nsql.Query) (nsql.Rowset, error) {
	var err error = nil
	var rtn = newRowset(nil).(*rowset)
  for i := 0; i < len(query.Rows); i++ {
  
  	var stmt = sqlFactoryInsert {
  		Table : query.Table,
  		Columns : make([]string, len(query.Rows[i]), len(query.Rows[i])),
  		Values : make([]interface{}, len(query.Rows[i]), len(query.Rows[i])),
  		Requested : make([]string, len(query.Requested), len(query.Requested)),
  	}
  	
  	var offset = 0
  	for _, key := range query.Requested {
  		stmt.Requested[offset] = key
  		offset++
  	}
  	
  	offset = 0
  	for key, value := range query.Rows[i] {
  		stmt.Columns[offset] = key
  		stmt.Values[offset] = value
  		offset++
  	}
  	
  	var sql, val = self.factory.Insert(&stmt)
  	var rows, rerr = self.query(sql, val)
  	if rerr != nil {
  		err = rerr
  		break;
  	} else {
  		rtn.MergeValues(&rows.Rowset, true)
  	}
  }
	return rtn, err
}
 
 // Perform an insert operation
func (self *sql) Select(query *nsql.Query) (nsql.Rowset, error) {
  var stmt = sqlFactorySelect {
  	Table : query.Table,
  	Columns : make([]string, len(query.Requested), len(query.Requested)),
  	Where : query.Constraints.Where,
  	Values : query.Constraints.Values,
  	Limit : query.Constraints.Limit,
  	Offset : query.Constraints.Offset,
  }
  
  // Null where statement not permitted
  if stmt.Where == "" {
  	stmt.Where = "TRUE"
  }
  	
  offset := 0
  for _, key := range query.Requested {
  	stmt.Columns[offset] = key
  	offset++
  }
  	
  var sql, val = self.factory.Select(&stmt)
  var rows, rerr = self.query(sql, val)
  
	return rows, rerr
}

/*============================================================================*
 * }}}
 *============================================================================*/
