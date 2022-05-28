package db

// Gets intersection of allowedColumns and columns
func GetColumns(allowedColumns []string, columns ...string) []string {
	var columnsToUse []string

	if len(columns) == 0 {
		return allowedColumns
	}

	// gets intersection of allowedColumns and columns
	for _, col := range columns {
		for _, allowedCol := range allowedColumns {
			if col == allowedCol {
				columnsToUse = append(columnsToUse, col)
			}
		}
	}

	return columnsToUse
}
