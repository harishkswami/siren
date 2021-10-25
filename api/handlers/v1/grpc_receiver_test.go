package v1

import (
	"context"
	"errors"
	sirenv1 "github.com/odpf/siren/api/proto/odpf/siren/v1"
	"github.com/odpf/siren/domain"
	"github.com/odpf/siren/mocks"
	"github.com/odpf/siren/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
	"testing"
	"time"
)

func TestGRPCServer_ListReceiver(t *testing.T) {
	configurations := make(map[string]interface{})
	configurations["foo"] = "bar"
	labels := make(map[string]string)
	labels["foo"] = "bar"
	dummyResult := []*domain.Receiver{
		{
			Id:             1,
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurations,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	t.Run("should return list of all receiver", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("ListReceivers").
			Return(dummyResult, nil).Once()

		res, err := dummyGRPCServer.ListReceivers(context.Background(), &emptypb.Empty{})
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetReceivers()))
		assert.Equal(t, uint64(1), res.GetReceivers()[0].GetId())
		assert.Equal(t, "foo", res.GetReceivers()[0].GetName())
		assert.Equal(t, "bar", res.GetReceivers()[0].GetType())
		assert.Equal(t, "bar", res.GetReceivers()[0].GetConfigurations().AsMap()["foo"])
		assert.Equal(t, "bar", res.GetReceivers()[0].GetLabels()["foo"])
	})

	t.Run("should return error code 13 if getting providers failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("ListReceivers").
			Return(nil, errors.New("random error"))

		res, err := dummyGRPCServer.ListReceivers(context.Background(), &emptypb.Empty{})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})

	t.Run("should return error code 13 if NewStruct conversion failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		configurations["foo"] = string([]byte{0xff})
		dummyResult := []*domain.Receiver{
			{
				Id:             1,
				Name:           "foo",
				Type:           "bar",
				Labels:         labels,
				Configurations: configurations,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
		}

		mockedReceiverService.
			On("ListReceivers").
			Return(dummyResult, nil)
		res, err := dummyGRPCServer.ListReceivers(context.Background(), &emptypb.Empty{})
		assert.Nil(t, res)
		assert.Equal(t, strings.Replace(err.Error(), "\u00a0", " ", -1),
			"rpc error: code = Internal desc = proto: invalid UTF-8 in string: \"\\xff\"")
	})
}

func TestGRPCServer_CreateReceiver(t *testing.T) {
	configurations := make(map[string]interface{})
	configurations["client_id"] = "foo"
	configurations["client_secret"] = "bar"
	configurations["auth_code"] = "foo"
	labels := make(map[string]string)
	labels["foo"] = "bar"

	configurationsData, _ := structpb.NewStruct(configurations)
	dummyReq := &sirenv1.CreateReceiverRequest{
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurationsData,
	}
	payload := &domain.Receiver{
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurations,
	}

	t.Run("Should create a slack receiver object", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("CreateReceiver", payload).
			Return(payload, nil).Once()

		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, "foo", res.GetName())
		assert.Equal(t, "slack", res.GetType())
		assert.Equal(t, "bar", res.GetLabels()["foo"])
		assert.Equal(t, "foo", res.GetConfigurations().AsMap()["client_id"])
	})

	t.Run("should return error code 3 if slack client_id configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		slackConfigurations["client_secret"] = "foo"
		slackConfigurations["auth_code"] = "foo"

		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.CreateReceiverRequest{
			Name:           "foo",
			Type:           "slack",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"client_id\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if slack client_secret configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		slackConfigurations["client_id"] = "foo"
		slackConfigurations["auth_code"] = "foo"

		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.CreateReceiverRequest{
			Name:           "foo",
			Type:           "slack",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"client_secret\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if slack auth_code configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		slackConfigurations["client_id"] = "foo"
		slackConfigurations["client_secret"] = "foo"

		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.CreateReceiverRequest{
			Name:           "foo",
			Type:           "slack",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"auth_code\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if pagerduty service_key configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.CreateReceiverRequest{
			Name:           "foo",
			Type:           "pagerduty",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"service_key\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if http url configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.CreateReceiverRequest{
			Name:           "foo",
			Type:           "http",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"url\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 13 if creating receiver failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("CreateReceiver", payload).
			Return(nil, errors.New("random error")).Once()

		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})

	t.Run("should return error code 3 if receiver is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}

		configurationsData, _ := structpb.NewStruct(configurations)
		dummyReq := &sirenv1.CreateReceiverRequest{
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err, "rpc error: code = InvalidArgument desc = receiver not supported")
		assert.Nil(t, res)
	})

	t.Run("should return error code 13 if NewStruct conversion failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}

		configurations["workspace"] = string([]byte{0xff})
		newPayload := &domain.Receiver{
			Name:           "foo",
			Type:           "slack",
			Labels:         labels,
			Configurations: configurations,
		}

		mockedReceiverService.
			On("CreateReceiver", mock.Anything).
			Return(newPayload, nil)
		res, err := dummyGRPCServer.CreateReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.Equal(t, strings.Replace(err.Error(), "\u00a0", " ", -1),
			"rpc error: code = Internal desc = proto: invalid UTF-8 in string: \"\\xff\"")
	})
}

