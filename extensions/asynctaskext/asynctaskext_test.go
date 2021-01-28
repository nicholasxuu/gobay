package asynctaskext

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/RichardKnop/machinery/v1/backends/result"
	"github.com/RichardKnop/machinery/v1/tasks"

	"github.com/shanbay/gobay"
)

var (
	task AsyncTaskExt
)

func init() {
	task = AsyncTaskExt{NS: "asynctask_"}

	app, _ := gobay.CreateApp(
		"../../testdata",
		"testing",
		map[gobay.Key]gobay.Extension{
			"asynctask": &task,
		},
	)
	if err := app.Init(); err != nil {
		log.Panic(err)
	}
}

func TestPushConsume(t *testing.T) {
	if err := task.RegisterWorkerHandlers(map[string]interface{}{
		"add": TaskAdd, "sub": TaskSub, "subCtx": TaskSubWithContext,
	}); err != nil {
		t.Error(err)
	}
	go func() {
		// use default queue
		if err := task.StartWorker("", 1); err != nil {
			t.Error(err)
		}
	}()
	go func() {
		if err := task.StartWorker("gobay.task_sub", 1); err != nil {
			t.Error(err)
		}
	}()

	signs := []*tasks.Signature{
		{
			Name: "add",
			Args: []tasks.Arg{ // use default queue
				{
					Type:  "int64",
					Value: 1,
				},
				{
					Type:  "int64",
					Value: 2,
				},
				{
					Type:  "int64",
					Value: 3,
				},
			},
		},
		{
			Name:       "sub",
			RoutingKey: "gobay.task_sub",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 7,
				},
				{
					Type:  "int64",
					Value: 1,
				},
			},
		},
		{
			Name:       "subCtx",
			RoutingKey: "gobay.task_sub",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 10,
				},
				{
					Type:  "int64",
					Value: 4,
				},
			},
		},
	}
	for _, sign := range signs {
		var (
			asyncResult *result.AsyncResult
			err         error
		)
		if sign.Name == "subCtx" {
			asyncResult, err = task.SendTaskWithContext(context.Background(), sign)
		} else {
			asyncResult, err = task.SendTask(sign)
		}
		if err != nil {
			t.Error(err)
		} else if results, err := asyncResult.Get(time.Millisecond * 5); err != nil {
			t.Error(err)
		} else if res, ok := results[0].Interface().(int64); !ok || res != 6 {
			t.Error("result error")
		}
	}
}
func TaskAdd(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func TaskSub(arg1, arg2 int64) (int64, error) {
	return arg1 - arg2, nil
}

func TaskSubWithContext(ctx context.Context, arg1, arg2 int64) (int64, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	return arg1 - arg2, nil
}
