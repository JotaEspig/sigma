package db

func GetColumns(columns ...string) interface{} {
	var columnsToUse interface{}

	columnsToUse = "*"
	if len(columns) != 0 {
		columnsToUse = columns
	}

	return columnsToUse
}
