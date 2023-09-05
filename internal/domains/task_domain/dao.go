package task_domain

import (
	"context"
	"os"
	"strconv"
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
}

func newDao(filePath string) *dao {

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

	_, err = os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return &dao{
		taskInfo: "task",
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
	}
}

func (dao *dao) getTaskInfo(ctx context.Context) string {
	return dao.taskInfo
}

func (dao *dao) updateTaskFile(ctx context.Context, value int) (err error) {

	dao.files.task.mut.Lock()

	defer dao.files.task.mut.Unlock()

	var fileBin []uint8
	var valueCounter int

	fileBin, err = os.ReadFile(dao.files.task.relativePath)
	if err != nil {
		return
	}

	if len(fileBin) == 0 {
		fileBin = []uint8{48}
	}

	valueCounter, err = strconv.Atoi(string(fileBin))
	if err != nil {
		return
	}

	fileBinNew := []uint8(strconv.Itoa(valueCounter + value))

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
