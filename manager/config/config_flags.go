// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "pod-application"), DefaultConfig.PodApplication, "Application name for managed pods")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "pod-namespace"), DefaultConfig.PodNamespace, "Namespace to use for managing FlytePropeller pods")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "pod-template-container-name"), DefaultConfig.PodTemplateContainerName, "The container name within the K8s PodTemplate name used to set FlyteWorkflow CRD labels selectors")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "pod-template-name"), DefaultConfig.PodTemplateName, "K8s PodTemplate name to use for starting FlytePropeller pods")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "pod-template-namespace"), DefaultConfig.PodTemplateNamespace, "Namespace where the k8s PodTemplate is located")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "scan-interval"), DefaultConfig.ScanInterval.String(), "Frequency to scan FlytePropeller pods and start / restart if necessary")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "shard.type"), DefaultConfig.ShardConfig.Type, "Shard implementation to use")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "shard.shard-count"), DefaultConfig.ShardConfig.ShardCount, "The number of shards to manage for a 'hash' shard type")
	return cmdFlags
}
