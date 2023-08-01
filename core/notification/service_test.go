package notification_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	saltlog "github.com/goto/salt/log"
	"github.com/goto/siren/core/alert"
	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/notification/mocks"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/plugins/queues"
	"github.com/stretchr/testify/mock"
)

const testPluginType = "test"

func TestService_CheckAndInsertIdempotency(t *testing.T) {
	var (
		scope = "test-scope"
		key   = "test-key"
	)
	testCases := []struct {
		name    string
		setup   func(*mocks.IdempotencyRepository)
		scope   string
		key     string
		wantErr bool
	}{
		{
			name: "should return error if idempotency exist and success",
			setup: func(ir *mocks.IdempotencyRepository) {
				ir.EXPECT().InsertOnConflictReturning(mock.AnythingOfType("*context.emptyCtx"), scope, key).Return(nil, errors.ErrConflict)
			},
			scope:   scope,
			key:     key,
			wantErr: true,
		},
		{
			name: "should return error if repository returning some error",
			setup: func(ir *mocks.IdempotencyRepository) {
				ir.EXPECT().InsertOnConflictReturning(mock.AnythingOfType("*context.emptyCtx"), scope, key).Return(nil, errors.New("some error"))
			},
			scope:   scope,
			key:     key,
			wantErr: true,
		},
		{
			name: "should return id and nil error if no idempotency exists",
			setup: func(ir *mocks.IdempotencyRepository) {
				ir.EXPECT().InsertOnConflictReturning(mock.AnythingOfType("*context.emptyCtx"), scope, key).Return(&notification.Idempotency{
					ID: 1,
				}, nil)
			},
			scope:   scope,
			key:     key,
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockIdempotencyRepository := new(mocks.IdempotencyRepository)

			if tc.setup != nil {
				tc.setup(mockIdempotencyRepository)
			}

			ns := notification.NewService(saltlog.NewNoop(), notification.Config{}, nil, nil, nil, notification.Deps{IdempotencyRepository: mockIdempotencyRepository})

			_, err := ns.CheckAndInsertIdempotency(context.Background(), tc.scope, tc.key)

			if (err != nil) != tc.wantErr {
				t.Errorf("NotificationService.CheckAndInsertIdempotency() error = %v, wantErr %v", err, tc.wantErr)
			}

			mockIdempotencyRepository.AssertExpectations(t)
		})
	}
}

