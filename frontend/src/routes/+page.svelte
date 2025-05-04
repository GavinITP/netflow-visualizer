<script lang="ts">
  import { onMount } from "svelte";
  import ProtocolChart from "../components/ProtocolChart.svelte";
  import StatCard from "../components/StatCard.svelte";
  import TotalFlowCountChart from "../components/TotalFlowCountChart.svelte";

  const BASE = import.meta.env.VITE_APP_BASE_URL;

  type OverviewStats = {
    total_bytes: number;
    total_packets: number;
    total_flow_count: number;
    anomaly_flow_count: number;
    protocol_distribution: Record<string, number>;
  };

  let stats: OverviewStats = {
    total_bytes: 0,
    total_packets: 0,
    total_flow_count: 0,
    anomaly_flow_count: 0,
    protocol_distribution: { tcp: 0, udp: 0, icmp: 0, other: 0 },
  };
  let error: string | null = null;

  let packetHistory: { time: string; count: number }[] = [];

  let netflows: any[] = [];

  onMount(() => {
    const netflowWs = new WebSocket(`ws://${BASE}/api/netflows`);
    netflowWs.addEventListener("message", (e) => {
      try {
        netflows = JSON.parse(e.data);
      } catch {}
    });
    netflowWs.addEventListener("error", () => {});

    const overviewWs = new WebSocket(`ws://${BASE}/api/overview`);
    overviewWs.addEventListener("message", (e) => {
      try {
        stats = JSON.parse(e.data);
        error = null;
      } catch {
        error = "Malformed overview data";
      }
    });
    overviewWs.addEventListener("error", () => {
      error = "Overview WebSocket error";
    });

    const flowWs = new WebSocket(`ws://${BASE}/api/overview-flow-count`);
    flowWs.addEventListener("message", (e) => {
      try {
        packetHistory = JSON.parse(e.data);
        error = null;
      } catch {
        error = "Malformed flow-history data";
      }
    });
    flowWs.addEventListener("error", () => {
      error = "Flow-history WebSocket error";
    });

    return () => {
      netflowWs.close();
      overviewWs.close();
      flowWs.close();
    };
  });

  interface StatCardData {
    title: string;
    value: number | string | null;
    iconPath: string;
    bgColor: string;
  }

  let statCardData: StatCardData[];
  $: statCardData = [
    {
      title: "Total Flow Count",
      value: stats.total_flow_count,
      iconPath: "total-flow-count.svg",
      bgColor: "#FEF9C3",
    },
    {
      title: "Anomaly Flow Count",
      value: stats.anomaly_flow_count,
      iconPath: "/anomaly-flow-count.svg",
      bgColor: "#F3E8FF",
    },
    {
      title: "Total Packets (dPkts)",
      value: stats.total_packets,
      iconPath: "/total-packets.svg",
      bgColor: "#DBEAFE",
    },
    {
      title: "Total Bytes (dOctets)",
      value: stats.total_bytes,
      iconPath: "/total-bytes.svg",
      bgColor: "#DCFCE7",
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
  <div
    class="h-[500px] rounded-md bg-white px-5 pb-10 pt-5 shadow-md transition-transform hover:scale-[1.01]"
  >
    <h2 class="text-xl font-semibold">Total Flow Count Chart</h2>
    <TotalFlowCountChart {packetHistory} />
  </div>

  <div
    class="h-[500px] rounded-md bg-white px-5 pb-10 pt-5 shadow-md transition-transform hover:scale-[1.01]"
  >
    <h2 class="text-xl font-semibold">Protocol Distribution</h2>
    <ProtocolChart {protocolData} />
  </div>
</section>
