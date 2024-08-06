package types

type ClientRequest struct {
	RedirectURIs string
}

type ClientResponse struct {
	ID           uint
	ClientID     string
	ClientSecret string
}

// type response struct {
// 	ID           uint
// 	ClientID     string
// 	ClientSecret string
// }
