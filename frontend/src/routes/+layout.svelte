<script lang="ts">
  import "../app.css";

  let { children } = $props();
</script>

<div class="flex min-h-screen bg-gray-100">
  <aside class="w-[90px] p-4 lg:w-[250px]">
    <div
      class="sticky top-4 min-h-[calc(100vh-32px)] rounded-lg bg-white p-6 shadow-lg"
    >
      <h2 class="mb-4 text-xl font-bold">Sidebar</h2>

      <ul>
        <li class="mb-2">
          <a href="#" class="">1</a>
        </li>
        <li class="mb-2">
          <a href="#" class="">2</a>
        </li>
        <li class="mb-2">
          <a href="#" class="">3</a>
        </li>
      </ul>
    </div>
  </aside>

  <main class="flex-1 p-4">
    {@render children()}
  </main>
</div>
