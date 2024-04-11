package infra

import "techbookfest16-sample/infra/dao"

type repManagerArgs struct {
	wb品目 dao.Wb品目
}

type repManagerOption func(*repManagerArgs)

func Wb品目(wb品目 dao.Wb品目) repManagerOption {
	return func(args *repManagerArgs) {
		args.wb品目 = wb品目
	}
}
