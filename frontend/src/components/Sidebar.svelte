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
    class="sticky top-4 flex min-h-[calc(100vh-32px)] flex-col justify-between rounded-lg bg-white shadow-md"
  >
    <div>
      <div
        class="flex h-16 items-center justify-center gap-2 border-b border-gray-200 px-1 lg:justify-start lg:px-6"
      >
        <img src="/logo.svg" alt="logo" class="h-7 w-7" />
        <h2 class="hidden text-2xl font-bold lg:block">StreamFlow</h2>
      </div>

      <ul class="mt-10 flex flex-col gap-2 px-1 lg:px-4">
        {#each menuItems as item (item)}
          <li>
            <SidebarItem {...item}></SidebarItem>
          </li>
        {/each}
      </ul>
    </div>

    <div
      class="flex cursor-pointer items-center gap-3 border-t border-gray-200 p-4 hover:bg-[#f2f2f2]"
    >
      <img src="/account.svg" alt="User" class="h-8 w-8 rounded-full" />
      <a href="/#" class="hidden overflow-hidden lg:block">
        <p class="truncate text-sm font-medium text-gray-800">
          Phumipat Chaiprasertsud
        </p>
        <p class="text-xs text-gray-500">Admin</p>
      </a>
    </div>
  </div>
</aside>
