# cf-reverse-service-lookup-plugin

cf-cli plugin to perform reverse lookups against service instance GUIDs

## background

Pairs nicely with problematic BOSH service instances 👌👩‍🍳

This plugin aims to map between how BOSH references service deployments, e.g. `service-instance_GUID`, and how they're referenced in Cloud Foundry. It can be frustrating when tools like the bosh-cli and Prometheus/Alertmanager start yelling about a service instance barfing, but you have to spend time going through this whole riggamarole of looking up the service to see which org/space it belongs to, etc.

With this plugin, as long as you're logged into the same Cloud Foundry installation as the `service-instance_GUID` you're trying to find, this will go through the whole riggamarole of looking up the org/space of the service instance _for_ you.

## usage

```sh
# get details on a service instance by it's GUID alone
cf reverse-service-lookup --service-guid 67e9e04e-8cc5-4744-8a5a-eb0a2d21c7ee

# passing BOSH's default 'service-instance_' prefix is acceptable, too
cf reverse-service-lookup --service-guid service-instance_67e9e04e-8cc5-4744-8a5a-eb0a2d21c7ee

# optionally, different presentation formats can be specified
cf reverse-service-lookup --service-guid xyz --format json (default)
cf reverse-service-lookup --service-guid xyz --format table
```

`--format json`:

```json
{
  "organization": {
    "name": "gershman"
  },
  "service": {
    "name": "tiny-turtle-sql",
    "space_guid": "4ba83909-cde1-448e-873f-d53b27391604"
  },
  "space": {
    "name": "dev",
    "organization_guid": "a297887a-f58f-4266-9ba2-186563eec13a"
  }
}
```

`--format table`:

```txt
+----------+-------+-----------------+
|   ORG    | SPACE |     SERVICE     |
+----------+-------+-----------------+
| gershman | dev   | tiny-turtle-sql |
+----------+-------+-----------------+
```

## installation

If you want to try it out, install it directly from [the github releases tab as follows](https://github.com/aegershman/cf-reverse-service-lookup-plugin/releases):

```sh
# osx 64bit
cf install-plugin -f https://github.com/aegershman/cf-reverse-service-lookup-plugin/releases/download/0.1.0/cf-reverse-service-lookup-plugin-darwin

# linux 64bit (32bit and ARM6 also available)
cf install-plugin -f https://github.com/aegershman/cf-reverse-service-lookup-plugin/releases/download/0.1.0/cf-reverse-service-lookup-plugin-amd64

# windows 64bit (32bit also available)
cf install-plugin -f https://github.com/aegershman/cf-reverse-service-lookup-plugin/releases/download/0.1.0/cf-reverse-service-lookup-plugin-windows-amd64.exe
```

## feedback welcome

Ideas, use-cases, opinions on architecture, doc, etc. all welcome
