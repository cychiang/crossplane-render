package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/cychiang/crossplane-render/cmd/render"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	logger := logging.NewLogrLogger(zap.New(zap.UseDevMode(true)))
	ctx := kong.Parse(&render.Cmd{},
		kong.Name("crossplane-render"),
		kong.Description("A command line tool for rendering crossplane compositions."),
		kong.BindTo(logger, (*logging.Logger)(nil)),
		kong.ConfigureHelp(kong.HelpOptions{
			FlagsLast:      true,
			Compact:        true,
			WrapUpperBound: 80,
		}),
		kong.UsageOnError(),
	)
	err := ctx.Run(logger)
	ctx.FatalIfErrorf(err)

	os.Exit(0)
}