<script lang="ts">
  const metrics = [
    { label: "Disk Space Used (Bytes)", value: "5242880000" },
    { label: "Index Size (Bytes)", value: "104857600" },
  ];
</script>

<div
  class="mt-4 w-full divide-y divide-gray-200 text-sm font-medium text-gray-800"
>
  {#each metrics as metric}
    <div class="flex items-center justify-between py-2">
      <span class="text-gray-600">{metric.label}</span>
      <span class="font-normal">{metric.value}</span>
    </div>
  {/each}
</div>
