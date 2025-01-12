package cli

import (
	"io/ioutil"
	"testing"
)

func testApp() *App {
	app := NewApp()
	app.Name = "greet"
	app.Flags = []Flag{
		&StringFlag{
			Name:      "socket",
			Aliases:   []string{"s"},
			Usage:     "some 'usage' text",
			Value:     "value",
			TakesFile: true,
		},
		&StringFlag{Name: "flag", Aliases: []string{"fl", "f"}},
		&BoolFlag{
			Name:    "another-flag",
			Aliases: []string{"b"},
			Usage:   "another usage text",
		},
	}
	app.Commands = []*Command{{
		Aliases: []string{"c"},
		Flags: []Flag{
			&StringFlag{
				Name:      "flag",
				Aliases:   []string{"fl", "f"},
				TakesFile: true,
			},
			&BoolFlag{
				Name:    "another-flag",
				Aliases: []string{"b"},
				Usage:   "another usage text",
			},
		},
		Name:  "config",
		Usage: "another usage test",
		Subcommands: []*Command{{
			Aliases: []string{"s", "ss"},
			Flags: []Flag{
				&StringFlag{Name: "sub-flag", Aliases: []string{"sub-fl", "s"}},
				&BoolFlag{
					Name:    "sub-command-flag",
					Aliases: []string{"s"},
					Usage:   "some usage text",
				},
			},
			Name:  "sub-config",
			Usage: "another usage test",
		}},
	}, {
		Aliases: []string{"i", "in"},
		Name:    "info",
		Usage:   "retrieve generic information",
	}, {
		Name: "some-command",
	}, {
		Name:   "hidden-command",
		Hidden: true,
	}}
	app.UsageText = "app [first_arg] [second_arg]"
	app.Usage = "Some app"
	app.Authors = []*Author{{Name: "Harrison", Email: "harrison@lolwut.com"}}
	return app
}

func expectFileContent(t *testing.T, file, expected string) {
	data, err := ioutil.ReadFile(file)
	expect(t, err, nil)
	expect(t, string(data), expected)
}

func TestToMarkdownFull(t *testing.T) {
	// Given
	app := testApp()

	// When
	res, err := app.ToMarkdown()

	// Then
	expect(t, err, nil)
	expectFileContent(t, "testdata/expected-doc-full.md", res)
}

func TestToMarkdownNoFlags(t *testing.T) {
	// Given
	app := testApp()
	app.Flags = nil

	// When
	res, err := app.ToMarkdown()

	// Then
	expect(t, err, nil)
	expectFileContent(t, "testdata/expected-doc-no-flags.md", res)
}

func TestToMarkdownNoCommands(t *testing.T) {
	// Given
	app := testApp()
	app.Commands = nil

	// When
	res, err := app.ToMarkdown()

	// Then
	expect(t, err, nil)
	expectFileContent(t, "testdata/expected-doc-no-commands.md", res)
}

func TestToMan(t *testing.T) {
	// Given
	app := testApp()

	// When
	res, err := app.ToMan()

	// Then
	expect(t, err, nil)
	expectFileContent(t, "testdata/expected-doc-full.man", res)
}
