package test

import (
	"errors"
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"github.com/night-sornram/employee-management/leave-management-service/repository/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLeaves(t *testing.T) {
	t.Run("Valid-GetLeaves", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("GetAll", mock.Anything).Return(repository.DataJson{}, nil)

		service := repository.NewLeaveService(mockRepo)
		_, err := service.GetLeaves(repository.Query{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetLeaves", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("GetAll", mock.Anything).Return(repository.DataJson{}, errors.New("invalid"))

		service := repository.NewLeaveService(mockRepo)
		_, err := service.GetLeaves(repository.Query{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetLeave(t *testing.T) {
	t.Run("Valid-GetLeave", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("GetByID", 1).Return(repository.Leave{}, nil)

		service := repository.NewLeaveService(mockRepo)
		_, err := service.GetLeave(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetLeave", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("GetByID", 1).Return(repository.Leave{}, errors.New("invalid"))

		service := repository.NewLeaveService(mockRepo)
		_, err := service.GetLeave(1)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreateLeave(t *testing.T) {
	t.Run("Valid-CreateLeave", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("Create", mock.Anything).Return(repository.Leave{}, nil)

		service := repository.NewLeaveService(mockRepo)
		_, err := service.CreateLeave(repository.Leave{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-CreateLeave", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("Create", mock.Anything).Return(repository.Leave{}, errors.New("invalid"))

		service := repository.NewLeaveService(mockRepo)
		_, err := service.CreateLeave(repository.Leave{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdateLeave(t *testing.T) {
	t.Run("Valid-UpdateLeave", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("Update", 1, mock.Anything).Return(repository.Leave{}, nil)

		service := repository.NewLeaveService(mockRepo)
		_, err := service.UpdateLeave(1, repository.Leave{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-UpdateLeave", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("Update", 1, mock.Anything).Return(repository.Leave{}, errors.New("invalid"))

		service := repository.NewLeaveService(mockRepo)
		_, err := service.UpdateLeave(1, repository.Leave{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDeleteLeave(t *testing.T) {
	t.Run("Valid-DeleteLeave", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("Delete", 1).Return(nil)

		service := repository.NewLeaveService(mockRepo)
		err := service.DeleteLeave(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-DeleteLeave", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("Delete", 1).Return(errors.New("invalid"))

		service := repository.NewLeaveService(mockRepo)
		err := service.DeleteLeave(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestUpdateStatusLeave(t *testing.T) {
	t.Run("Valid-UpdateStatus-Approve", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("UpdateStatus", 1, mock.Anything).Return(repository.Leave{}, nil)
		mockRepo.On("GetByID", 1).Return(repository.Leave{}, nil)

		mockUpdateStatus := repository.LeaveStatus{
			Status:         "approve",
			ManagerOpinion: "OK, approve",
		}

		service := repository.NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.NoError(t, err)
	})

	t.Run("Valid-UpdateStatus-Reject", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("UpdateStatus", 1, mock.Anything).Return(repository.Leave{}, nil)
		mockRepo.On("GetByID", 1).Return(repository.Leave{}, nil)

		mockUpdateStatus := repository.LeaveStatus{
			Status:         "reject",
			ManagerOpinion: "reject",
		}

		service := repository.NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.NoError(t, err)
	})

	t.Run("Invalid-UpdateStatus-ID-Notfound", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("UpdateStatus", 1, mock.Anything).Return(repository.Leave{}, nil)
		mockRepo.On("GetByID", 1).Return(repository.Leave{}, errors.New("invalid"))

		mockUpdateStatus := repository.LeaveStatus{
			Status:         "approve",
			ManagerOpinion: "OK, approve",
		}

		service := repository.NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
	//t.Run("Invalid-UpdateStatus-PostFail", func(t *testing.T) {
	//	mockRepo := &mockLeaveRepo{
	//		UpdateStatusFunc: func(id int, leave Leave) (Leave, error) {
	//			return Leave{}, nil
	//		},
	//		GetByIDFunc: func(id int) (Leave, error) {
	//			return Leave{}, nil
	//		},
	//	}
	//	mockPostAttendance := func(payload repository.Attendance) error {
	//		return errors.New("failed to post attendance")
	//	}
	//
	//	mockUpdateStatus := LeaveStatus{
	//		Status:         "approve",
	//		ManagerOpinion: "OK, approve",
	//	}
	//
	//	service := NewLeaveService(mockRepo)
	//	originalPostAttendance := PostAttendance
	//	PostAttendance = mockPostAttendance
	//	defer func() {
	//		PostAttendance = originalPostAttendance
	//	}()
	//	_, err := service.UpdateStatus(1, mockUpdateStatus)
	//	assert.Error(t, err)
	//	//assert.Equal(t, "invalid", err.Error())
	//})
	t.Run("Invalid-UpdateStatus", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("UpdateStatus", 1, mock.Anything).Return(repository.Leave{}, errors.New("invalid"))
		mockRepo.On("GetByID", 1).Return(repository.Leave{}, nil)

		mockUpdateStatus := repository.LeaveStatus{
			Status:         "approve",
			ManagerOpinion: "OK, approve",
		}

		service := repository.NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetAllMeLeaves(t *testing.T) {
	t.Run("Valid-GetAllMeLeaves", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("GetAllMe", mock.Anything, "E0001").Return(repository.DataJson{}, nil)

		service := repository.NewLeaveService(mockRepo)
		_, err := service.GetAllMe(repository.Query{}, "E0001")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetAllMeLeaves", func(t *testing.T) {
		mockRepo := new(mocks.LeaveRepository)
		mockRepo.On("GetAllMe", mock.Anything, "E0001").Return(repository.DataJson{}, errors.New("invalid"))

		service := repository.NewLeaveService(mockRepo)
		_, err := service.GetAllMe(repository.Query{}, "E0001")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
