package utils

import (
	"github.com/sony/sonyflake"
	"time"
)

//雪花算法
func CreateSnowFlake() uint64 {
	t:=time.Unix(0,0)
	settings:=sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
		CheckMachineID: checkMachineID,
	}
	sf:=sonyflake.NewSonyflake(settings)
	id,_:=sf.NextID()
	return id
}
func getMachineID()(uint16,error){
	return uint16(1),nil
}

func checkMachineID(machineID uint16)bool{
	return true
}

