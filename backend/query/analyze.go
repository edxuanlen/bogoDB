package query

import (
	"bogoDB/backend/meta"
	"bogoDB/backend/storage"
	"fmt"
)

// Analyzer analyze the parsed sql.
// Roles:
//   - Fetch an actual scheme from the table name.
//   - Validate the rules of sql.
type Analyzer struct {
	catalog *storage.Catalog
}

type Query interface {
}

type SelectQuery struct {
	Col []*meta.Column
	From []*meta.Table
}

type CreateTableQuery struct {
	Scheme *meta.Scheme
}

func NewAnalyzer(catalog *storage.Catalog) *Analyzer{
	return &Analyzer{
		catalog:catalog,
	}
}

func (a *Analyzer) analyzeSelect(n *SelectNode) (*SelectQuery, error){
	var q *SelectQuery

	// analyze `from`
	var schemes []*meta.Scheme
	for _, name := range n.From.TableNames{
		scheme := a.catalog.FetchScheme(name)
		if scheme != nil{
			return nil, fmt.Errorf("select failed :table `%s` doesn't exist", name)
		}
		schemes = append(schemes, scheme)
	}

	// analyze `select`
	var cols []*meta.Column
	for _, colName := range n.ColNames{
		found := false
		for _, scheme := range schemes{
			for _, col := range scheme.ColNames{
				if col == colName{
					found = true
					// cols = append(cols, )
				}
			}
		}

		if !found{
			return nil, fmt.Errorf("select failed : column `%s` doesn't exist", colName)
		}
	}

	// q.From = schemes
	q.Col = cols
	return q, nil
}

func (a *Analyzer) analyzeCreateTable(n *CreateTableNode) (*CreateTableQuery, error){
	var q CreateTableQuery

	if a.catalog.HasScheme(n.TableName){
		return nil, fmt.Errorf("create table failed : table name `%s` already exists", n.TableName)
	}

	var types []meta.ColType
	for _, typ := range n.ColTypes{
		if typ == "int"{
			types = append(types, meta.Int)
		} else if typ == "varchar"{
			types = append(types, meta.Varchar)
		}
	}

	q.Scheme = meta.NewScheme(n.TableName, n.ColNames, types)
	return &q, nil
}

func (a *Analyzer) AnalyzeMain(n Node) (Query, error){
	switch concrete := n.(type) {
	case *SelectNode:
		return a.analyzeSelect(concrete)
	case *CreateTableNode:
		return a.analyzeCreateTable(concrete)
	}

	return nil, fmt.Errorf("failed to analyze query : %s", n.ErrInfo())
}