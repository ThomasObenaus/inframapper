package tfstate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var invalidData = `{;[}]}`

var emptyData = `{}`

var emptyStateData = `{
	"version": 3,
	"terraform_version": "0.11.7",
	"serial": 3,
	"lineage": "3e6f20dc-3dfa-b8df-882b-1ccbbfe9c46c",
	"modules": [
			{
					"path": [
							"root"
					],
					"outputs": {},
					"resources": {},
					"depends_on": []
			}
	]
}`

var instanceAndSgState = `{
	"version": 3,
	"terraform_version": "0.11.7",
	"serial": 2,
	"lineage": "3e6f20dc-3dfa-b8df-882b-1ccbbfe9c46c",
	"modules": [
			{
					"path": [
							"root"
					],
					"outputs": {},
					"resources": {
							"aws_instance.hello_world": {
									"type": "aws_instance",
									"depends_on": [
											"aws_security_group.sg_hwld"
									],
									"primary": {
											"id": "i-07914cc8cd9ca4825",
											"attributes": {
													"ami": "ami-7c412f13",
													"associate_public_ip_address": "true",
													"availability_zone": "eu-central-1b",
													"credit_specification.#": "1",
													"credit_specification.0.cpu_credits": "standard",
													"disable_api_termination": "false",
													"ebs_block_device.#": "0",
													"ebs_optimized": "false",
													"ephemeral_block_device.#": "0",
													"get_password_data": "false",
													"iam_instance_profile": "",
													"id": "i-07914cc8cd9ca4825",
													"instance_state": "running",
													"instance_type": "t2.micro",
													"ipv6_addresses.#": "0",
													"key_name": "",
													"monitoring": "false",
													"network_interface.#": "0",
													"network_interface_id": "eni-6e47ea43",
													"password_data": "",
													"placement_group": "",
													"primary_network_interface_id": "eni-6e47ea43",
													"private_dns": "ip-172-31-78-180.eu-central-1.compute.internal",
													"private_ip": "172.31.78.180",
													"public_dns": "ec2-52-59-187-225.eu-central-1.compute.amazonaws.com",
													"public_ip": "52.59.187.225",
													"root_block_device.#": "1",
													"root_block_device.0.delete_on_termination": "true",
													"root_block_device.0.iops": "100",
													"root_block_device.0.volume_id": "vol-02e3157ba6fa17914",
													"root_block_device.0.volume_size": "8",
													"root_block_device.0.volume_type": "gp2",
													"security_groups.#": "1",
													"security_groups.22753008": "Test-SG",
													"source_dest_check": "true",
													"subnet_id": "subnet-84dc4ef9",
													"tags.%": "0",
													"tenancy": "default",
													"volume_tags.%": "0",
													"vpc_security_group_ids.#": "1",
													"vpc_security_group_ids.1527081739": "sg-20f1484c"
											},
											"meta": {
													"e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
															"create": 600000000000,
															"delete": 1200000000000,
															"update": 600000000000
													},
													"schema_version": "1"
											},
											"tainted": false
									},
									"deposed": [],
									"provider": "provider.aws"
							},
							"aws_security_group.sg_hwld": {
									"type": "aws_security_group",
									"depends_on": [
											"data.aws_vpc.default"
									],
									"primary": {
											"id": "sg-20f1484c",
											"attributes": {
													"arn": "arn:aws:ec2:eu-central-1:307557990628:security-group/sg-20f1484c",
													"description": "Managed by Terraform",
													"egress.#": "0",
													"id": "sg-20f1484c",
													"ingress.#": "0",
													"name": "Test-SG",
													"owner_id": "307557990628",
													"revoke_rules_on_delete": "false",
													"tags.%": "0",
													"vpc_id": "vpc-ff5fec97"
											},
											"meta": {
													"e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
															"create": 600000000000,
															"delete": 600000000000
													},
													"schema_version": "1"
											},
											"tainted": false
									},
									"deposed": [],
									"provider": "provider.aws"
							},
							"data.aws_vpc.default": {
									"type": "aws_vpc",
									"depends_on": [],
									"primary": {
											"id": "vpc-ff5fec97",
											"attributes": {
													"cidr_block": "172.31.0.0/16",
													"cidr_block_associations.#": "1",
													"cidr_block_associations.0.association_id": "vpc-cidr-assoc-fb4be992",
													"cidr_block_associations.0.cidr_block": "172.31.0.0/16",
													"cidr_block_associations.0.state": "associated",
													"default": "true",
													"dhcp_options_id": "dopt-d3be72bb",
													"enable_dns_hostnames": "true",
													"enable_dns_support": "true",
													"id": "vpc-ff5fec97",
													"instance_tenancy": "default",
													"state": "available",
													"tags.%": "1",
													"tags.Name": "default"
											},
											"meta": {},
											"tainted": false
									},
									"deposed": [],
									"provider": "provider.aws"
							}
					},
					"depends_on": []
			}
	]
}
`

func TestParse_Fail(t *testing.T) {
	tfstate, err := Parse([]byte(invalidData))
	if err == nil {
		assert.Fail(t, "Parse was expected to fail")
	}
	if tfstate != nil {
		assert.Fail(t, "Parse was expected to fail")
	}
}

func TestParse(t *testing.T) {
	tfstate, err := Parse([]byte(instanceAndSgState))
	if err != nil {
		assert.Failf(t, "Parse failed: %s", err.Error())
	}
	require.NotNil(t, tfstate, "tfstate is nil")

	// modules
	require.NotEmpty(t, tfstate.Modules, "No modules found")

	module := tfstate.Modules[0]
	resources := module.Resources

	// modules
	require.Len(t, resources, 3, "Expected 3 resources")

	// instance
	instance := resources["aws_instance.hello_world"]
	require.NotNil(t, instance)
	assert.Equal(t, "aws_instance", instance.Type)
	require.NotNil(t, instance.Primary)
	assert.Equal(t, "i-07914cc8cd9ca4825", instance.Primary.ID)
}
