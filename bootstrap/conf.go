package bootstrap

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"wetees.com/domain"
	"wetees.com/internal/configloader"
)

func NewConf() *domain.Config {
	appConfigPath := "./config/config.yaml"

	loader := configloader.New(appConfigPath,
		// before config loaded hook
		func(c *domain.Config, loader *viper.Viper) {
			c.LoadDefaultValue(loader)
		},
		// after config loaded hooks
		func(c *domain.Config, loader *viper.Viper) {
			// validate configuration
			if err := c.Validate(); err != nil {
				log.Fatal().Msgf("Error validating configuration: %s", err)
			}
		},
	)

	// get config and expand conf variable to configuration array
	conf := loader.GetConfig()

	return &conf
}
