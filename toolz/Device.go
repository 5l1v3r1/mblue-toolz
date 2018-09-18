package toolz

import (
	"github.com/mame82/mblue-toolz/dbusHelper"
)

const dbusIfaceDevice = "org.bluez.Device1"

const (
	PropDeviceAddress          = "Address"          //readonly, string -> net.HardwareAddr
	PropDeviceAddressType      = "AddressType"      //readonly, string
	PropDeviceName             = "Name"             //readonly, optional, string
	PropDeviceIcon             = "Icon"             //readonly, optional, string
	PropDeviceClass            = "Class"            //readonly, optional, uint32
	PropDeviceAppearance       = "Appearance"       //readonly, optional, uint16
	PropDeviceUUIDs            = "UUIDs"            //readonly, optional, []string
	PropDevicePaired           = "Paired"           //readonly, bool
	PropDeviceConnected        = "Connected"        //readonly, bool
	PropDeviceTrusted          = "Trusted"          //readwrite, bool
	PropDeviceBlocked          = "Blocked"          //readwrite, bool
	PropDeviceAlias            = "Alias"            //readwrite, string
	PropDeviceAdapter          = "Adapter"          //readonly, ObjectPath
	PropDeviceLegacyPairing    = "LegacyPairing"    //readonly, bool
	PropDeviceModalias         = "Modalias"         //readonly, optional, string
	PropDeviceRSSI             = "RSSI"             //readonly, optional, uint16
	PropDeviceTxPower          = "TxPower"          //readonly, optional, uint16
	PropDeviceManufacturerData = "ManufacturerData" //readonly, optional, map[???]???
	PropDeviceServiceData      = "ServiceData"      //readonly, optional, map[string][]byte ??
	PropDeviceServicesResolved = "ServicesResolved" //readonly, bool
	PropDeviceAdvertisingFlags = "AdvertisingFlags" //readonly, experimental, []byte
	PropDeviceAdvertisingData  = "AdvertisingData"  //readonly, experimental, map[uint8][]byte ???
)

type Device1 struct {
	c *dbusHelper.Client
}

func (d *Device1) Close() {
	// closes CLients DBus connection
	d.c.Disconnect()
}

func (d *Device1) Connect() error {
	call, err := d.c.Call("Connect")
	if err != nil {
		return err
	}
	return call.Err
}

func (d *Device1) Disconnect() error {
	call, err := d.c.Call("Disconnect")
	if err != nil {
		return err
	}
	return call.Err
}

func (d *Device1) Pair() error {
	call, err := d.c.Call("Pair")
	if err != nil {
		return err
	}
	return call.Err
}

func (d *Device1) CancelPairing() error {
	call, err := d.c.Call("CancelPairing")
	if err != nil {
		return err
	}
	return call.Err
}


/* Properties */
func (d *Device1) GetTrusted() (res bool, err error) {
	val, err := d.c.GetProperty(PropDeviceTrusted)
	if err != nil {
		return
	}
	return val.Value().(bool), nil
}

func (d *Device1) SetTrusted(val bool) (err error) {
	return d.c.SetProperty(PropDeviceTrusted, val)
}

func (d *Device1) GetBlocked() (res bool, err error) {
	val, err := d.c.GetProperty(PropDeviceBlocked)
	if err != nil {
		return
	}
	return val.Value().(bool), nil
}

func (d *Device1) SetBlocked(val bool) (err error) {
	return d.c.SetProperty(PropDeviceBlocked, val)
}




func Device(devicePath string) (res *Device1, err error) {
	/*
	exists, err := Exists(deviceName)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, eDoesntExist
	}
	*/

	res = &Device1{
		c: dbusHelper.NewClient(dbusHelper.SystemBus, "org.bluez", dbusIfaceDevice, devicePath),
	}
	return
}
/*
func Exists(adapterName string) (exists bool, err error) {
	om, err := dbusHelper.NewObjectManager()
	if err != nil {
		return
	}
	defer om.Close()

	objs := om.GetManagedObjects()
	opath := dbus.ObjectPath("/org/bluez/" + adapterName)
	dev, exists := objs[opath]
	if !exists {
		return
	}
	_, exists = dev[dbusIfaceAdapter]
	return
}
*/