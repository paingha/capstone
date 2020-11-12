provider "aws" {
  profile = "default"
  region  = "us-east-2"
}

resource "aws_instance" "mailservice" {
  ami           = "ami-2757f631"
  instance_type = "t2.micro"
}
