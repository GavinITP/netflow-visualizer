<script lang="ts">
  import NetworkChart from "../components/NetworkChart.svelte";
  import ProtocolChart from "../components/ProtocolChart.svelte";
  import StatCard from "../components/StatCard.svelte";
  import AnomalyTable from "../components/AnomalyTable.svelte";

  interface StatCardData {
    title: string;
    value: number | string;
    iconPath: string;
    bgColor: string;
  }

  const statCardData: StatCardData[] = [
    {
      title: "Total Packets",
      value: 123456,
      iconPath: "/total-packets.svg",
      bgColor: "#DBEAFE",
    },
    {
      title: "Packets / Second",
      value: 2345,
      iconPath: "/traffic-rate.svg",
      bgColor: "#DCFCE7",
    },
    {
      title: "Active Alerts",
      value: 6,
      iconPath: "/crisis-alert.svg",
      bgColor: "#FEF9C3",
    },
    {
      title: "Uptime",
      value: "04:32:12",
      iconPath: "/uptime.svg",
      bgColor: "#F3E8FF",
    },
  ];
</script>

<h1 class="mb-4 text-4xl font-extrabold">Network Overview</h1>
<p class="text-gray">
  Discover the essentials of network traffic analysis and monitoring
</p>

<div
  class="my-10 grid grid-cols-1 justify-between gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4"
>
  {#each statCardData as data (data.title)}
    <StatCard {...data} />
  {/each}
</div>

<section class="my-10 grid grid-cols-1 gap-6 lg:grid-cols-[3fr_2fr]">
  <!-- Anomaly Network Graph -->
  <div
    class="h-[360px] rounded-md bg-white px-5 pb-10 pt-5 shadow-md transition-transform hover:scale-[1.01]"
  >
    <h2 class="text-xl font-semibold">Anomaly Network Graph</h2>

    <NetworkChart />
  </div>

  <!-- Protocol Distribution -->
  <div
    class="h-[360px] rounded-md bg-white px-5 pb-10 pt-5 shadow-md transition-transform hover:scale-[1.01]"
  >
    <h2 class="text-xl font-semibold">Protocol Distribution</h2>

    <ProtocolChart />
  </div>
</section>

<!-- Anomaly Network Events -->
<div class="rounded-md bg-white p-5 shadow-md">
  <h2 class="text-xl font-semibold">Anomaly Network Events</h2>

  <AnomalyTable />
</div>
