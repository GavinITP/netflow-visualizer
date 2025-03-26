<script lang="ts">
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
  Detailed statistics from the InfluxDB time-series database
</p>

<div
  class="my-10 grid grid-cols-1 justify-between gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4"
>
  {#each statCardData as data (data.title)}
    <StatCard {...data} />
  {/each}
</div>
