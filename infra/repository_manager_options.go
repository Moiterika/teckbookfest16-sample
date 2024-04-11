package infra

import "techbookfest16-sample/infra/dao"

type repManagerArgs struct {
	wb品目  dao.Wb品目
	wb仕入品 dao.Wb品目仕入品
	wb製造品 dao.Wb品目製造品
	wb受払  dao.Wb受払
}

type repManagerOption func(*repManagerArgs)

func Wb品目(wb品目 dao.Wb品目) repManagerOption {
	return func(args *repManagerArgs) {
		args.wb品目 = wb品目
	}
}
func Wb仕入品(wb仕入品 dao.Wb品目仕入品) repManagerOption {
	return func(args *repManagerArgs) {
		args.wb仕入品 = wb仕入品
	}
}
func Wb製造品(wb製造品 dao.Wb品目製造品) repManagerOption {
	return func(args *repManagerArgs) {
		args.wb製造品 = wb製造品
	}
}
func Wb受払(wb受払 dao.Wb受払) repManagerOption {
	return func(args *repManagerArgs) {
		args.wb受払 = wb受払
	}
}
