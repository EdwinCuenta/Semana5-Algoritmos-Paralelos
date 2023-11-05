


//token y fingerprint
variable "fingerprint_key" {
    type        = string
    description = "key generada con ssh local"
    default     = "6c:ee:fd:83:10:64:2d:01:c4:3b:91:68:f5:f6:dc:77"
}

variable "tokenapi" {
    type        = string
    description = "token de digitalOcean"
    default     = "dop_v1_a8d460095ccd18af23df28b88bf67f0a1fce8f079a00c928087dde22f2a87c0b"
}


//keys
variable "public_key" {
    type = string
    description = "Public Key"
    default = "C:\\Users\\EDWIN-PC\\.ssh\\id_rsa.pub"
}
variable "private_key" {
    type = string
    description = "Private Key"
    default =  "C:\\Users\\EDWIN-PC\\.ssh\\id_rsa"
}