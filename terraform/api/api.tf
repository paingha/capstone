provider "aws" {
  profile = "default"
  region  = "us-east-2"
}

resource "aws_instance" "api" {
  ami           = "ami-2757f631"
  instance_type = "t2.micro"
}
