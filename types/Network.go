package types

type Network struct {
	Name   string `json:"name" sql:"name"`
	RpcURL string `json:"rpcURL" sql:"rpc_url"`
}
