package services

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"track_service/config"
	"track_service/mocks/mock_repositories"
	"track_service/models"

	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func setup() (*trackService, *mock_repositories.MockTrackRepository) {
	ctrl := gomock.NewController(nil)
	mockRepo := mock_repositories.NewMockTrackRepository(ctrl)

	cfg := &config.Config{
		RequestTimeout: time.Second * 2,
	}

	return NewTrackService(mockRepo, cfg).(*trackService), mockRepo
}

func TestUpload_ErrorInGetByTitleAndUserId(t *testing.T) {
	service, mockRepo := setup()

	metadata := models.TrackMetadata{
		Title:  "test-title",
		UserID: "test-user",
	}

	mockRepo.EXPECT().GetByTitleAndUserId(metadata.Title, metadata.UserID).Return(nil, errors.New("some error"))

	_, err := service.Upload(context.Background(), metadata)
	st, _ := status.FromError(err)
	if st.Code() != codes.Internal {
		t.Fatalf("Expected error with code %v, got %v", codes.Internal, st.Code())
	}
}

func TestUpload_TrackAlreadyExists(t *testing.T) {
	service, mockRepo := setup()

	metadata := models.TrackMetadata{
		Title:  "test-title",
		UserID: "test-user",
	}

	mockRepo.EXPECT().GetByTitleAndUserId(metadata.Title, metadata.UserID).Return(&models.Track{}, nil)

	_, err := service.Upload(context.Background(), metadata)
	st, _ := status.FromError(err)
	if st.Code() != codes.AlreadyExists {
		t.Fatalf("Expected error with code %v, got %v", codes.AlreadyExists, st.Code())
	}
}

func TestUpload_ErrorInCreate(t *testing.T) {
	service, mockRepo := setup()

	metadata := models.TrackMetadata{
		Title:  "test-title",
		UserID: "test-user",
	}

	mockRepo.EXPECT().GetByTitleAndUserId(metadata.Title, metadata.UserID).Return(nil, gorm.ErrRecordNotFound)
	mockRepo.EXPECT().Create(gomock.Any()).Return(errors.New("create error"))

	_, err := service.Upload(context.Background(), metadata)
	st, _ := status.FromError(err)
	if st.Code() != codes.Internal {
		t.Fatalf("Expected error with code %v, got %v", codes.Internal, st.Code())
	}
}

func TestUpload_Success(t *testing.T) {
	service, mockRepo := setup()

	metadata := models.TrackMetadata{
		Title:  "test-title",
		UserID: "test-user",
		URL:    "test-url",
	}

	mockRepo.EXPECT().GetByTitleAndUserId(metadata.Title, metadata.UserID).Return(nil, gorm.ErrRecordNotFound)
	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

	resp, err := service.Upload(context.Background(), metadata)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}

	if resp.URL != metadata.URL {
		t.Fatalf("Expected URL to be %v, got %v", metadata.URL, resp.URL)
	}

	if resp.UserID != metadata.UserID {
		t.Fatalf("Expected UserID to be %v, got %v", metadata.UserID, resp.UserID)
	}
}

func TestGetInfoById_Success(t *testing.T) {
	service, mockRepo := setup()

	trackId := "123"
	expectedTrack := &models.Track{
		ID:     123,
		Title:  "Test Title",
		Artist: "Test Artist",
	}
	mockRepo.EXPECT().GetById(trackId).Return(expectedTrack, nil)

	resp, err := service.GetInfoById(context.Background(), models.TrackIdRequest{Id: trackId})
	assert.NoError(t, err)
	assert.Equal(t, expectedTrack.Title, resp.Title)
	assert.Equal(t, expectedTrack.Artist, resp.Artist)
}

func TestGetInfoById_NotFound(t *testing.T) {
	service, mockRepo := setup()

	trackId := "123"
	mockRepo.EXPECT().GetById(trackId).Return(nil, status.Error(codes.NotFound, "track not found"))

	_, err := service.GetInfoById(context.Background(), models.TrackIdRequest{Id: trackId})
	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
}

func TestEditInfo_TrackNotFound(t *testing.T) {
	service, mockRepo := setup()

	mockRepo.EXPECT().GetById("someTrackID").Return(nil, errors.New("track not found"))

	req := models.EditTrackRequest{TrackId: "someTrackID"}
	err := service.EditInfo(context.Background(), req)

	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
	assert.Contains(t, err.Error(), "track not found")
}

func TestEditInfo_UpdateError(t *testing.T) {
	service, mockRepo := setup()

	track := &models.Track{
		Title: "Some Title",
	}
	mockRepo.EXPECT().GetById("someTrackID").Return(track, nil)

	mockRepo.EXPECT().Update(gomock.Any()).Return(errors.New("failed to update track info"))

	req := models.EditTrackRequest{TrackId: "someTrackID"}
	err := service.EditInfo(context.Background(), req)

	assert.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Contains(t, err.Error(), "failed to update track info")
}

func TestEditInfo_Success(t *testing.T) {
	service, mockRepo := setup()

	track := &models.Track{
		Title: "Some Title",
	}
	mockRepo.EXPECT().GetById("someTrackID").Return(track, nil)

	mockRepo.EXPECT().Update(gomock.Any()).Return(nil)

	req := models.EditTrackRequest{TrackId: "someTrackID"}
	err := service.EditInfo(context.Background(), req)

	assert.NoError(t, err)
}

func TestDeleteById_Success(t *testing.T) {
	service, mockRepo := setup()

	mockRepo.EXPECT().Delete("123").Return(nil)

	err := service.DeleteById(context.Background(), models.TrackIdRequest{Id: "123"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteById_ErrorDeleting(t *testing.T) {
	service, mockRepo := setup()

	mockRepo.EXPECT().Delete("123").Return(status.Error(codes.Internal, "failed to delete track"))

	err := service.DeleteById(context.Background(), models.TrackIdRequest{Id: "123"})
	if st, ok := status.FromError(err); !ok || st.Code() != codes.Internal || st.Message() != "failed to delete track" {
		t.Fatalf("expected status error with code Internal and message 'failed to delete track', got %v", err)
	}
}

func TestFindAll_Success(t *testing.T) {
	service, mockRepo := setup()

	expectedTracks := []*models.Track{
		{
			ID:     1,
			Title:  "Song1",
			Artist: "Artist1",
			Album:  "Album1",
			Genre:  "Genre1",
			URL:    "URL1",
			UserID: "UserID1",
		},
		{
			ID:     2,
			Title:  "Song2",
			Artist: "Artist2",
			Album:  "Album2",
			Genre:  "Genre2",
			URL:    "URL2",
			UserID: "UserID2",
		},
	}

	mockRepo.EXPECT().FindAll().Return(expectedTracks, nil)

	res, err := service.FindAll(context.Background())
	assert.Nil(t, err)

	assert.Len(t, res, 2)
	assert.Equal(t, "Song1", res[0].Title)
	assert.Equal(t, "Song2", res[1].Title)
}

func TestFindAll_Error(t *testing.T) {
	service, mockRepo := setup()

	mockRepo.EXPECT().FindAll().Return(nil, status.Error(codes.Internal, "internal error"))

	_, err := service.FindAll(context.Background())
	assert.NotNil(t, err)

	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Equal(t, "internal error", status.Convert(err).Message())
}

func TestStatus(t *testing.T) {
	service, _ := setup()

	tests := []struct {
		name   string
		output string
		err    error
	}{
		{
			name:   "valid status",
			output: "ok",
			err:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			statusTest, err := service.Status(context.Background())
			assert.Equal(t, test.output, statusTest)
			assert.Equal(t, test.err, err)
		})
	}
}
