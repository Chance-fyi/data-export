package boot

import "data-export/app/service"

func initConsumer() {
	go service.ExportQueueConsumer(10)
}
