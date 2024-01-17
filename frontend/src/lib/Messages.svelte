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
let selectedFileName: string;
let filename_list = [];
let audioFileBlob;

onMount(async () => {
  const filenameList = pb.collection('users').getList(1, 200, {
    sort: '-created'
  });
  filename_list = (await filenameList).items;
});

onDestroy(() => {
});

async function addAudioFile(){
}

async function createRecording() {
  const formData = new URLSearchParams();
  formData.append("filename", selectedFileName);
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
 
<div class="messages">
  {#each filename_list as filename (filename.id)}
    <div class="msg">
      <div>
        <p class="msg-text"></p>
      <img
        class="avatar"
        src={`https://avatars.dicebear.com/api/identicon/test.svg`}
        alt="avatar"
        width="40px"
        />  
      <b>@{filename}</b>
      </div>
    </div>
  {/each}

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
  <textarea placeholder="Filename" type="text" bind:value={selectedFileName} />
  <textarea placeholder="Message" type="text" bind:value={newMessage} />
  <button type="submit">></button>
</form>
<!-- add send file thingy -->
</div>

