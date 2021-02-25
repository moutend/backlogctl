package app

import (
	"os"
	"path/filepath"
	"time"

	"github.com/moutend/backlogctl/internal/cache"
	"github.com/moutend/backlogctl/internal/migrate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCommand = &cobra.Command{
	Use:                "backlogctl",
	Short:              "backlogctl - command line backlog client",
	SilenceUsage:       true,
	PersistentPreRunE:  rootPersistentPreRunE,
	PersistentPostRunE: rootPersistentPostRunE,
}

func rootPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	viper.AutomaticEnv()

	if path, _ := cmd.Flags().GetString("config"); path != "" {
		viper.SetConfigFile(path)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	space := viper.GetString("BACKLOG_SPACE")
	token := viper.GetString("BACKLOG_TOKEN")

	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	cachePath := filepath.Join(wd, ".backlog")
	db3Path := filepath.Join(cachePath, space+".db3")

	os.Mkdir(cachePath, 0755)

	if err := migrate.Setup(db3Path); err != nil {
		return err
	}
	if err := cache.Setup(space, token, db3Path); err != nil {
		return err
	}

	return nil
}

func rootPersistentPostRunE(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	RootCommand.PersistentFlags().BoolP("debug", "d", false, "enable debug output")
	RootCommand.PersistentFlags().BoolP("skip-fetch", "s", false, "skip fetch")
	RootCommand.PersistentFlags().StringP("config", "c", "", "path to configuration file")
	RootCommand.PersistentFlags().DurationP("timeout", "t", 5*time.Minute, "timeout")
}
