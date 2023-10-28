load('ext://restart_process', 'docker_build_with_restart')

local_resource(
  'backend-compile',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app ./',
  dir="./backend",
  deps=['./backend/main.go'])

docker_build_with_restart(
  'backend',
  './backend',
  entrypoint="/app/bin/app",
  only=[
    './bin',
  ],
  live_update=[
    sync('./backend/bin', '/app/bin'),
  ],
)

k8s_yaml('k8s/backend.yaml')
k8s_resource('backend', port_forwards=3001,
             resource_deps=['backend-compile'])

docker_build("tempo", "./tempo")
k8s_yaml('k8s/tempo.yaml')
k8s_resource('tempo', port_forwards='3200:3200')

docker_build("prometheus", "./prometheus")
k8s_yaml('k8s/prometheus.yaml')
k8s_resource('prometheus', port_forwards=9090)

# k8s_yaml('k8s/jaeger.yaml')
# k8s_resource('jaeger', port_forwards='16686:16686')
