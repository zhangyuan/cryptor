package main

import (
	"log"
	"os"

	"cryptor/pkg"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cryptor",
		Usage: "encrypt and decrypt file",
		Commands: []*cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "mode",
						Usage:   "AES mode",
						Value:   "gcm",
						Aliases: []string{"m"},
					},
					&cli.StringFlag{
						Name:     "password",
						Usage:    "Password",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "salt",
						Usage:    "Salt",
						Aliases:  []string{"s"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "input-file",
						Usage:    "Path to plaintext file",
						Aliases:  []string{"i"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output-file",
						Usage:    "Path to ciphertext file",
						Aliases:  []string{"o"},
						Required: true,
					},
				},
				Usage: "Encrypt a file",
				Action: func(cCtx *cli.Context) error {
					mode := cCtx.String("mode")
					password := cCtx.String("password")
					salt := cCtx.String("salt")
					inputFile := cCtx.String("input-file")
					outputFile := cCtx.String("output-file")
					return pkg.Encrypt(mode, password, salt, inputFile, outputFile)
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "mode",
						Usage:   "AES mode",
						Value:   "gcm",
						Aliases: []string{"m"},
					},
					&cli.StringFlag{
						Name:     "password",
						Usage:    "Password",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "salt",
						Usage:    "Salt",
						Aliases:  []string{"s"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "input-file",
						Usage:    "Path to ciphertext file",
						Aliases:  []string{"i"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output-file",
						Usage:    "Path to plaintext file",
						Aliases:  []string{"o"},
						Required: true,
					},
				},
				Usage: "Decrypt a file",
				Action: func(cCtx *cli.Context) error {
					mode := cCtx.String("mode")
					password := cCtx.String("password")
					salt := cCtx.String("salt")
					inputFile := cCtx.String("input-file")
					outputFile := cCtx.String("output-file")
					return pkg.Decrypt(mode, password, salt, inputFile, outputFile)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
