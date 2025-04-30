<script lang="ts">
  import AnomalyTable from "../../components/AnomalyTable.svelte";

  let ip = "";
  let recentCount = "50";
  let port = "";
  let protocol = "";

  let tableData: any[] = [];
  let loading = false;
  let error: string | null = null;

  const BASE = import.meta.env.VITE_APP_BASE_URL;
  const TIMEOUT_MS = 20000;
  const MAX_ROWS = 1000;

  async function loadTable() {
    loading = true;
    error = null;
    tableData = [];

    const params = new URLSearchParams();
    if (ip) params.set("search", ip);
    if (recentCount) params.set("recent_count", recentCount);
    if (port) params.set("port", port);
    if (protocol) params.set("protocol", protocol);

    const url = `http://${BASE}/api/anomaly-logs?${params.toString()}`;
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), TIMEOUT_MS);

    try {
      const res = await fetch(url, { signal: controller.signal });
      clearTimeout(timeoutId);

      if (!res.ok) {
        error = `Server error: ${res.status}`;
        return;
      }

      const data = await res.json();
      if (Array.isArray(data)) {
        if (data.length > MAX_ROWS) {
          tableData = data.slice(0, MAX_ROWS);
          // error = `Response too large; showing first ${MAX_ROWS} rows.`;
        } else {
          tableData = data;
        }
      } else {
        error = "Unexpected response format";
      }
    } catch (e: any) {
      clearTimeout(timeoutId);
      if (e.name === "AbortError") {
        error = "Request timed out";
      } else {
        error = `Fetch error: ${e.message}`;
      }
    } finally {
      loading = false;
    }
  }
</script>

<h1 class="mb-4 text-4xl font-extrabold">Anomaly Traffic Logs</h1>
<p class="text-gray">Access and analyze historical NetFlow traffic data</p>

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
      >Recent Count</label
    >
    <select
      id="recent-count"
      bind:value={recentCount}
      class="rounded-lg border border-gray-400 px-3 py-2 focus:border-blue-300 focus:ring"
    >
      <option value="50">50</option>
      <option value="100">100</option>
      <option value="500">500</option>
      <option value="1000">1000</option>
    </select>
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

<div class="rounded-md bg-white p-5 shadow-md">
  <h2 class="mb-4 text-xl font-semibold">History</h2>

  {#if loading}
    <p class="py-4 text-center">Loading…</p>
  {:else if tableData.length === 0}
    <p class="py-4 text-center text-gray-600">No results found</p>
  {:else if error}
    <p class="py-4 text-center text-red-600">{error}</p>
  {:else}
    <AnomalyTable {tableData} />
  {/if}
</div>
