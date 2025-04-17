<script lang="ts">
  import NetworkChart from "../components/NetworkChart.svelte";
  import ProtocolChart from "../components/ProtocolChart.svelte";
  import StatCard from "../components/StatCard.svelte";
  import AnomalyTable from "../components/AnomalyTable.svelte";

  import { onMount } from "svelte";

  const BASE = import.meta.env.VITE_APP_BASE_URL;

  type NetflowStats = {
    total_packets: number;
    packets_per_second: number;
    active_alert: number;
    uptime: {
      hours: number;
      minutes: number;
      seconds: number;
    };
    protocol_distribution: Record<string, number>;
  };

  interface StatCardData {
    title: string;
    value: number | string;
    iconPath: string;
    bgColor: string;
  }

  let stats: NetflowStats = {
    total_packets: 0,
    packets_per_second: 0,
    active_alert: 0,
    uptime: { hours: 0, minutes: 0, seconds: 0 },
    protocol_distribution: {},
  };

  let error: string | null = null;

  let protocolData: Record<string, number> = {};

  let netflows: any[] = [];

  onMount(() => {
    const ws = new WebSocket(`ws://${BASE}/api/netflow-stats`);
    fetch(`http://${BASE}/api/netflows`)
      .then((res) => res.json())
      .then((data) => {
        netflows = data;
        console.log("netflows", data);
      })
      .catch(() => {});

    ws.addEventListener("message", (e) => {
      try {
        stats = JSON.parse(e.data);
        error = null;
      } catch {
        error = "Malformed stats data";
      }
    });

    ws.addEventListener("error", () => {
      error = "WebSocket error";
    });

    return () => {
      ws.close();
    };
  });

  let statCardData: StatCardData[];

  $: statCardData = [
    {
      title: "Total Packets",
      value: stats.total_packets,
      iconPath: "/total-packets.svg",
      bgColor: "#DBEAFE",
    },
    {
      title: "Packets / Second",
      value: stats.packets_per_second,
      iconPath: "/traffic-rate.svg",
      bgColor: "#DCFCE7",
    },
    {
      title: "Active Alerts",
      value: stats.active_alert,
      iconPath: "/crisis-alert.svg",
      bgColor: "#FEF9C3",
    },
    {
      title: "Uptime",
      value: `${stats.uptime.hours
        .toString()
        .padStart(2, "0")}:${stats.uptime.minutes
        .toString()
        .padStart(2, "0")}:${stats.uptime.seconds.toString().padStart(2, "0")}`,
      iconPath: "/uptime.svg",
      bgColor: "#F3E8FF",
    },
  ];

  $: protocolData = { ...stats.protocol_distribution };
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
    class="h-[360px] rounded-md bg-white px-5 pt-5 pb-10 shadow-md transition-transform hover:scale-[1.01]"
  >
    <h2 class="text-xl font-semibold">Anomaly Network Graph</h2>

    <NetworkChart />
  </div>

  <!-- Protocol Distribution -->
  <div
    class="h-[360px] rounded-md bg-white px-5 pt-5 pb-10 shadow-md transition-transform hover:scale-[1.01]"
  >
    <h2 class="text-xl font-semibold">Protocol Distribution</h2>

    <ProtocolChart {protocolData} />
  </div>
</section>

<!-- Anomaly Network Events -->
<div class="rounded-md bg-white p-5 shadow-md">
  <h2 class="text-xl font-semibold">Anomaly Network Events</h2>

  <AnomalyTable tableData={netflows.slice(0, 10)} />
</div>
