package v1beta1

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/odpf/salt/log"
	"github.com/odpf/siren/pkg/errors"
	sirenv1beta1 "github.com/odpf/siren/proto/odpf/siren/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	newrelic *newrelic.Application
	logger   log.Logger
	sirenv1beta1.UnimplementedSirenServiceServer
	templateService     TemplateService
	ruleService         RuleService
	alertService        AlertService
	providerService     ProviderService
	namespaceService    NamespaceService
	receiverService     ReceiverService
	subscriptionService SubscriptionService
}

func NewGRPCServer(
	nr *newrelic.Application, logger log.Logger,
	templateService TemplateService,
	ruleService RuleService,
	alertService AlertService,
	providerService ProviderService,
	namespaceService NamespaceService,
	receiverService ReceiverService,
	subscriptionService SubscriptionService) *GRPCServer {
	return &GRPCServer{
		newrelic:            nr,
		logger:              logger,
		templateService:     templateService,
		ruleService:         ruleService,
		alertService:        alertService,
		providerService:     providerService,
		namespaceService:    namespaceService,
		receiverService:     receiverService,
		subscriptionService: subscriptionService,
	}
}

func (s *GRPCServer) generateRPCErr(e error) error {
	s.logger.Error("grpc error", "err", errors.Verbose(e))
	err := errors.E(e)

	var code codes.Code
	switch {
	case errors.Is(err, errors.ErrNotFound):
		code = codes.NotFound

	case errors.Is(err, errors.ErrConflict):
		code = codes.AlreadyExists

	case errors.Is(err, errors.ErrInvalid):
		code = codes.InvalidArgument

	default:
		code = codes.Internal
	}
	return status.Error(code, err.Error())
}