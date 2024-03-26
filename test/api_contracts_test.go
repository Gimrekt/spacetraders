/*
SpaceTraders API

Testing ContractsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ContractsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContractsAPIService AcceptContract", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contractId string

		resp, httpRes, err := apiClient.ContractsAPI.AcceptContract(context.Background(), contractId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContractsAPIService DeliverContract", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contractId string

		resp, httpRes, err := apiClient.ContractsAPI.DeliverContract(context.Background(), contractId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContractsAPIService FulfillContract", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contractId string

		resp, httpRes, err := apiClient.ContractsAPI.FulfillContract(context.Background(), contractId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContractsAPIService GetContract", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contractId string

		resp, httpRes, err := apiClient.ContractsAPI.GetContract(context.Background(), contractId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContractsAPIService GetContracts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContractsAPI.GetContracts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
