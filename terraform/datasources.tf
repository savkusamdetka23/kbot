data "goog_compute_instance_group" "node_instance_groups" {
  self_link = google_container_cluster.demo.node_pool[0].managed_instance_groups_urls[0]
}

data "google_compute_instance" "nodes" {
  for_each  = toset(data.goog_compute_instance_group.node_instance_groups.instances[*])
  self_link = each.key
}