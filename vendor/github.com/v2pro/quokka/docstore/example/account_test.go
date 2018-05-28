package example

import (
	"testing"
	"github.com/v2pro/quokka/docstore"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/docstore/runtime"
)

func defineAccount() {
	docstore.Entity("Account", `
	struct Entity {
		1: i64 amount
		2: string account_type
	}
	`).Command("create", `
	function handle(entity, req) {
		entity.amount = req.amount;
		entity.account_type = req.account_type;
		return {};
	}`, `
	struct Request {
		1: i64 amount
		2: string account_type
	}
	struct Response {
	}
	`).Command("charge", `
	function handle(entity, req) {
		if (entity.account_type == 'vip') {
			if (entity.amount - req.charge < -10) {
				throw 'vip account can not below -10';
			}
		} else {
			if (entity.amount - req.charge < 0) {
				throw 'normal account can not below 0';
			}
		}
		entity.amount -= req.charge;
		return {remaining_amount: entity.amount};
	}`, `
	struct Request {
		1: i64 charge
	}
	struct Response {
		1: i64 remaining_amount
	}
	`)
}

func Test_account_minimum_amount_rule(t *testing.T) {
	defineAccount()
	execAndExpectSuccess(t, "/Account/create",
		"EntityId", "123", "CommandRequest", runtime.NewObject("amount", 100, "account_type", "vip"))
	execAndExpectSuccess(t, "/Account/charge",
		"EntityId", "123", "CommandRequest", runtime.NewObject("charge", 10))
	execAndExpectError(t, "/Account/charge", docstore.ErrBusinessRuleViolated,
		"EntityId", "123", "CommandRequest", runtime.NewObject("charge", 101))
}

func Test_charge_idempotence(t *testing.T) {
	defineAccount()
	should := require.New(t)
	execAndExpectSuccess(t, "/Account/create",
		"EntityId", "123", "CommandRequest", runtime.NewObject("amount", 100, "account_type", "vip"))
	resp := execAndExpectSuccess(t, "/Account/charge",
		"EntityId", "123", "CommandId", "xcvf", "CommandRequest", runtime.NewObject("charge", 10))
	should.Equal(90, jsoniter.Get(resp, "data", "remaining_amount").ToInt())
	resp = execAndExpectSuccess(t, "/Account/charge",
		"EntityId", "123", "CommandId", "xcvf", "CommandRequest", runtime.NewObject("charge", 10))
	should.Equal(90, jsoniter.Get(resp, "data", "remaining_amount").ToInt())
}