func TestGRPCServer_GetReceiver(t *testing.T) {
	configurations := make(map[string]interface{})
	configurations["foo"] = "bar"
	labels := make(map[string]string)
	labels["foo"] = "bar"

	receiverId := uint64(1)
	dummyReq := &sirenv1.GetReceiverRequest{
		Id: 1,
	}
	payload := &domain.Receiver{
		Name:           "foo",
		Type:           "bar",
		Labels:         labels,
		Configurations: configurations,
	}

	t.Run("should return a receiver", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("GetReceiver", receiverId).
			Return(payload, nil).Once()

		res, err := dummyGRPCServer.GetReceiver(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, "foo", res.GetName())
		assert.Equal(t, "bar", res.GetType())
		assert.Equal(t, "bar", res.GetLabels()["foo"])
		assert.Equal(t, "bar", res.GetConfigurations().AsMap()["foo"])
	})

	t.Run("should return error code 5 if no receiver found", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("GetReceiver", receiverId).
			Return(nil, nil).Once()

		res, err := dummyGRPCServer.GetReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = NotFound desc = receiver not found")
	})

	t.Run("should return error code 13 if getting receiver failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("GetReceiver", receiverId).
			Return(payload, errors.New("random error")).Once()

		res, err := dummyGRPCServer.GetReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})

	t.Run("should return error code 13 if NewStruct conversion of configuration failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}

		configurations["foo"] = string([]byte{0xff})
		payload := &domain.Receiver{
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurations,
		}

		mockedReceiverService.
			On("GetReceiver", receiverId).
			Return(payload, nil)
		res, err := dummyGRPCServer.GetReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.Equal(t, strings.Replace(err.Error(), "\u00a0", " ", -1),
			"rpc error: code = Internal desc = proto: invalid UTF-8 in string: \"\\xff\"")
	})

	t.Run("should return error code 13 if data NewStruct conversion of data failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		data := make(map[string]interface{})
		data["channels"] = string([]byte{0xff})
		payload := &domain.Receiver{
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurations,
			Data:           data,
		}

		mockedReceiverService.
			On("GetReceiver", receiverId).
			Return(payload, nil)
		res, err := dummyGRPCServer.GetReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.Equal(t, strings.Replace(err.Error(), "\u00a0", " ", -1),
			"rpc error: code = Internal desc = proto: invalid UTF-8 in string: \"\\xff\"")
	})
}

