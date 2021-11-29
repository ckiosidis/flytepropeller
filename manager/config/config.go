package config

import (
	"time"

	"github.com/flyteorg/flytestdlib/config"
)

//go:generate pflags Config --default-var=DefaultConfig

var (
	DefaultConfig = &Config{
		PodApplication:           "flytepropeller",
		PodNamespace:             "flyte",
		PodTemplateContainerName: "flytepropeller",
		PodTemplateName:          "flytepropeller-template",
		PodTemplateNamespace:     "flyte",
		ScanInterval: config.Duration{
			Duration: 10 * time.Second,
		},
		ShardConfig: ShardConfig{
			Type:       "hash",
			ShardCount: 3,
		},
	}

	configSection = config.MustRegisterSection("manager", DefaultConfig)
)

type ShardType = string

const (
	DomainShardType  ShardType = "domain"
	ProjectShardType ShardType = "project"
	HashShardType    ShardType = "hash"
)

// Configuration for defining shard replicas when using project or domain shard types
type PerShardMappingsConfig struct {
	IDs []string `json:"ids" pflag:",The list of ids to be managed"`
}

// Configuration for the FlytePropeller sharding strategy
type ShardConfig struct {
	Type             ShardType                `json:"type" pflag:"\"hash\",Shard implementation to use"`
	PerShardMappings []PerShardMappingsConfig `json:"per-shard-mapping" pflag:"-"`
	ShardCount       int                      `json:"shard-count" pflag:"\"3\",The number of shards to manage for a 'hash' shard type"`
}

// Configuration for the FlytePropeller Manager instance
type Config struct {
	PodApplication           string          `json:"pod-application" pflag:"\"flytepropeller\",Application name for managed pods"`
	PodNamespace             string          `json:"pod-namespace" pflag:"\"flyte\",Namespace to use for managing FlytePropeller pods"`
	PodTemplateContainerName string          `json:"pod-template-container-name" pflag:"\"flytepropeller\",The container name within the K8s PodTemplate name used to set FlyteWorkflow CRD labels selectors"`
	PodTemplateName          string          `json:"pod-template-name" pflag:"\"flytepropeller-template\",K8s PodTemplate name to use for starting FlytePropeller pods"`
	PodTemplateNamespace     string          `json:"pod-template-namespace" pflag:"\"flyte\",Namespace where the k8s PodTemplate is located"`
	ScanInterval             config.Duration `json:"scan-interval" pflag:"\"10s\",Frequency to scan FlytePropeller pods and start / restart if necessary"`
	ShardConfig              ShardConfig     `json:"shard" pflag:",Configure the shard strategy for this manager"`
}

func GetConfig() *Config {
	return configSection.GetConfig().(*Config)
}
