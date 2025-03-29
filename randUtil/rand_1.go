package randUtil

import (
	"math/rand"
	"strconv"

	"gocommutils/timeUtil"
)

var (
	Rand    = _rand{}
	RandPtr = &Rand
)

type _rand struct {
}

func (r _rand) GetRand() *rand.Rand {
	return rand.New(rand.NewSource(timeUtil.Time.NowUnixNano()))
}

func (r _rand) GenOrderNo() string {
	rr := r.GetRand()
	rd := rr.Intn(89999) + 10000
	return strconv.Itoa(int(timeUtil.Time.NowUnix()*100000) + rd)
}

// ***********************************************************
// 生成随机的字符串 包含大写字母/数字
// ***********************************************************

// 设计一个随机种子相加
var radonParam int64 = 0

func (ra _rand) RandomString(size int) string {
	radonParam += 1
	if radonParam > 1000000000 {
		radonParam = 0
	}
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(radonParam + timeUtil.Time.GetTimeStampMilsecd()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// ***********************************************************
// 生成随机的字符串 只包含小写字母
// ***********************************************************
func (ra _rand) RandomStringMix(size int) string {
	radonParam += 1
	if radonParam > 1000000000 {
		radonParam = 0
	}
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(radonParam + timeUtil.Time.GetTimeStampMilsecd()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// *****************************************************************************
// 产生int类型的随机数,在对应范围内
// *****************************************************************************

func (ra _rand) RandomInt(cknum int) int {
	if cknum < 1 {
		return 0
	}
	radonParam += 1
	if radonParam > 1000000000 {
		radonParam = 0
	}
	randExecute := rand.New(rand.NewSource(radonParam + timeUtil.Time.GetTimeStampMilsecd()))
	return randExecute.Intn(cknum)
}
