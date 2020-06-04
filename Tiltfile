k8s_yaml('tilt-deployment.yaml')

custom_build(
'telliottio/envserver',
('make build'),
[
    'cmd',
    'Dockerfile',
    'Makefile',
    'go.mod'
],
tag="tilt"
)

k8s_resource('envserver', port_forwards='8080:8080')