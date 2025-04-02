package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/FT1006/blogator/internal/config"
	"github.com/FT1006/blogator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func newState(cfg *config.Config, db *database.Queries) *state {
	return &state{
		config: cfg,
		db:     db,
	}
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// get the DB URL from config
	dbURL := cfg.DBUrl

	// open the database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	// defer closing the database connection
	defer db.Close()

	// create the database queries
	dbQueries := database.New(db)

	// create the state
	currentState := newState(&cfg, dbQueries)

	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments provided")
		os.Exit(1) // Exit with non-zero status code to indicate an error
	}

	var cmdMap commands
	cmdMap.CommandMap = make(map[string]func(*state, command) error)
	cmdMap.register("login", handlerLogin)
	cmdMap.register("register", handlerRegister)
	cmdMap.register("reset", handlerReset)
	cmdMap.register("users", handlerUsers)
	cmdMap.register("agg", handlerAgg)
	cmdMap.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmdMap.register("feeds", handlerFeeds)
	cmdMap.register("follow", middlewareLoggedIn(handlerFollow))
	cmdMap.register("following", middlewareLoggedIn(handlerFollowing))
	cmdMap.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmdMap.register("browse", handlerBrowse)

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmdMap.run(currentState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
