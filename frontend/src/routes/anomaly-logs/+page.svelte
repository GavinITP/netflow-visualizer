<script lang="ts">
  import AnomalyTable from "../../components/AnomalyTable.svelte";

  let ip = "";
  let recentCount = "50";
  let port = "";
  let protocol = "";

  let tableData: any[] = [];
  let loading = false;
  let error: string | null = null;

  // Validation state
  let ipError = false;
  let portError = false;

  const BASE = import.meta.env.VITE_APP_BASE_URL;
  const TIMEOUT_MS = 20000;
  const MAX_ROWS = 1000;

  function isValidIp(ip: string): boolean {
    const regex =
      /^(25[0-5]|2[0-4][0-9]|1?[0-9]{1,2})(\.(25[0-5]|2[0-4][0-9]|1?[0-9]{1,2})){3}$/;
    return regex.test(ip);
  }

  function ipToInt(ip: string): number {
    return (
      ip
        .split(".")
        .reduce((acc, octet) => (acc << 8) + parseInt(octet, 10), 0) >>> 0
    );
  }

  function intToIp(int: number | string): string {
    const num = typeof int === "string" ? parseInt(int, 10) : int;
    return [
      (num >>> 24) & 255,
      (num >>> 16) & 255,
      (num >>> 8) & 255,
      num & 255,
    ].join(".");
  }

  async function loadTable() {
    loading = false;
    error = null;
    tableData = [];

    ipError = false;
    portError = false;

    // Validate IP
    if (ip && !isValidIp(ip)) {
      ipError = true;
      error = "Invalid IP address format.";
      return;
    }

    // Validate Port
    const portNumber = Number(port);
    if (port && (isNaN(portNumber) || portNumber < 1 || portNumber > 65535)) {
      portError = true;
      error = "Port must be a number between 1 and 65535.";
      return;
    }

    loading = true;

    const params = new URLSearchParams();
    if (ip) {
      const ipInt = ipToInt(ip);
      params.set("search", ipInt.toString());
    }
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
        const processed = data.map((entry) => ({
          ...entry,
          srcaddr: intToIp(entry.srcaddr),
          dstaddr: intToIp(entry.dstaddr),
          nexthop: intToIp(entry.nexthop),
        }));

        tableData = processed.slice(0, MAX_ROWS);
      } else {
        error = "Unexpected response format";
      }
    } catch (e: any) {
      clearTimeout(timeoutId);
      error =
        e.name === "AbortError"
          ? "Request timed out"
          : `Fetch error: ${e.message}`;
    } finally {
      loading = false;
    }
  }
</script>

<h1 class="mb-4 text-4xl font-extrabold">Anomaly Traffic Logs</h1>
<p class="text-gray">
  Access and analyze historical NetFlow traffic data from the 10 most recent CSV
  files
</p>

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
      class="rounded-lg border px-3 py-2 focus:ring
        {ipError
        ? 'border-red-500 focus:border-red-500'
        : 'border-gray-400 focus:border-blue-300'}"
    />
    {#if ipError}
      <span class="mt-1 text-sm text-red-600"
        >Please enter a valid IPv4 address.</span
      >
    {/if}
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
      class="rounded-lg border px-3 py-2 focus:ring
        {portError
        ? 'border-red-500 focus:border-red-500'
        : 'border-gray-400 focus:border-blue-300'}"
    />
    {#if portError}
      <span class="mt-1 text-sm text-red-600"
        >Port must be between 1 and 65535.</span
      >
    {/if}
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
  {:else if error}
    <p class="py-4 text-center text-red-600">{error}</p>
  {:else if tableData.length === 0}
    <p class="py-4 text-center text-gray-600">No results found</p>
  {:else}
    <AnomalyTable {tableData} />
  {/if}
</div>
