package types

import (
	"testing"

	decimal "github.com/shopspring/decimal"
)

func TestProrate(t *testing.T) {
	type args struct {
		a  Amount
		bs []ProrationBasis[Code単位]
	}
	tests := []struct {
		name    string
		args    args
		wantAs  []Amount
		wantErr bool
	}{
		// Add test cases here.
		{
			name: "通常按分",
			args: args{
				a: Amount{
					val: decimal.NewFromInt32(200),
					cur: Jpy,
				},
				bs: []ProrationBasis[Code単位]{
					Quantity{val: decimal.NewFromInt32(40), unit: "kg"},
					Quantity{val: decimal.NewFromInt32(30), unit: "kg"},
					Quantity{val: decimal.NewFromInt32(20), unit: "kg"},
					Quantity{val: decimal.NewFromInt32(10), unit: "kg"},
				},
			},
			wantAs: []Amount{
				{val: decimal.NewFromInt32(80), cur: Jpy},
				{val: decimal.NewFromInt32(60), cur: Jpy},
				{val: decimal.NewFromInt32(40), cur: Jpy},
				{val: decimal.NewFromInt32(20), cur: Jpy},
			},
			wantErr: false,
		},
		{
			name: "通常按分。端数誤差あり",
			args: args{
				a: Amount{
					val: decimal.NewFromInt32(1000),
					cur: Jpy,
				},
				bs: []ProrationBasis[Code単位]{
					Quantity{val: decimal.NewFromInt32(100), unit: "kg"},
					Quantity{val: decimal.NewFromInt32(100), unit: "kg"},
					Quantity{val: decimal.NewFromInt32(100), unit: "kg"},
				},
			},
			wantAs: []Amount{
				{val: decimal.NewFromInt32(334), cur: Jpy},
				{val: decimal.NewFromInt32(333), cur: Jpy},
				{val: decimal.NewFromInt32(333), cur: Jpy},
			},
			wantErr: false,
		},
		{
			name: "通常按分。按分なし",
			args: args{
				a: Amount{
					val: decimal.NewFromInt32(1000),
					cur: Jpy,
				},
				bs: []ProrationBasis[Code単位]{
					Quantity{val: decimal.NewFromInt32(25), unit: "kg"},
				},
			},
			wantAs: []Amount{
				{val: decimal.NewFromInt32(1000), cur: Jpy},
			},
			wantErr: false,
		},
		{
			name: "エラー。計算不能",
			args: args{
				a: Amount{
					val: decimal.NewFromInt32(1000),
					cur: Jpy,
				},
				bs: make([]ProrationBasis[Code単位], 0),
			},
			wantErr: true,
		},
		{
			name: "エラー。計算不能",
			args: args{
				a: Amount{
					val: decimal.NewFromInt32(1000),
					cur: Jpy,
				},
				bs: nil,
			},
			wantErr: true,
		},
		{
			name: "エラー。按分基準の単位違い",
			args: args{
				a: Amount{
					val: decimal.NewFromInt32(1000),
					cur: Jpy,
				},
				bs: []ProrationBasis[Code単位]{
					Quantity{val: decimal.NewFromInt32(100), unit: "kg"},
					Quantity{val: decimal.NewFromInt32(100), unit: "CS"},
					Quantity{val: decimal.NewFromInt32(100), unit: "L"},
				},
			},
			wantErr: true,
		},
		{
			name: "エラー。按分時に0除算発生",
			args: args{
				a: Amount{
					val: decimal.NewFromInt32(1000),
					cur: Jpy,
				},
				bs: []ProrationBasis[Code単位]{
					Quantity{val: decimal.NewFromInt32(0), unit: "kg"},
					Quantity{val: decimal.NewFromInt32(100), unit: "kg"},
					Quantity{val: decimal.NewFromInt32(-100), unit: "kg"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotZList, err := Prorate(tt.args.a, tt.args.bs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc金額按分() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotZList, tt.wantZList) {
			// 	t.Errorf("Calc金額按分() = %v, want %v", gotZList, tt.wantZList)
			// }
			if len(gotZList) != len(tt.wantAs) {
				t.Errorf("len(zList) = %v, len(want) %v", len(gotZList), len(tt.wantAs))
			}

			for i := range gotZList {
				//fmt.Println(gotZList[i].String())
				if gotZList[i].String() != tt.wantAs[i].String() {
					t.Errorf("got = %s, want= %s", gotZList[i].String(), tt.wantAs[i].String())
				}
			}

		})
	}
}
