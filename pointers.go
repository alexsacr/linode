package linode

// Int is a convenience function for creating an integer pointer.
func Int(i int) *int {
	return &i
}

// String is a convenience function for creating a string pointer.
func String(s string) *string {
	return &s
}

// Bool is a convenience function for creating a bool pointer.
func Bool(b bool) *bool {
	return &b
}
