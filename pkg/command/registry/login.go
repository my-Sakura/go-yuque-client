package registry

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewLoginCommand creates a new `yuque login` command
func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Log in yuque application",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			return runLogin(args[0])
		},
	}

	return cmd
}

func runLogin(token string) error {
	viper.Set("token", token)

	return viper.WriteConfig()
}
