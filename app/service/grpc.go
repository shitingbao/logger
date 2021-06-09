package service

import (
	"context"
	"encoding/json"
	"logger/app/api"
	"logger/app/model"
	"logger/lib/snsq"
	"time"

	"google.golang.org/grpc/peer"
)

func (s *Service) SendOneLog(ctx context.Context, l *api.RequestLogMessages) (*api.RespondLogRes, error) {
	p, _ := peer.FromContext(ctx)

	m := &snsq.NsqMes{
		Sys:   l.Sys,
		Level: l.Level,
		Msg: snsq.NsqMesContent{
			Host:    p.Addr.Network() + ":" + p.Addr.String(),
			Topic:   l.Msg.Topic,
			Content: l.Msg.Content,
			Version: l.Version,
			LogTime: l.LogTime,
		},
	}
	if err := s.LogSeverNsq.LogPulish(m); err != nil {
		return nil, err
	}
	return &api.RespondLogRes{
		Code: 10000,
		Msg:  "success",
	}, nil
}

func (s *Service) SendManyIdenticalLog(ctx context.Context, l *api.RequestLogIdenticalMessageList) (*api.RespondLogRes, error) {
	p, _ := peer.FromContext(ctx)
	for _, v := range l.Msg {
		m := &snsq.NsqMes{
			Sys:   l.Sys,
			Level: l.Level,
			Msg: snsq.NsqMesContent{
				Host:    p.Addr.Network(),
				Topic:   v.Topic,
				Content: v.Content,
				Version: l.Version,
				LogTime: l.LogTime,
			},
		}
		s.LogSeverNsq.LogPulish(m)
	}
	return &api.RespondLogRes{
		Code: 10000,
		Msg:  "success",
	}, nil
}

func (s *Service) SendManyDifferentLog(ctx context.Context, l *api.RequestLogDifferentMessageList) (*api.RespondLogRes, error) {
	p, _ := peer.FromContext(ctx)
	for _, v := range l.Msg {
		m := &snsq.NsqMes{
			Sys:   v.Sys,
			Level: v.Level,
			Msg: snsq.NsqMesContent{
				Host:    p.Addr.Network(),
				Topic:   v.Msg.Topic,
				Content: v.Msg.Content,
				Version: v.Version,
				LogTime: v.LogTime,
			},
		}

		s.LogSeverNsq.LogPulish(m)
	}
	return &api.RespondLogRes{
		Code: 10000,
		Msg:  "success",
	}, nil
}

func (s *Service) LogFind(ctx context.Context, in *api.RequestLogFindParam) (*api.RespondLogFindList, error) {
	var err error
	timeStart, timeStop := time.Time{}, time.Time{}
	if in.LogStartTime != "" {
		timeStart, err = time.Parse("2006-01-02 15:04:05", in.LogStartTime)
		if err != nil {
			return nil, err
		}
	}
	if in.LogEndTime != "" {
		timeStop, err = time.Parse("2006-01-02 15:04:05", in.LogEndTime)
		if err != nil {
			return nil, err
		}
	}

	order := model.ArgOrder{
		OrderField: in.Order.OrderField,
		OrderVal:   int(in.Order.OrderVal),
	}
	arg := &model.ArgLogCondition{
		LogSys:       in.LogSys,
		LogStartTime: timeStart,
		LogEndTime:   timeStop,
		LogLevel:     in.LogLevel,
		Topic:        in.Topic,
		Content:      in.Content,
		Page:         in.Page,
		PageSize:     in.PageSize,
		Order:        order,
	}
	msg, err := s.Find(arg)
	if err != nil {
		return nil, err
	}
	resMsg, err := json.Marshal(msg)
	return &api.RespondLogFindList{
		Msg: string(resMsg),
	}, err
}
