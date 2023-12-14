terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.52.0"
    }
  }
}

variable "GOOGLE_PROJECT" {
  type        = string
  description = "GCP project name"
}

variable "GOOGLE_REGION" {
  type        = string
  default     = "us-central1-c"
  description = "GCP region name"
}

provider "google" {
  project = var.GOOGLE_PROJECT
  region  = var.GOOGLE_REGION
}

variable "GKE_MACHINE_TYPE" {
  type        = string
  default     = "g1-small"
  description = "machine type"
}

variable "GKE_NUM_NODES" {
  type        = number
  default     = 2
  description = "node poll"

}

resource "google_container_cluster" "demo" {
  name     = "demo-cluster"
  location = var.GOOGLE_REGION

  remove_default_node_pool = true
  initial_node_count       = 1

  workload_identity_config {
    workload_pool = "${var.GOOGLE_PROJECT}.svc.id.goog"
  }

  node_config {
    workload_metadata_config {
      mode = "GKE_METADATA"
    }
  }
}
resource "google_container_node_pool" "demo" {
  name       = "main"
  project    = google_container_cluster.demo.project
  cluster    = google_container_cluster.demo.name
  location   = google_container_cluster.demo.location
  node_count = var.GKE_NUM_NODES


  node_config {
    machine_type = var.GKE_MACHINE_TYPE
  }
}