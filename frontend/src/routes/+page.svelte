<script lang="ts">
  import StatCard from "../components/StatCard.svelte";

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
