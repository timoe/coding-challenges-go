# Queues

The queue package implememts two basic queues, one with FiFo semantics and a prioritized queue. 

[FiFo queue](fifo-queue.go) uses a circular linked list, [priority queue](priority-queue.go) creates a max heap when new items get queued to determine the one with the highest priority.