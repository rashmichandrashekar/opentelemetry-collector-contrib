// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
)

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for resourcedetectionprocessor/heroku resource attributes.
type ResourceAttributesConfig struct {
	CloudProvider                  ResourceAttributeConfig `mapstructure:"cloud.provider"`
	HerokuAppID                    ResourceAttributeConfig `mapstructure:"heroku.app.id"`
	HerokuDynoID                   ResourceAttributeConfig `mapstructure:"heroku.dyno.id"`
	HerokuReleaseCommit            ResourceAttributeConfig `mapstructure:"heroku.release.commit"`
	HerokuReleaseCreationTimestamp ResourceAttributeConfig `mapstructure:"heroku.release.creation_timestamp"`
	ServiceInstanceID              ResourceAttributeConfig `mapstructure:"service.instance.id"`
	ServiceName                    ResourceAttributeConfig `mapstructure:"service.name"`
	ServiceVersion                 ResourceAttributeConfig `mapstructure:"service.version"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		CloudProvider: ResourceAttributeConfig{
			Enabled: true,
		},
		HerokuAppID: ResourceAttributeConfig{
			Enabled: true,
		},
		HerokuDynoID: ResourceAttributeConfig{
			Enabled: true,
		},
		HerokuReleaseCommit: ResourceAttributeConfig{
			Enabled: true,
		},
		HerokuReleaseCreationTimestamp: ResourceAttributeConfig{
			Enabled: true,
		},
		ServiceInstanceID: ResourceAttributeConfig{
			Enabled: true,
		},
		ServiceName: ResourceAttributeConfig{
			Enabled: true,
		},
		ServiceVersion: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}
