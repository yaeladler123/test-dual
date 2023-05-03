# hippo-service
This repository is a template to create new Services on the CloudBees Platform.

## Working with Other Services

You've got a service and now you actually want to use it.
You have to register in a few different places for it to be visible to other services in the platform

### API Gateway

To see status: [api-gateway application-properties.erb](https://github.com/calculi-corp/api-gateway/blob/main/ci/config/application-properties.erb)

Add the service to the downstream grpc resources [api-gateway main.go](https://github.com/calculi-corp/api-gateway/blob/main/main.go#L55) and to the known list of endpoints [main.go](https://github.com/calculi-corp/api-gateway/blob/main/main.go#L81)

You also need to register the server in [router.go](https://github.com/calculi-corp/api-gateway/blob/main/server/router.go#L18)

And finally call it in [server.go](https://github.com/calculi-corp/api-gateway/blob/main/server/server.go#L130)

More details can be found in the `api-gateway` [README file](https://github.com/calculi-corp/api-gateway/).

### GRPC-hostflags

You can follow along with all the changes made in this [PR](https://github.com/calculi-corp/grpc-hostflags/pull/23).

The most important thing is picking a new port not used by anyone else.
Check the order in serviceflags.go and just pick the next in order.

### API

Register/create all the protos for your service here.
You'll also compile them with protoc (using `make generate` in the `api` repository), and then reference the resulting `.pb.go` and `-grpc.pb.go` from your code.

### DB Service
While you don't need to register anything with db-service, you're going to need to understand a lot of that service when it's getting used.
Please reference the user guide in that service's repo for hints and tips about working with the service.
