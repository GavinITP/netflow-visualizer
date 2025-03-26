<script lang="ts">
  import Logs from "../../components/Logs.svelte";
  import StatCard from "../../components/StatCard.svelte";

  interface StatCardData {
    title: string;
    value: number | string;
    iconPath: string;
    bgColor: string;
  }

  const statCardData: StatCardData[] = [
    {
      title: "Total Series Count",
      value: 3456789,
      iconPath: "/total-series-count.svg",
      bgColor: "#DBEAFE",
    },
    {
      title: "Measurement Count",
      value: 3,
      iconPath: "/measurement-count.svg",
      bgColor: "#DCFCE7",
    },
    {
      title: "Retention Policies",
      value: "30 Days",
      iconPath: "/policy.svg",
      bgColor: "#FEF9C3",
    },
    {
      title: "Number of Databases",
      value: 2,
      iconPath: "/database.svg",
      bgColor: "#F3E8FF",
    },
  ];
</script>

<h1 class="mb-4 text-4xl font-extrabold">Database Statistics</h1>
<p class="text-gray">
  Monitor detailed statistics from the InfluxDB time-series database
</p>

<div
  class="my-10 grid grid-cols-1 justify-between gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4"
>
  {#each statCardData as data (data.title)}
    <StatCard {...data} />
  {/each}
</div>

<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
  <div class="rounded-md bg-white p-5 shadow-md">
    <h2 class="text-xl font-semibold">Recent Database Logs</h2>

    <Logs />
  </div>

  <div class="flex flex-col gap-6">
    <div class="rounded-md bg-white p-5 shadow-md">
      <h2 class="text-xl font-semibold">Query Performance</h2>
    </div>

    <div class="rounded-md bg-white p-5 shadow-md">
      <h2 class="text-xl font-semibold">Storage Usage</h2>
    </div>
  </div>
</div>
