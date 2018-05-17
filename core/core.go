package core

func Init() {
	var configPath string // config path
	var port int

	// Parse command line flags
	SetupAndParseFlags(&configPath, &port)

	routes := ParseConfig(configPath)

	ServeRoutes(routes, port)
}
