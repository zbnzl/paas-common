package clients

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewPodHttpClient, NewPodGrpcClient)

func init() {

}
