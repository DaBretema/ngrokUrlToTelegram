package main

type ngrok struct {
	tries int
	url   string
}

func newNg() *ngrok {
	ng := ngrok{tries: 0}
	ng.getURL()
	return &ng
}

func (ng *ngrok) getURL() {

	// Exit after some tries
	// if ng.tries >= _Tries {
	// 	errxit(_NgrokBug)
	// }

	// If fails try again (sometimes uri response is slow)
	ng.tries++
	defer recovWithFunc(ng.getURL)

	// Get ngrok url
	ng.url = ngFromURI(ngTunnels).Tunnels[0].PublicURL
}
