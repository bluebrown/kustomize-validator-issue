package main

import (
	"os"

	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/kio"
)

func main() {
	framework.Execute(proc(), &kio.ByteReadWriter{
		Reader:                os.Stdin,
		Writer:                os.Stdout,
		OmitReaderAnnotations: true,
		KeepReaderAnnotations: false,
		PreserveSeqIndent:     false,
		NoWrap:                true,
	})
}

func proc() framework.ResourceListProcessor {
	return framework.ResourceListProcessorFunc(func(rl *framework.ResourceList) error {
		return nil
	})
}
