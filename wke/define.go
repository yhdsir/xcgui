package wke

type WkeProxyType uint32

const (
	WKE_PROXY_NONE WkeProxyType = iota
	WKE_PROXY_HTTP
	WKE_PROXY_SOCKS4
	WKE_PROXY_SOCKS4A
	WKE_PROXY_SOCKS5
	WKE_PROXY_SOCKS5HOSTNAME
)

type WkeLoadingResult uint32

const (
	WKE_LOADING_SUCCEEDED WkeLoadingResult = iota
	WKE_LOADING_FAILED
	WKE_LOADING_CANCELED
)