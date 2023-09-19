package utils

import "strconv"

func ConvertTodoIdToUint(todoId string) (uint, error) {
	id, err := strconv.ParseUint(todoId, 10, 64)
	return uint(id), err
}
