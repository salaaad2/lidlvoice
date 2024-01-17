#!/bin/bash
# by salade


# Variables
ServerURL="http://127.0.0.1:8090"
ChatRoomPath="/api/currentChatRoom"
PostArgument="qrCodeUrl"
SA=("amazing", "breathtaking", "masterclass", "classic", "base")


for item in "${SA[@]}"; do
    echo 'curl --data "$PostArgument=$item" $ServerURL$ChatRoomPath'
    curl --data "$PostArgument=$item" $ServerURL$ChatRoomPath
    echo " "
    sleep 0.3
done
