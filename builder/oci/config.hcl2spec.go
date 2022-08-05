// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package oci

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string                `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string                `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion         *string                `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug               *bool                  `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool                  `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string                `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string      `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Type                      *string                `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect        *string                `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                   *string                `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                   *int                   `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername               *string                `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword               *string                `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName            *string                `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string                `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType   *string                `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits   *int                   `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                []string               `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys    *bool                  `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos               []string               `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile         *string                `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile        *string                `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                    *bool                  `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                *string                `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout            *string                `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth              *bool                  `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool                  `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int                   `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost            *string                `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort            *int                   `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool                  `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string                `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword        *string                `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive     *bool                  `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string                `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile *string                `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod     *string                `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost              *string                `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort              *int                   `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername          *string                `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword          *string                `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string                `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string                `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string               `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string               `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey              []byte                 `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey             []byte                 `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                 *string                `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword             *string                `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                 *string                `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy              *bool                  `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                 *int                   `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout              *string                `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL               *bool                  `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure             *bool                  `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM              *bool                  `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	InstancePrincipals        *bool                  `mapstructure:"use_instance_principals" cty:"use_instance_principals" hcl:"use_instance_principals"`
	SkipCreateImage           *bool                  `mapstructure:"skip_create_image" required:"false" cty:"skip_create_image" hcl:"skip_create_image"`
	AccessCfgFile             *string                `mapstructure:"access_cfg_file" cty:"access_cfg_file" hcl:"access_cfg_file"`
	AccessCfgFileAccount      *string                `mapstructure:"access_cfg_file_account" cty:"access_cfg_file_account" hcl:"access_cfg_file_account"`
	UserID                    *string                `mapstructure:"user_ocid" cty:"user_ocid" hcl:"user_ocid"`
	TenancyID                 *string                `mapstructure:"tenancy_ocid" cty:"tenancy_ocid" hcl:"tenancy_ocid"`
	Region                    *string                `mapstructure:"region" cty:"region" hcl:"region"`
	Fingerprint               *string                `mapstructure:"fingerprint" cty:"fingerprint" hcl:"fingerprint"`
	KeyFile                   *string                `mapstructure:"key_file" cty:"key_file" hcl:"key_file"`
	PassPhrase                *string                `mapstructure:"pass_phrase" cty:"pass_phrase" hcl:"pass_phrase"`
	UsePrivateIP              *bool                  `mapstructure:"use_private_ip" cty:"use_private_ip" hcl:"use_private_ip"`
	SecurityTokenFilePath     *string                `mapstructure:"security_token_file" cty:"security_token_file" hcl:"security_token_file"`
	AvailabilityDomain        *string                `mapstructure:"availability_domain" cty:"availability_domain" hcl:"availability_domain"`
	CompartmentID             *string                `mapstructure:"compartment_ocid" cty:"compartment_ocid" hcl:"compartment_ocid"`
	BaseImageID               *string                `mapstructure:"base_image_ocid" cty:"base_image_ocid" hcl:"base_image_ocid"`
	BaseImageFilter           *FlatListImagesRequest `mapstructure:"base_image_filter" cty:"base_image_filter" hcl:"base_image_filter"`
	ImageName                 *string                `mapstructure:"image_name" cty:"image_name" hcl:"image_name"`
	ImageCompartmentID        *string                `mapstructure:"image_compartment_ocid" cty:"image_compartment_ocid" hcl:"image_compartment_ocid"`
	LaunchMode                *string                `mapstructure:"image_launch_mode" cty:"image_launch_mode" hcl:"image_launch_mode"`
	NicAttachmentType         *string                `mapstructure:"nic_attachment_type" cty:"nic_attachment_type" hcl:"nic_attachment_type"`
	InstanceName              *string                `mapstructure:"instance_name" cty:"instance_name" hcl:"instance_name"`
	InstanceTags              map[string]string      `mapstructure:"instance_tags" cty:"instance_tags" hcl:"instance_tags"`
	InstanceDefinedTagsJson   *string                `mapstructure:"instance_defined_tags_json" required:"false" cty:"instance_defined_tags_json" hcl:"instance_defined_tags_json"`
	Shape                     *string                `mapstructure:"shape" cty:"shape" hcl:"shape"`
	ShapeConfig               *FlatFlexShapeConfig   `mapstructure:"shape_config" cty:"shape_config" hcl:"shape_config"`
	BootVolumeSizeInGBs       *int64                 `mapstructure:"disk_size" cty:"disk_size" hcl:"disk_size"`
	Metadata                  map[string]string      `mapstructure:"metadata" cty:"metadata" hcl:"metadata"`
	UserData                  *string                `mapstructure:"user_data" cty:"user_data" hcl:"user_data"`
	UserDataFile              *string                `mapstructure:"user_data_file" cty:"user_data_file" hcl:"user_data_file"`
	SubnetID                  *string                `mapstructure:"subnet_ocid" cty:"subnet_ocid" hcl:"subnet_ocid"`
	CreateVnicDetails         *FlatCreateVNICDetails `mapstructure:"create_vnic_details" cty:"create_vnic_details" hcl:"create_vnic_details"`
	Tags                      map[string]string      `mapstructure:"tags" cty:"tags" hcl:"tags"`
	DefinedTagsJson           *string                `mapstructure:"defined_tags_json" required:"false" cty:"defined_tags_json" hcl:"defined_tags_json"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":          &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":      &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":      &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                  &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":  &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":         &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file": &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":               &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"use_instance_principals":      &hcldec.AttrSpec{Name: "use_instance_principals", Type: cty.Bool, Required: false},
		"skip_create_image":            &hcldec.AttrSpec{Name: "skip_create_image", Type: cty.Bool, Required: false},
		"access_cfg_file":              &hcldec.AttrSpec{Name: "access_cfg_file", Type: cty.String, Required: false},
		"access_cfg_file_account":      &hcldec.AttrSpec{Name: "access_cfg_file_account", Type: cty.String, Required: false},
		"user_ocid":                    &hcldec.AttrSpec{Name: "user_ocid", Type: cty.String, Required: false},
		"tenancy_ocid":                 &hcldec.AttrSpec{Name: "tenancy_ocid", Type: cty.String, Required: false},
		"region":                       &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"fingerprint":                  &hcldec.AttrSpec{Name: "fingerprint", Type: cty.String, Required: false},
		"key_file":                     &hcldec.AttrSpec{Name: "key_file", Type: cty.String, Required: false},
		"pass_phrase":                  &hcldec.AttrSpec{Name: "pass_phrase", Type: cty.String, Required: false},
		"use_private_ip":               &hcldec.AttrSpec{Name: "use_private_ip", Type: cty.Bool, Required: false},
		"security_token_file":          &hcldec.AttrSpec{Name: "security_token_file", Type: cty.String, Required: false},
		"availability_domain":          &hcldec.AttrSpec{Name: "availability_domain", Type: cty.String, Required: false},
		"compartment_ocid":             &hcldec.AttrSpec{Name: "compartment_ocid", Type: cty.String, Required: false},
		"base_image_ocid":              &hcldec.AttrSpec{Name: "base_image_ocid", Type: cty.String, Required: false},
		"base_image_filter":            &hcldec.BlockSpec{TypeName: "base_image_filter", Nested: hcldec.ObjectSpec((*FlatListImagesRequest)(nil).HCL2Spec())},
		"image_name":                   &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_compartment_ocid":       &hcldec.AttrSpec{Name: "image_compartment_ocid", Type: cty.String, Required: false},
		"image_launch_mode":            &hcldec.AttrSpec{Name: "image_launch_mode", Type: cty.String, Required: false},
		"nic_attachment_type":          &hcldec.AttrSpec{Name: "nic_attachment_type", Type: cty.String, Required: false},
		"instance_name":                &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"instance_tags":                &hcldec.AttrSpec{Name: "instance_tags", Type: cty.Map(cty.String), Required: false},
		"instance_defined_tags_json":   &hcldec.AttrSpec{Name: "instance_defined_tags_json", Type: cty.String, Required: false},
		"shape":                        &hcldec.AttrSpec{Name: "shape", Type: cty.String, Required: false},
		"shape_config":                 &hcldec.BlockSpec{TypeName: "shape_config", Nested: hcldec.ObjectSpec((*FlatFlexShapeConfig)(nil).HCL2Spec())},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"metadata":                     &hcldec.AttrSpec{Name: "metadata", Type: cty.Map(cty.String), Required: false},
		"user_data":                    &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":               &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"subnet_ocid":                  &hcldec.AttrSpec{Name: "subnet_ocid", Type: cty.String, Required: false},
		"create_vnic_details":          &hcldec.BlockSpec{TypeName: "create_vnic_details", Nested: hcldec.ObjectSpec((*FlatCreateVNICDetails)(nil).HCL2Spec())},
		"tags":                         &hcldec.AttrSpec{Name: "tags", Type: cty.Map(cty.String), Required: false},
		"defined_tags_json":            &hcldec.AttrSpec{Name: "defined_tags_json", Type: cty.String, Required: false},
	}
	return s
}

// FlatCreateVNICDetails is an auto-generated flat version of CreateVNICDetails.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatCreateVNICDetails struct {
	AssignPublicIp      *bool             `mapstructure:"assign_public_ip" required:"false" cty:"assign_public_ip" hcl:"assign_public_ip"`
	DefinedTagsJson     *string           `mapstructure:"defined_tags_json" required:"false" cty:"defined_tags_json" hcl:"defined_tags_json"`
	DisplayName         *string           `mapstructure:"display_name" required:"false" cty:"display_name" hcl:"display_name"`
	FreeformTags        map[string]string `mapstructure:"tags" required:"false" cty:"tags" hcl:"tags"`
	HostnameLabel       *string           `mapstructure:"hostname_label" required:"false" cty:"hostname_label" hcl:"hostname_label"`
	NsgIds              []string          `mapstructure:"nsg_ids" required:"false" cty:"nsg_ids" hcl:"nsg_ids"`
	PrivateIp           *string           `mapstructure:"private_ip" required:"false" cty:"private_ip" hcl:"private_ip"`
	SkipSourceDestCheck *bool             `mapstructure:"skip_source_dest_check" required:"false" cty:"skip_source_dest_check" hcl:"skip_source_dest_check"`
	SubnetId            *string           `mapstructure:"subnet_id" required:"false" cty:"subnet_id" hcl:"subnet_id"`
}

// FlatMapstructure returns a new FlatCreateVNICDetails.
// FlatCreateVNICDetails is an auto-generated flat version of CreateVNICDetails.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*CreateVNICDetails) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatCreateVNICDetails)
}

// HCL2Spec returns the hcl spec of a CreateVNICDetails.
// This spec is used by HCL to read the fields of CreateVNICDetails.
// The decoded values from this spec will then be applied to a FlatCreateVNICDetails.
func (*FlatCreateVNICDetails) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"assign_public_ip":       &hcldec.AttrSpec{Name: "assign_public_ip", Type: cty.Bool, Required: false},
		"defined_tags_json":      &hcldec.AttrSpec{Name: "defined_tags_json", Type: cty.String, Required: false},
		"display_name":           &hcldec.AttrSpec{Name: "display_name", Type: cty.String, Required: false},
		"tags":                   &hcldec.AttrSpec{Name: "tags", Type: cty.Map(cty.String), Required: false},
		"hostname_label":         &hcldec.AttrSpec{Name: "hostname_label", Type: cty.String, Required: false},
		"nsg_ids":                &hcldec.AttrSpec{Name: "nsg_ids", Type: cty.List(cty.String), Required: false},
		"private_ip":             &hcldec.AttrSpec{Name: "private_ip", Type: cty.String, Required: false},
		"skip_source_dest_check": &hcldec.AttrSpec{Name: "skip_source_dest_check", Type: cty.Bool, Required: false},
		"subnet_id":              &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
	}
	return s
}

// FlatFlexShapeConfig is an auto-generated flat version of FlexShapeConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatFlexShapeConfig struct {
	Ocpus       *float32 `mapstructure:"ocpus" required:"false" cty:"ocpus" hcl:"ocpus"`
	MemoryInGBs *float32 `mapstructure:"memory_in_gbs" required:"false" cty:"memory_in_gbs" hcl:"memory_in_gbs"`
}

// FlatMapstructure returns a new FlatFlexShapeConfig.
// FlatFlexShapeConfig is an auto-generated flat version of FlexShapeConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*FlexShapeConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatFlexShapeConfig)
}

// HCL2Spec returns the hcl spec of a FlexShapeConfig.
// This spec is used by HCL to read the fields of FlexShapeConfig.
// The decoded values from this spec will then be applied to a FlatFlexShapeConfig.
func (*FlatFlexShapeConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"ocpus":         &hcldec.AttrSpec{Name: "ocpus", Type: cty.Number, Required: false},
		"memory_in_gbs": &hcldec.AttrSpec{Name: "memory_in_gbs", Type: cty.Number, Required: false},
	}
	return s
}

// FlatListImagesRequest is an auto-generated flat version of ListImagesRequest.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatListImagesRequest struct {
	CompartmentId          *string `mapstructure:"compartment_id" cty:"compartment_id" hcl:"compartment_id"`
	DisplayName            *string `mapstructure:"display_name" cty:"display_name" hcl:"display_name"`
	DisplayNameSearch      *string `mapstructure:"display_name_search" cty:"display_name_search" hcl:"display_name_search"`
	OperatingSystem        *string `mapstructure:"operating_system" cty:"operating_system" hcl:"operating_system"`
	OperatingSystemVersion *string `mapstructure:"operating_system_version" cty:"operating_system_version" hcl:"operating_system_version"`
	Shape                  *string `mapstructure:"shape" cty:"shape" hcl:"shape"`
}

// FlatMapstructure returns a new FlatListImagesRequest.
// FlatListImagesRequest is an auto-generated flat version of ListImagesRequest.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ListImagesRequest) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatListImagesRequest)
}

// HCL2Spec returns the hcl spec of a ListImagesRequest.
// This spec is used by HCL to read the fields of ListImagesRequest.
// The decoded values from this spec will then be applied to a FlatListImagesRequest.
func (*FlatListImagesRequest) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"compartment_id":           &hcldec.AttrSpec{Name: "compartment_id", Type: cty.String, Required: false},
		"display_name":             &hcldec.AttrSpec{Name: "display_name", Type: cty.String, Required: false},
		"display_name_search":      &hcldec.AttrSpec{Name: "display_name_search", Type: cty.String, Required: false},
		"operating_system":         &hcldec.AttrSpec{Name: "operating_system", Type: cty.String, Required: false},
		"operating_system_version": &hcldec.AttrSpec{Name: "operating_system_version", Type: cty.String, Required: false},
		"shape":                    &hcldec.AttrSpec{Name: "shape", Type: cty.String, Required: false},
	}
	return s
}
