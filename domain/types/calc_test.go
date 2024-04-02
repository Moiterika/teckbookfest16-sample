package types

import (
	"testing"

	decimal "github.com/shopspring/decimal"
)

func TestProrate(t *testing.T) {
	type args struct {
		a  Amount
		bs []ProrationBasis[Unit]
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
					val:  decimal.NewFromInt32(200),
					unit: Jpy,
				},
				bs: []ProrationBasis[Unit]{
					Quantity{val: decimal.NewFromInt32(40), unit: 1},
					Quantity{val: decimal.NewFromInt32(30), unit: 1},
					Quantity{val: decimal.NewFromInt32(20), unit: 1},
					Quantity{val: decimal.NewFromInt32(10), unit: 1},
				},
			},
			wantAs: []Amount{
				{val: decimal.NewFromInt32(80), unit: Jpy},
				{val: decimal.NewFromInt32(60), unit: Jpy},
				{val: decimal.NewFromInt32(40), unit: Jpy},
				{val: decimal.NewFromInt32(20), unit: Jpy},
			},
			wantErr: false,
		},
		{
			name: "通常按分。端数誤差あり",
			args: args{
				a: Amount{
					val:  decimal.NewFromInt32(1000),
					unit: Jpy,
				},
				bs: []ProrationBasis[Unit]{
					Quantity{val: decimal.NewFromInt32(100), unit: 1},
					Quantity{val: decimal.NewFromInt32(100), unit: 1},
					Quantity{val: decimal.NewFromInt32(100), unit: 1},
				},
			},
			wantAs: []Amount{
				{val: decimal.NewFromInt32(334), unit: Jpy},
				{val: decimal.NewFromInt32(333), unit: Jpy},
				{val: decimal.NewFromInt32(333), unit: Jpy},
			},
			wantErr: false,
		},
		{
			name: "通常按分。按分なし",
			args: args{
				a: Amount{
					val:  decimal.NewFromInt32(1000),
					unit: Jpy,
				},
				bs: []ProrationBasis[Unit]{
					Quantity{val: decimal.NewFromInt32(25), unit: 1},
				},
			},
			wantAs: []Amount{
				{val: decimal.NewFromInt32(1000), unit: Jpy},
			},
			wantErr: false,
		},
		{
			name: "エラー。計算不能",
			args: args{
				a: Amount{
					val:  decimal.NewFromInt32(1000),
					unit: Jpy,
				},
				bs: make([]ProrationBasis[Unit], 0),
			},
			wantErr: true,
		},
		{
			name: "エラー。計算不能",
			args: args{
				a: Amount{
					val:  decimal.NewFromInt32(1000),
					unit: Jpy,
				},
				bs: nil,
			},
			wantErr: true,
		},
		{
			name: "エラー。按分基準の単位違い",
			args: args{
				a: Amount{
					val:  decimal.NewFromInt32(1000),
					unit: Jpy,
				},
				bs: []ProrationBasis[Unit]{
					Quantity{val: decimal.NewFromInt32(100), unit: 1},
					Quantity{val: decimal.NewFromInt32(100), unit: 2},
					Quantity{val: decimal.NewFromInt32(100), unit: 3},
				},
			},
			wantErr: true,
		},
		{
			name: "エラー。按分時に0除算発生",
			args: args{
				a: Amount{
					val:  decimal.NewFromInt32(1000),
					unit: Jpy,
				},
				bs: []ProrationBasis[Unit]{
					Quantity{val: decimal.NewFromInt32(0), unit: 1},
					Quantity{val: decimal.NewFromInt32(100), unit: 1},
					Quantity{val: decimal.NewFromInt32(-100), unit: 1},
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
