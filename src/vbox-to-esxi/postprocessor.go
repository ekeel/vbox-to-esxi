package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	// StagingDir string `mapstructure:"staging_directory"`
	VMName    string `mapstructure:"vm_name"`
	OutputDir string `mapstructure:"output_dir"`
	OVFDir    string `mapstructure:"ovf_dir"`
	OVFFIle   string `mapstructure:"ovf_file"`
	NAT       string `mapstructure:"nat"`
	Datastore string `mapstructure:"datastore"`

	ctx interpolate.Context
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)
	if err != nil {
		return err
	}

	// if p.config.StagingDir == "" {
	// 	p.config.StagingDir = "output-virtualbox-iso"
	// }

	if p.config.VMName == "" {
		return fmt.Errorf("vm_name is required")
	}

	if p.config.OutputDir == "" {
		p.config.OutputDir = "output-esxi/"
	}

	if p.config.OVFDir == "" && p.config.OVFFIle != "" {
		p.config.OVFDir = strings.Split(p.config.OVFFIle, "/")[0] + "/"
	} else {
		return fmt.Errorf("could not automatically determine ovf_dir. please provide ovf_dir")
	}

	if p.config.OVFFIle == "" {
		// p.config.OVFFIle = "output-virtualbox-iso/packer-ubuntu-18.04-amd64.ovf"
		return fmt.Errorf("ovf_file is required")
	}

	if p.config.NAT == "" {
		p.config.NAT = fmt.Sprintf("--net:\"NAT=%s\"", "N/A")
	}

	if p.config.Datastore == "" {
		p.config.Datastore = fmt.Sprintf("-ds=%s", "N/A")
	}

	// if _, err := os.Stat(p.config.StagingDir); err == nil {
	// 	return fmt.Errorf("staging directory '%s' already exists", p.config.StagingDir)
	// }

	return nil
}

func (p *PostProcessor) PostProcess(ctx context.Context, ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, bool, error) {
	if artifact.BuilderId() != "mitchellh.virtualbox" {
		err := fmt.Errorf("unknown artifact type: %s\ncan only import from virtualbox artifacts", artifact.BuilderId())
		return nil, false, false, err
	}

	err := os.MkdirAll(p.config.OutputDir, 0755)
	if err != nil {
		return nil, false, false, fmt.Errorf("failed to create output directory: %s", err)
	}

	err = ConvertOVFFile(p.config.OVFFIle, path.Join(p.config.OutputDir, p.config.VMName+".ovf"))
	if err != nil {
		return nil, false, false, fmt.Errorf("failed to update the ovf file: %s", err)
	}

	err = copy(path.Join(p.config.OVFDir, p.config.VMName+"-disk001.vmdk"), path.Join(p.config.OutputDir, p.config.VMName+"-disk001.vmdk"))
	if err != nil {
		return nil, false, false, fmt.Errorf("failed to copy the VMDK file: %s", err)
	}

	err = copy(path.Join(p.config.OVFDir, p.config.VMName+".ovf"), path.Join(p.config.OutputDir, p.config.VMName+".2.ovf"))
	if err != nil {
		return nil, false, false, fmt.Errorf("failed to copy the OVF2 file: %s", err)
	}

	return nil, true, true, nil
}

func copy(src, dst string) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed while opening source file %s: %s", src, err)
	}
	defer source.Close()

	dest, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to open destination file %s: %s", dst, err)
	}
	defer dest.Close()

	_, err = io.Copy(dest, source)
	return err
}
