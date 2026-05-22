package http

import "fmt"

// return Validator::make($data, [
//
//	'name' => ['required', 'string', 'min:2', 'max:35'],
//	'email' => ['required', 'string', 'email', 'max:190', 'unique:users'],
//	'password' => ['required', 'string', 'min:8'],
//
// ]);
func validMin(attr, val string, min int) (string, bool) {
	if len(val) >= min {
		return "", true
	}
	return fmt.Sprintf("The %s must be at least %d characters.", attr, min), false
}

func validMax(attr, val string, max int) (string, bool) {
	if len(val) <= max {
		return "", true
	}
	return fmt.Sprintf("The %s must not be greater than %d characters.", attr, max), false
}

func validRequired(attr, val string) (string, bool) {
	if val != "" {
		return "", true
	}
	return fmt.Sprintf("The %s field is required.", attr), false
}

func validUnique(table, field, val, attr string) (string, bool) {
	// TODO: make query

	return fmt.Sprintf("The %s has already been taken.", attr), false
}
