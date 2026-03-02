# \AuthenticationAPI

All URIs are relative to *https://api.aerocall.ai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthSignin**](AuthenticationAPI.md#AuthSignin) | **Post** /auth/signin | Sign in user
[**AuthSignup**](AuthenticationAPI.md#AuthSignup) | **Post** /auth/signup | Sign up new user



## AuthSignin

> AuthResponse AuthSignin(ctx).AuthSigninRequest(authSigninRequest).Execute()

Sign in user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authSigninRequest := *openapiclient.NewAuthSigninRequest("Email_example", "Password_example") // AuthSigninRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthSignin(context.Background()).AuthSigninRequest(authSigninRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthSignin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSignin`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthSignin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSigninRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authSigninRequest** | [**AuthSigninRequest**](AuthSigninRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSignup

> AuthResponse AuthSignup(ctx).AuthSignupRequest(authSignupRequest).Execute()

Sign up new user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authSignupRequest := *openapiclient.NewAuthSignupRequest("user@example.com", "SecurePass123!") // AuthSignupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthSignup(context.Background()).AuthSignupRequest(authSignupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthSignup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSignup`: AuthResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthSignup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authSignupRequest** | [**AuthSignupRequest**](AuthSignupRequest.md) |  | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

