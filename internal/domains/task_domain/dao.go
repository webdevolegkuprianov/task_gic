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

	var dataBin []uint8
	var valueCounter int

	dataBin, err = os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if len(dataBin) == 0 {
		dataBin = []uint8{48}
	}

	valueCounter, err = strconv.Atoi(string(dataBin))
	if err != nil {
		panic(err)
	}

	go func() {
		ch <- valueCounter
	}()

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
		chTask: ch,
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

	newValueCounter := valueCounter + value

	fileBinNew := []uint8(strconv.Itoa(newValueCounter))

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

	dao.chTask <- newValueCounter

	return

}
