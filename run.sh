#!/bin/bash

export http_proxy=http://localhost:8123
export https_proxy=http://localhost:8123

"$@"
