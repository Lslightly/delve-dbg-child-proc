#!/bin/bash -x
go build -gcflags "-N -l"
cd utils
go build -gcflags "-N -l"
cd ../


	