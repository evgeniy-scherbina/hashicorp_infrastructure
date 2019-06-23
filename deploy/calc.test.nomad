job "example" {
  datacenters = ["dc1"]

  type = "service"

  update {
    max_parallel = 1
    min_healthy_time = "10s"
    healthy_deadline = "3m"
    progress_deadline = "10m"
    auto_revert = false
    canary = 0
  }

  migrate {
    max_parallel = 1
    health_check = "checks"
    min_healthy_time = "10s"
    healthy_deadline = "5m"
  }

  group "addition" {
    count = 1

    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

//    ephemeral_disk {
//      size = 5
//    }

    task "addition" {
      driver = "docker"

      config {
        image = "scherbina/hashicorp-infrastructure-services:v0.0.1"
        port_map {
          grpc = 9090
        }

        args = [
          "addition-service"
        ]
      }

      resources {
        cpu    = 500 # 500 MHz
        memory = 256 # 256MB
        network {
          mbits = 10
          port "grpc" {}
        }
      }
//
//      service {
//        name = "redis-cache"
//        tags = ["global", "cache"]
//        port = "db"
//        check {
//          name     = "alive"
//          type     = "tcp"
//          interval = "10s"
//          timeout  = "2s"
//        }
//      }
    }
  }
}