package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type App struct {
	Distributors map[string]Distributor
	Places       *Places
}

func main() {
	data := Read_file("cities.csv")
	p := LoadPlace(data)
	app := App{
		Distributors: make(map[string]Distributor),
	}

	app.Places = p

	r := bufio.NewReader(os.Stdin)
	fmt.Println("====================Distribution System====================")
	fmt.Println("        Please consult README.md for commands")
	for {
		fmt.Print("> ")
		input, _ := r.ReadString(byte('\n'))
		input, _ = strings.CutSuffix(input, "\n")
		if input == "exit" {
			break
		}

		cmds := strings.Fields(input)
		if cmds[0] == "make" {
			if len(cmds) == 2 {
				dist_name := strings.TrimSpace(cmds[1])
				if dist_name == "" {
					fmt.Println("This command needs a name for distributor")
					continue
				}
				err := app.New_Distributor(dist_name)
				if err != nil {
					fmt.Println(err.Error())
				}
				continue
			}
			if len(cmds) == 4 && cmds[2] == "<" {
				err := app.New_Dist_With_Parent(cmds[3], cmds[1])
				if err != nil {
					fmt.Println(err.Error())
				}
				continue
			}
			fmt.Println("invalid command")
			continue
		}

		if cmds[0] == "for" {
			if len(cmds) < 3 {
				fmt.Println("Invalid command")
				continue
			}
			dist := cmds[1]
			action := cmds[2]

			if action == "list" {
				err := app.Get_Permissions(dist)
				if err != nil {
					fmt.Println(err.Error())
				}
				continue
			}

			place_codes := strings.Split(cmds[3], ",")
			if action == "include" {
				for _, v := range place_codes {
					err := app.Include_Permissions(dist, v)
					if err != nil {
						fmt.Println(err.Error())
					}
					continue
				}
			} else if action == "exclude" {
				for _, v := range place_codes {
					err := app.Exclude_Permissions(dist, v)
					if err != nil {
						fmt.Println(err.Error())
					}
					continue
				}
			} else {
				fmt.Println("invalid command")
				continue
			}
			continue
		}
		fmt.Println("invalid command")
	}
}
