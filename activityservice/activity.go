package activityservice

import "context"

// ActivityService abstraction
type ActivityService interface {
	InsertLog(ctx context.Context, request *RequestInsertLog) (response ResponseLog, code int, err error)
	GetLog(ctx context.Context, param *LogParam) (response []ResponseLog, totaldata int, err error)
	GetLogByID(ctx context.Context, param *LogParam) (response ResponseLog, totaldata int, err error)
}
