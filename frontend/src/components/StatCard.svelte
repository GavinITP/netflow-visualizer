<script lang="ts">
  const { title, value, iconPath, bgColor } = $props();

  const formatValue = (val: string | number) => {
    return typeof val === "number" ? new Intl.NumberFormat().format(val) : val;
  };
</script>

<div
  class="w-full transform rounded-md bg-white p-5 shadow-md transition-transform hover:scale-[1.02] sm:max-w-[400px]"
>
  <div class="flex items-center gap-4">
    <div
      class="flex h-[48px] w-[48px] items-center justify-center rounded-full"
      style={"background-color: " + bgColor}
    >
      <img src={iconPath} alt={title} class="h-5 w-5" />
    </div>

    <div class="flex flex-col justify-between">
      <p class="text-gray text-sm font-bold">{title}</p>
      <p class="text-xl font-bold">{formatValue(value)}</p>
    </div>
  </div>
</div>
