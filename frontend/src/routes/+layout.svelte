<script lang="ts">
  import "../app.css";
  import Sidebar from "../components/Sidebar.svelte";

  let { children } = $props();
</script>

<div class="mx-auto flex min-h-screen max-w-[1950px]">
  <Sidebar />

  <main class="flex-1 px-4 py-10">
    {@render children()}
  </main>
</div>
