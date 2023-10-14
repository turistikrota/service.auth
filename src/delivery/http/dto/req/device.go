package req

import "github.com/turistikrota/service.auth/src/app/command"

type DeviceRequest struct {
	DeviceUUID string `params:"device_uuid" validate:"required,uuid"`
}

func (d *DeviceRequest) ToDestroyCommand(uuid string) command.SessionDestroyCommand {
	return command.SessionDestroyCommand{
		UserUUID:   uuid,
		DeviceUUID: d.DeviceUUID,
	}
}

func (d *DeviceRequest) ToDestroyOthersCommand(uuid string) command.SessionDestroyOthersCommand {
	return command.SessionDestroyOthersCommand{
		UserUUID:   uuid,
		DeviceUUID: d.DeviceUUID,
	}
}

func (r *request) Device() *DeviceRequest {
	return &DeviceRequest{}
}
