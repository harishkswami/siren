package v1beta1

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/odpf/salt/log"
	"github.com/odpf/siren/core/subscription"
	"github.com/odpf/siren/internal/server/v1beta1/mocks"
	"github.com/stretchr/testify/assert"
	sirenv1beta1 "go.buf.build/odpf/gw/odpf/proton/odpf/siren/v1beta1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGRPCServer_ListSubscriptions(t *testing.T) {
	configuration := make(map[string]string)
	configuration["foo"] = "bar"
	match := make(map[string]string)
	match["foo"] = "bar"

	t.Run("should return list of all subscriptions", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		dummyResult := []*subscription.Subscription{
			{
				Id:        1,
				Urn:       "foo",
				Namespace: 1,
				Receivers: []subscription.ReceiverMetadata{{Id: 1, Configuration: configuration}},
				Match:     match,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		mockedSubscriptionService.EXPECT().ListSubscriptions(context.Background()).Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.ListSubscriptions(context.Background(), &emptypb.Empty{})
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetSubscriptions()))
		assert.Equal(t, uint64(1), res.GetSubscriptions()[0].GetId())
		assert.Equal(t, "bar", res.GetSubscriptions()[0].GetMatch()["foo"])
	})

	t.Run("should return error code 13 if getting subscriptions fails", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		mockedSubscriptionService.EXPECT().ListSubscriptions(context.Background()).
			Return(nil, errors.New("random error")).Once()
		res, err := dummyGRPCServer.ListSubscriptions(context.Background(), &emptypb.Empty{})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})
}

func TestGRPCServer_GetSubscription(t *testing.T) {
	configuration := make(map[string]string)
	configuration["foo"] = "bar"
	match := make(map[string]string)
	match["foo"] = "bar"

	t.Run("should return a subscription", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		dummyResult := &subscription.Subscription{
			Id:        1,
			Urn:       "foo",
			Namespace: 1,
			Receivers: []subscription.ReceiverMetadata{{Id: 1, Configuration: configuration}},
			Match:     match,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockedSubscriptionService.EXPECT().GetSubscription(context.Background(), uint64(1)).Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.GetSubscription(context.Background(), &sirenv1beta1.GetSubscriptionRequest{Id: 1})
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), res.GetId())
		assert.Equal(t, "bar", res.GetMatch()["foo"])
	})

	t.Run("should return error code 5 if subscriptions not found", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		mockedSubscriptionService.EXPECT().GetSubscription(context.Background(), uint64(1)).Return(nil, nil).Once()
		res, err := dummyGRPCServer.GetSubscription(context.Background(), &sirenv1beta1.GetSubscriptionRequest{Id: 1})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = NotFound desc = subscription not found")
	})

	t.Run("should return error code 13 if getting subscription fails", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		mockedSubscriptionService.EXPECT().GetSubscription(context.Background(), uint64(1)).
			Return(nil, errors.New("random error")).Once()
		res, err := dummyGRPCServer.GetSubscription(context.Background(), &sirenv1beta1.GetSubscriptionRequest{Id: 1})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})
}

func TestGRPCServer_CreateSubscription(t *testing.T) {
	configuration := make(map[string]string)
	configuration["foo"] = "bar"
	match := make(map[string]string)
	match["foo"] = "baz"

	payload := &subscription.Subscription{
		Namespace: 1,
		Urn:       "foo",
		Receivers: []subscription.ReceiverMetadata{{Id: 1, Configuration: configuration}},
		Match:     match,
	}

	t.Run("should create a subscription", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		dummyResult := &subscription.Subscription{
			Id:        1,
			Urn:       "foo",
			Namespace: 10,
			Receivers: []subscription.ReceiverMetadata{{Id: 1, Configuration: configuration}},
			Match:     match,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockedSubscriptionService.EXPECT().CreateSubscription(context.Background(), payload).Return(nil).
			Run(func(ctx context.Context, s *subscription.Subscription) {
				*s = *dummyResult
			}).Once()
		res, err := dummyGRPCServer.CreateSubscription(context.Background(), &sirenv1beta1.CreateSubscriptionRequest{
			Namespace: 1,
			Urn:       "foo",
			Receivers: []*sirenv1beta1.ReceiverMetadata{{Id: 1, Configuration: configuration}},
			Match:     match,
		})
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), res.GetId())
		assert.Equal(t, uint64(10), res.GetNamespace())
		assert.Equal(t, 1, len(res.GetMatch()))
		assert.Equal(t, "baz", res.GetMatch()["foo"])
		assert.Equal(t, "bar", res.GetReceivers()[0].GetConfiguration()["foo"])
	})

	t.Run("should return error code 13 if creating subscriptions fails", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}

		mockedSubscriptionService.EXPECT().CreateSubscription(context.Background(), payload).
			Return(errors.New("random error")).Once()
		res, err := dummyGRPCServer.CreateSubscription(context.Background(), &sirenv1beta1.CreateSubscriptionRequest{
			Namespace: 1,
			Urn:       "foo",
			Receivers: []*sirenv1beta1.ReceiverMetadata{{Id: 1, Configuration: configuration}},
			Match:     match,
		})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})
}

