module "gke_cluster" {
  source         = "github.com/savkusamdetka23/tf-google-gke-cluster"
  GOOGLE_REGION  = var.GOOGLE_REGION
  GOOGLE_PROJECT = var.GOOGLE_PROJECT
  GKE_NUM_NODES  = var.GKE_NUM_NODES
}

resource "google_container_cluster" "demo" {
  name     = "demo-cluster"
  location = "us-central1-c"

  remove_default_node_pool = true
  initial_node_count       = 1

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