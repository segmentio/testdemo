package chantest

type QueueConsumer struct {
	c    chan string
	errC chan error
}
