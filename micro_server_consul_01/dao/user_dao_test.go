package dao

import (
	"fmt"
	"micro_server/models"
	"reflect"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/14 上午12:50
 * @Desc:
 */

func TestUserDao_GetUserByUserId(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Users
		wantErr bool
	}{
		{
			args:args{userId:1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewUserDao()
			got, err := d.GetUserByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			fmt.Printf("got: %+v", got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}