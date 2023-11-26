package main

import (
	"fmt"
	"log"
)

// VM 정의
type VM struct {
	ID   string
	Name string
}

// Inventory 정의
type Inventory struct {
	Vms []VM
}

// NCP API 호출
func deleteMemberServerImageInstances(vmID string) error {
	// NCP API 호출 로직
	return nil
}

// CR 삭제 (CR 정보를 기반으로 Inventory 및 NCP에서 삭제)
func delete(crID string, inventory *Inventory) error {
	// Inventory에서 해당 CR에 대한 VM 정보 찾기 및 삭제
	var vmIDToDelete string
	for i, vm := range inventory.Vms {
		if vm.ID == crID {
			vmIDToDelete = vm.ID
			// 해당 VM을 Inventory에서 삭제
			inventory.Vms = append(inventory.Vms[:i], inventory.Vms[i+1:]...)
			break
		}
	}

	if vmIDToDelete == "" {
		return fmt.Errorf("VM with ID %s not found in the inventory", crID)
	}

	// NCP에서 VM 삭제
	err := deleteMemberServerImageInstances(vmIDToDelete)
	if err != nil {
		return fmt.Errorf("failed to delete VM from NCP: %v", err)
	}

	return nil
}

func main() {
	// 사용자가 입력한 CR ID
	crIDToDelete := "123"

	// VM 정보를 관리하는 Inventory 생성
	inventory := &Inventory{
		Vms: []VM{
			{ID: "123", Name: "VM1"},
			{ID: "456", Name: "VM2"},
			// 기존의 VM 정보들
		},
	}

	// CR 삭제
	err := delete(crIDToDelete, inventory)
	if err != nil {
		log.Fatal(err)
	}

	// 사용자에게 결과를 로그로 출력
	fmt.Printf("CR with ID %s successfully deleted. Inventory updated:\n", crIDToDelete)
	for _, vm := range inventory.Vms {
		fmt.Printf("VM ID: %s, VM Name: %s\n", vm.ID, vm.Name)
	}
}