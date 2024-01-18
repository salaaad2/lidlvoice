# LIDL_voice

![how it looks running](./screen.png)

## How to run :

get 5-whatever long second extracts of voices.

## create a conda virtualenv named tts

```bash
conda create -n tts python=3.9
conda activate tts
python -m pip install TTS
conda deactivate
```

## Run LidlVoice
```bash
# open two terminals:
#1:
cd frontend 
npm i
npm run dev
#2:
cd backend
go build
./lidlvoice serve 
```