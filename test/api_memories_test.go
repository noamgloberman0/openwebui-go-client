/*
FastAPI

Testing MemoriesAPIService

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

func Test_openapi_MemoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MemoriesAPIService AddMemoryApiV1MemoriesAddPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MemoriesAPI.AddMemoryApiV1MemoriesAddPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MemoriesAPIService DeleteMemoryByIdApiV1MemoriesMemoryIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var memoryId string

		resp, httpRes, err := apiClient.MemoriesAPI.DeleteMemoryByIdApiV1MemoriesMemoryIdDelete(context.Background(), memoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MemoriesAPIService DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MemoriesAPI.DeleteMemoryByUserIdApiV1MemoriesDeleteUserDelete(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MemoriesAPIService GetEmbeddingsApiV1MemoriesEfGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MemoriesAPI.GetEmbeddingsApiV1MemoriesEfGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MemoriesAPIService GetMemoriesApiV1MemoriesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MemoriesAPI.GetMemoriesApiV1MemoriesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MemoriesAPIService QueryMemoryApiV1MemoriesQueryPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MemoriesAPI.QueryMemoryApiV1MemoriesQueryPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MemoriesAPIService ResetMemoryFromVectorDbApiV1MemoriesResetPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MemoriesAPI.ResetMemoryFromVectorDbApiV1MemoriesResetPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MemoriesAPIService UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var memoryId string

		resp, httpRes, err := apiClient.MemoriesAPI.UpdateMemoryByIdApiV1MemoriesMemoryIdUpdatePost(context.Background(), memoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
