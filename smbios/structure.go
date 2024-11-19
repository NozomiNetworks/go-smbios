// Copyright 2017-2018 NozomiNetworks.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package smbios

// A Header is a Structure's header.
type Header struct {
	Type   uint8
	Length uint8
	Handle uint16
}

//SystemEnclosure struct to fetch the chassis type information

type SystemEnclosure struct {
	Manufacturer byte

	Type byte
}

// A Structure is an SMBIOS structure.
type Structure struct {
	Header     Header
	Formatted  []byte
	Strings    []string
	SystemInfo SystemInfo
}

type SMBIOSSystemInfo struct {
	Manufacturer byte
	ProductName  byte
	Version      byte
	SN           byte
	UUID         [16]byte
	WakeUpType   byte
	SKUNumber    byte
	Family       byte
}

type SMBIOSBaseboardInfo struct {
	Manufacturer       byte
	Product            byte
	Version            byte
	SerialNumber       byte
	AssetTag           byte
	FeatureFlags       byte
	LocationInChassis  byte
	ChassisHandle      uint16
	BoardType          byte
	ObjectHandlesCount byte
	ObjectHandles      uintptr
}

type BaseboardInfo struct {
	Manufacturer string
	Product      string
	Version      string
	SerialNumber string
}

//FIXME and use review
type SMBIOSMemoryInfo struct {
	MemArrayHandle          uint16
	MemErrorInfoHandle      uint16
	TotalWidth              uint16
	DataWidth               uint16
	Size                    uint16
	FormFactor              byte
	DeviceSet               byte
	DeviceLocator           byte
	BankLocator             byte
	MemType                 byte
	TypeDetail              uint16
	Speed                   uint16
	SerialNumber            byte
	Manufacturer            byte
	AssetTag                byte
	PartNumber              byte
	Attribute               byte
	ExtendedSize            uint32
	ConfiguredMemClockSpeed uint16
	MinVoltage              uint16
	MaxVoltage              uint16
	ConfiguredVoltage       uint16
}

type MemoryInfoRead struct {
	MemArrayHandle          uint16
	MemErrorInfoHandle      uint16
	TotalWidth              uint16
	DataWidth               uint16
	Size                    uint16
	FormFactor              byte
	DeviceSet               byte
	DeviceLocator           byte
	BankLocator             byte
	MemType                 byte
	TypeDetail              uint16
	Speed                   uint16
	Manufacturer            byte
	SerialNumber            byte
	AssetTag                byte
	PartNumber              byte
	Attribute               byte
	ExtendedSize            uint32
	ConfiguredMemClockSpeed uint16
	MinVoltage              uint16
	MaxVoltage              uint16
	ConfiguredVoltage       uint16
}

type BIOSInfoRead struct {
	Vendor      byte
	Version     byte
	StartAddSeg uint16
	ReleaseDate byte
}

type BIOSInfo struct {
	Vendor      string
	Version     string
	BiosVersion string
	ReleaseDate string
}

type SystemInfoRead struct {
	Manufacturer string
	ProductName  string
	Version      string
	SerialNumber string
}

type SMBIOSProcessorInfo struct {
	SocketDesignation     byte
	ProcessorType         byte
	ProcessorFamily       byte
	ProcessorManufacturer byte
	ProcessorID           [8]byte
	ProcessorVersion      byte
}

type SMBIOSProcessorType struct {
	SocketDesignation        byte
	ProcessorType            byte
	ProcessorFamily          byte
	ProcessorManufacturer    byte
	ProcessorID              [4]uint16
	ProcessorVersion         byte
	Voltage                  byte
	ExternalClock            uint16
	MaxSpeed                 uint16
	CurrentSpeed             uint16
	Status                   byte
	ProcessorUpgrade         byte
	L1CacheHandle            uint16
	L2CacheHandle            uint16
	L3CacheHandle            uint16
	SerialNumber             byte
	AssetTag                 byte
	PartNumber               byte
	CoreCount                byte
	CoreEnabled              byte
	ThreadCount              byte
	ProcessorCharacteristics uint16
	ProcessorFamily2         uint16
	CoreCount2               uint16
	COreEnabled2             uint16
	ThreadCount2             uint16
}

type Processor struct {
	ID              string
	Family          int
	Product         string
	CoreCount       int
	ClockSpeedInMHz int
}

type SystemInfo struct {
	SystemManufacturerRef, SystemProductName, BiosSerial, VirtualMachineUUID,
	MotherboardAdapter, Memory, ProcessorType, ProcessorID string
	BiosInfo        *BIOSInfo
	BaseboardInfo   *BaseboardInfo
	Processors      []*Processor
	PhyMemory       []*PhysicalMemory
	SystemEnclosure *SystemEnclosure
}

type PhysicalMemory struct {
	Manufacturer, SerialNumber string
	SizeInBytes                uint64
	TotalWidth                 uint64
	DataWidth                  uint64
}

type SMBIOSVersion struct {
	Major, Minor, Revision int
}
