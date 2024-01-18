# separator: https://voice.ai/tools/vocal-remover

import sys

from TTS.api import TTS
tts = TTS("tts_models/multilingual/multi-dataset/xtts_v2", gpu=False)

if len(sys.argv) == 2:
    output_file = sys.argv[1]
elif len(sys.argv) == 3:
    output_file = sys.argv[1]
    message = sys.argv[2]
elif len(sys.argv) == 4:
    output_file = sys.argv[1]
    message = sys.argv[2]
    source_speaker = sys.argv[3]
else:
    message="Votez cocarde étudiante pour le cyclone nationaliste dans les universités"
    output_file="output.wav"
    source_speaker = "/home/salade/docs/projects/condavoice-online/backend/resources"

# generate speech by cloning a voice using default settings
tts.tts_to_file(text=message,
                file_path=output_file,
                speaker_wav=source_speaker,
                language="fr")
