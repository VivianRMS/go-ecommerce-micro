package account_test

import (
	v1 "github.com/VivianRMS/go-ecommerce-micro/account/gen/go/account/v1"
	v2 "github.com/VivianRMS/go-ecommerce-micro/account/gen/go/account/v2"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestProtoCompatibility(t *testing.T) {
	v1Msg := &v1.Account{Id: "1", Name: "Vivian"}
	data, _ := proto.Marshal(v1Msg)

	var v2Msg v2.Account
	_ = proto.Unmarshal(data, &v2Msg)
	t.Logf("[v1 ➡ v2] Parsed: %+v", v2Msg)

	v2Back := &v2.Account{Id: "2", Name: "Test", Type: v2.AccountType_ACCOUNT_TYPE_PERSONAL}
	data2, _ := proto.Marshal(v2Back)

	var v1Back v1.Account
	_ = proto.Unmarshal(data2, &v1Back)
	t.Logf("[v2 ➡ v1] Parsed: %+v", v1Back)
}
