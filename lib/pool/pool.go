package pool

import (
	"math"

	"github.com/panjf2000/ants"
)

const DEFAULT_ANTS_POOL_SIZE = math.MaxInt32

func NewAntsPool() *ants.Pool {
	pool, err := ants.NewPool(DEFAULT_ANTS_POOL_SIZE) //新建一个pool对象，其他同上
	if err != nil {
		panic(err)
	}
	return pool
}
