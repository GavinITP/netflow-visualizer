<script lang="ts">
  import SidebarItem from "./SidebarItem.svelte";

  const menuItems = [
    { href: "/", title: "Overview", iconSrc: "/home.svg" },
    {
      href: "/anomaly-logs",
      title: "Anomaly Traffic Logs",
      iconSrc: "/anomaly-logs.svg",
    },
    {
      href: "/db-stats",
      title: "Database Stats",
      iconSrc: "/database-stats.svg",
    },
  ];
</script>

<aside
  class="w-[90px] p-4 transition-all duration-300 ease-in-out lg:w-[280px]"
>
  <div
    class="sticky top-4 min-h-[calc(100vh-32px)] rounded-lg bg-white shadow-md"
  >
    <div
      class="flex h-16 items-center justify-center gap-2 px-1 lg:justify-start lg:px-6"
    >
      <img src="/logo.svg" alt="logo" class="h-7 w-7" />
      <h2 class="hidden text-2xl font-bold lg:block">StreamFlow</h2>
    </div>

    <div class="h-[1px] w-full bg-[#E8E8E8]"></div>

    <ul class="mt-10 flex flex-col gap-2 px-1 lg:px-4">
      {#each menuItems as item (item)}
        <li>
          <SidebarItem {...item}></SidebarItem>
        </li>
      {/each}
    </ul>
  </div>
</aside>
