package workers

type Worker interface {
	Start(queueName string)
}
