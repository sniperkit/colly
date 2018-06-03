package models

type ExportConfig struct {
	// Disabled sets the status of all export configuration values
	Disabled bool `default:'false' json:'enabled' yaml:'enabled' toml:'enabled' xml:'enabled' ini:'enabled'`

	// Encoding sets..
	Encoding string `default:'csv' json:'encoding' yaml:'encoding' toml:'encoding' xml:'encoding' ini:'encoding'`

	// Format sets..
	UseTemplate string `json:'format' yaml:'format' toml:'format' xml:'format' ini:'format'`

	// PrefixPath sets..
	PrefixPath string `default:'./shared' json:'prefix_path' yaml:'prefix_path' toml:'prefix_path' xml:'prefixPath' ini:'prefixPath'`

	// ExportDir
	ExportDir string `default:'./storage/export' json:'export_dir' yaml:'export_dir' toml:'export_dir' xml:'exportDir' ini:'exportDir'`

	// ForceDir specifies that the program will try to create missing storage directories.
	EnsureDirs bool `default:'true' json:'ensure_dir' yaml:'ensure_dir' toml:'ensure_dir' xml:'ensureDirs' ini:'ensureDirs'`

	// ForceDirRecursive specifies that the program will try to create missing storage directories recursively.
	ForceDirRecursive bool `default:'true' json:'ensure_dir_recursively' yaml:'ensure_dir_recursively' toml:'ensure_dir_recursively' xml:'ensureDirRecursively' ini:'ensureDirRecursively'`
}
