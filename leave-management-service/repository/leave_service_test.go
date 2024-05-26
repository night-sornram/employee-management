package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockLeaveRepo struct {
	GetAllFunc       func() ([]Leave, error)
	GetByIDFunc      func(id int) (Leave, error)
	CreateFunc       func(leave Leave) (Leave, error)
	UpdateFunc       func(id int, leave Leave) (Leave, error)
	DeleteFunc       func(id int) error
	UpdateStatusFunc func(id int, leave Leave) (Leave, error)
	GetAllMeFunc     func(eid string) ([]Leave, error)
}

func (m *mockLeaveRepo) GetAll() ([]Leave, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, errors.New("not implemented")
}

func (m *mockLeaveRepo) GetByID(id int) (Leave, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return Leave{}, errors.New("not implemented")
}

func (m *mockLeaveRepo) Create(leave Leave) (Leave, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(leave)
	}
	return Leave{}, errors.New("not implemented")
}

func (m *mockLeaveRepo) Update(id int, leave Leave) (Leave, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, leave)
	}
	return Leave{}, errors.New("not implemented")
}

func (m *mockLeaveRepo) Delete(id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return errors.New("not implemented")
}

func (m *mockLeaveRepo) UpdateStatus(id int, leave Leave) (Leave, error) {
	if m.UpdateStatusFunc != nil {
		return m.UpdateStatusFunc(id, leave)
	}
	return Leave{}, errors.New("not implemented")
}

func (m *mockLeaveRepo) GetAllMe(eid string) ([]Leave, error) {
	if m.GetAllMeFunc != nil {
		return m.GetAllMeFunc(eid)
	}
	return []Leave{}, errors.New("not implemented")
}

func TestGetLeaves(t *testing.T) {
	t.Run("Valid-GetLeaves", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			GetAllFunc: func() ([]Leave, error) {
				return []Leave{}, nil
			},
		}

		service := NewLeaveService(mockRepo)
		_, err := service.GetLeaves()
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetLeaves", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			GetAllFunc: func() ([]Leave, error) {
				return []Leave{}, errors.New("invalid")
			},
		}

		service := NewLeaveService(mockRepo)
		_, err := service.GetLeaves()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetLeave(t *testing.T) {
	t.Run("Valid-GetLeave", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, nil
			},
		}

		service := NewLeaveService(mockRepo)
		_, err := service.GetLeave(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetLeave", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, errors.New("invalid")
			},
		}

		service := NewLeaveService(mockRepo)
		_, err := service.GetLeave(1)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreateLeave(t *testing.T) {
	t.Run("Valid-CreateLeave", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			CreateFunc: func(leave Leave) (Leave, error) {
				return Leave{}, nil
			},
		}
		service := NewLeaveService(mockRepo)
		_, err := service.CreateLeave(Leave{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-CreateLeave", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			CreateFunc: func(leave Leave) (Leave, error) {
				return Leave{}, errors.New("invalid")
			},
		}
		service := NewLeaveService(mockRepo)
		_, err := service.CreateLeave(Leave{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestUpdateLeave(t *testing.T) {
	t.Run("Valid-UpdateLeave", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			UpdateFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, nil
			},
		}
		service := NewLeaveService(mockRepo)
		_, err := service.UpdateLeave(1, Leave{})
		assert.NoError(t, err)
	})
	t.Run("Invalid-UpdateLeave", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			UpdateFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, errors.New("invalid")
			},
		}
		service := NewLeaveService(mockRepo)
		_, err := service.UpdateLeave(1, Leave{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDeleteLeave(t *testing.T) {
	t.Run("Valid-DeleteLeave", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			DeleteFunc: func(id int) error {
				return nil
			},
		}
		service := NewLeaveService(mockRepo)
		err := service.DeleteLeave(1)
		assert.NoError(t, err)
	})
	t.Run("Invalid-DeleteLeave", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			DeleteFunc: func(id int) error {
				return errors.New("invalid")
			},
		}
		service := NewLeaveService(mockRepo)
		err := service.DeleteLeave(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("Valid-UpdateStatus-Approve", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			UpdateStatusFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, nil
			},
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, nil
			},
		}

		mockUpdateStatus := LeaveStatus{
			Status:         "approve",
			ManagerOpinion: "OK, approve",
		}

		service := NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.NoError(t, err)
	})

	t.Run("Valid-UpdateStatus-Reject", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			UpdateStatusFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, nil
			},
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, nil
			},
		}

		mockUpdateStatus := LeaveStatus{
			Status:         "reject",
			ManagerOpinion: "reject",
		}

		service := NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.NoError(t, err)
	})

	t.Run("Invalid-UpdateStatus-ID-Notfound", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			UpdateStatusFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, nil
			},
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, errors.New("invalid")
			},
		}

		mockUpdateStatus := LeaveStatus{
			Status:         "approve",
			ManagerOpinion: "OK, approve",
		}

		service := NewLeaveService(mockRepo)
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
		mockRepo := &mockLeaveRepo{
			UpdateStatusFunc: func(id int, leave Leave) (Leave, error) {
				return Leave{}, errors.New("invalid")
			},
			GetByIDFunc: func(id int) (Leave, error) {
				return Leave{}, nil
			},
		}

		mockUpdateStatus := LeaveStatus{
			Status:         "approve",
			ManagerOpinion: "OK, approve",
		}

		service := NewLeaveService(mockRepo)
		_, err := service.UpdateStatus(1, mockUpdateStatus)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetMyLeaves(t *testing.T) {
	t.Run("Valid-GetMyLeaves", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			GetAllMeFunc: func(eid string) ([]Leave, error) {
				return []Leave{}, nil
			},
		}
		service := NewLeaveService(mockRepo)
		_, err := service.GetAllMe("E0001")
		assert.NoError(t, err)
	})
	t.Run("Invalid-GetMyLeaves", func(t *testing.T) {
		mockRepo := &mockLeaveRepo{
			GetAllMeFunc: func(eid string) ([]Leave, error) {
				return []Leave{}, errors.New("invalid")
			},
		}
		service := NewLeaveService(mockRepo)
		_, err := service.GetAllMe("E0001")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
