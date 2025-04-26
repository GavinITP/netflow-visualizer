<script lang="ts">
  import { onMount } from "svelte";
  import NetworkChart from "../components/TotalFlowCountChart.svelte";
  import ProtocolChart from "../components/ProtocolChart.svelte";
  import StatCard from "../components/StatCard.svelte";
  import TotalFlowCountChart from "../components/TotalFlowCountChart.svelte";

  const BASE = import.meta.env.VITE_APP_BASE_URL;

  type NetflowStats = {
    total_packets: number;
    packets_per_second: number;
    active_alert: number;
    uptime: { hours: number; minutes: number; seconds: number };
    protocol_distribution: Record<string, number>;
  };

  let stats: NetflowStats = {
    total_packets: 0,
    packets_per_second: 0,
    active_alert: 0,
    uptime: { hours: 0, minutes: 0, seconds: 0 },
    protocol_distribution: {
      TCP: 0,
      UDP: 0,
      ICMP: 0,
      Others: 0,
    },
  };
  let error: string | null = null;

  let packetHistory: { time: string; count: number }[] = [
    // { time: "12:00:00", count: 12 },
    // { time: "12:01:00", count: 19 },
    // { time: "12:02:00", count: 7 },
    // { time: "12:03:00", count: 14 },
    // { time: "12:04:00", count: 22 },
    // { time: "12:05:00", count: 18 },
    // { time: "12:06:00", count: 25 },
    // { time: "12:07:00", count: 9 },
    // { time: "12:08:00", count: 16 },
    // { time: "12:09:00", count: 20 },
  ];

  let netflows: any[] = [];

  onMount(() => {
    fetch(`http://${BASE}/api/netflows`)
      .then((res) => res.json())
      .then((data) => {
        netflows = data;
      })
      .catch(() => {});

    const ws = new WebSocket(`ws://${BASE}/api/netflow-stats`);

    ws.addEventListener("message", (e) => {
      try {
        stats = JSON.parse(e.data);
        error = null;

        const now = new Date().toLocaleTimeString();
        packetHistory = [
          ...packetHistory,
          {
            time: now,
            count: stats.packets_per_second,
          },
        ].slice(-10);
      } catch {
        error = "Malformed stats data";
      }
    });

    ws.addEventListener("error", () => {
      error = "WebSocket error";
    });

    return () => ws.close();
  });

  interface StatCardData {
    title: string;
    value: number | string;
    iconPath: string;
    bgColor: string;
  }

  $: statCardData = [
    {
      title: "Total Packets (dPkts)",
      value: stats.total_packets,
      iconPath: "/total-packets.svg",
      bgColor: "#DBEAFE",
    },
    {
      title: "Total Bytes (dOctets)",
      value: stats.packets_per_second,
      iconPath: "/traffic-rate.svg",
      bgColor: "#DCFCE7",
    },
    {
      title: "Total Flow Count",
      value: stats.active_alert,
      iconPath: "/crisis-alert.svg",
      bgColor: "#FEF9C3",
    },
    {
      title: "Anomaly Flow Count",
      value: null,
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
  <!-- Total Flow Count Chart -->
  <div
    class="h-[500px] rounded-md bg-white px-5 pt-5 pb-10 shadow-md transition-transform hover:scale-[1.01]"
  >
    <h2 class="text-xl font-semibold">Total Flow Count Chart</h2>

    <TotalFlowCountChart {packetHistory} />
  </div>

  <!-- Protocol Distribution -->
  <div
    class="h-[500px] rounded-md bg-white px-5 pt-5 pb-10 shadow-md transition-transform hover:scale-[1.01]"
  >
    <h2 class="text-xl font-semibold">Protocol Distribution</h2>

    <ProtocolChart {protocolData} />
  </div>
</section>
