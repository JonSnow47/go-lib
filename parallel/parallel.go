package parallel

import (
	"errors"
	"log"
	"sync"
)

type DoWorkPieceFunc func(piece int)

var (
	invalidParam = errors.New("Invalid Param")
)

// Parallelize is a very simple framework that allow for parallelizing
// N independent pieces of work.
func Parallelize(workers, pieces int, doWorkPiece DoWorkPieceFunc, additionalHandlers ...func(interface{})) error {
	if workers <= 0 || pieces <= 0 || doWorkPiece == nil {
		log.Println(invalidParam)
		return invalidParam
	}

	toProcess := make(chan int, pieces)
	for i := 0; i < pieces; i++ {
		toProcess <- i
	}
	close(toProcess)

	if pieces < workers {
		workers = pieces
	}

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			//defer utilruntime.HandleCrash(additionalHandlers...)
			defer wg.Done()
			for piece := range toProcess {
				doWorkPiece(piece)
			}
		}()
	}
	wg.Wait()
	return nil
}
