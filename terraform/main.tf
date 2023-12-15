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