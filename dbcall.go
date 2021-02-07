package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx"
)

var purchase struct {
	cardNo  string
	expDate string
	amt     int
	mcode   string
	curr    string
}

var conn *pgx.Conn

func main() {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

	start := time.Now()

	var err error
	// conn, err = pgx.Connect(context.Background(), "postgresql://postgres:xmbd2311@localhost")
	conn, err = pgx.Connect(context.Background(), "postgres://root:admin@127.0.0.1:41801/movr?sslmode=require")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(" conn time ", time.Since(start))
	connTime := time.Now()

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "list":
		err = listTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list tasks: %v\n", err)
			os.Exit(1)
		}

	case "add":
		err = addTask(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to add task: %v\n", err)
			os.Exit(1)
		}

	case "update":
		n, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into int: %v\n", err)
			os.Exit(1)
		}
		err = updateTask(int(n), os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to update task: %v\n", err)
			os.Exit(1)
		}

	case "remove":
		n, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into int: %v\n", err)
			os.Exit(1)
		}
		err = removeTask(int(n))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to remove task: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
		printHelp()
		os.Exit(1)
	}
	fmt.Println(" completion time ", time.Since(connTime))

}

func listTasks() error {
	rows, _ := conn.Query(context.Background(), "select * from tasks")

	for rows.Next() {
		var id int
		var description string
		err := rows.Scan(&id, &description)
		if err != nil {
			return err
		}
		fmt.Printf("%d. %s\n", id, description)
	}

	return rows.Err()
}

func addTask(description string) error {
	_, err := conn.Exec(context.Background(), "insert into tasks(description) values($1)", description)
	return err
}

func updateTask(itemNum int, description string) error {
	_, err := conn.Exec(context.Background(), "update tasks set description=$1 where id=$2", description, itemNum)
	return err
}

func removeTask(itemNum int) error {
	_, err := conn.Exec(context.Background(), "delete from tasks where id=$1", itemNum)
	return err
}

func printHelp() {
	fmt.Print(`Todo pgx demo
Usage:
  todo list
  todo add task
  todo update task_num item
  todo remove task_num
Example:
  todo add 'Learn Go'
  todo list
`)
}
