module github.com/JohnVicke/klask-backend

go 1.18

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.5.3
	github.com/JohnVicke/klask-backend/cmd/healthcheck v1.0.0-beta.1
)

require (
	cloud.google.com/go/functions v1.4.0 // indirect
	github.com/cloudevents/sdk-go/v2 v2.10.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
)

replace github.com/JohnVicke/klask-backend/cmd/helloworld => ./cmd/helloworld

replace github.com/JohnVicke/klask-backend/cmd/healthcheck => ./cmd/healthcheck
