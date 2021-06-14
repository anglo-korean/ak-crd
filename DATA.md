# The data I think I need

This is the data/ information I think I need in order to do an ank.

(This data is pretty much a reimagining of the helm charts I usually start with, and I reckon I can get almost all of my usual apps built this way)

## Common schema

This stuff will be used by any of our resources

1. `name: string` -> the base name of an application, interpolates with `env` and `version` to create resource names.
1. `container: string` -> the container to install, as a `text/template` - useful for setting stuff like tags from versions or whatever. Means the resource controller can be built with a useful default
1. `env: string` -> the environment this deployment lives in
1. `version: string` -> the version, as semver. Can be either with or without leading `v`, I don't care (take that helm!)
1. `uniqueish: bool` -> whether to error when running create where a name is in use (in prod, for instance, maybe creating multiple resources of the same name is an error)
1. `singleton: bool` -> create only one replica, useful for conmplex tasks
1. `contact: string` -> where errors, alerts, and so on should go
1. `repo: string` -> location of repo 'thing' lives in
1. `arbitratry_config: interface{}` -> just some stuff to dump into the default config map. Absolutely anything. It won't clobber the config map if it exists, but will overwrite any overlapping fields
1. `storage_family: enum` -> one of ['low', 'cpuBound', 'memBound', 'greedy'] - these configure things like resource requests and limits

### Optionals

If absent, then parent resources get skipped

1. `storage.claim: string` -> optional storage claim (needs to be pre-created)
1. `storage.mount_point: string` -> mount point for claim
1. `preinstall: string` -> a container to run pre-install (to check envs, maybe- has access to all the secrets and configmaps)
1. `postinstall: string` -> a container to run post-install (stuff like running migrations, or notifications perhaps)


## Specific schema

### cron.stable.anglo-korean.github.io

Cronjob configs

1. `cronjob.schedule: string` -> schedule on which to run a job

### http.stable.anglo-korean.github.io and grpc.stable.anglo-korean.github.io

These do pretty much the same thing, with the exception of how requests are routed, and some healthcheck stuff

1. `service.url_template: string` -> go `text/template` of how to generate the ingress url for this service. Has access to all of the vars in the resource (so you could set: `{{ .Version}}.{{ .Name }}.{{ .Env }}.exampple.com` and know where to expect an ingress)
1. `ingress.annotations: map[string]string` -> only annotation overrides we support, are for ingresses in case we have multiple ingress controllers, or need to set bodies, or whatever
