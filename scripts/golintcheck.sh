#!/bin/bash

golint -set_exit_status ./...
exit $?