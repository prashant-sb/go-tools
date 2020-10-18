package main

import (
	"flag"
	"fmt"

	log "github.com/golang/glog"
	uinfo "github.com/prashant-sb/go-utils/userinfo/users"
)

// CLI Flags:
//
// -list -user <username>   : List specific user schema
// -list                    : List all system users
// -create -from <json>	    : Create user from given json schema file
// -delete -user <username> : Deletes user by username
var (
	list   = flag.Bool("list", false, "Lists the system users")
	create = flag.Bool("create", false, "Creates the system user")
	delete = flag.Bool("delete", false, "Deletes the system user")

	user = flag.String("user", "", "List specific system user")
	from = flag.String("from", "", "Json configuration for create user")
)

func main() {
	flag.Parse()

	switch {
	case *list:
		// Get the user details
		if *user != "" {
			ui := uinfo.NewUserOps()

			u, err := ui.Get(*user)
			if err != nil {
				log.Error(err.Error())
				return
			}

			jsonUser, err := uinfo.Decode(u)
			if err != nil {
				log.Error("Error in decode ", jsonUser)
				return
			}

			fmt.Printf("%+v\n", jsonUser)
		} else {
			// List all users
			ul := uinfo.NewUserList()
			ulist, err := ul.Get()
			if err != nil {
				log.Error("Error in Get User list: ", err)
				return
			}

			jsonList, err := uinfo.Decode(ulist)
			if err != nil {
				log.Error("Error in decode: ", err)
				return
			}

			fmt.Printf("%v\n", jsonList)
		}

	case *create:
		// Add user from json User Schema
		if *from != "" {
			ui := uinfo.NewUserOps()
			userName, err := ui.AddUser(*from)
			if err != nil {
				log.Error(err.Error())
				return
			}
			fmt.Printf("User %s added\n", userName)
		}

	case *delete:
		// Deletes user by Username
		if *user != "" {
			ui := uinfo.NewUserOps()
			if _, err := ui.DeleteUser(*user); err != nil {
				log.Error(err.Error())
				return
			}
			fmt.Printf("%s user deleted.\n", *user)
		}

	default:
		// Prints usage in all other cases.
		flag.Usage()
	}
}
