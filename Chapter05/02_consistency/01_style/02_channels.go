package _1_style

func ChannelExample() {
	// vanilla
	signal := make(chan struct{})
	performShutdown(signal)
	<-signal

	// with Ch
	signalCh := make(chan struct{})
	performShutdown(signalCh)
	<-signalCh
}

func performShutdown(signal chan struct{}) {
	defer close(signal)

	// implementation removed
}
