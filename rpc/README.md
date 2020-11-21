## RPC Related Library

### How to use

If the call success with err == nil, it will mutate the `res` variable.
```
var result domain.User
err := GetUserByID(&rpc.ValInt{Value: int64(1)}, &result)


func GetUserByID(req *rpc.ValInt, res *domain.User) error {
	err := rpc.CallRPC(config.serverURL, config.ServiceName+".GetUserByID", req, &res)
	return err
}
```
