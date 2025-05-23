/*
FastAPI

Testing OpenaiAPIService

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

func Test_openapi_OpenaiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OpenaiAPIService GenerateChatCompletionOpenaiChatCompletionsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OpenaiAPI.GenerateChatCompletionOpenaiChatCompletionsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService GetConfigOpenaiConfigGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OpenaiAPI.GetConfigOpenaiConfigGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService GetModelsOpenaiModelsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OpenaiAPI.GetModelsOpenaiModelsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService GetModelsOpenaiModelsUrlIdxGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var urlIdx int32

		resp, httpRes, err := apiClient.OpenaiAPI.GetModelsOpenaiModelsUrlIdxGet(context.Background(), urlIdx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService ProxyOpenaiPathPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.OpenaiAPI.ProxyOpenaiPathPut(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService ProxyOpenaiPathPut_1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.OpenaiAPI.ProxyOpenaiPathPut_0(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService ProxyOpenaiPathPut_2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.OpenaiAPI.ProxyOpenaiPathPut_1(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService ProxyOpenaiPathPut_3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.OpenaiAPI.ProxyOpenaiPathPut_2(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService SpeechOpenaiAudioSpeechPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OpenaiAPI.SpeechOpenaiAudioSpeechPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService UpdateConfigOpenaiConfigUpdatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OpenaiAPI.UpdateConfigOpenaiConfigUpdatePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenaiAPIService VerifyConnectionOpenaiVerifyPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OpenaiAPI.VerifyConnectionOpenaiVerifyPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
