package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bbengfort/ipseity"
	"github.com/bbengfort/ipseity/pb"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	// Load the .env file if it exists
	godotenv.Load()

	// Instantiate the command line application
	app := cli.NewApp()
	app.Name = "ipseity"
	app.Version = ipseity.PackageVersion
	app.Usage = "server that hands out a monotonically increasing identity"

	// Define commands available to application
	app.Commands = []cli.Command{
		{
			Name:     "serve",
			Usage:    "run an ipseity identity server",
			Action:   serve,
			Category: "server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "a, addr",
					Usage: "address for the server to listen on",
					Value: ":3265",
				},
				cli.StringFlag{
					Name:  "s, stype",
					Usage: "type of server to initialize",
					Value: "simple",
				},
				cli.DurationFlag{
					Name:  "u, uptime",
					Usage: "only run for a specified amount of time",
					Value: 0,
				},
			},
		},
		{
			Name:     "get",
			Usage:    "get the next identity from the server",
			Action:   get,
			Category: "client",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "a, addr",
					Usage: "address of the identity server",
					Value: "localhost:3265",
				},
				cli.StringFlag{
					Name:  "k, key",
					Usage: "a key to associate with the identity",
					Value: "",
				},
			},
		},
		{
			Name:     "bench",
			Usage:    "run a benchmark for the given number of clients",
			Action:   bench,
			Category: "client",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "a, addr",
					Usage: "address of the identity server",
					Value: "localhost:3265",
				},
				cli.IntFlag{
					Name:  "c, clients",
					Usage: "number of concurrent clients",
					Value: 4,
				},
				cli.IntFlag{
					Name:  "m, messages",
					Usage: "messages per client to send to server",
					Value: 5000,
				},
				cli.StringFlag{
					Name:  "s, stype",
					Usage: "expected type of server running",
					Value: "",
				},
			},
		},
	}

	// Run the CLI program
	app.Run(os.Args)
}

//===========================================================================
// Server Functions
//===========================================================================

func serve(c *cli.Context) error {

	server, err := ipseity.New(c.String("stype"))
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	if uptime := c.Duration("uptime"); uptime > 0 {
		time.AfterFunc(uptime, func() {
			os.Exit(0)
		})
	}

	if err := server.Listen(c.String("addr")); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}

//===========================================================================
// Client Functions
//===========================================================================

func get(c *cli.Context) error {
	client, err := ipseity.NewClient(c.String("addr"))
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	req := &pb.IdentityRequest{Key: c.String("key")}
	rep, err := client.Next(context.Background(), req)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Printf("%X\n", rep.Identity)
	return nil
}

func bench(c *cli.Context) error {
	addr := c.String("addr")
	if addr == "" {
		return cli.NewExitError("must specify an address to connect to", 1)
	}

	bench, err := ipseity.NewBenchmark(c.Int("clients"), c.Int("messages"), addr, c.String("stype"))
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(bench.String())
	return nil
}
