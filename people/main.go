package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	_ "people/docs"
	"people/repositories"
	routes "people/routes"
	"people/services"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}


func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "content-type")

			if req.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(w, req)
		})
}
func main() {
	log.Println("Starting application...")

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConn, err := sql.Open(`mysql`, connection)

	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	r := mux.NewRouter()

	userRepo := repositories.NewUserRepository(dbConn)
	groupRepo := repositories.NewGroupRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userService := services.NewUserService(userRepo, timeoutContext)
	groupService := services.NewGroupService(groupRepo, timeoutContext)

	routes.NewUserHandler(r, userService)
	routes.NewGroupHandler(r, groupService)

	r.Use(CORSMiddleware)

	log.Println("Listening on port: " + viper.GetString(`server.port`))

	log.Fatal(http.ListenAndServe(viper.GetString(`server.port`), r))
}