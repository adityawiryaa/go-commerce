package main

import (
	"fmt"
	"go-commerce/pkg/config"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var command = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CT version v1.0.0.0")
	},
}

func commandMigrationStatus(cfg *config.Migration) *cobra.Command {
	c := &cobra.Command{
		Use:   "migrations:status",
		Short: "Chek migrations status",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("status")
		},
	}
	return c
}

func commandMigrationCreate(cfg *config.Migration) *cobra.Command {
	fmt.Println(cfg)
	var table, create string
	c := &cobra.Command{
		Use:   "migrations:create",
		Short: "Create SQL file migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("arg 0 is required")
				os.Exit(1)
			}
			if len(table) < 1 && len(create) < 1 {
				fmt.Println("create or table param required")
				os.Exit(1)
			}
			param := args[0]
			version := int(time.Now().Unix())
			filename := strconv.Itoa(version) + "_" + param
			f1, _ := os.Create(cfg.Path + "/" + filename + ".up.sql")
			f2, _ := os.Create(cfg.Path + "/" + filename + ".down.sql")
			log.Println(f1)
			if len(table) > 0 {
				f1.WriteString(fmt.Sprintf("ALTER TABLE %s\nALTER COLUMN %s TYPE %s;", table, "column_name", "column_definition"))
				f2.WriteString(fmt.Sprintf("ALTER TABLE %s\nALTER COLUMN %s TYPE %s;", table, "column_name", "column_definition"))
			} else {
				f1.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s \n(\n);", create))
				f2.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %s;", create))
			}
		},
	}
	c.Flags().StringVar(&table, "table", "", "table name of migrations")
	c.Flags().StringVar(&create, "create", "table_name", "table name of migrations")
	return c
}

func commandMigrationUp(cfg *config.Migration) *cobra.Command {
	c := &cobra.Command{
		Use:   "migrations:up",
		Short: "Migrate up database",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrate.New(fmt.Sprintf("file://%s", cfg.Path), cfg.Database)
			if err != nil {
				log.Fatal(err)
			}
			if err := m.Up(); err != nil {
				log.Fatal(err)
			}
		},
	}
	return c
}

func commandMigrationDown(cfg *config.Migration) *cobra.Command {
	var input int
	c := &cobra.Command{
		Use:   "migrations:down",
		Short: "Migrate down database",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("down")
		},
	}
	c.Flags().IntVar(&input, "step", 1, "Number of step down migrations")
	return c
}

func init() {
	cfg, err := config.Migrations()
	if err != nil {
		panic(err)
	}

	command.AddCommand(commandMigrationStatus(cfg))
	command.AddCommand(commandMigrationCreate(cfg))
	command.AddCommand(commandMigrationUp(cfg))
	command.AddCommand(commandMigrationDown(cfg))
}

func main() {
	err := command.Execute()
	if err != nil {
		log.Println(err.Error())
	}
}
