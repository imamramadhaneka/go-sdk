package sdk

import (
	"sync"

	"bitbucket.org/sehatQ/go-sdk/activityservice"
	"bitbucket.org/sehatQ/go-sdk/digiqare"
	"bitbucket.org/sehatQ/go-sdk/mqservice"
	"bitbucket.org/sehatQ/go-sdk/netcore"
	"bitbucket.org/sehatQ/go-sdk/paymentservice"
	"bitbucket.org/sehatQ/go-sdk/telemed"
	"bitbucket.org/sehatQ/go-sdk/userservice"
)

// Option func type
type Option func(*sdkInstance)

// SetMqService option func
func SetMqService(mqservice mqservice.MqService) Option {
	return func(s *sdkInstance) {
		s.mqservice = mqservice
	}
}

// SetDigiqare option func
func SetDigiqare(digiqare digiqare.Digiqare) Option {
	return func(s *sdkInstance) {
		s.digiqare = digiqare
	}
}

// SetTelemed option func
func SetTelemed(telemed telemed.Telemed) Option {
	return func(s *sdkInstance) {
		s.telemed = telemed
	}
}

// SetActivityService option func
func SetActivityService(activityservice activityservice.ActivityService) Option {
	return func(s *sdkInstance) {
		s.activityservice = activityservice
	}
}

// SetUserService option func
func SetUserService(userservice userservice.UserService) Option {
	return func(s *sdkInstance) {
		s.userservice = userservice
	}
}

// SetNetcore option func
func SetNetcore(netcore netcore.Netcore) Option {
	return func(s *sdkInstance) {
		s.netcore = netcore
	}
}

// SetPaymentService option func
func SetPaymentService(paymentservice paymentservice.PaymentService) Option {
	return func(s *sdkInstance) {
		s.paymentservice = paymentservice
	}
}

// SDK instance abstraction
type SDK interface {
	MqService() mqservice.MqService
	Digiqare() digiqare.Digiqare
	Telemed() telemed.Telemed
	ActivityService() activityservice.ActivityService
	UserService() userservice.UserService
	Netcore() netcore.Netcore
	PaymentService() paymentservice.PaymentService
}

// sdkInstance instance
type sdkInstance struct {
	mqservice       mqservice.MqService
	digiqare        digiqare.Digiqare
	telemed         telemed.Telemed
	activityservice activityservice.ActivityService
	userservice     userservice.UserService
	netcore         netcore.Netcore
	paymentservice  paymentservice.PaymentService
}

func (s *sdkInstance) MqService() mqservice.MqService {
	return s.mqservice
}

func (s *sdkInstance) Digiqare() digiqare.Digiqare {
	return s.digiqare
}

func (s *sdkInstance) Telemed() telemed.Telemed {
	return s.telemed
}

func (s *sdkInstance) ActivityService() activityservice.ActivityService {
	return s.activityservice
}

func (s *sdkInstance) UserService() userservice.UserService {
	return s.userservice
}

func (s *sdkInstance) Netcore() netcore.Netcore {
	return s.netcore
}

func (s *sdkInstance) PaymentService() paymentservice.PaymentService {
	return s.paymentservice
}

var (
	sdk  SDK
	once sync.Once
)

// SetGlobalSDK constructor with each service option.
/*
MqService
DigiQare
Telemed
ActivityService
UserService
Netcore
*/
func SetGlobalSDK(opts ...Option) {
	s := new(sdkInstance)

	for _, o := range opts {
		o(s)
	}
	once.Do(func() {
		sdk = s
	})
}

// GetSDK get global sdk instance
func GetSDK() SDK {
	return sdk
}
