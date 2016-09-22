package virtualmachine

import (
	"fmt"
	"encoding/xml"
	"github.com/megamsys/opennebula-go/api"
	"strconv"
)

type Vnc struct {
	VmId string
	T    *api.Rpc
	VM   *VM `xml:"VM"`
}

type VM struct {
	Id             string          `xml:"ID"`
	Name           string          `xml:"NAME"`
	VmTemplate     *VmTemplate     `xml:"TEMPLATE"`
	HistoryRecords *HistoryRecords `xml:"HISTORY_RECORDS"`
}

type VmTemplate struct {
	Graphics *Graphics `xml:"GRAPHICS"`
	Context  *Context  `xml:"CONTEXT"`
}

type Context struct {
	VMIP string `xml:"ETH0_IP"`
}

type HistoryRecords struct {
	History *History `xml:"HISTORY"`
}
type History struct {
	HostName string `xml:"HOSTNAME"`
}

type Graphics struct {
	Port string `xml:"PORT"`
}

func (v *Vnc) GetVm() (*VM, error) {
		fmt.Println("**********%%%%%%%%%%%%%%%%%*****  1")
	intstr, _ := strconv.Atoi(v.VmId)
	args := []interface{}{v.T.Key, intstr}
	onevm, err := v.T.Call(api.VM_INFO, args)
	defer v.T.Client.Close()
	if err != nil {
		return nil, err
	}
	fmt.Println("**********%%%%%%%%%%%%%%%%%*****  2")


	xmlVM := &VM{}
	if err = xml.Unmarshal([]byte(onevm), xmlVM); err != nil {
		return nil, err
	}
		fmt.Println("**********%%%%%%%%%%%%%%%%%*****",xmlVM)


	return xmlVM, err
}

func (u *VM) GetPort() string {
	return u.VmTemplate.Graphics.Port
}

func (u *VM) GetHostIp() string {
	return u.HistoryRecords.History.HostName
}

func (u *VM) GetVMIP() string {
	return u.VmTemplate.Context.VMIP
}
