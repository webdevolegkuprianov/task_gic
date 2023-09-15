package task_domain

import (
	"context"
	"encoding/json"
	"os"
	"sync"
)

type dao struct {
	taskInfo string
	files    struct {
		task struct {
			relativePath string
			filePath     string
			mut          *sync.Mutex
		}
	}
	chTask chan int
}

func newDao(filePath string, ch chan int) *dao {

	var err error
	var currentPath string
	var path string

	currentPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	path = currentPath + "/" + filePath

	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {

			_, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}

		} else {
			panic(err)
		}
	}

	return &dao{
		taskInfo: data,
		files: struct {
			task struct {
				relativePath string
				filePath     string
				mut          *sync.Mutex
			}
		}{
			task: struct {
				relativePath string
				filePath     string
				mut          *sync.Mutex
			}{
				relativePath: path,
				filePath:     filePath,
				mut:          new(sync.Mutex),
			},
		},
		chTask: ch,
	}
}

func (dao *dao) getTaskInfo(ctx context.Context, data []byte) string {
	return dao.taskInfo
}

func (dao *dao) updateDurationFile(ctx context.Context, duration int64) (err error) {

	dao.files.task.mut.Lock()

	defer dao.files.task.mut.Unlock()

	var fileBin []uint8

	fileBin, err = os.ReadFile(dao.files.task.relativePath)
	if err != nil {
		return
	}

	var model *durationModel

	if len(fileBin) == 0 {

		model = &durationModel{
			Quantity: 1,
			Duration: duration,
		}

	} else {

		if err = json.Unmarshal(fileBin, &model); err != nil {
			return
		}

		model.Duration += duration
		model.Quantity++

	}

	var fileBinNew []uint8

	fileBinNew, err = json.Marshal(model)
	if err != nil {
		return
	}

	var f *os.File

	f, err = os.Create(dao.files.task.filePath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write(fileBinNew)
	if err != nil {
		return
	}

	return

}
