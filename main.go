package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jpanderson91/blog-aggregator/internal/config"
	"github.com/jpanderson91/blog-aggregator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}


func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	dbQueries := database.New(db)
	programState := state{
		cfg: &cfg,
		db:  dbQueries,
	}
cmds := commands{
    registeredCommands: make(map[string]func(*state, command) error),
}
cmds.register("login", handlerLogin)
cmds.register("register", handlerRegister)
cmds.register("reset", handlerReset)
cmds.register("users", handlerUsers)
cmds.register("agg", handlerAgg)
cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
cmds.register("feeds", handlerFeeds)
cmds.register("follow", middlewareLoggedIn(handlerFollow))
cmds.register("following", middlewareLoggedIn(handlerFollowing))
cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
cmds.register("browse", middlewareLoggedIn(handlerBrowse))
//Check if len(os.Args) < 2 - if so, call log.Fatal with an error message
//Extract the command name from os.Args[1]
//Extract the arguments from os.Args[2:]
//Create a command instance and run it with cmds.run
if len(os.Args) < 2 {
	log.Fatal("no command provided")
}
commandName := os.Args[1]
commandArgs := os.Args[2:]
cmd := command{
	Name: commandName,
	Args: commandArgs,
}
err = cmds.run(&programState, cmd)
if err != nil {
	log.Fatalf("error running command: %v", err)
}
}

