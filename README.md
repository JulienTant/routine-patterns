# Golang routing patterns

* [generator](https://github.com/JulienTant/routine-patterns/tree/master/generator) : don't wait for the data to be ready to process it
* [future](https://github.com/JulienTant/routine-patterns/tree/master/future) : allow to do an async operation (get data from an HTTP request) without blocking us. At some point, we get get the data or just wait for it to be ready
* [worker-pool](https://github.com/JulienTant/routine-patterns/tree/master/worker-pool) : have multiple routine ready to handle tasks put into a single channel