func TestService_Dispatch(t *testing.T) {
	tests := []struct {
		name    string
		n       notification.Notification
		setup   func(notification.Notification, *mocks.Repository, *mocks.LogService, *mocks.AlertService, *mocks.Queuer, *mocks.Dispatcher)
		wantErr bool
	}{
		{
			name:    "should return error if notification type is unknown",
			n:       notification.Notification{},
			wantErr: true,
		},
		{
			name: "should return error if repository return error",
			n: notification.Notification{
				Type: notification.TypeSubscriber,
				Labels: map[string]string{
					"k1": "v1",
				},
			},
			setup: func(n notification.Notification, r *mocks.Repository, _ *mocks.LogService, _ *mocks.AlertService, _ *mocks.Queuer, _ *mocks.Dispatcher) {
				r.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if dispatcher service return error",
			n: notification.Notification{
				Type: notification.TypeSubscriber,
				Labels: map[string]string{
					"k1": "v1",
				},
			},
			setup: func(n notification.Notification, r *mocks.Repository, _ *mocks.LogService, _ *mocks.AlertService, _ *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().PrepareMessage(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, nil, false, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if dispatcher service return empty results",
			n: notification.Notification{
				Type: notification.TypeSubscriber,
				Labels: map[string]string{
					"k1": "v1",
				},
			},
			setup: func(n notification.Notification, r *mocks.Repository, _ *mocks.LogService, _ *mocks.AlertService, _ *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().PrepareMessage(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, nil, false, nil)
			},
			wantErr: true,
		},
		{
			name: "should return error if log notifications return error",
			n: notification.Notification{
				Type: notification.TypeSubscriber,
				Labels: map[string]string{
					"k1": "v1",
				},
			},
			setup: func(n notification.Notification, r *mocks.Repository, l *mocks.LogService, _ *mocks.AlertService, _ *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().PrepareMessage(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.Message{{ID: "123"}}, []log.Notification{{ReceiverID: 123}}, false, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("log.Notification")).Return(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if update alerts silence status return error",
			n: notification.Notification{
				Type: notification.TypeSubscriber,
				Labels: map[string]string{
					"k1": "v1",
				},
			},
			setup: func(n notification.Notification, r *mocks.Repository, l *mocks.LogService, a *mocks.AlertService, _ *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().PrepareMessage(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.Message{{ID: "123"}}, []log.Notification{{ReceiverID: 123}}, false, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
				a.EXPECT().UpdateSilenceStatus(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("bool"), mock.AnythingOfType("bool")).Return(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if enqueue return error",
			n: notification.Notification{
				Type: notification.TypeSubscriber,
				Labels: map[string]string{
					"k1": "v1",
				},
			},
			setup: func(n notification.Notification, r *mocks.Repository, l *mocks.LogService, a *mocks.AlertService, q *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().PrepareMessage(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.Message{{ID: "123"}}, []log.Notification{{ReceiverID: 123}}, false, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
				a.EXPECT().UpdateSilenceStatus(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("bool"), mock.AnythingOfType("bool")).Return(nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return no error if enqueue success",
			n: notification.Notification{
				Type: notification.TypeSubscriber,
				Labels: map[string]string{
					"k1": "v1",
				},
			},
			setup: func(n notification.Notification, r *mocks.Repository, l *mocks.LogService, a *mocks.AlertService, q *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Create(mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("notification.Notification")).Return(n, nil)
				d.EXPECT().PrepareMessage(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.Message{{ID: "123"}}, []log.Notification{{ReceiverID: 123}}, false, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
				a.EXPECT().UpdateSilenceStatus(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("bool"), mock.AnythingOfType("bool")).Return(nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mockQueuer       = new(mocks.Queuer)
				mockRepository   = new(mocks.Repository)
				mockDispatcher   = new(mocks.Dispatcher)
				mockLogService   = new(mocks.LogService)
				mockAlertService = new(mocks.AlertService)
			)

			if tt.setup != nil {
				tt.setup(tt.n, mockRepository, mockLogService, mockAlertService, mockQueuer, mockDispatcher)
			}

			mockQueuer.EXPECT().Type().Return(queues.KindPostgres.String())
			s := notification.NewService(
				saltlog.NewNoop(),
				notification.Config{},
				mockRepository,
				mockQueuer,
				nil,
				notification.Deps{
					AlertService:              mockAlertService,
					LogService:                mockLogService,
					DispatchReceiverService:   mockDispatcher,
					DispatchSubscriberService: mockDispatcher,
				},
			)
			if err := s.Dispatch(context.TODO(), tt.n); (err != nil) != tt.wantErr {
				t.Errorf("Service.Dispatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_BuildFromAlerts(t *testing.T) {
	tests := []struct {
		name      string
		alerts    []alert.Alert
		firingLen int
		want      []notification.Notification
		errString string
	}{

		{
			name:      "should return empty notification if alerts slice is empty",
			errString: "empty alerts",
		},
		{
			name: "should properly return notification (same annotations are joined by newline and different labels are splitted into two notifications)",
			alerts: []alert.Alert{
				{
					ID:           14,
					ProviderID:   1,
					NamespaceID:  1,
					ResourceName: "test-alert-host-1",
					MetricName:   "test-alert",
					MetricValue:  "15",
					Severity:     "WARNING",
					Rule:         "test-alert-template",
					Labels:       map[string]string{"lk1": "lv1"},
					Annotations:  map[string]string{"ak1": "akv1"},
					Status:       "FIRING",
				},
				{
					ID:           15,
					ProviderID:   1,
					NamespaceID:  1,
					ResourceName: "test-alert-host-2",
					MetricName:   "test-alert",
					MetricValue:  "16",
					Severity:     "WARNING",
					Rule:         "test-alert-template",
					Labels:       map[string]string{"lk1": "lv1", "lk2": "lv2"},
					Annotations:  map[string]string{"ak1": "akv1"},
					Status:       "FIRING",
				},
				{
					ID:           16,
					ProviderID:   1,
					NamespaceID:  1,
					ResourceName: "test-alert-host-2",
					MetricName:   "test-alert",
					MetricValue:  "16",
					Severity:     "WARNING",
					Rule:         "test-alert-template",
					Labels:       map[string]string{"lk1": "lv1", "lk2": "lv2"},
					Annotations:  map[string]string{"ak1": "akv11", "ak2": "akv2"},
					Status:       "FIRING",
				},
			},
			firingLen: 2,
			want: []notification.Notification{
				{
					NamespaceID: 1,
					Type:        notification.TypeSubscriber,
					Data: map[string]any{
						"generator_url":     "",
						"num_alerts_firing": 2,
						"status":            "FIRING",
						"ak1":               "akv1",
					},
					Labels: map[string]string{
						"lk1": "lv1",
					},
					UniqueKey: "ignored",
					Template:  template.ReservedName_SystemDefault,
					AlertIDs:  []int64{14},
				},
				{
					NamespaceID: 1,
					Type:        notification.TypeSubscriber,

					Data: map[string]any{
						"generator_url":     "",
						"num_alerts_firing": 2,
						"status":            "FIRING",
						"ak1":               "akv1\nakv11",
						"ak2":               "akv2",
					},
					Labels: map[string]string{
						"lk1": "lv1",
						"lk2": "lv2",
					},
					UniqueKey: "ignored",
					Template:  template.ReservedName_SystemDefault,
					AlertIDs:  []int64{15, 16},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := notification.NewService(
				saltlog.NewNoop(),
				notification.Config{},
				nil,
				nil,
				nil,
				notification.Deps{},
			)
			got, err := s.BuildFromAlerts(tt.alerts, tt.firingLen, time.Time{})
			if (err != nil) && (err.Error() != tt.errString) {
				t.Errorf("BuildTypeReceiver() error = %v, wantErr %s", err, tt.errString)
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(notification.Notification{}, "ID", "UniqueKey")); diff != "" {
				t.Errorf("BuildFromAlerts() got diff = %v", diff)
			}
		})
	}
}
