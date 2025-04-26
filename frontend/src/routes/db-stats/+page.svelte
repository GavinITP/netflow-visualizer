<script lang="ts">
  import { onMount } from "svelte";
  import Logs from "../../components/Logs.svelte";
  import StatCard from "../../components/StatCard.svelte";
  import StorageUsageTable from "../../components/StorageUsageTable.svelte";
  import ConfigParams from "../../components/ConfigParams.svelte";

  const BASE = import.meta.env.VITE_APP_BASE_URL;

  type DBStats = {
    page_size: number;
    page_count: number;
    db_size_bytes: number;
    freelist_count: number;
    cache_size: number;
    synchronous: number;
    schema_version: number;
    user_version: number;
    foreign_keys: number;
    journal_mode: string;
    integrity_check: string;
    db_abs_path: string;
    open_connections: number;
  };

  let stats: DBStats = {
    page_size: 0,
    page_count: 0,
    db_size_bytes: 0,
    freelist_count: 0,
    cache_size: 0,
    synchronous: 0,
    schema_version: 0,
    user_version: 0,
    foreign_keys: 0,
    journal_mode: "",
    integrity_check: "",
    db_abs_path: "",
    open_connections: 0,
  };
  let error: string | null = null;
  let lastCheckDate = "";

  onMount(() => {
    const ws = new WebSocket(`ws://${BASE}/api/db-stats`);

    ws.addEventListener("open", () => {});

    ws.addEventListener("message", (e) => {
      try {
        stats = JSON.parse(e.data) as DBStats;
        lastCheckDate = new Date().toLocaleString("en-US", {
          month: "long",
          day: "numeric",
          year: "numeric",
          hour: "2-digit",
          minute: "2-digit",
          hour12: true,
        });

        error = null;
      } catch {
        error = "Failed to parse DB stats";
      }
    });

    ws.addEventListener("error", () => {
      error = "WebSocket error";
    });

    return () => {
      ws.close();
    };
  });

  $: databaseSize = stats.db_size_bytes;
  $: cacheSize = stats.cache_size;
  $: freelistCount = stats.freelist_count;
  $: openConnections = stats.open_connections;
  $: synchronousMode = stats.synchronous;
  $: schemaVersion = stats.schema_version;
  $: userVersion = stats.user_version;
  $: foreignKeyEnforce = stats.foreign_keys;
  $: journalMode = stats.journal_mode;
  $: integrityStatus = stats.integrity_check;
  $: databaseFilePath = stats.db_abs_path;

  interface StatCardData {
    title: string;
    value: number | string;
    iconPath: string;
    bgColor: string;
  }

  let statCardData: StatCardData[];
  $: statCardData = [
    {
      title: "Database Size",
      value: databaseSize,
      iconPath: "/total-series-count.svg",
      bgColor: "#DBEAFE",
    },
    {
      title: "Cache Size",
      value: cacheSize,
      iconPath: "/measurement-count.svg",
      bgColor: "#DCFCE7",
    },
    {
      title: "Freelist Pages",
      value: freelistCount,
      iconPath: "/policy.svg",
      bgColor: "#FEF9C3",
    },
    {
      title: "Open Connections",
      value: openConnections,
      iconPath: "/connection.svg",
      bgColor: "#F3E8FF",
    },
  ];
</script>

<h1 class="mb-4 text-4xl font-extrabold">Database Statistics</h1>
<p class="text-gray">Monitor detailed statistics from the SQLite database</p>

<div class="my-6 flex items-center gap-4">
  <span class="mb-1 font-semibold text-gray-500">Database Path:</span>
  <span class="flex items-center">
    <span
      class="flex-grow overflow-x-auto rounded-md bg-gray-100 p-2 font-mono text-sm whitespace-nowrap text-blue-600"
      id="file-path-display">{databaseFilePath}</span
    >
  </span>
</div>

<div
  class="mb-10 grid grid-cols-1 justify-between gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4"
>
  {#each statCardData as data (data.title)}
    <StatCard {...data} />
  {/each}
</div>

<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
  <div class="rounded-md bg-white p-5 shadow-md">
    <h2 class="text-xl font-semibold">Recent SQL Query Logs</h2>

    <Logs />
  </div>

  <div class="flex flex-col gap-6">
    <div class="rounded-md bg-white p-5 shadow-md">
      <h2 class="text-xl font-semibold">Configuration Parameters</h2>

      <ConfigParams
        {synchronousMode}
        {schemaVersion}
        {userVersion}
        {foreignKeyEnforce}
        {journalMode}
      />
    </div>

    <div class="rounded-md bg-white p-5 shadow-md">
      <h2 class="text-xl font-semibold">Integrity Check</h2>

      <StorageUsageTable status={integrityStatus} {lastCheckDate} />
    </div>
  </div>
</div>
