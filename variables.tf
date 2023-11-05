


//token y fingerprint
variable "fingerprint_key" {
    type        = string
    description = "key generada con ssh local"
}

variable "tokenapi" {
    type        = string
    description = "token de digitalOcean"
}


//keys
variable "public_key" {
    type = string
    description = "Public Key"
    
}
variable "private_key" {
    type = string
    description = "Private Key"
}