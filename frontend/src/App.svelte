<!---
  LIDL_voice        ( //    
  svelte 'main'     ( )/    
  by salade         )(/     
 ________________  ( /)     
()__)____________)))))   :^}
--->

<script lang="ts">
/*******************************************************************************/
/* Logic                                                                       */
/*******************************************************************************/
import Login from "./lib/Login.svelte";
import Generator from "./lib/Generator.svelte";
import { currentUser } from "./lib/pocketbase";
</script>

<!---***************************************************************************/
/* Page Structure                                                              */
/****************************************************************************--->
<h1>LIDL_voice</h1>

<p>xtts-license: <a href="https://coqui.ai/cpml">license</a></p>

<Login />

{#if $currentUser}
  <Generator />
{/if}
