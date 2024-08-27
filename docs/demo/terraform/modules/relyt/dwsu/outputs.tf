output "dwsu_id" {
  value = relyt_dwsu.dwsu.id
}

output "host" {
  value = relyt_dwsu.dwsu.endpoints[1].host
}