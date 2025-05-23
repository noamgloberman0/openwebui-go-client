/*
FastAPI

Testing ChatsAPIService

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

func Test_openapi_ChatsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChatsAPIService AddTagByIdAndTagNameApiV1ChatsIdTagsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.AddTagByIdAndTagNameApiV1ChatsIdTagsPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService ArchiveAllChatsApiV1ChatsArchiveAllPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.ArchiveAllChatsApiV1ChatsArchiveAllPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService ArchiveChatByIdApiV1ChatsIdArchivePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.ArchiveChatByIdApiV1ChatsIdArchivePost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService CloneChatByIdApiV1ChatsIdClonePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.CloneChatByIdApiV1ChatsIdClonePost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService CloneSharedChatByIdApiV1ChatsIdCloneSharedPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.CloneSharedChatByIdApiV1ChatsIdCloneSharedPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService CreateNewChatApiV1ChatsNewPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.CreateNewChatApiV1ChatsNewPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService DeleteAllTagsByIdApiV1ChatsIdTagsAllDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.DeleteAllTagsByIdApiV1ChatsIdTagsAllDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService DeleteAllUserChatsApiV1ChatsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.DeleteAllUserChatsApiV1ChatsDelete(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService DeleteChatByIdApiV1ChatsIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.DeleteChatByIdApiV1ChatsIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService DeleteSharedChatByIdApiV1ChatsIdShareDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.DeleteSharedChatByIdApiV1ChatsIdShareDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService DeleteTagByIdAndTagNameApiV1ChatsIdTagsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.DeleteTagByIdAndTagNameApiV1ChatsIdTagsDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetAllUserChatsInDbApiV1ChatsAllDbGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetAllUserChatsInDbApiV1ChatsAllDbGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetAllUserTagsApiV1ChatsAllTagsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetAllUserTagsApiV1ChatsAllTagsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetArchivedSessionUserChatListApiV1ChatsArchivedGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetArchivedSessionUserChatListApiV1ChatsArchivedGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetChatByIdApiV1ChatsIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.GetChatByIdApiV1ChatsIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetChatTagsByIdApiV1ChatsIdTagsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.GetChatTagsByIdApiV1ChatsIdTagsGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetChatsByFolderIdApiV1ChatsFolderFolderIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.ChatsAPI.GetChatsByFolderIdApiV1ChatsFolderFolderIdGet(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetPinnedStatusByIdApiV1ChatsIdPinnedGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.GetPinnedStatusByIdApiV1ChatsIdPinnedGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetSessionUserChatListApiV1ChatsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetSessionUserChatListApiV1ChatsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetSessionUserChatListApiV1ChatsListGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetSessionUserChatListApiV1ChatsListGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetSharedChatByIdApiV1ChatsShareShareIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var shareId string

		resp, httpRes, err := apiClient.ChatsAPI.GetSharedChatByIdApiV1ChatsShareShareIdGet(context.Background(), shareId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetUserArchivedChatsApiV1ChatsAllArchivedGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetUserArchivedChatsApiV1ChatsAllArchivedGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetUserChatListByTagNameApiV1ChatsTagsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetUserChatListByTagNameApiV1ChatsTagsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetUserChatListByUserIdApiV1ChatsListUserUserIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.ChatsAPI.GetUserChatListByUserIdApiV1ChatsListUserUserIdGet(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetUserChatsApiV1ChatsAllGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetUserChatsApiV1ChatsAllGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService GetUserPinnedChatsApiV1ChatsPinnedGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.GetUserPinnedChatsApiV1ChatsPinnedGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService ImportChatApiV1ChatsImportPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.ImportChatApiV1ChatsImportPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService PinChatByIdApiV1ChatsIdPinPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.PinChatByIdApiV1ChatsIdPinPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService SearchUserChatsApiV1ChatsSearchGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatsAPI.SearchUserChatsApiV1ChatsSearchGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService ShareChatByIdApiV1ChatsIdSharePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.ShareChatByIdApiV1ChatsIdSharePost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService UpdateChatByIdApiV1ChatsIdPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.UpdateChatByIdApiV1ChatsIdPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatsAPIService UpdateChatFolderIdByIdApiV1ChatsIdFolderPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ChatsAPI.UpdateChatFolderIdByIdApiV1ChatsIdFolderPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
