package vego

import (
	"io/ioutil"
	"os"
	"path/filepath"

	vbg "github.com/uruddarraju/virtualbox-go"
)

// CreateVM function to create a new virtualbox VM.
func CreateVM() {
	// setup temp directory, that will be used to cache different VM related files during the creation of the VM.
	dirName, err := ioutil.TempDir("", "vbm")
	if err != nil {
		t.Errorf("Tempdir creation failed %v", err)
	}
	defer os.RemoveAll(dirName)

	vb := vbg.NewVBox(vbg.Config{
		BasePath: dirName,
	})

	disk1 := vbg.Disk{
		Path:   filepath.Join(dirName, "disk1.vdi"),
		Format: VDI,
		SizeMB: 10,
	}

	err = vb.CreateDisk(&disk1)
	if err != nil {
		t.Errorf("CreateDisk failed %v", err)
	}

	vm := &vbg.VirtualMachine{}
	vm.Spec.Name = "testvm1"
	vm.Spec.OSType = Linux64
	vm.Spec.CPU.Count = 2
	vm.Spec.Memory.SizeMB = 1000
	vm.Spec.Disks = []vbg.Disk{disk1}

	err = vb.CreateVM(vm)
	if err != nil {
		t.Fatalf("Failed creating vm %v", err)
	}

	err = vb.RegisterVM(vm)
	if err != nil {
		t.Fatalf("Failed registering vm")
	}
}
