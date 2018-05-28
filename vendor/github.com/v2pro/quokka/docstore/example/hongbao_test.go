package example

import (
	"github.com/v2pro/quokka/docstore"
	"testing"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"fmt"
)

func init() {
docstore.Entity("Hongbao", `
struct taken {
	1: i64 taken_at
	2: i64 taken_amount // unit fen
}
struct Entity {
	1: i64 total_count
	2: i64 total_amount // unit fen
	3: i64 remaining_amount // unit fen
	4: map<string, taken> takens
}
`).Command("create", `
function handle(entity, req) {
	entity.total_count = req.count;
	entity.total_amount = req.amount;
	entity.remaining_amount = req.amount;
	entity.takens = {};
	return {};
}`, `
struct Request {
	1: float64 amount
	2: i64 count
}
struct Response {
}
`).Command("take", `
function handle(entity, req) {
	if (entity.takens[req.username]) {
		throw 'one user can not take twice'
	}
	var takens_count = Object.keys(entity.takens).length;
	if (takens_count == entity.total_count) {
		throw 'nothing left'
	}
	if (takens_count == entity.total_count - 1) {
		// last one, take all
		var taken_amount = entity.remaining_amount;
		entity.remaining_amount = 0;
		entity.takens[req.username] = {taken_at: Date.now(), taken_amount: taken_amount};
		return {taken_amount: taken_amount};
	}
	var taken_amount = calcRandomAmount(entity.remaining_amount, entity.total_amount, entity.total_count);
	entity.remaining_amount -= taken_amount;
	entity.takens[req.username] = {
		taken_at: Date.now(),
		taken_amount: taken_amount
	};
	return {taken_amount: taken_amount};
}
function calcRandomAmount(remaining_amount, total_amount, total_count) {
	var cap_amount = 2 * total_amount / total_count;
	if (remaining_amount < cap_amount) {
		cap_amount = remaining_amount;
	}
	return Math.floor(randBetween(1, cap_amount));
}
function randBetween(min, max) {
	return Math.random() * (max - min) + min;
}`, `
struct Request {
	1: string username
}
struct Response {
	1: float64 taken_amount
}
`)
}

func Test_take_hongbao(t *testing.T) {
	should := require.New(t)
	execAndExpectSuccess(t, "/Hongbao/create",
		"EntityId", "123", "CommandRequest", runtime.NewObject("amount", 1314, "count", 3))
	resp := execAndExpectSuccess(t, "/Hongbao/take",
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "tom"))
	taken1 := jsoniter.Get(resp, "data", "taken_amount").ToInt()
	resp = execAndExpectError(t, "/Hongbao/take", docstore.ErrBusinessRuleViolated,
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "tom"))
	should.Equal("one user can not take twice", jsoniter.Get(resp, "errmsg").ToString())
	resp = execAndExpectSuccess(t, "/Hongbao/take",
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "jerry"))
	taken2 := jsoniter.Get(resp, "data", "taken_amount").ToInt()
	resp = execAndExpectSuccess(t, "/Hongbao/take",
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "donald"))
	taken3 := jsoniter.Get(resp, "data", "taken_amount").ToInt()
	should.Equal(1314, taken1+taken2+taken3)
	resp = execAndExpectError(t, "/Hongbao/take", docstore.ErrBusinessRuleViolated,
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "lily"))
	should.Equal("nothing left", jsoniter.Get(resp, "errmsg").ToString())
	resp = queryAndExpectSuccess(t, "/entities/Hongbao/123")
	fmt.Println(string(resp))
}
