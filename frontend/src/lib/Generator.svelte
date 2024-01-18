<!---
  LIDL_voice        ( //    
  message handler   ( )/    
  by salade         )(/     
 ________________  ( /)     
()__)____________)))))   :^}
--->

<script lang="ts">
/*******************************************************************************/
/* Logic                                                                       */
/*******************************************************************************/
import { onMount, onDestroy } from 'svelte';
import { currentUser, pb } from './pocketbase';

let newMessage: string;
let selectedVoice: string;
let voice_list = [];
let audioFileBlob;

onMount(async () => {
  const voiceList = pb.collection('voices').getList(1, 200, {
    sort: '-created'
  });
  voice_list = (await voiceList).items;
  console.log(voice_list);
});

onDestroy(() => {
});

async function addAudioFile(){
}

async function createRecording() {
  const formData = new URLSearchParams();
  formData.append("voice", selectedVoice);
  formData.append("providedText", newMessage);
  //console.log(parameter_name + ": " + parameter_value);

  const response = await fetch(pb.baseUrl + "/api/createRecording", {
    method: 'POST',
    headers: {
      'Content-type': 'application/x-www-form-urlencoded',
    },
    body: formData.toString(),
  });
  if (!response.ok)
  {
    throw new Error('Network response NOK');
  }
  audioFileBlob = await response.blob();
}

</script>

<!---***************************************************************************/
/* Page Structure                                                              */
/****************************************************************************--->
 
<div class="available-voices">
  <p class="msg-text"> Voix Disponibles : <br>
  {#each voice_list as voice (voice.id)}
      {voice.voice_name},   
  {/each}
  </p>
</div>

<div class="mediaPlayer">
  {#if audioFileBlob}
  <audio controls>
    <source src={URL.createObjectURL(audioFileBlob)} type="audio/mp3" />
  </audio>
  {/if}
</div>

<div class="msg_submit">
<form on:submit|preventDefault={createRecording}>
  <label for="dropdown"> Select a voice: </label>
  <select id="dropdown" bind:value={selectedVoice}>
  {#each voice_list as voice (voice.id)}
  <option value={voice.voice_name}>{voice.voice_name}</option>
  {/each}
</select>
  <textarea placeholder="Message" type="text" bind:value={newMessage} />
  <button type="submit">></button>
</form>
<!-- add send file thingy -->
</div>

