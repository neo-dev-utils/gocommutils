package randUtil

import (
	"math/rand"

	engname "github.com/yelinaung/eng-name"

	"gocommutils/timeUtil"
)

func RandEnName() string {
	seed := timeUtil.Time.NowUnixNano()
	randName := engname.New(seed)
	rr := rand.New(rand.NewSource(timeUtil.Time.NowUnixNano()))
	if rr.Intn(1) > 0 {
		return randName.GetMenName()
	}
	return randName.GetWomenName()
}

func RandEnMenName() string {
	seed := timeUtil.Time.NowUnixNano()
	randName := engname.New(seed)
	return randName.GetMenName()
}

func RandEnWomenName() string {
	seed := timeUtil.Time.NowUnixNano()
	randName := engname.New(seed)
	return randName.GetWomenName()
}
