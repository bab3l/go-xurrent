// collision-hints prints JSON discovery steps for duplicate identifier errors (read-only).
//
//	go run ./cmd/collision-hints -entity team -field name
//	go run ./cmd/collision-hints -list
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/xurrent/go-xurrent/pkg/collisionhints"
)

func main() {
	entity := flag.String("entity", "", "entity kind: team, site, service, person, product, organization, configuration_item, sla")
	field := flag.String("field", "", "duplicate field from 422 (default: name)")
	list := flag.Bool("list", false, "list supported entity kinds and exit")
	flag.Parse()

	if *list {
		for _, k := range []collisionhints.Kind{
			collisionhints.KindTeam,
			collisionhints.KindSite,
			collisionhints.KindService,
			collisionhints.KindPerson,
			collisionhints.KindProduct,
			collisionhints.KindOrganization,
			collisionhints.KindCI,
			collisionhints.KindSLA,
		} {
			fmt.Println(string(k))
		}
		os.Exit(0)
	}

	if *entity == "" {
		fmt.Fprintln(os.Stderr, "usage: collision-hints -entity <kind> [-field <field>]")
		fmt.Fprintln(os.Stderr, "        collision-hints -list")
		flag.PrintDefaults()
		os.Exit(2)
	}

	k, err := collisionhints.ParseKind(*entity)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	plan, err := collisionhints.DetectIdentifierCollision(k, *field)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(plan); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
