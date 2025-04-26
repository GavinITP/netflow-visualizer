<script lang="ts">
  import { onMount } from "svelte";
  import AnomalyTable from "../../components/AnomalyTable.svelte";

  let ip = "";
  let recentCount = "";
  let port = "";
  let protocol = "";

  let tableData: any[] = [];

  async function loadTable() {
    const params = new URLSearchParams();
    if (ip) params.set("search", ip);
    if (recentCount) params.set("recent_count", recentCount);
    if (port) params.set("port", port);
    if (protocol) params.set("protocol", protocol);

    const url = `http://${import.meta.env.VITE_APP_BASE_URL}/api/netflows?${params.toString()}`;
    const res = await fetch(url);
    if (res.ok) {
      tableData = await res.json();
    } else {
      console.error("Failed to load netflows", res.status);
      tableData = [];
    }
  }

  onMount(loadTable);
</script>

<h1 class="mb-4 text-4xl font-extrabold">Anomaly Traffic Logs</h1>
<p class="text-gray">Access and analyze historical NetFlow traffic data</p>

<!-- FILTER FORM -->
<form
  class="my-10 flex flex-wrap gap-4 rounded-md bg-white p-4 shadow-md"
  on:submit|preventDefault={loadTable}
>
  <!-- IP Address -->
  <div class="flex min-w-[200px] flex-1 flex-col">
    <label for="ip" class="mb-1 text-sm font-medium">IP Address</label>
    <input
      id="ip"
      type="text"
      bind:value={ip}
      placeholder="e.g. 192.168.1.1"
      class="rounded-lg border border-gray-400 px-3 py-2 focus:border-blue-300 focus:ring"
    />
  </div>

  <!-- Recent Count -->
  <div class="flex w-32 flex-col">
    <label for="recent-count" class="mb-1 text-sm font-medium"
      >Recent count</label
    >
    <input
      id="recent-count"
      type="number"
      bind:value={recentCount}
      placeholder="e.g. 1000000"
      class="rounded-lg border border-gray-400 px-3 py-2 focus:border-blue-300 focus:ring"
    />
  </div>

  <!-- Port -->
  <div class="flex w-32 flex-col">
    <label for="port" class="mb-1 text-sm font-medium">Port</label>
    <input
      id="port"
      type="number"
      bind:value={port}
      placeholder="e.g. 80"
      class="rounded-lg border border-gray-400 px-3 py-2 focus:border-blue-300 focus:ring"
    />
  </div>

  <!-- Protocol -->
  <div class="flex w-32 flex-col">
    <label for="protocol" class="mb-1 text-sm font-medium">Protocol</label>
    <select
      id="protocol"
      bind:value={protocol}
      class="rounded-lg border border-gray-400 px-3 py-2 focus:border-blue-300 focus:ring"
    >
      <option value="">All</option>
      <option value="TCP">TCP</option>
      <option value="UDP">UDP</option>
      <option value="OTHERS">Others</option>
    </select>
  </div>

  <!-- Search Button -->
  <div class="flex items-end">
    <button
      type="submit"
      class="cursor-pointer rounded-lg bg-black px-4 py-2 text-white hover:bg-gray-700"
    >
      Search
    </button>
  </div>
</form>

<!-- RESULTS TABLE -->
<div class="rounded-md bg-white p-5 shadow-md">
  <h2 class="mb-4 text-xl font-semibold">History</h2>
  <AnomalyTable {tableData} />
</div>
