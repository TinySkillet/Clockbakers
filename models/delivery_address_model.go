package models

import (
	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/google/uuid"
)

type DeliveryAddress struct {
	ID      uuid.UUID `json:"id"`
	Address string    `json:"address" validate:"required"`
	UserID  uuid.UUID `json:"user_id" validate:"required"`
}

func DBDeliveryAddressToAddress(dbAddr database.DeliveryAddress) DeliveryAddress {
	return DeliveryAddress{
		ID:      dbAddr.ID,
		Address: dbAddr.Address,
		UserID:  dbAddr.UserID,
	}
}

func DBDeliveryAddressesToAddresses(dbAddrs []database.DeliveryAddress) []DeliveryAddress {
	addrs := make([]DeliveryAddress, len(dbAddrs))
	for i, dbAddr := range dbAddrs {
		addrs[i] = DBDeliveryAddressToAddress(dbAddr)
	}
	return addrs
}