func TestGRPCServer_UpdateReceiver(t *testing.T) {
	configurations := make(map[string]interface{})
	configurations["client_id"] = "foo"
	configurations["client_secret"] = "bar"
	configurations["auth_code"] = "foo"

	labels := make(map[string]string)
	labels["foo"] = "bar"

	configurationsData, _ := structpb.NewStruct(configurations)
	dummyReq := &sirenv1.UpdateReceiverRequest{
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurationsData,
	}
	payload := &domain.Receiver{
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurations,
	}

	t.Run("should update receiver object", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("UpdateReceiver", payload).
			Return(payload, nil).Once()

		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, "foo", res.GetName())
		assert.Equal(t, "slack", res.GetType())
		assert.Equal(t, "bar", res.GetLabels()["foo"])
		assert.Equal(t, "foo", res.GetConfigurations().AsMap()["client_id"])
	})

	t.Run("should return error code 3 if slack client_id configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		slackConfigurations["client_secret"] = "foo"
		slackConfigurations["auth_code"] = "foo"

		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.UpdateReceiverRequest{
			Name:           "foo",
			Type:           "slack",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"client_id\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if slack client_secret configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		slackConfigurations["client_id"] = "foo"
		slackConfigurations["auth_code"] = "foo"

		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.UpdateReceiverRequest{
			Name:           "foo",
			Type:           "slack",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"client_secret\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if slack auth_code configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		slackConfigurations["client_id"] = "foo"
		slackConfigurations["client_secret"] = "foo"

		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.UpdateReceiverRequest{
			Name:           "foo",
			Type:           "slack",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"auth_code\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if pagerduty service_key configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		slackConfigurations["client_id"] = "foo"
		slackConfigurations["client_secret"] = "foo"

		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.UpdateReceiverRequest{
			Name:           "foo",
			Type:           "pagerduty",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"service_key\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if http url configuration is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		slackConfigurations := make(map[string]interface{})
		slackConfigurations["client_id"] = "foo"
		slackConfigurations["client_secret"] = "foo"

		configurationsData, _ := structpb.NewStruct(slackConfigurations)
		dummyReq := &sirenv1.UpdateReceiverRequest{
			Name:           "foo",
			Type:           "http",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err,
			"rpc error: code = InvalidArgument desc = No value supplied for required configurations map key \"url\"")
		assert.Nil(t, res)
	})

	t.Run("should return error code 3 if receiver is missing", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}

		configurationsData, _ := structpb.NewStruct(configurations)
		dummyReq := &sirenv1.UpdateReceiverRequest{
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurationsData,
		}

		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.EqualError(t, err, "rpc error: code = InvalidArgument desc = receiver not supported")
		assert.Nil(t, res)
	})

	t.Run("should return error code 13 if updating receiver failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("UpdateReceiver", payload).
			Return(nil, errors.New("random error"))

		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})

	t.Run("should return error code 13 if NewStruct conversion failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		configurations["foo"] = string([]byte{0xff})
		newPayload := &domain.Receiver{
			Name:           "foo",
			Type:           "bar",
			Labels:         labels,
			Configurations: configurations,
		}

		mockedReceiverService.
			On("UpdateReceiver", mock.Anything).
			Return(newPayload, nil)
		res, err := dummyGRPCServer.UpdateReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.Equal(t, strings.Replace(err.Error(), "\u00a0", " ", -1),
			"rpc error: code = Internal desc = proto: invalid UTF-8 in string: \"\\xff\"")
	})
}

func TestGRPCServer_DeleteReceiver(t *testing.T) {
	providerId := uint64(10)
	dummyReq := &sirenv1.DeleteReceiverRequest{
		Id: uint64(10),
	}

	t.Run("should delete receiver object", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("DeleteReceiver", providerId).
			Return(nil).Once()

		res, err := dummyGRPCServer.DeleteReceiver(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, "", res.String())
	})

	t.Run("should return error code 13 if deleting receiver failed", func(t *testing.T) {
		mockedReceiverService := &mocks.ReceiverService{}
		dummyGRPCServer := GRPCServer{
			container: &service.Container{
				ReceiverService: mockedReceiverService,
			},
			logger: zaptest.NewLogger(t),
		}
		mockedReceiverService.
			On("DeleteReceiver", providerId).
			Return(errors.New("random error")).Once()

		res, err := dummyGRPCServer.DeleteReceiver(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})
}