<script lang="ts">
  import { onMount } from "svelte";
  import AnomalyTable from "../../components/AnomalyTable.svelte";

  let ip = $state("");
  let port = $state("");
  let protocol = $state("");
  let dateFrom = $state("");
  let dateTo = $state("");

  let tableData: any[] = $state([]);

  async function loadTable() {
    const params = new URLSearchParams();
    if (ip) params.set("search", ip);
    if (port) params.set("port", port);
    if (protocol) params.set("search", protocol);
    if (dateFrom) params.set("from", new Date(dateFrom).toISOString());
    if (dateTo) params.set("to", new Date(dateTo).toISOString());

    const res = await fetch(
      `http://${import.meta.env.VITE_APP_BASE_URL}/api/netflows?${params.toString()}`,
    );
    if (res.ok) {
      tableData = await res.json();
    } else {
      console.error("Failed to load netflows", res.status);
      tableData = [];
    }
  }

  onMount(() => {
    loadTable();
  });
</script>

<h1 class="mb-4 text-4xl font-extrabold">Anomaly Traffic Logs</h1>
<p class="text-gray">Access and analyze historical NetFlow traffic data</p>

<div class="my-10 flex flex-wrap gap-4 rounded-md bg-white p-4 shadow-md">
  <div class="flex w-full flex-wrap gap-4 md:w-full lg:w-2/5 lg:flex-nowrap">
    <!-- IP Address Input -->
    <div class="flex w-full flex-col lg:w-2/3">
      <label for="ip" class="mb-1 text-sm font-medium">IP Address</label>
      <input
        id="ip"
        type="text"
        bind:value={ip}
        placeholder="Enter IP address"
        class="rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-300 focus:ring focus:outline-none"
      />
    </div>

    <!-- Port Input -->
    <div class="flex w-full flex-col lg:w-1/3">
      <label for="port" class="mb-1 text-sm font-medium">Port</label>
      <input
        id="port"
        type="number"
        bind:value={port}
        placeholder="Enter port number"
        class="rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-300 focus:ring focus:outline-none"
      />
    </div>
  </div>

  <!-- Protocol Selector -->
  <div class="flex w-full flex-col lg:w-1/6">
    <label for="protocol" class="mb-1 text-sm font-medium">Protocol</label>
    <select
      id="protocol"
      bind:value={protocol}
      class="rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-300 focus:ring focus:outline-none"
    >
      <option value="">All</option>
      <option value="http">TCP</option>
      <option value="https">UDP</option>
    </select>
  </div>

  <!-- Date Range Group -->
  <div class="flex w-full flex-wrap gap-4 md:w-full lg:w-1/3 lg:flex-nowrap">
    <!-- Date From -->
    <div class="flex w-full flex-col lg:w-1/2">
      <label for="dateFrom" class="mb-1 text-sm font-medium">Date From</label>
      <input
        id="dateFrom"
        type="date"
        bind:value={dateFrom}
        class="rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-300 focus:ring focus:outline-none"
      />
    </div>

    <!-- Date To -->
    <div class="flex w-full flex-col lg:w-1/2">
      <label for="dateTo" class="mb-1 text-sm font-medium">Date To</label>
      <input
        id="dateTo"
        type="date"
        bind:value={dateTo}
        class="rounded-lg border border-gray-300 px-3 py-2 focus:border-blue-300 focus:ring focus:outline-none"
      />
    </div>
  </div>
</div>

<!-- History  -->
<div class="rounded-md bg-white p-5 shadow-md">
  <h2 class="text-xl font-semibold">History</h2>

  <AnomalyTable {tableData} />
</div>
