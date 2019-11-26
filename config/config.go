package config

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func init() {

}

type Config struct {
	KingctlConf KingctlConf `yaml:"config"`
}

type KingctlConf struct {
	Server    SERVER    `yaml:"server"`
	Cmd       CMD       `yaml:"cmd"`
	CloudInfo CloudINFO `yaml:"cloudInfo"`
}

type SERVER struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type CMD struct {
	StartCmd string `yaml:"startCmd"`
	StopCmd  string `yaml:"stopCmd"`
	SavePath string `yaml:"savePath"`
	WorkPath string `yaml:"workPath"`
}

type CloudINFO struct {
	Placement           PlacementInfo `yaml:"placementInfo"`
	DataDisk            string        `yaml:"dataDisk"`
	ImageID             string        `yaml:"imageID"`
	ImageName           string        `yaml:"imageName"`
	SystemDisk          SystemDisk    `yaml:"systemDisk"`
	EnhancedService     string        `yaml:"enhancedService"`
	HostName            string        `yaml:"hostName"`
	UserData            string        `yaml:"userData"`
	VirtualPrivateCloud string        `yaml:"virtualPrivateCloud"`
	InternetAccessible  string        `yaml:"internetAccessible"`
	InstanceCount       int           `yaml:"instanceCount"`
	InstanceName        string        `yaml:"instanceName"`
	SecurityGroupID     string        `yaml:"securityGroupID"`
}

type PlacementInfo struct {
	Zone      string `yaml:"zone"`
	ProjectID string `yaml:"projectId"`
}

type SystemDisk struct {
	DiskType string `yaml:"diskType"`
	DiskSize string `yaml:"diskSize"`
}

func ReadConfigFile(path string) ([]byte, error) {
	conf, err := readConfigFile(path)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	byts, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	// fmt.Println(string(byts))
	return byts, nil
}

func readConfigFile(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil

}
