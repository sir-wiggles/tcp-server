package server

import (
	"log"
	"strconv"
	"strings"
)

type VincenzoDataHandler struct{}

func (v VincenzoDataHandler) Handle(data string) error {
	data = strings.Trim(data, "\n")
	fields := strings.Split(data, ",")
	switch len(fields) {
	case 2:
		return v.DockRack(fields...)
	case 3:
		return v.SlotInventory(fields...)
	}
	return nil
}

func (v VincenzoDataHandler) DockRack(fields ...string) error {
	ids, err := stringToIDs(fields)
	if err != nil {
		return err
	}
	rackID := ids[0]
	dockID := ids[1]
	log.Printf("docking rack %d into dock %d", rackID, dockID)
	return nil
}

func (v VincenzoDataHandler) SlotInventory(fields ...string) error {
	ids, err := stringToIDs(fields)
	if err != nil {
		return err
	}
	inventoryID := ids[0]
	rackID := ids[1]
	slotID := ids[2]
	log.Printf("slotting inventory %d into slot %d of rack %d", inventoryID, slotID, rackID)

	return nil
}

func stringToIDs(ids []string) ([]int64, error) {
	values := make([]int64, 0, len(ids))
	for _, id := range ids {
		if value, err := strconv.ParseInt(id, 10, 64); err != nil {
			return nil, err
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}
