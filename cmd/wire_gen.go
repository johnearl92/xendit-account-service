// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"github.com/gorilla/mux"
	"github.com/johnearl92/xendit-account-service/internal/db"
	"github.com/johnearl92/xendit-account-service/internal/handler"
	"github.com/johnearl92/xendit-account-service/internal/server"
	"github.com/johnearl92/xendit-account-service/internal/service"
	"github.com/johnearl92/xendit-account-service/internal/store"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func createServer() (*server.APIServer, error) {
	config := ProvideServerConfig()
	router := mux.NewRouter()
	dbConfig := ProvideDBConfig()
	gormDB, err := db.NewConn(dbConfig)
	if err != nil {
		return nil, err
	}
	accountStore := store.NewAccountStore(gormDB)
	xenditService := service.NewXenditService(accountStore)
	xenditHandler := handler.NewXenditHandler(xenditService)
	v := ProvideXenditHandlers(xenditHandler)
	apiServer := server.NewAPIServer(config, router, v)
	return apiServer, nil
}

// wire.go:

// ProvideServerConfig server configuration
func ProvideServerConfig() *server.Config {
	return server.NewAPIServerConfig(viper.GetString("server.host"), viper.GetInt("server.port"), viper.GetString("server.spec"), viper.GetStringSlice("server.cors.allowedOrigins"), viper.GetStringSlice("server.cors.allowedHeaders"), viper.GetStringSlice("server.cors.allowedMethods"))
}

// ProvideDBConfig db configuration
func ProvideDBConfig() *db.DBConfig {
	return db.NewDBConfig(viper.GetString("db.host"), viper.GetInt("db.port"), viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.name"), viper.GetInt("db.pool.minOpen"), viper.GetInt("db.pool.maxOpen"), viper.GetBool("db.migrate"), viper.GetBool("db.logMode"))
}

// ProvideXenditHandlers handler injection
func ProvideXenditHandlers(p *handler.XenditHandler) []handler.Routable {
	return []handler.Routable{p}
}
