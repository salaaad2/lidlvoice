#!/bin/bash

eval "$(conda shell.bash hook)"
echo "arg1: $1 arg2: $2 arg3: $3"
conda activate /home/salade/.conda/envs/tts
python resources/main.py "$1" "$2" `pwd`/"$3"
conda deactivate