func TestGRPCServer_UpdateSubscription(t *testing.T) {
	configuration := make(map[string]string)
	configuration["foo"] = "bar"
	match := make(map[string]string)
	match["foo"] = "baz"
	payload := &subscription.Subscription{
		Id:        1,
		Namespace: 10,
		Urn:       "foo",
		Receivers: []subscription.ReceiverMetadata{{Id: 1, Configuration: configuration}},
		Match:     match,
	}

	t.Run("should update a subscription", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		dummyResult := &subscription.Subscription{
			Id:        1,
			Urn:       "foo",
			Namespace: 10,
			Receivers: []subscription.ReceiverMetadata{{Id: 1, Configuration: configuration}},
			Match:     match,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockedSubscriptionService.EXPECT().UpdateSubscription(context.Background(), payload).Return(nil).
			Run(func(ctx context.Context, s *subscription.Subscription) {
				*s = *dummyResult
			}).Once()
		res, err := dummyGRPCServer.UpdateSubscription(context.Background(), &sirenv1beta1.UpdateSubscriptionRequest{
			Id:        1,
			Namespace: 10,
			Urn:       "foo",
			Receivers: []*sirenv1beta1.ReceiverMetadata{{Id: 1, Configuration: configuration}},
			Match:     match,
		})
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), res.GetId())
		assert.Equal(t, uint64(10), res.GetNamespace())
		assert.Equal(t, 1, len(res.GetMatch()))
		assert.Equal(t, "baz", res.GetMatch()["foo"])
		assert.Equal(t, "bar", res.GetReceivers()[0].GetConfiguration()["foo"])
	})

	t.Run("should return error code 13 if creating subscriptions fails", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		mockedSubscriptionService.EXPECT().UpdateSubscription(context.Background(), payload).Return(errors.New("random error")).Once()
		res, err := dummyGRPCServer.UpdateSubscription(context.Background(), &sirenv1beta1.UpdateSubscriptionRequest{
			Id:        1,
			Namespace: 10,
			Urn:       "foo",
			Receivers: []*sirenv1beta1.ReceiverMetadata{{Id: 1, Configuration: configuration}},
			Match:     match,
		})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
	})

	t.Run("should return error code 5 for bad requests", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}
		mockedSubscriptionService.EXPECT().UpdateSubscription(context.Background(), payload).Return(
			errors.New(`violates unique constraint "urn_provider_id_unique"`)).Once()
		res, err := dummyGRPCServer.UpdateSubscription(context.Background(), &sirenv1beta1.UpdateSubscriptionRequest{
			Id:        1,
			Namespace: 10,
			Urn:       "foo",
			Receivers: []*sirenv1beta1.ReceiverMetadata{{Id: 1, Configuration: configuration}},
			Match:     match,
		})
		assert.Nil(t, res)
		assert.EqualError(t, err, `rpc error: code = InvalidArgument desc = urn and provider pair already exist`)
	})
}

func TestGRPCServer_DeleteSubscription(t *testing.T) {
	t.Run("should delete a subscription", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}

		mockedSubscriptionService.EXPECT().DeleteSubscription(context.Background(), uint64(1)).Return(nil).Once()
		res, err := dummyGRPCServer.DeleteSubscription(context.Background(), &sirenv1beta1.DeleteSubscriptionRequest{Id: 1})
		assert.Nil(t, err)
		assert.Equal(t, &emptypb.Empty{}, res)
	})

	t.Run("should return error code 13 if deleting subscription fails", func(t *testing.T) {
		mockedSubscriptionService := &mocks.SubscriptionService{}
		dummyGRPCServer := GRPCServer{
			subscriptionService: mockedSubscriptionService,
			logger:              log.NewNoop(),
		}

		mockedSubscriptionService.EXPECT().DeleteSubscription(context.Background(), uint64(1)).Return(errors.New("random error")).Once()
		res, err := dummyGRPCServer.DeleteSubscription(context.Background(), &sirenv1beta1.DeleteSubscriptionRequest{Id: 1})
		assert.EqualError(t, err, "rpc error: code = Internal desc = random error")
		assert.Nil(t, res)
	})
}