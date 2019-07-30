package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"net/http"
	"net/rpc"
)

// Action is an action such asn an alert, triggered by a given asset status
type Action struct {
}

// AssetStatus is the current status of a given asset, based on an evaluation of 1 or more conditions
type AssetStatus struct {
	StatusName string
	AssetName  string
	Evaluation []struct {
		ConditionID   int32
		Operator      string
		Value         string
		ChainOperator string
	}
	StatusValue string
	LastUpdated time.Time
}

// Condition is a set of rules evaluated against device values. E.g. temperature < 3 && temp_delta <2, including last changed date
type Condition struct {
	ConditionID   int32
	ConditionName string
	DeviceName    string
	Evaluation    []struct {
		Operand       string
		Operator      string
		Value         string
		ChainOperator string
	}
	ConditionValue string
	LastUpdated    time.Time
}

// Device is a struct representing the device values, including deltas and other calculated values
type Device struct {
	ClientID int32
	DeviceID int32
	Payload  []struct {
		Key        string
		Value      string
		PastValue1 string
		PastValue2 string
		PastValue3 string
		PastValue4 string
		MaxValue   string
		MinValue   string
		DeltaValue string
		MeanValue  string
		ModalValue string
	}
}

// Payload is expected to be a string of JSON from TTN
type Payload struct {
	ClientID      int32
	DeviceID      int32
	PayloadString string
}

// RACE stands for Rules Actions Conditions Engine - a kind of rules engine thing
type RACE int

// NewMsg receives a new payload from a sensor and uses it to update the state map
func (t *RACE) NewMsg(args *Payload, reply *int) (err error) {
	fmt.Println(args.PayloadString)

	// Update our values for this device
	if err := updateDeviceValues(args.ClientID, args.DeviceID, args.PayloadString); err != nil {
		fmt.Println(err)
		return err
	}
	// Identify what conditions evaluate this device
	conditions, err := listConditionsForDevice(args.ClientID, args.DeviceID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Evaluate all of those conditions
	for _, condition := range conditions {
		if err := evaluateCondition(condition); err != nil {
			fmt.Println(err)
			return err
		}
	}
	// Identify what asset statuses depend on those updated conditions
	assetStatuses, err := listAssetStatusesForConditions(conditions)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Update those asset statuses
	for _, assetStatus := range assetStatuses {
		if err := evaluateAssetStatus(assetStatus); err != nil {
			fmt.Println(err)
			return err
		}
	}
	// Identify which alerts / actions are triggered by changes in asset status
	actions, err := listActionsForAssetStatuses(assetStatuses)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Trigger those alerts / actions
	for idx, action := range actions {
		fmt.Printf("TODO: trigger %d actions for %v\n", idx, action)
	}

	*reply = 99
	return nil
}

func updateDeviceValues(clientID int32, deviceID int32, payloadString string) (err error) {
	fmt.Printf("TODO: update device values for %v\n", payloadString)
	return
}

func listConditionsForDevice(clientID int32, deviceID int32) (conditions []Condition, err error) {
	fmt.Printf("TODO: list conditions for device %v\n", deviceID)
	return
}

func evaluateCondition(condition Condition) (err error) {
	fmt.Printf("TODO: evaluate condition %v\n", condition)
	return
}

func listAssetStatusesForConditions(conditions []Condition) (assetStatuses []AssetStatus, err error) {
	fmt.Println("TODO: list asset statuses for the conditions")
	return
}

func evaluateAssetStatus(assetStatus AssetStatus) (err error) {
	fmt.Printf("TODO: uevaluate asset status: %v\n", assetStatus)
	return
}

func listActionsForAssetStatuses(assetStatuses []AssetStatus) (actions []Action, err error) {
	fmt.Println("TODO: list actions for the asset statuses")
	return
}

func main() {
	race := new(RACE)
	rpc.Register(race)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	defer l.Close()
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	// Make some devices
	Devices := make([]Device, 2)
	Devices[0] = Device{
		ClientID: 1,
		DeviceID: 48,
	}

	Devices[1] = Device{
		ClientID: 0,
		DeviceID: 49,
	}

	for {
		// Just hang around waiting for something to happen...
	}
}
