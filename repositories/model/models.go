package model

// Models Add the new model here so auto migration can reference to it
var Models = []interface{}{
	&Employee{},
	&Account{},
	&Auth{},
	&LeaveBalances{},
	&LeaveRequest{},
	&LeaveType{},
	&PasswordReset{},
}
