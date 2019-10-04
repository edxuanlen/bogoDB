package meta

// Scheme is the definition of tables and columns
type Scheme struct {
	TblName string `json:"tblName"`
	ColNames []string `json:"colNames"`
	ColTypes []ColType `json:"colTypes"`
}

func NewScheme(tblName string, colNames []string, colTypes []ColType) *Scheme{
	return &Scheme{
		TblName:tblName,
		ColNames:colNames,
		ColTypes:colTypes,
	}
}

type Table struct {
	name string
	columns []Column
}

type Column struct {
	name string
	ctype string
}

type ColType uint8

const(
	Int ColType = iota
	Varchar
)

func (c ColType) String() string{
	if c == Int{
		return "int"
	}

	if c == Varchar{
		return "varchar"
	}

	return "undefined"
}
