output "node_ip" {
    value = [for node in data.data.google_compute_instance.nodes : node.network_interface[0].access_config[0].nat_ip]
}