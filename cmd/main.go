package main

import (
	"flag"
	"log"

	"school/internal/api/router"
	"school/internal/config"
	"school/internal/database"
)

type args struct {
	env *string
}

// @title									school
// @version									1.0
// @description								school Specification
// @BasePath								/api
//
// @securitydefinitions.oauth2.password		OAuth2Password
//
// @tokenUrl								http://srvch025166003:8080/auth/realms/dfi_realm/protocol/openid-connect/token
// @in										header
// @name									Authorization
// @description								Authorization
//
// @securityDefinitions.apikey				Bearer
// @in										header
// @name									Authorization
// @description								Type "Bearer " followed by a space and JWT token.
func main() {
	args := parseArgs()
	config.LoadEnv(*args.env)

	err := database.Connect()
	if err != nil {
		log.Fatalf("error db connect: %v", err)
	}

	r := router.NewRouter()
	r.SetupRouter()
	r.Run()
}

func parseArgs() (args args) {
	args.env = flag.String("env", ".env", "path to environment file")
	flag.Parse()
	return args
}
