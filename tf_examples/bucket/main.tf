# Create an instance on AWS --> we need the provider for aws resources
provider "aws" {
  region  = "eu-central-1"
  profile = "playground"
}

resource "aws_s3_bucket" "mybucket" {
  bucket = "my-test-bucket-1727819191"
  acl    = "private"
}
