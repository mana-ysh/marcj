package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli"
)

const (
	EOSString = "EOS"
	Splitter  = "\t"
)

func main() {
	app := cli.NewApp()
	app.Name = "marcj"
	app.Usage = "Morphological Analysis Result Converter for Japanese"
	app.Version = "0.1.0"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "tool, t",
			Value: "sudachi",
			Usage: "Analyzer tool you used",
		},

		cli.BoolFlag{
			Name:  "normalize, n",
			Usage: "Flag for normalization form",
		},
	}

	// action
	app.Action = func(c *cli.Context) error {
		switch c.String("tool") {
		case "sudachi":
			runSudachi(bufio.NewReader(os.Stdin), c.Bool("normalize"))
		case "mecabuni":
			runMeCabUniDic(bufio.NewReader(os.Stdin), c.Bool("normalize"))
		case "mecabipa":
			runMeCabIPADic(bufio.NewReader(os.Stdin), c.Bool("normalize"))
		default:
			runDefault() // raise error
		}
		return nil
	}

	app.Run(os.Args)
}

func runSudachi(reader *bufio.Reader, normflag bool) {
	sent := []string{}
	for {
		input, err := reader.ReadString('\n')

		if err != nil && err == io.EOF {
			break
		}

		input = strings.TrimRight(input, "\n")

		if input == EOSString {
			fmt.Println(strings.Join(sent, " "))
			sent = []string{}
		} else {
			info := strings.Split(input, Splitter)
			if len(info) != 3 {
				fmt.Println("File may be broken or Invalid format")
				os.Exit(1)
			}
			surface := ""
			if normflag {
				surface += info[2]
			} else {
				surface += info[0]
			}
			sent = append(sent, surface)
		}
	}
}

func runMeCabUniDic(reader *bufio.Reader, normflag bool) {
	sent := []string{}
	for {
		input, err := reader.ReadString('\n')

		if err != nil && err == io.EOF {
			break
		}

		input = strings.TrimRight(input, "\n")

		if input == EOSString {
			fmt.Println(strings.Join(sent, " "))
			sent = []string{}
		} else {
			info := strings.Split(input, Splitter)
			if len(info) != 7 {
				fmt.Println("File may be broken or Invalid format")
				os.Exit(1)
			}
			surface := ""
			if normflag {
				fmt.Println("Normalized form not supported yet")
				os.Exit(1)
				surface += info[2]
			} else {
				surface += info[0]
			}
			sent = append(sent, surface)
		}
	}
}

func runMeCabIPADic(reader *bufio.Reader, normflag bool) {
	sent := []string{}
	for {
		input, err := reader.ReadString('\n')

		if err != nil && err == io.EOF {
			break
		}

		input = strings.TrimRight(input, "\n")

		if input == EOSString {
			fmt.Println(strings.Join(sent, " "))
			sent = []string{}
		} else {
			info := strings.Split(input, Splitter)
			if len(info) != 2 {
				fmt.Println("File may be broken or Invalid format")
				os.Exit(1)
			}
			surface := ""
			if normflag {
				fmt.Println("Normalized form not supported yet")
				os.Exit(1)
				surface += info[2]
			} else {
				surface += info[0]
			}
			sent = append(sent, surface)
		}
	}
}

func runDefault() {
	fmt.Println("Not Available")
	os.Exit(1)
}
