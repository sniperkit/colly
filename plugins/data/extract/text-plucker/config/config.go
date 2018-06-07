package config

// Config specifies parameters for plucking content
type Config struct {

	// Name sets the key in the returned map, after completion
	Name string

	// Activators must be found in order, before capturing commences
	Activators []string

	// Permanent set the number of activators that stay permanently (counted from left to right)
	Permanent int

	// Deactivator restarts capturing
	Deactivator string

	// Finisher trigger the end of capturing this pluck
	Finisher string // finishes capturing this pluck

	// Limit specifies the number of times capturing can occur
	Limit int

	// Sanitize enables the html stripping
	Sanitize bool

	// Maximum set the number of characters for a capture
	Maximum int

	//-- End
}
