variable "GOOGLE_PROJECT" {
  type        = string
  description = "GCP project name"
}

variable "GOOGLE_REGION" {
  type        = string
  default     = "us-central1-c"
  description = "GCP region name"
}

variable "GKE_MACHINE_TYPE" {
  type        = string
  default     = "g1-small"
  description = "machine type"
}

variable "GKE_NUM_NODES" {
  type        = number
  default     = 2
  description = "node pool count"
}

variable "TELE_TOKEN " {
  type        = string
  default     = 2
  description = "6723316123dd8:AAGiiWtqo8RNSasdv6CA2hRAX1eby621316XiXbc"
}