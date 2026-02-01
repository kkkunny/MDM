package runner

type RunnerType func() (<-chan struct{}, <-chan error)

var Runners []RunnerType
