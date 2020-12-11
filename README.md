# cf-reverse-service-lookup-plugin

cf-cli plugin to perform reverse lookups against service instance GUIDs

## background

Pairs nicely with problematic BOSH service instances 👌👩‍🍳

This plugin aims to map between how BOSH references service deployments, e.g. `service-instance_GUID`, and how they're referenced in Cloud Foundry. It can be frustrating when tools like the bosh-cli and Prometheus/Alertmanager start yelling about a service instance barfing, but you have to spend time going through this whole riggamarole of looking up the service to see which org/space it belongs to, etc.

With this plugin, as long as you're logged into the same Cloud Foundry installation as the `service-instance_GUID` you're trying to find, this will go through the whole riggamarole of looking up the org/space of the service instance _for_ you.

## usage

```sh
# get details on a service instance by it's GUID alone
cf rsl -s 4c463943-d421-4f6f-8501-247fba95882d

# passing BOSH's default 'service-instance_' prefix is acceptable, too
cf rsl -s service-instance_4c463943-d421-4f6f-8501-247fba95882d

# you can pass '-s service-instance_GUID' multiple times
cf rsl -s service-instance_4c463943-d421-4f6f-8501-247fba95882d -s bbaa77df-52e7-4d6a-8c86-d07a7c93ab82

# optionally, multiple different presentation formats can be specified
cf rsl -s xyz --format plain-text (default)
cf rsl -s xyz --format table
cf rsl -s xyz --format json

# or both, why not
cf rsl -s xyz --format table --format json
```

`--format json`:

```json
[
  {
    "organization": {
      "name": "grundlework"
    },
    "service": {
      "guid": "d6bb8908-a8f8-46b9-9c21-3069cdb939ef",
      "name": "small-redis",
      "space_guid": "8a90a486-6a40-4709-b65f-efe6705bdc8f"
    },
    "space": {
      "name": "scratchpad",
      "organization_guid": "d63feced-41d1-44d7-ba9b-8d062975313b"
    }
  }
]
```

`--format table`:

```txt
+--------------------------------------+-------------+-------------+------------+
|             SERVICE GUID             |   SERVICE   |     ORG     |   SPACE    |
+--------------------------------------+-------------+-------------+------------+
| d6bb8908-a8f8-46b9-9c21-3069cdb939ef | small-redis | grundlework | scratchpad |
| a23c6626-2e42-484d-8664-0f6008316f09 | other-redis | otherorgxyz | 123xyspace |
+--------------------------------------+-------------+-------------+------------+
```

## installation

If you want to try it out, install it directly from [the github releases tab as follows](https://github.com/aegershman/cf-reverse-service-lookup-plugin/releases):

```sh
# osx 64bit
cf install-plugin -f https://github.com/aegershman/cf-reverse-service-lookup-plugin/releases/download/0.7.0/cf-reverse-service-lookup-plugin-darwin

# linux 64bit (32bit and ARM6 also available)
cf install-plugin -f https://github.com/aegershman/cf-reverse-service-lookup-plugin/releases/download/0.7.0/cf-reverse-service-lookup-plugin-amd64

# windows 64bit (32bit also available)
cf install-plugin -f https://github.com/aegershman/cf-reverse-service-lookup-plugin/releases/download/0.7.0/cf-reverse-service-lookup-plugin-windows-amd64.exe
```

## updating and releasing

- `go get -u all`
- `go mod tidy`
- update the plugin version in `main.go`
- update the `README` install-plugin section to reference the new upcoming release version
- `git tag 0.7.0` -- or whatever version, of course
- `git push origin --tags`
- `export GITHUB_TOKEN="xyzabc"`
- `goreleaser release`

## feedback welcome

Ideas, use-cases, opinions on architecture, doc, etc. all welcome
