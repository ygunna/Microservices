package main

import (
	"fmt"
	"os"

	"crossent/micro/studio/studiocmd"
	"github.com/jessevdk/go-flags"
	"github.com/alexedwards/scs"
	"crossent/micro/studio/domain"
)

// cloudfoundry/diego-release/src/code.cloudfoundry.org/auctioneer/cmd/auctioneer/main.go

func main() {
	// Initialize a new encrypted-cookie based session manager and store it in a global
	// variable. In a real application, you might inject the session manager as a
	// dependency to your handlers instead. The parameter to the NewCookieManager()
	// function is a 32 character long random key, which is used to encrypt and
	// authenticate the session cookies.
	domain.SessionManager = *scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

	cmd := &microcmd.MicroCommand{}

	parser := flags.NewParser(cmd, flags.Default)
	parser.NamespaceDelimiter = "-"

	//cmd.WireDynamicFlags(parser.Command)

	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	err = cmd.Execute(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
