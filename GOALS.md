# Project Goals

To create a kubernetes Custom Resource (*something* - I don't quite follow the terminology at time of writing) which creates:

1. A deployment
1. An ingress instance (and can do magic for stuff like gRPC)
1. A default configmap (or accepts an array of them, perhaps from terraform?) which are added to the env
1. Ditto with secrets- or maybe we add a series of secrets via terraform also? (I think a secret would be better)

Maybe with configmap and secrets, we do what we've done in helm, and just create one per 'thing' with a predictable name that things like terraform can use.

Each deployment would need to only create these things if they don't exist.

(I also don't care about other 'thing's which need files or whatever- none of my apps do)

A few years ago I created a project that used Jupyter to provide dynamic runbooks. The idea was that each deployment was tracked by an informer that templated a notebook. During outages we could visit that notebook, run through it, and figure out problems.

I like this idea of generating some files somewhere from deployments, for stuff like this. I will also, then, use a config map and some binary entries containing templated stuff like a notebook or something. Magic.

I'd like to look at writing some extra healthcheck stuff- we know that gRPC healthchecks for ingresses aren't great, but maybe we can set some extra state like "I'm doing more than I want to be doing- scale me up!", and then maybe we can scale down if everything is happy for a while. Dunno. Soft/hard scaling healthchecks- that's definitely something that'd be cool.
