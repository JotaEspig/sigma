package db

// Gets intersection of allowedColumns and columns. if columns is empty,
// it returns allowedColumns
func GetColumns(allowedColumns []string, columns ...string) []string {
	if len(columns) == 0 {
		return allowedColumns
	}

	var result []string
	for _, c := range columns {
		for _, ac := range allowedColumns {
			if c == ac {
				result = append(result, c)
			}
		}
	}

	return result
}
