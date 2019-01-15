workflow "Build and test" {
  on = "push"
  resolves = ["Test"]
}

action "Test" {
  uses = "docker://golang:1.11"
  runs = "go"
  args = "test -v ./..."
}
