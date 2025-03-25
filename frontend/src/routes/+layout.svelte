<script lang="ts">
  import "../app.css";

  let { children } = $props();
</script>

<div class="flex min-h-screen bg-gray-100">
  <div>
    <div class="min-h-full w-[250px] bg-red-500">test sidebar</div>
  </div>

  <main class="flex-1">
    {@render children()}
  </main>
</div>
