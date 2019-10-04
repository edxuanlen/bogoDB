package query

import (
	"bogoDB/backend/meta"
	"bogoDB/backend/storage"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAnalyzeCreateTable(t *testing.T){
	n := &CreateTableNode{
		TableName:"users",
		ColNames:[]string{"id"},
		ColTypes:[]string{"int"},
	}

	ctg := storage.NewEmtpyCatalog()
	analyzer := NewAnalyzer(ctg)
	q, err := analyzer.analyzeCreateTable(n)

	if err != nil{
		log.Fatal(err)
	}

	assert.Equal(t, "users", q.Scheme.TblName)
	assert.Equal(t, "id", q.Scheme.ColNames[0])
	assert.Equal(t, meta.Int, q.Scheme.ColTypes[0])
}
