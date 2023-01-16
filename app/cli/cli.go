package cli

import (
	"fmt"
	"os"

	"github.com/jtefteller/copilot_cli/app/copilot"
	"github.com/jtefteller/copilot_cli/utility"
	"github.com/urfave/cli/v2"
)

func Run(args []string) error {
	app := &cli.App{
		Name:                 "copilot",
		Usage:                "CLI for OpenAI's Copilot API",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "prompt",
				Aliases:  []string{"p"},
				Usage:    "Prompt for completion",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			r := utility.Request{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				APIKey: os.Getenv("OPENAI_API_KEY"),
			}
			prompt := c.String("prompt")
			cc := &copilot.CompletionConfig{}
			cd := cc.Default(prompt)
			rsp, err := cd.Completion(&r)
			if err != nil {
				return err
			}
			fmt.Println(rsp.Choices[0].Text)
			return nil
		},
	}

	return app.Run(args)
}
