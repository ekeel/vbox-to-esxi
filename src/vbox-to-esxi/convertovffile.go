package main

import (
	"io/ioutil"
	"strings"
)

var (
	ide0Text = `<Item>
	<rasd:Address>0</rasd:Address>
	<rasd:Caption>ideController0</rasd:Caption>
	<rasd:Description>IDE Controller</rasd:Description>
	<rasd:ElementName>ideController0</rasd:ElementName>
	<rasd:InstanceID>3</rasd:InstanceID>
	<rasd:ResourceSubType>PIIX4</rasd:ResourceSubType>
	<rasd:ResourceType>5</rasd:ResourceType>
  </Item>`

	ide0Fill = `<Item>
	<rasd:Address>0</rasd:Address>
	<rasd:Caption>SCSIController</rasd:Caption>
	<rasd:Description>SCSI Controller</rasd:Description>
	<rasd:ElementName>SCSIController</rasd:ElementName>
	<rasd:InstanceID>3</rasd:InstanceID>
	<rasd:ResourceSubType>lsilogic</rasd:ResourceSubType>
	<rasd:ResourceType>5</rasd:ResourceType>
  </Item>`

	strgCntrlText = `<StorageControllers>
	<StorageController name="IDE Controller" type="PIIX4" PortCount="2" useHostIOCache="true" Bootable="true">
	  <AttachedDevice type="HardDisk" hotpluggable="false" port="0" device="0">
		<Image uuid="{5b06e4e5-e76f-4e08-b493-a8e4c5a2889f}"/>
	  </AttachedDevice>
	</StorageController>
  </StorageControllers>`

	strgCntrlFill = `<StorageControllers>
	<StorageController name="SCSIController" type="lsilogic" PortCount="2" useHostIOCache="true" Bootable="true">
	  <AttachedDevice type="HardDisk" hotpluggable="false" port="0" device="0">
		<Image uuid="{5b06e4e5-e76f-4e08-b493-a8e4c5a2889f}"/>
	  </AttachedDevice>
	</StorageController>
  </StorageControllers>`
)

func ConvertOVFFile(file string, output string) error {
	ovfContent, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// addressMatchGroup := AddressRegex.FindStringSubmatch(string(ovfContent))

	// address := addressMatchGroup[1]

	// ovfContent = VirtualSystemTypeRegex.ReplaceAll(ovfContent, []byte("<vssd:VirtualSystemType>vmx-07</vssd:VirtualSystemType>"))
	// ovfContent = CaptionRegex.ReplaceAll(ovfContent, []byte("<rasd:Caption>SCSIController"+address+"</rasd:Caption>"))
	// ovfContent = DescriptionRegex.ReplaceAll(ovfContent, []byte("<rasd:Description>SCSI Controller</rasd:Description>"))
	// ovfContent = ElementNameRegex.ReplaceAll(ovfContent, []byte("<rasd:ElementName>SCSIController"+address+"</rasd:ElementName>"))
	// ovfContent = ResourceSubTypeRegex.ReplaceAll(ovfContent, []byte("<rasd:ResourceSubType>lsilogic</rasd:ResourceSubType>"))
	// ovfContent = ResourceTypeRegex.ReplaceAll(ovfContent, []byte("<rasd:ResourceType>6</rasd:ResourceType>"))

	ovfResult := strings.ReplaceAll(string(ovfContent), ide0Text, ide0Fill)
	ovfResult = strings.ReplaceAll(ovfResult, strgCntrlText, strgCntrlFill)

	err = ioutil.WriteFile(output, []byte(ovfResult), 0644)
	if err != nil {
		panic(err)
	}

	return err
}
