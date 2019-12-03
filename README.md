# cf-service-reverse-lookup-plugin

cf-cli plugin to perform reverse lookups against service instance GUIDs

## background

Pairs nicely with problematic BOSH service instances 👌👩‍🍳

This plugin aims to map between how BOSH references service deployments, e.g. `service-instance_GUID`, and how they're referenced in Cloud Foundry. It can be frustrating when tools like the bosh-cli and Prometheus/Alertmanager start yelling about a service instance barfing, but you have to spend time going through this whole riggamarole of looking up the service to see which org/space it belongs to, etc.

With this plugin, as long as you're logged into the same Cloud Foundry installation as the `service-instance_GUID` you're trying to find, this will go through the whole riggamarole of looking up the org/space of the service instance _for_ you.

## usage

```sh
# get details on a service instance by it's GUID alone
cf service-reverse-lookup --service-guid 67e9e04e-8cc5-4744-8a5a-eb0a2d21c7ee

# passing BOSH's default 'service-instance_' prefix is acceptable, too
cf service-reverse-lookup --service-guid service-instance_67e9e04e-8cc5-4744-8a5a-eb0a2d21c7ee
```

currently there is only a `json` format:

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

## feedback welcome

Ideas, use-cases, opinions on architecture, doc, etc. all welcome
