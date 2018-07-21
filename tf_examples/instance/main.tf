# Create an instance on AWS --> we need the provider for aws resources
provider "aws" {
  region  = "eu-central-1"
  profile = "playground"
}

resource "aws_instance" "hello_world" {
  ami                    = "ami-7c412f13"                       # Ubuntu Server 16.04 LTS (HVM)
  instance_type          = "t2.micro"
  vpc_security_group_ids = ["${aws_security_group.sg_hwld.id}"]
}

data "aws_vpc" "default" {
  default = true
}

resource "aws_security_group" "sg_hwld" {
  vpc_id = "${data.aws_vpc.default.id}"
  name   = "Test-SG"
}
