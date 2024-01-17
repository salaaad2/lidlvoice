<script lang="ts">
/*******************************************************************************/
/* Logic                                                                       */
/*******************************************************************************/
import { currentUser, pb } from './pocketbase';

let username: string;
let password: string;

async function login() {
  await pb.collection('users').authWithPassword(username, password);
}

function signOut() {
  pb.authStore.clear();
}
</script>

<!---***************************************************************************/
/* Page Structure                                                              */
/****************************************************************************--->
{#if $currentUser}
<p>
  Signed in as {$currentUser.username}
  <button on:click={signOut}>Sign Out</button>
</p>
{:else}
<form on:submit|preventDefault>
  <input
    placeholder="Username"
    type="text"
    bind:value={username}
  />

  <input 
    placeholder="Password" 
    type="password" 
    bind:value={password} 
  />

  <button on:click={login}>Login</button>
</form>
{/if}
