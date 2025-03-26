package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$uqzXWhGLSDgO/COjzCOsMOW2cJt/uynr8HxwNfINt018Fdr4eLwIG", "123456"))
}
