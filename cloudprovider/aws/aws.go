package aws

import (
	"bytes"
	"fmt"

	"github.com/go-ini/ini"
	"github.com/rancher/types/apis/management.cattle.io/v3"
)

type CloudProvider struct {
	Config *v3.AWSCloudProvider
	Name   string
}

const (
	AWSCloudProviderName = "aws"
)

func GetInstance() *CloudProvider {
	return &CloudProvider{}
}

func (p *CloudProvider) Init(cloudProviderConfig v3.CloudProvider) error {
	if cloudProviderConfig.AWSCloudProvider == nil {
		return fmt.Errorf("AWS Cloud Provider Config is empty")
	}
	p.Name = AWSCloudProviderName
	p.Config = cloudProviderConfig.AWSCloudProvider
	return nil
}

func (p *CloudProvider) GetName() string {
	return p.Name
}

func (p *CloudProvider) GenerateCloudConfigFile() (string, error) {
	iniFile := ini.Empty()
	section, err := iniFile.NewSection("Global")
	if err != nil {
		return "", err
	}
	if err = section.ReflectFrom(p.Config); err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	_, err = iniFile.WriteTo(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
