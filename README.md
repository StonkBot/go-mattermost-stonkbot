![Stonks](img/stonks.jpg)

This is a Mattermost bot which can react with emojis on posts. It is mainly used to
emphasize won deals so they stand out more.

## Installation and Usage

### Preparations

Add the stonks emoji as a custom emoji in Mattermost if wanted. Or use any other built-in or custom emoji you want.

Generate a personal access token in Mattermost. Refer to the [Mattermost documentation](https://docs.mattermost.com/developer/personal-access-tokens.html#contents).
Get your user id, either through a sysadmin or I dunno, just ask an admin :)

### Manual installation

Download the latest release from [GitHub Stonksbot releases](https://github.com/StonkBot/go-mattermost-stonkbot/releases/latest)
to a location you like. Then extract the binaries and the sample configuration for the archive.

You need the create a mattermost token. Then, configure your bot with this information. Add Channels and Emojis as you need.

Now you're ready to start your bot: `./go-mattermost-stonksbot`

### Installation from package

Download the package from [GitHub Stonksbot releases](https://github.com/StonkBot/go-mattermost-stonkbot/releases/latest)
in the desired format and install it.

Create the config file. An example is installed in `/etc/stonks/config.toml.example`.
Copy the example to one of the following paths and add your mattermost token and id.

* /etc/stonks/config.toml
* $HOME/.config/stonks/config.toml
* ./config.toml

Now start or enable the service

```bash
systemctl --user start stonkbot
```


## TODOs

* Make configurable stonks string
* Create packages: deb, rpm and aur

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

Login to your Mattermost instance at http://localhost:8065 with the user `bill`.


### Build packages

Install goreleaser and nfpm. Now build snapshot packages for dev testing

```
goreleaser --skip-publish --rm-dist --snapshot
```
