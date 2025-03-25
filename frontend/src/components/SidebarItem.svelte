<script lang="ts">
  import { page } from "$app/state";

  const { href, title, iconSrc } = $props();
</script>

<a
  {href}
  class="flex h-[46px] w-full items-center gap-4 rounded-md px-4 hover:bg-[#f2f2f2]"
  class:bg-blue={page.url.pathname === href}
>
  <img
    class="text-gray"
    class:text-gray-900={page.url.pathname === href}
    src={iconSrc}
    alt={title}
  />
  <h3
    class="text-gray hidden capitalize lg:block"
    class:text-gray-900={page.url.pathname === href}
  >
    {title}
  </h3>
</a>
