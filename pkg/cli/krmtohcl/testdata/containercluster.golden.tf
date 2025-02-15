resource "google_container_cluster" "twenty_namespaces" {
  addons_config {
    network_policy_config {
      disabled = true
    }
  }

  cluster_autoscaling {
    enabled = false
  }

  cluster_ipv4_cidr = "10.12.0.0/14"

  cluster_telemetry {
    type = "ENABLED"
  }

  database_encryption {
    state = "DECRYPTED"
  }

  location = "us-central1-c"

  master_auth {
    client_certificate_config {
      issue_client_certificate = false
    }
  }

  name    = "twenty-namespaces"
  network = "projects/my-project/global/networks/default"

  network_policy {
    enabled  = false
    provider = "PROVIDER_UNSPECIFIED"
  }

  networking_mode = "ROUTES"

  node_config {
    disk_size_gb = 100
    disk_type    = "pd-standard"
    image_type   = "COS"
    machine_type = "n1-standard-1"

    metadata = {
      disable-legacy-endpoints = "true"
    }

    oauth_scopes    = ["https://www.googleapis.com/auth/devstorage.read_only", "https://www.googleapis.com/auth/logging.write", "https://www.googleapis.com/auth/monitoring", "https://www.googleapis.com/auth/service.management.readonly", "https://www.googleapis.com/auth/servicecontrol", "https://www.googleapis.com/auth/trace.append"]
    service_account = "default"

    shielded_instance_config {
      enable_integrity_monitoring = true
    }

    workload_metadata_config {
      node_metadata = "GKE_METADATA_SERVER"
    }
  }

  node_version = "1.16.15-gke.1600"

  pod_security_policy_config {
    enabled = false
  }

  project = "my-project"

  release_channel {
    channel = "UNSPECIFIED"
  }

  subnetwork = "projects/my-project/regions/us-central1/subnetworks/default"

  workload_identity_config {
    workload_pool = "my-project.svc.id.goog"
  }
}
# terraform import google_container_cluster.twenty_namespaces my-project/us-central1-c/twenty-namespaces
