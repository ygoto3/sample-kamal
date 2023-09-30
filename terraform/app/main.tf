provider "digitalocean" {
  token = "${var.do_token}"
}

data "digitalocean_ssh_key" "ssh" {
  name = "${var.ssh_key_name}"
}

module "module_kamal" {
  source = "../modules/digitalocean"

  ssh_key_id = "${data.digitalocean_ssh_key.ssh.id}"
  project_name = "${var.project_name}"
}
