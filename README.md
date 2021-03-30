## Dev setup

Start a mattermost container and create a test team as well as two users (one bot and one for you)


```
podman run -dt -p 8065:8065 mattermost/mattermost-preview

podman exec ${container_name} mattermost --conf=/mm/mattermost/config/config_docker.json team create --name botsample --display_name "Sample Bot playground" --email "admin@example.com"

podman exec ${container_name} mattermost --conf=/mm/mattermost/config/config_docker.json user create --email="bot@example.com" --password="Password1!" --username="samplebot"
podman exec ${container_name} mattermost --conf=/mm/mattermost/config/config_docker.json user create --email="bill@example.com" --password="Password1!" --username="bill"
podman exec ${container_name} mattermost --conf=/mm/mattermost/config/config_docker.json roles system_admin bill
podman exec ${container_name} mattermost --conf=/mm/mattermost/config/config_docker.json team add botsample samplebot bill
podman exec ${container_name} mattermost --conf=/mm/mattermost/config/config_docker.json user verify samplebot
```