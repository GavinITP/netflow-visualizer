<script lang="ts">
  let {
    synchronousMode,
    schemaVersion,
    userVersion,
    foreignKeyEnforce,
    journalMode,
  } = $props();

  let metrics = $derived([
    { label: "Synchronous Mode", value: String(synchronousMode) },
    { label: "Schema Version", value: String(schemaVersion) },
    { label: "User Version", value: String(userVersion) },
    { label: "Foreign Key Enforcement", value: String(foreignKeyEnforce) },
    { label: "Journal Mode", value: String(journalMode).toUpperCase() },
  ]);
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
