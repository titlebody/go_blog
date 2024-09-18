package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	HashPwd("123456")
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$VM/q/HCAEHucXKWGaflQJeEBxoMWBTv4czC.wQmYVN/yh46Qiloz2", "123456"))

}
