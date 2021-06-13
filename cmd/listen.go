package cmd

import (
	"context"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/svix/svix-cli/relay"
)

type listenCmd struct {
	cmd *cobra.Command
}

func newListenCmd() *listenCmd {
	lc := &listenCmd{}
	lc.cmd = &cobra.Command{
		Use:   "listen [localURL]",
		Short: "Forwards webhook requests a port on localhost",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			urlStr := args[0]
			url, err := url.Parse(urlStr)
			if err != nil {
				return fmt.Errorf("invalid url %s", urlStr)
			}
			client := relay.NewClient(url, &relay.ClientOptions{
				DisableSecurity: viper.GetBool("relay_disable_security"),
				RelayDebugUrl:   viper.GetString("relay_debug_url"),
			})
			client.Listen(context.Background())
			return nil
		},
	}
	return lc
}
