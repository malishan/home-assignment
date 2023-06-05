package impl

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/malishan/home-assignment/model"
	"github.com/malishan/home-assignment/utils/errors"
	"github.com/malishan/home-assignment/utils/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

func (impl *DbServiceImpl) PingStatus(ctx context.Context) *errors.Error {
	customErr := &errors.Error{
		StatusCode: http.StatusInternalServerError,
		Code:       errors.InternalServerErrorCode,
		Message:    errors.InternalServerError,
	}

	if impl.DBClient == nil {
		customErr.Details = "pingStatus db - nil client"
		return customErr
	}

	if err := impl.DBClient.PingContext(ctx); err != nil {
		customErr.Details = "pingStatus db - ping failed-" + err.Error()
		return customErr
	}

	return nil
}

func (impl *DbServiceImpl) InsertUserActivityRecord(ctx context.Context, id, userId string, record []*model.BoredApiResponse) *errors.Error {

	customErr := &errors.Error{
		StatusCode: http.StatusInternalServerError,
		Code:       errors.InternalServerErrorCode,
		Message:    errors.InternalServerError,
	}

	if impl.DBClient == nil {
		customErr.Details = "insertUserActivityRecord db - nil client"
		return customErr
	}
	if len(record) == 0 || id == "" {
		customErr.Details = "insertUserActivityRecord db - invalid func args"
		return customErr
	}

	valueStrings := make([]string, 0)
	valueArgs := make([]interface{}, 0)
	currentTime := time.Now()

	count := 0
	for _, eachRecord := range record {
		valueStrings = append(valueStrings, fmt.Sprintf("($%v, $%v, $%v, $%v, $%v)", count+1, count+2, count+3, count+4, count+5))
		valueArgs = append(valueArgs, id)
		valueArgs = append(valueArgs, userId)
		valueArgs = append(valueArgs, eachRecord.Key)
		valueArgs = append(valueArgs, eachRecord.Activity)
		valueArgs = append(valueArgs, currentTime)
		count += 5
	}

	smt := `INSERT INTO user_activity_record(id, user_id, key, activity, created_at) VALUES %s`
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	timer := prometheus.NewTimer(metrics.PsqlTiming.WithLabelValues("user_activity_record", "InsertUserActivityRecord"))
	defer timer.ObserveDuration()

	trx, err := impl.DBClient.BeginTx(ctx, nil)
	if err != nil {
		customErr.Details = "insertUserActivityRecord db - begin trx failed-" + err.Error()
		return customErr
	}

	_, err = trx.ExecContext(ctx, smt, valueArgs...)
	if err != nil {
		trx.Rollback()
		customErr.Details = "insertUserActivityRecord db - execution failed-" + err.Error()
		return customErr
	}

	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		customErr.Details = "insertUserActivityRecord db - commit failed-" + err.Error()
		return customErr
	}

	return nil
}

func (impl *DbServiceImpl) FetchUserActivityRecord(ctx context.Context) ([]*model.SchedulerResponse, *errors.Error) {

	customErr := &errors.Error{
		StatusCode: http.StatusInternalServerError,
		Code:       errors.InternalServerErrorCode,
		Message:    errors.InternalServerError,
	}

	if impl.DBClient == nil {
		customErr.Details = "fetchUserActivityRecord db - nil client"
		return nil, customErr
	}

	query := `select count(*) as total, key from user_activity_record group by key order by total desc`
	timer := prometheus.NewTimer(metrics.PsqlTiming.WithLabelValues("user_activity_record", "FetchUserActivityRecord"))
	rslt, err := impl.DBClient.QueryContext(ctx, query)
	timer.ObserveDuration()
	if err != nil {
		customErr.Details = "fetchUserActivityRecord db - query failed-" + err.Error()
		return nil, customErr
	}

	defer rslt.Close()

	records := make([]*model.SchedulerResponse, 0)

	for rslt.Next() {
		record := model.SchedulerResponse{}
		if err := rslt.Scan(&record.Count, &record.Key); err != nil {
			customErr.Details = "fetchUserActivityRecord db - scan failed-" + err.Error()
			return nil, customErr
		}
		records = append(records, &record)
	}

	return records, nil
}
