<script lang="ts">
  import "../app.css";
  import Sidebar from "../components/Sidebar.svelte";

  let { children } = $props();
</script>

<div class="flex min-h-screen bg-gray-100">
  <Sidebar />

  <main class="flex-1 px-4 pt-10">
    {@render children()}
  </main>
</div>
