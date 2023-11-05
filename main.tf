terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
    }
  }
}

provider "digitalocean" {
  token = var.tokenapi
}



resource "digitalocean_droplet" "crud" {
  image    = "docker-20-04"
  name     = "peliculasserver"
  region   = "nyc3"
  size     = "s-1vcpu-1gb"
  ssh_keys = [var.fingerprint_key]

  connection {
    host        = self.ipv4_address
    user        = "root"
    type        = "ssh"
    private_key = file(var.private_key)
    timeout     = "2m"
  }

  provisioner "remote-exec" {
    inline = [
    "sudo apt-get update",
    "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
    "sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"",
    "sudo apt-get update",
    "sudo apt-get install -y docker-ce",
    "docker --version",
    "sudo apt install -y git",
    "git clone https://github.com/EdwinCuenta/Semana5-Algoritmos-Paralelos.git",
    "cd Semana5-Algoritmos-Paralelos/conf-docker-bd",
    "curl -L https://github.com/docker/compose/releases/download/v2.22.0/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose",
    "chmod +x /usr/local/bin/docker-compose",
    "docker-compose --version",
    "docker network create mynetwork",
    "docker-compose up -d --build",
    "docker container ls",
    "cd ../aplicacion",
    "curl -L https://github.com/docker/compose/releases/download/v2.22.0/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose",
    "chmod +x /usr/local/bin/docker-compose",
    "docker-compose --version",
    "docker-compose up -d --build",
    "docker container ls",
    "sudo apt-get install -y postgresql-client",
    "sudo ufw allow 1339",
    "sudo ufw allow 443",
    "sudo ufw allow 5432",
    "sudo ufw allow 5433",
    "sudo ufw allow 8080"
  ]
}


}


